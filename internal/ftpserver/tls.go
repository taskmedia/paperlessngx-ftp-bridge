package ftpserver

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// DefaultTLSCertDir is the fixed, hardcoded path where an operator-supplied
// TLS Secret (tls.crt/tls.key, the kubernetes.io/tls shape) is mounted, per
// ADR-0003. Its presence is the toggle between the existing-secret path and
// the self-signed default; the path itself is not configurable, since the
// chart controls both the mount and the app binary in the same release.
const DefaultTLSCertDir = "/etc/ftp/tls"

const selfSignedCertValidity = 365 * 24 * time.Hour

// tlsCertSource supplies the TLS certificate offered to AUTH TLS clients.
// If certDir contains a tls.crt/tls.key pair it is loaded and re-checked on
// every call to GetCertificate, picking up a cert-manager renewal without a
// restart; otherwise a self-signed certificate generated once at
// construction is used for the lifetime of the process.
type tlsCertSource struct {
	certDir string

	selfSigned *tls.Certificate

	mu            sync.Mutex
	loaded        *tls.Certificate
	loadedModTime time.Time
}

// newTLSCertSource builds a tlsCertSource for certDir. publicHost is used as
// the self-signed certificate's SAN (ADR-0003 ties it to the configured PASV
// public host). A cert/key pair present in certDir but unreadable or invalid
// is a fatal error here, by design: an operator-supplied secret that fails
// to load is a genuine misconfiguration, not a transient condition.
func newTLSCertSource(certDir, publicHost string) (*tlsCertSource, error) {
	selfSigned, err := generateSelfSignedCert(publicHost)
	if err != nil {
		return nil, fmt.Errorf("generating self-signed TLS certificate: %w", err)
	}

	src := &tlsCertSource{certDir: certDir, selfSigned: selfSigned}

	if _, err := os.Stat(src.certPath()); err == nil {
		if _, err := src.loadFromDisk(); err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("checking for existing TLS secret in %q: %w", certDir, err)
	}

	return src, nil
}

// GetCertificate returns the certificate to offer the current connection,
// reloading it from disk first if certDir's tls.crt has changed since the
// last call.
func (s *tlsCertSource) GetCertificate() (*tls.Certificate, error) {
	info, err := os.Stat(s.certPath())
	if err != nil {
		if os.IsNotExist(err) {
			return s.selfSigned, nil
		}

		return nil, fmt.Errorf("checking for existing TLS secret in %q: %w", s.certDir, err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.loaded != nil && info.ModTime().Equal(s.loadedModTime) {
		return s.loaded, nil
	}

	return s.loadFromDisk()
}

// loadFromDisk reads and parses the cert/key pair in certDir, caching the
// result and the file's mtime. Callers hold s.mu, except during
// newTLSCertSource's initial startup check.
func (s *tlsCertSource) loadFromDisk() (*tls.Certificate, error) {
	info, err := os.Stat(s.certPath())
	if err != nil {
		return nil, fmt.Errorf("checking for existing TLS secret in %q: %w", s.certDir, err)
	}

	cert, err := tls.LoadX509KeyPair(s.certPath(), s.keyPath())
	if err != nil {
		return nil, fmt.Errorf("loading TLS secret from %q: %w", s.certDir, err)
	}

	s.loaded = &cert
	s.loadedModTime = info.ModTime()

	return s.loaded, nil
}

func (s *tlsCertSource) certPath() string {
	return filepath.Join(s.certDir, "tls.crt")
}

func (s *tlsCertSource) keyPath() string {
	return filepath.Join(s.certDir, "tls.key")
}

// generateSelfSignedCert builds a fresh, in-memory-only ECDSA P-256
// certificate (~1yr validity) with its SAN set to host, per ADR-0003.
func generateSelfSignedCert(host string) (*tls.Certificate, error) {
	certPEM, keyPEM, err := generateSelfSignedCertPEM(host)
	if err != nil {
		return nil, err
	}

	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, fmt.Errorf("parsing generated self-signed certificate: %w", err)
	}

	return &cert, nil
}

// generateSelfSignedCertPEM generates a self-signed ECDSA P-256 certificate
// and returns its PEM-encoded cert and key, for both production use and
// test fixtures that stand in for a mounted cert-manager Secret.
func generateSelfSignedCertPEM(host string) (certPEM, keyPEM []byte, err error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("generating ECDSA key: %w", err)
	}

	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, nil, fmt.Errorf("generating certificate serial number: %w", err)
	}

	now := time.Now()
	template := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      pkix.Name{CommonName: host},
		DNSNames:     []string{host},
		NotBefore:    now.Add(-time.Hour),
		NotAfter:     now.Add(selfSignedCertValidity),
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IsCA:         true,
	}

	der, err := x509.CreateCertificate(rand.Reader, template, template, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, fmt.Errorf("creating self-signed certificate: %w", err)
	}

	keyDER, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling ECDSA private key: %w", err)
	}

	certPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	keyPEM = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER})

	return certPEM, keyPEM, nil
}
