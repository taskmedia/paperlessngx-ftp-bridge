package ftpserver

import (
	"errors"
	"testing"
)

type fakeUploader struct {
	calls    int
	filename string
	data     []byte
	err      error
}

func (f *fakeUploader) Upload(filename string, data []byte) error {
	f.calls++
	f.filename = filename
	f.data = append([]byte(nil), data...)

	return f.err
}

func TestUploadTransferClosePostsBufferedBytes(t *testing.T) {
	uploader := &fakeUploader{}
	transfer := newUploadTransfer("scan.pdf", uploader)

	if _, err := transfer.Write([]byte("hello ")); err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	if _, err := transfer.Write([]byte("world")); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	if err := transfer.Close(); err != nil {
		t.Fatalf("Close() error = %v, want nil", err)
	}

	if uploader.calls != 1 {
		t.Fatalf("Upload() called %d times, want 1", uploader.calls)
	}
	if uploader.filename != "scan.pdf" {
		t.Errorf("Upload() filename = %q, want %q", uploader.filename, "scan.pdf")
	}
	if string(uploader.data) != "hello world" {
		t.Errorf("Upload() data = %q, want %q", uploader.data, "hello world")
	}
}

func TestUploadTransferCloseSurfacesUploadError(t *testing.T) {
	wantErr := errors.New("paperless-ngx unreachable")
	uploader := &fakeUploader{err: wantErr}
	transfer := newUploadTransfer("scan.pdf", uploader)

	if _, err := transfer.Write([]byte("data")); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	if err := transfer.Close(); !errors.Is(err, wantErr) {
		t.Fatalf("Close() error = %v, want %v", err, wantErr)
	}

	if uploader.calls != 1 {
		t.Fatalf("Upload() called %d times, want 1", uploader.calls)
	}
}

func TestUploadTransferAbortedTransferSkipsUpload(t *testing.T) {
	uploader := &fakeUploader{}
	transfer := newUploadTransfer("scan.pdf", uploader)

	if _, err := transfer.Write([]byte("partial")); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	transfer.TransferError(errors.New("client aborted"))

	if err := transfer.Close(); err != nil {
		t.Fatalf("Close() error = %v, want nil", err)
	}

	if uploader.calls != 0 {
		t.Fatalf("Upload() called %d times, want 0", uploader.calls)
	}
}

func TestUploadTransferReadIsUnsupported(t *testing.T) {
	transfer := newUploadTransfer("scan.pdf", &fakeUploader{})

	if _, err := transfer.Read(make([]byte, 4)); !errors.Is(err, errReadNotSupported) {
		t.Fatalf("Read() error = %v, want %v", err, errReadNotSupported)
	}
}
