package ftpserver

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/drakkan/goftp"

	"github.com/taskmedia/paperlessngx-ftp-bridge/internal/paperless"
)

// startTestServer brings up the embedded FTP server on an ephemeral port
// (both control and PASV data channels stay unset, per ADR-0005) and tears
// it down on test cleanup.
func startTestServer(t *testing.T, cfg Config) string {
	t.Helper()

	cfg.ListenAddr = "127.0.0.1:0"
	srv := NewServer(cfg)

	if err := srv.Listen(); err != nil {
		t.Fatalf("Listen() error = %v", err)
	}

	go func() {
		_ = srv.Serve()
	}()

	t.Cleanup(func() {
		_ = srv.Stop()
	})

	return srv.Addr()
}

func dialTestClient(t *testing.T, addr, user, pass string) *goftp.Client {
	t.Helper()

	client, err := goftp.DialConfig(goftp.Config{
		User:               user,
		Password:           pass,
		ConnectionsPerHost: 1,
		Timeout:            5 * time.Second,
	}, addr)
	if err != nil {
		t.Fatalf("Dial() error = %v", err)
	}

	t.Cleanup(func() { _ = client.Close() })

	return client
}

func TestIntegrationAllowedUploadReachesPaperless(t *testing.T) {
	var gotFilename string
	var gotBody []byte

	paperlessServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		w.WriteHeader(http.StatusOK)
	}))
	defer paperlessServer.Close()

	uploader := paperless.NewClient(paperlessServer.URL, "", "")
	addr := startTestServer(t, Config{Username: "scanner", Password: "secret", Uploader: uploader})
	client := dialTestClient(t, addr, "scanner", "secret")

	err := client.Store("scan.pdf", bytes.NewReader([]byte("%PDF-1.4 fake pdf bytes")))
	if err != nil {
		t.Fatalf("Store() error = %v, want success", err)
	}

	if gotFilename != "scan.pdf" {
		t.Errorf("paperless-ngx received filename %q, want %q", gotFilename, "scan.pdf")
	}
	if string(gotBody) != "%PDF-1.4 fake pdf bytes" {
		t.Errorf("paperless-ngx received body %q, want %q", gotBody, "%PDF-1.4 fake pdf bytes")
	}
}

func TestIntegrationDisallowedExtensionRejectedBeforeHTTPCall(t *testing.T) {
	called := false

	paperlessServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	}))
	defer paperlessServer.Close()

	uploader := paperless.NewClient(paperlessServer.URL, "", "")
	addr := startTestServer(t, Config{Username: "scanner", Password: "secret", Uploader: uploader})
	client := dialTestClient(t, addr, "scanner", "secret")

	err := client.Store("scan.exe", bytes.NewReader([]byte("not a pdf")))
	if err == nil {
		t.Fatalf("Store() error = nil, want an FTP error reply for a disallowed extension")
	}

	if called {
		t.Errorf("paperless-ngx was called for a disallowed extension, want no HTTP call at all")
	}
}

func TestIntegrationPaperlessErrorSurfacesAsFTPError(t *testing.T) {
	paperlessServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer paperlessServer.Close()

	uploader := paperless.NewClient(paperlessServer.URL, "", "")
	addr := startTestServer(t, Config{Username: "scanner", Password: "secret", Uploader: uploader})
	client := dialTestClient(t, addr, "scanner", "secret")

	err := client.Store("scan.pdf", bytes.NewReader([]byte("%PDF-1.4 fake pdf bytes")))
	if err == nil {
		t.Fatalf("Store() error = nil, want an FTP error reply when paperless-ngx fails")
	}
}
