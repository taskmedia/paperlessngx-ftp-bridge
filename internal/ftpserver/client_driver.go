package ftpserver

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	ftpserverlib "github.com/fclairamb/ftpserverlib"
	"github.com/spf13/afero"
)

// errExtensionNotAllowed is returned by GetHandle when a fresh STOR's
// extension isn't in the configured allowlist.
var errExtensionNotAllowed = errors.New("ftpserver: extension not allowed")

// clientDriver is the ftpserverlib.ClientDriver for one FTP session. It
// backs CWD/MKD/STAT/LIST with a fresh, in-memory, per-connection
// filesystem that is discarded on disconnect (ADR-0001) and intercepts
// fresh STOR uploads via GetHandle, routing them to the configured
// Uploader instead of the in-memory filesystem.
type clientDriver struct {
	afero.Fs

	uploader   Uploader
	allowedExt map[string]struct{}
}

func newClientDriver(uploader Uploader, allowedExtensions []string) *clientDriver {
	allowed := make(map[string]struct{}, len(allowedExtensions))
	for _, ext := range allowedExtensions {
		allowed[strings.ToLower(ext)] = struct{}{}
	}

	return &clientDriver{
		Fs:         afero.NewMemMapFs(),
		uploader:   uploader,
		allowedExt: allowed,
	}
}

// GetHandle implements ftpserverlib.ClientDriverExtentionFileTransfer. A
// fresh STOR (no REST offset) of an allowed extension is routed to an
// uploadTransfer; a fresh STOR of a disallowed extension is rejected here,
// eagerly, before any bytes are buffered or any HTTP call is made. Every
// other shape (RETR, resume, APPE) falls through to the empty in-memory
// filesystem, which fails naturally with "file does not exist".
func (d *clientDriver) GetHandle(name string, flags int, offset int64) (ftpserverlib.FileTransfer, error) {
	if !isFreshUpload(flags, offset) {
		// Strip O_CREATE: RETR, resume, and APPE must fail with the
		// filesystem's natural "file does not exist" error (ADR-0001)
		// rather than silently creating an empty file in the in-memory fs.
		return d.OpenFile(name, flags&^os.O_CREATE, 0o644) //nolint:gosec // in-memory fs, mode is advisory only
	}

	ext := strings.ToLower(filepath.Ext(name))
	if _, ok := d.allowedExt[ext]; !ok {
		return nil, fmt.Errorf("%w: %q", errExtensionNotAllowed, ext)
	}

	return newUploadTransfer(filepath.Base(name), d.uploader), nil
}

// isFreshUpload reports whether flags/offset describe a brand new STOR:
// write access, file creation, no resume offset, and no append flag.
func isFreshUpload(flags int, offset int64) bool {
	const writeCreate = os.O_WRONLY | os.O_CREATE

	return offset == 0 && flags&writeCreate == writeCreate && flags&os.O_APPEND == 0
}

var _ ftpserverlib.ClientDriver = (*clientDriver)(nil)
var _ ftpserverlib.ClientDriverExtentionFileTransfer = (*clientDriver)(nil)
