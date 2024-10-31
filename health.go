package main

import (
	"crypto/tls"
	"fmt"
	log "log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jlaffaye/ftp"
)

var (
	lastResults      = make([]bool, 10)
	lastResultsIndex = 0
	lastResultsMutex sync.Mutex
)

func startHealthCheckServer() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if isHealthy() {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "OK")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "NOT OK")
		}
	})

	log.Info("Starting health check server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error("Failed to start health check server", "error", err)
	}
}

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

func updateLastResults(success bool) {
	lastResultsMutex.Lock()
	defer lastResultsMutex.Unlock()

	lastResults[lastResultsIndex] = success
	lastResultsIndex = (lastResultsIndex + 1) % len(lastResults)
}

func isHealthy() bool {
	lastResultsMutex.Lock()
	defer lastResultsMutex.Unlock()

	for _, result := range lastResults {
		if !result {
			return false
		}
	}
	return true
}
