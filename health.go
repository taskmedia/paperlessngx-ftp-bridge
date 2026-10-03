package main

import (
	"fmt"
	log "log/slog"
	"net/http"
)

// startHealthCheckServer serves a minimal, always-OK /healthz endpoint.
// The old rolling-window probe dialed the external FTP client this bridge
// used to poll; that client is gone now that the bridge is itself the FTP
// endpoint. A proper liveness/readiness split for the embedded server is
// tracked separately (ADR-0004).
func startHealthCheckServer() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, "OK")
	})

	log.Info("Starting health check server on :8080/healthz")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error("Failed to start health check server", "error", err)
	}
}
