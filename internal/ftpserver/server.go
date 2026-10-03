package ftpserver

import ftpserverlib "github.com/fclairamb/ftpserverlib"

// NewServer constructs the embedded FTP server described by cfg. Callers
// are responsible for calling Listen/Serve (or ListenAndServe) on the
// result.
func NewServer(cfg Config) *ftpserverlib.FtpServer {
	return ftpserverlib.NewFtpServer(newMainDriver(cfg))
}
