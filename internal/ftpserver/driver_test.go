package ftpserver

import (
	"errors"
	"testing"

	ftpserverlib "github.com/fclairamb/ftpserverlib"
)

func TestMainDriverAuthUser(t *testing.T) {
	cfg := Config{
		Accounts: []Account{
			{Username: "scanner", Password: "secret"},
			{Username: "second-scanner", Password: "other-secret"},
		},
		Uploader:   &fakeUploader{},
		TLSCertDir: t.TempDir(),
	}
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
		{desc: "correct credentials, first account", user: "scanner", pass: "secret", wantErr: false},
		{desc: "correct credentials, second account", user: "second-scanner", pass: "other-secret", wantErr: false},
		{desc: "wrong password", user: "scanner", pass: "wrong", wantErr: true},
		{desc: "wrong username", user: "someone-else", pass: "secret", wantErr: true},
		{desc: "credentials swapped across accounts", user: "scanner", pass: "other-secret", wantErr: true},
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

func TestMainDriverGetSettingsPASVRange(t *testing.T) {
	tests := []struct {
		desc      string
		portMin   int
		portMax   int
		wantEphem bool
		wantStart int
		wantEnd   int
	}{
		{desc: "range configured", portMin: 50000, portMax: 50019, wantStart: 50000, wantEnd: 50019},
		{desc: "range left unset means ephemeral ports", wantEphem: true},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			driver, err := newMainDriver(Config{PASVPortMin: tt.portMin, PASVPortMax: tt.portMax, PublicHost: "ftp.example.org", TLSCertDir: t.TempDir()})
			if err != nil {
				t.Fatalf("newMainDriver() error = %v, want nil", err)
			}

			settings, err := driver.GetSettings()
			if err != nil {
				t.Fatalf("GetSettings() error = %v, want nil", err)
			}

			if tt.wantEphem {
				if settings.PassiveTransferPortRange != nil {
					t.Fatalf("PassiveTransferPortRange = %v, want nil (ephemeral ports)", settings.PassiveTransferPortRange)
				}

				return
			}

			portRange, ok := settings.PassiveTransferPortRange.(*ftpserverlib.PortRange)
			if !ok || portRange == nil {
				t.Fatalf("PassiveTransferPortRange = %v, want a *PortRange", settings.PassiveTransferPortRange)
			}
			if portRange.Start != tt.wantStart || portRange.End != tt.wantEnd {
				t.Errorf("PassiveTransferPortRange = {%d, %d}, want {%d, %d}", portRange.Start, portRange.End, tt.wantStart, tt.wantEnd)
			}
			if settings.PublicHost != "ftp.example.org" {
				t.Errorf("PublicHost = %q, want %q", settings.PublicHost, "ftp.example.org")
			}
		})
	}
}
