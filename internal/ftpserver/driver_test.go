package ftpserver

import (
	"errors"
	"testing"
)

func TestMainDriverAuthUser(t *testing.T) {
	cfg := Config{Username: "scanner", Password: "secret", Uploader: &fakeUploader{}}
	driver := newMainDriver(cfg)

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

func TestMainDriverGetTLSConfigReturnsNil(t *testing.T) {
	driver := newMainDriver(Config{})

	tlsConfig, err := driver.GetTLSConfig()
	if err != nil {
		t.Fatalf("GetTLSConfig() error = %v, want nil", err)
	}
	if tlsConfig != nil {
		t.Fatalf("GetTLSConfig() = %v, want nil", tlsConfig)
	}
}
