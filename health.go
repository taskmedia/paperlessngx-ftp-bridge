package main

import (
	"crypto/tls"
	log "log/slog"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jlaffaye/ftp"
)

func readinessProbe(config Config) bool {
	// Check FTP server
	conn, err := ftp.Dial(
		config.ftpHost,
		ftp.DialWithTimeout(5*time.Second),
		ftp.DialWithExplicitTLS(&tls.Config{
			InsecureSkipVerify: true,
		}))
	if err != nil {
		log.Error("Failed to connect to FTP server during readiness probe", "error", err)
		return false
	}
	defer conn.Quit()

	// Check Paperless server
	client := resty.New()
	resp, err := client.R().Get(config.paperlessURL)
	if err != nil || resp.IsError() {
		log.Error("Failed to connect to Paperless server during readiness probe", "error", err)
		return false
	}

	return true
}
