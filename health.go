package main

import (
	"crypto/tls"
	"fmt"
	log "log/slog"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jlaffaye/ftp"
)

const (
	// Percentage limit for the number of false results to determine health status
	unhealthyPercentage = 0.5
)

var (
	lastResults      = make([]bool, 10)
	lastResultsIndex = 0
	lastResultsMutex sync.Mutex
)

func init() {
	// start with last results all true to avoid false health check
	for i := range lastResults {
		lastResults[i] = true
	}
}

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

	falseCount := 0
	for _, result := range lastResults {
		if !result {
			falseCount++
		}
	}

	threshold := int(math.Floor(float64(len(lastResults)) * unhealthyPercentage))

	return falseCount <= threshold
}
