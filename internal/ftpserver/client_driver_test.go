package ftpserver

import (
	"os"
	"testing"
)

func TestClientDriverGetHandleFreshUpload(t *testing.T) {
	tests := []struct {
		desc    string
		name    string
		wantErr bool
	}{
		{desc: "allowed extension lowercase", name: "/scan.pdf", wantErr: false},
		{desc: "allowed extension uppercase", name: "/scan.PDF", wantErr: false},
		{desc: "allowed extension mixed case", name: "/scan.Pdf", wantErr: false},
		{desc: "disallowed extension", name: "/scan.exe", wantErr: true},
		{desc: "no extension", name: "/scan", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			uploader := &fakeUploader{}
			driver := newClientDriver(uploader, []string{".pdf"})

			handle, err := driver.GetHandle(tt.name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("GetHandle() error = nil, want error")
				}
				if uploader.calls != 0 {
					t.Errorf("Upload() called before Close(), want no call yet")
				}

				return
			}

			if err != nil {
				t.Fatalf("GetHandle() error = %v, want nil", err)
			}
			if handle == nil {
				t.Fatalf("GetHandle() handle = nil, want non-nil")
			}
			if uploader.calls != 0 {
				t.Errorf("Upload() called before any bytes were written")
			}
		})
	}
}

func TestClientDriverGetHandleFallsThroughForNonFreshUpload(t *testing.T) {
	tests := []struct {
		desc   string
		flags  int
		offset int64
	}{
		{desc: "RETR (read)", flags: os.O_RDONLY, offset: 0},
		{desc: "resume STOR with REST offset", flags: os.O_WRONLY | os.O_CREATE, offset: 10},
		{desc: "APPE", flags: os.O_WRONLY | os.O_CREATE | os.O_APPEND, offset: 0},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			driver := newClientDriver(&fakeUploader{}, []string{".pdf"})

			_, err := driver.GetHandle("/missing.pdf", tt.flags, tt.offset)
			if err == nil {
				t.Fatalf("GetHandle() error = nil, want \"file does not exist\" from the empty MemMapFs")
			}
			if !os.IsNotExist(err) {
				t.Fatalf("GetHandle() error = %v, want a not-exist error", err)
			}
		})
	}
}

func TestClientDriverRemoveFallsThroughToEmptyFs(t *testing.T) {
	driver := newClientDriver(&fakeUploader{}, []string{".pdf"})

	err := driver.Remove("/missing.pdf")
	if err == nil {
		t.Fatalf("Remove() error = nil, want \"file does not exist\" from the empty MemMapFs")
	}
	if !os.IsNotExist(err) {
		t.Fatalf("Remove() error = %v, want a not-exist error", err)
	}
}
