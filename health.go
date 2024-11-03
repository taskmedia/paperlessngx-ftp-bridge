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
	// allowed percentage of unhealthy results in last n results
	unhealthyPercentage = 0.5
	evaluatedResults    = 10
)

var (
	lastResults      = make([]bool, evaluatedResults)
	lastResultsIndex = 0
	lastResultsMutex sync.Mutex
	threshold        = int(math.Floor(float64(evaluatedResults) * unhealthyPercentage))
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
			log.Debug("Health check passed")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "OK")
		} else {
			log.Warn("Health check failed")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "NOT OK")
		}
	})

	log.Info("Starting health check server on :8080/healthz")
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
	defer quitFTPConnection(conn)

	// Attempt to log in to the FTP server
	err = conn.Login(config.ftpUsername, config.ftpPassword)
	if err != nil {
		log.Error("Failed to login to FTP server during readiness probe", "error", err)
		return false
	}

	// Check Paperless server
	client := resty.New()
	resp, err := client.R().
		SetBasicAuth(config.paperlessUser, config.paperlessPassword).
		Get(config.paperlessApiURL)
	if err != nil || (resp.StatusCode() != 405 && resp.IsError()) {
		log.Error("Failed to connect to Paperless server during readiness probe", "error", err)
		return false
	}

	return true
}

func updateLastResults(success bool) {
	lastResultsMutex.Lock()
	defer lastResultsMutex.Unlock()

	lastResults[lastResultsIndex] = success
	lastResultsIndex = (lastResultsIndex + 1) % evaluatedResults
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

	return falseCount <= threshold
}
