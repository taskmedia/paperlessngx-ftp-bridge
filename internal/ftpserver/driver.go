package ftpserver

import (
	"crypto/tls"
	"errors"

	ftpserverlib "github.com/fclairamb/ftpserverlib"
)

// errInvalidCredentials is returned by AuthUser when the supplied
// username/password don't match the configured account.
var errInvalidCredentials = errors.New("ftpserver: invalid username or password")

// mainDriver implements ftpserverlib.MainDriver. Account configuration is
// intentionally minimal (a single username/password pair) and TLS is not
// yet offered; both are widened by follow-on tickets.
type mainDriver struct {
	cfg Config
}

func newMainDriver(cfg Config) *mainDriver {
	return &mainDriver{cfg: cfg}
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

// GetTLSConfig returns no TLS configuration: explicit FTPS is added by a
// follow-on ticket (ADR-0003).
func (d *mainDriver) GetTLSConfig() (*tls.Config, error) {
	return nil, nil
}

var _ ftpserverlib.MainDriver = (*mainDriver)(nil)
