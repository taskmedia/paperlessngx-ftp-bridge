package paperless

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientUpload(t *testing.T) {
	tests := []struct {
		desc       string
		statusCode int
		dialErr    bool
		wantErr    bool
	}{
		{
			desc:       "paperless-ngx accepts the document",
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			desc:       "paperless-ngx returns a client error",
			statusCode: http.StatusBadRequest,
			wantErr:    true,
		},
		{
			desc:       "paperless-ngx returns a server error",
			statusCode: http.StatusInternalServerError,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			var gotUser, gotPass string
			var gotFilename string
			var gotBody []byte

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotUser, gotPass, _ = r.BasicAuth()

				if err := r.ParseMultipartForm(1 << 20); err != nil {
					t.Fatalf("parsing multipart form: %v", err)
				}

				file, header, err := r.FormFile("document")
				if err != nil {
					t.Fatalf("reading uploaded file: %v", err)
				}
				defer func() { _ = file.Close() }()

				gotFilename = header.Filename
				gotBody, err = io.ReadAll(file)
				if err != nil {
					t.Fatalf("reading uploaded file content: %v", err)
				}

				w.WriteHeader(tt.statusCode)
			}))
			defer server.Close()

			client := NewClient(server.URL, "someuser", "somepass")

			err := client.Upload("scan.pdf", []byte("pdf-bytes"))

			if tt.wantErr && err == nil {
				t.Fatalf("Upload() error = nil, want error")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("Upload() error = %v, want nil", err)
			}

			if gotUser != "someuser" || gotPass != "somepass" {
				t.Errorf("Upload() sent basic auth (%q, %q), want (%q, %q)", gotUser, gotPass, "someuser", "somepass")
			}
			if gotFilename != "scan.pdf" {
				t.Errorf("Upload() sent filename %q, want %q", gotFilename, "scan.pdf")
			}
			if string(gotBody) != "pdf-bytes" {
				t.Errorf("Upload() sent body %q, want %q", gotBody, "pdf-bytes")
			}
		})
	}
}
