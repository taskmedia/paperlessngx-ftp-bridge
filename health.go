package main

import (
	"fmt"
	log "log/slog"
	"net/http"
	"sync/atomic"
	"time"
)

// DefaultPaperlessCheckInterval is the interval between background
// paperless-ngx reachability checks used when no override is configured.
const DefaultPaperlessCheckInterval = 60 * time.Second

// ftpStatus reports whether the embedded FTP listener is currently bound
// and accepting connections. *ftpserverlib.FtpServer satisfies this via its
// Addr method (empty once Stop()'d or before Listen() has run).
type ftpStatus interface {
	Addr() string
}

// paperlessPinger checks paperless-ngx reachability.
type paperlessPinger interface {
	Ping() error
}

// readinessChecker implements the liveness/readiness split described in
// ADR-0004: liveness reflects only the FTP listener's bound state;
// readiness additionally reflects the latest cached result of a background
// paperless-ngx reachability check, run on its own goroutine so a slow or
// down paperless-ngx never blocks a probe request.
type readinessChecker struct {
	ftp       ftpStatus
	paperless paperlessPinger
	interval  time.Duration

	paperlessReachable atomic.Bool
}

func newReadinessChecker(ftp ftpStatus, paperless paperlessPinger, interval time.Duration) *readinessChecker {
	return &readinessChecker{ftp: ftp, paperless: paperless, interval: interval}
}

// run checks paperless-ngx reachability immediately, then on every tick of
// c.interval. It blocks, so callers should run it in its own goroutine. An
// initial failure is non-fatal (ADR-0004): isReady simply reports not-ready
// until the first check here succeeds.
func (c *readinessChecker) run() {
	c.checkPaperless()

	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.checkPaperless()
	}
}

func (c *readinessChecker) checkPaperless() {
	err := c.paperless.Ping()
	c.paperlessReachable.Store(err == nil)

	if err != nil {
		log.Warn("paperless-ngx reachability check failed", "error", err)
	}
}

// isLive reports whether the FTP listener is bound and accepting
// connections.
func (c *readinessChecker) isLive() bool {
	return c.ftp.Addr() != ""
}

// isReady reports whether the server is both live and paperless-ngx was
// reachable on the latest background check.
func (c *readinessChecker) isReady() bool {
	return c.isLive() && c.paperlessReachable.Load()
}

// startHealthCheckServer serves /healthz (liveness) and /readyz (readiness)
// per ADR-0004.
func startHealthCheckServer(checker *readinessChecker) {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthzHandler(checker))
	mux.HandleFunc("/readyz", readyzHandler(checker))

	log.Info("Starting health check server on :8080 (/healthz, /readyz)")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Error("Failed to start health check server", "error", err)
	}
}

func healthzHandler(checker *readinessChecker) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		writeHealthResponse(w, checker.isLive(), "FTP listener is not bound")
	}
}

func readyzHandler(checker *readinessChecker) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		writeHealthResponse(w, checker.isReady(), "not ready")
	}
}

func writeHealthResponse(w http.ResponseWriter, ok bool, notOKMessage string) {
	if !ok {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = fmt.Fprintln(w, notOKMessage)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "OK")
}
