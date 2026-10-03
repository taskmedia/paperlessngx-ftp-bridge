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

	// PublicHost is used as the self-signed TLS certificate's SAN. A
	// follow-on ticket also advertises it as the PASV public host.
	// Defaults to "localhost".
	PublicHost string

	// TLSCertDir is checked for an existing tls.crt/tls.key pair (the
	// kubernetes.io/tls Secret shape) to offer for AUTH TLS instead of the
	// self-signed default. Defaults to DefaultTLSCertDir. Presence of a
	// readable pair is the only toggle; there is no separate enable flag.
	TLSCertDir string

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

func (c Config) publicHost() string {
	if c.PublicHost == "" {
		return "localhost"
	}

	return c.PublicHost
}

func (c Config) tlsCertDir() string {
	if c.TLSCertDir == "" {
		return DefaultTLSCertDir
	}

	return c.TLSCertDir
}
