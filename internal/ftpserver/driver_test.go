package ftpserver

import (
	"errors"
	"testing"
)

func TestMainDriverAuthUser(t *testing.T) {
	cfg := Config{Username: "scanner", Password: "secret", Uploader: &fakeUploader{}, TLSCertDir: t.TempDir()}
	driver, err := newMainDriver(cfg)
	if err != nil {
		t.Fatalf("newMainDriver() error = %v, want nil", err)
	}

	tests := []struct {
		desc    string
		user    string
		pass    string
		wantErr bool
	}{
		{desc: "correct credentials", user: "scanner", pass: "secret", wantErr: false},
		{desc: "wrong password", user: "scanner", pass: "wrong", wantErr: true},
		{desc: "wrong username", user: "someone-else", pass: "secret", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			clientDriver, err := driver.AuthUser(nil, tt.user, tt.pass)

			if tt.wantErr {
				if !errors.Is(err, errInvalidCredentials) {
					t.Fatalf("AuthUser() error = %v, want %v", err, errInvalidCredentials)
				}

				return
			}

			if err != nil {
				t.Fatalf("AuthUser() error = %v, want nil", err)
			}
			if clientDriver == nil {
				t.Fatalf("AuthUser() clientDriver = nil, want non-nil")
			}
		})
	}
}

func TestMainDriverGetTLSConfigOffersSelfSignedCertByDefault(t *testing.T) {
	driver, err := newMainDriver(Config{TLSCertDir: t.TempDir(), PublicHost: "ftp.example.org"})
	if err != nil {
		t.Fatalf("newMainDriver() error = %v, want nil", err)
	}

	tlsConfig, err := driver.GetTLSConfig()
	if err != nil {
		t.Fatalf("GetTLSConfig() error = %v, want nil", err)
	}
	if tlsConfig == nil || len(tlsConfig.Certificates) == 0 {
		t.Fatalf("GetTLSConfig() = %v, want a config offering a self-signed certificate", tlsConfig)
	}
}

func TestNewMainDriverFailsFastOnInvalidExistingCert(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir+"/tls.crt", "not a certificate")
	writeFile(t, dir+"/tls.key", "not a key")

	if _, err := newMainDriver(Config{TLSCertDir: dir}); err == nil {
		t.Fatalf("newMainDriver() error = nil, want a fatal error for an invalid TLS secret")
	}
}
