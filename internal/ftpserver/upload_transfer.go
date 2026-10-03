package ftpserver

import (
	"bytes"
	"errors"
	"io"
)

// errReadNotSupported is returned by uploadTransfer.Read: this handle only
// ever backs a STOR (upload), never a RETR (download).
var errReadNotSupported = errors.New("ftpserver: upload transfer does not support reading")

// uploadTransfer implements ftpserverlib.FileTransfer and
// ftpserverlib.FileTransferError for a single fresh STOR. It buffers the
// incoming bytes in memory and forwards them to an Uploader on Close,
// unless the transfer was interrupted first.
type uploadTransfer struct {
	filename string
	uploader Uploader

	buf         bytes.Buffer
	transferErr error
}

func newUploadTransfer(filename string, uploader Uploader) *uploadTransfer {
	return &uploadTransfer{filename: filename, uploader: uploader}
}

// Write buffers the uploaded bytes.
func (t *uploadTransfer) Write(p []byte) (int, error) {
	return t.buf.Write(p)
}

// Read is never expected to be called for an upload handle.
func (t *uploadTransfer) Read(_ []byte) (int, error) {
	return 0, errReadNotSupported
}

// Seek is a no-op: GetHandle only ever hands out an uploadTransfer for a
// fresh STOR at offset 0, and ftpserverlib skips calling Seek in that case
// (it only seeks when resuming at a non-zero REST offset).
func (t *uploadTransfer) Seek(_ int64, _ int) (int64, error) {
	return 0, nil
}

// TransferError records that the transfer was interrupted (e.g. an ABOR or a
// dropped data connection), so Close knows to skip the paperless-ngx POST.
func (t *uploadTransfer) TransferError(err error) {
	t.transferErr = err
}

// Close forwards the buffered upload to paperless-ngx, unless the transfer
// was interrupted. A failed upload is returned as an error so ftpserverlib
// surfaces it as an FTP error reply to the client.
func (t *uploadTransfer) Close() error {
	if t.transferErr != nil {
		return nil
	}

	return t.uploader.Upload(t.filename, t.buf.Bytes())
}

var _ io.Reader = (*uploadTransfer)(nil)
var _ io.Writer = (*uploadTransfer)(nil)
var _ io.Seeker = (*uploadTransfer)(nil)
var _ io.Closer = (*uploadTransfer)(nil)
