package ftpserver

import (
	"crypto/tls"
	"errors"
	"fmt"

	ftpserverlib "github.com/fclairamb/ftpserverlib"
)

// errInvalidCredentials is returned by AuthUser when the supplied
// username/password don't match the configured account.
var errInvalidCredentials = errors.New("ftpserver: invalid username or password")

// mainDriver implements ftpserverlib.MainDriver. Account configuration is
// intentionally minimal (a single username/password pair); it is widened by
// a follow-on ticket.
type mainDriver struct {
	cfg  Config
	cert *tlsCertSource
}

// newMainDriver builds a mainDriver, eagerly constructing (or loading) the
// TLS certificate it will offer. An operator-supplied cert/key pair that
// exists but fails to load is returned as an error here, by design, so
// callers can treat it as a fatal startup error rather than a per-connection
// one.
func newMainDriver(cfg Config) (*mainDriver, error) {
	cert, err := newTLSCertSource(cfg.tlsCertDir(), cfg.publicHost())
	if err != nil {
		return nil, fmt.Errorf("preparing TLS certificate: %w", err)
	}

	return &mainDriver{cfg: cfg, cert: cert}, nil
}

// GetSettings returns the embedded server's settings. PassiveTransferPortRange
// is left unset so passive data connections use ephemeral ports, matching
// the minimal/placeholder PASV configuration for this ticket.
func (d *mainDriver) GetSettings() (*ftpserverlib.Settings, error) {
	return &ftpserverlib.Settings{
		ListenAddr:          d.cfg.listenAddr(),
		DefaultTransferType: ftpserverlib.TransferTypeBinary,
	}, nil
}

// ClientConnected sends the welcome message for a newly connected client.
func (d *mainDriver) ClientConnected(_ ftpserverlib.ClientContext) (string, error) {
	return "paperlessngx-ftp-bridge", nil
}

// ClientDisconnected is a no-op: per-session state lives entirely in the
// ClientDriver returned by AuthUser and is garbage collected once dropped.
func (d *mainDriver) ClientDisconnected(_ ftpserverlib.ClientContext) {}

// AuthUser authenticates against the single configured account and, on
// success, returns a ClientDriver backed by a fresh in-memory filesystem.
func (d *mainDriver) AuthUser(_ ftpserverlib.ClientContext, user, pass string) (ftpserverlib.ClientDriver, error) {
	if user != d.cfg.Username || pass != d.cfg.Password {
		return nil, errInvalidCredentials
	}

	return newClientDriver(d.cfg.Uploader, d.cfg.allowedExtensions()), nil
}

// GetTLSConfig returns the certificate to offer for AUTH TLS: either the
// operator-supplied secret mounted at cfg.tlsCertDir(), reloaded whenever it
// changes on disk, or the self-signed default (ADR-0003). AUTH TLS is always
// offered this way; there is no enable/disable switch.
func (d *mainDriver) GetTLSConfig() (*tls.Config, error) {
	cert, err := d.cert.GetCertificate()
	if err != nil {
		return nil, fmt.Errorf("loading TLS certificate: %w", err)
	}

	return &tls.Config{Certificates: []tls.Certificate{*cert}}, nil
}

var _ ftpserverlib.MainDriver = (*mainDriver)(nil)
