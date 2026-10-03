// Package ftpserver embeds an FTP server (github.com/fclairamb/ftpserverlib)
// that intercepts every upload and forwards it to paperless-ngx instead of
// persisting it anywhere.
package ftpserver

// DefaultListenAddr is the control-channel address used when Config.ListenAddr
// is left empty.
const DefaultListenAddr = ":2121"

// DefaultAllowedExtensions is the extension allowlist used when
// Config.AllowedExtensions is left empty.
var DefaultAllowedExtensions = []string{".pdf"}

// Uploader forwards an uploaded file's content to paperless-ngx.
type Uploader interface {
	Upload(filename string, data []byte) error
}

// Account is one FTP username/password pair accepted by the server.
type Account struct {
	Username string
	Password string
}

// Config holds the settings needed to construct the embedded FTP server.
type Config struct {
	// ListenAddr is the control-channel listen address, e.g. ":2121".
	// Defaults to DefaultListenAddr.
	ListenAddr string

	// Accounts lists the FTP username/password pairs accepted by the
	// server. Adding a second or third account is a pure append — no code
	// change is required.
	Accounts []Account

	// AllowedExtensions lists the file extensions (including the leading
	// dot, e.g. ".pdf") accepted by STOR. Matching is case-insensitive.
	// Defaults to DefaultAllowedExtensions.
	AllowedExtensions []string

	// PublicHost is the IP/host advertised to clients for PASV data
	// connections. Required for passive mode to work behind NAT; there is
	// no default or auto-inference.
	PublicHost string

	// PASVPortMin and PASVPortMax bound the port range advertised for
	// passive data connections. When either is zero, passive connections
	// use ephemeral ports instead.
	PASVPortMin int
	PASVPortMax int

	// Uploader receives the bytes of every accepted upload.
	Uploader Uploader
}

func (c Config) listenAddr() string {
	if c.ListenAddr == "" {
		return DefaultListenAddr
	}

	return c.ListenAddr
}

func (c Config) allowedExtensions() []string {
	if len(c.AllowedExtensions) == 0 {
		return DefaultAllowedExtensions
	}

	return c.AllowedExtensions
}
