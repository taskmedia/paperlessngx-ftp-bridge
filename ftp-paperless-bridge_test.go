package main

import (
	"reflect"
	"testing"

	"github.com/taskmedia/paperlessngx-ftp-bridge/internal/ftpserver"
)

func TestLoadFTPAccounts(t *testing.T) {
	tests := []struct {
		desc string
		env  []string
		want []ftpserver.Account
	}{
		{
			desc: "single account",
			env:  []string{"FTP_ACCOUNT_SCANNER=secret", "PAPERLESS_URL=http://paperless"},
			want: []ftpserver.Account{{Username: "SCANNER", Password: "secret"}},
		},
		{
			desc: "multiple accounts sorted by username",
			env:  []string{"FTP_ACCOUNT_SCANNER2=secret2", "FTP_ACCOUNT_SCANNER1=secret1"},
			want: []ftpserver.Account{
				{Username: "SCANNER1", Password: "secret1"},
				{Username: "SCANNER2", Password: "secret2"},
			},
		},
		{
			desc: "no accounts configured",
			env:  []string{"PAPERLESS_URL=http://paperless"},
			want: nil,
		},
		{
			desc: "value containing an equals sign is preserved",
			env:  []string{"FTP_ACCOUNT_SCANNER=sec=ret"},
			want: []ftpserver.Account{{Username: "SCANNER", Password: "sec=ret"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := loadFTPAccounts(tt.env)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadFTPAccounts() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestAtoiOrZero(t *testing.T) {
	tests := []struct {
		desc string
		raw  string
		want int
	}{
		{desc: "valid integer", raw: "50000", want: 50000},
		{desc: "empty string", raw: "", want: 0},
		{desc: "not an integer", raw: "not-a-number", want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := atoiOrZero(tt.raw); got != tt.want {
				t.Errorf("atoiOrZero(%q) = %d, want %d", tt.raw, got, tt.want)
			}
		})
	}
}
