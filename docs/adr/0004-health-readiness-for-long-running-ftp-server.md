# Liveness/readiness model for the long-running embedded FTP server

The old `health.go` had one probe shape fitted to a cron job: a blocking synchronous `readinessProbe` (dial the external FTP server + hit the Paperless-ngx API) run once before the loop started, feeding a single `/healthz` endpoint backed by a rolling window (50% of the last 10 cron-run results). With the embedded FTP server, the process is a persistent listener, not a periodic job, so none of those pieces still fit: there's no "run" to produce a rolling result, and the thing being probed (an always-on listener plus a downstream HTTP dependency) has two different failure semantics that the old single endpoint conflated.

We split the HTTP health server into two endpoints with different consequences:

- `/healthz` (liveness) reflects only whether the embedded FTP listener is still bound and accepting connections. If it isn't, the process is useless and Kubernetes should restart it.
- `/readyz` (readiness) reflects the FTP listener's state **and** the last cached result of a background Paperless-ngx reachability check. If Paperless is unreachable, the pod should drop out of rotation but must not be restarted — restarting the FTP listener fixes nothing about a downstream outage and would needlessly drop any in-flight FTP sessions.

The Paperless-ngx check runs in its own background goroutine on a configurable interval (default 60s, surfaced as a Helm value / env var rather than hardcoded) and caches only the latest result — no rolling window. The interval itself is the debounce; a threshold-over-N-samples on top of that added complexity the old cron-based signal needed but this model doesn't.

On startup, a TLS certificate load failure is still fatal (`os.Exit(1)`) — it's a genuine misconfiguration, not a transient condition. An initial Paperless-ngx check failure is not: the FTP listener starts regardless, and `/readyz` simply reports not-ready until the first background check succeeds. This matches a long-running service rather than the old one-shot-job assumption that readiness must be proven before anything starts.

Considered alternative: keep a single `/healthz` endpoint covering both conditions. Rejected — conflating "restart me" with "don't route to me yet" would cause unnecessary restarts during ordinary, recoverable Paperless-ngx downtime.

## Consequences

The Helm chart needs two separate probe entries (`livenessProbe` pointed at `/healthz`, `readinessProbe` at `/readyz`) instead of one. The rolling-window/threshold code in the old `health.go` (`lastResults`, `unhealthyPercentage`, `evaluatedResults`) is removed entirely rather than adapted.
