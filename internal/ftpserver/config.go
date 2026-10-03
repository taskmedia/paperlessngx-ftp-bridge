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

// Config holds the minimal settings needed to construct the embedded FTP
// server. Account and PASV-range configuration is intentionally minimal here
// pending a dedicated ticket to widen it to a multi-account, Helm-driven
// schema.
type Config struct {
	// ListenAddr is the control-channel listen address, e.g. ":2121".
	// Defaults to DefaultListenAddr.
	ListenAddr string

	// Username and Password authenticate the single FTP account accepted by
	// the server.
	Username string
	Password string

	// AllowedExtensions lists the file extensions (including the leading
	// dot, e.g. ".pdf") accepted by STOR. Matching is case-insensitive.
	// Defaults to DefaultAllowedExtensions.
	AllowedExtensions []string

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
