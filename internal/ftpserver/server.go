package ftpserver

import ftpserverlib "github.com/fclairamb/ftpserverlib"

// NewServer constructs the embedded FTP server described by cfg. Callers
// are responsible for calling Listen/Serve (or ListenAndServe) on the
// result. An error here means the TLS certificate configured via
// cfg.TLSCertDir couldn't be prepared and is a fatal startup condition
// (ADR-0003); callers should treat it as such.
func NewServer(cfg Config) (*ftpserverlib.FtpServer, error) {
	driver, err := newMainDriver(cfg)
	if err != nil {
		return nil, err
	}

	return ftpserverlib.NewFtpServer(driver), nil
}
