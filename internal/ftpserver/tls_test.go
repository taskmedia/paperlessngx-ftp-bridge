package ftpserver

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewTLSCertSourceGeneratesSelfSignedWhenDirEmpty(t *testing.T) {
	dir := t.TempDir()

	src, err := newTLSCertSource(dir, "ftp.example.org")
	if err != nil {
		t.Fatalf("newTLSCertSource() error = %v, want nil", err)
	}

	cert, err := src.GetCertificate()
	if err != nil {
		t.Fatalf("GetCertificate() error = %v, want nil", err)
	}
	assertCertSAN(t, cert, "ftp.example.org")
}

func TestNewTLSCertSourceFailsFastOnInvalidExistingCert(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "tls.crt"), "not a certificate")
	writeFile(t, filepath.Join(dir, "tls.key"), "not a key")

	_, err := newTLSCertSource(dir, "ftp.example.org")
	if err == nil {
		t.Fatalf("newTLSCertSource() error = nil, want a fatal load error for an invalid cert/key pair")
	}
}

func TestNewTLSCertSourceLoadsExistingCertWhenPresent(t *testing.T) {
	dir := t.TempDir()
	writeSelfSignedCertFiles(t, dir, "existing.example.org")

	src, err := newTLSCertSource(dir, "ftp.example.org")
	if err != nil {
		t.Fatalf("newTLSCertSource() error = %v, want nil", err)
	}

	cert, err := src.GetCertificate()
	if err != nil {
		t.Fatalf("GetCertificate() error = %v, want nil", err)
	}
	// The mounted cert's SAN wins over the self-signed fallback's.
	assertCertSAN(t, cert, "existing.example.org")
}

func TestGetCertificateReloadsOnMtimeChange(t *testing.T) {
	dir := t.TempDir()
	writeSelfSignedCertFiles(t, dir, "first.example.org")

	src, err := newTLSCertSource(dir, "ftp.example.org")
	if err != nil {
		t.Fatalf("newTLSCertSource() error = %v, want nil", err)
	}

	first, err := src.GetCertificate()
	if err != nil {
		t.Fatalf("GetCertificate() error = %v, want nil", err)
	}
	assertCertSAN(t, first, "first.example.org")

	// Simulate a renewed secret: new content, mtime must move forward so
	// the filesystem sees a change even on coarse mtime resolutions.
	time.Sleep(10 * time.Millisecond)
	writeSelfSignedCertFiles(t, dir, "renewed.example.org")

	second, err := src.GetCertificate()
	if err != nil {
		t.Fatalf("GetCertificate() error = %v (after rotation), want nil", err)
	}
	assertCertSAN(t, second, "renewed.example.org")
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()

	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("WriteFile(%q) error = %v", path, err)
	}
}

// writeSelfSignedCertFiles writes a freshly generated self-signed
// tls.crt/tls.key pair (SAN = host) into dir, as a stand-in for a mounted
// cert-manager Secret.
func writeSelfSignedCertFiles(t *testing.T, dir, host string) {
	t.Helper()

	certPEM, keyPEM, err := generateSelfSignedCertPEM(host)
	if err != nil {
		t.Fatalf("generateSelfSignedCertPEM() error = %v", err)
	}

	writeFile(t, filepath.Join(dir, "tls.crt"), string(certPEM))
	writeFile(t, filepath.Join(dir, "tls.key"), string(keyPEM))
}

func assertCertSAN(t *testing.T, cert *tls.Certificate, wantHost string) {
	t.Helper()

	if cert == nil {
		t.Fatalf("certificate = nil, want a certificate with SAN %q", wantHost)
	}

	leaf, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		t.Fatalf("ParseCertificate() error = %v", err)
	}

	for _, name := range leaf.DNSNames {
		if name == wantHost {
			return
		}
	}

	t.Fatalf("certificate SANs = %v, want to contain %q", leaf.DNSNames, wantHost)
}
