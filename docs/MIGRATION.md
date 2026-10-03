# Migration notes

Breaking-change upgrade guides for this chart, newest first.

## v2: embedded FTP server

Spec: [#81](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/81).

### TL;DR

The bridge is now the FTP server itself (`scanner -> ftp-paperless-bridge -> paperless-ngx`) instead of polling an external one. Point your scanner(s) at the chart's Service, rewrite `values.yaml` per the table below, and make sure your cluster can satisfy a `LoadBalancer` Service. No dual-mode flag exists — this is a full replacement.

### Removed

| Old `values.yaml` key | Old env var | Replacement |
|---|---|---|
| `ftp.host` | `FTP_HOST` | none — scanners connect to the chart's own Service |
| `ftp.user` | `FTP_USERNAME` | `ftp.accounts[].username` |
| `ftp.password` / `ftp.passwordExistingSecret` | `FTP_PASSWORD` | `ftp.accounts[].password` / `ftp.accounts[].passwordExistingSecret` |
| `ftp.path` | `FTP_PATH` | none — uploads are forwarded in-flight, no remote directory |
| `interval` (top-level) | `CRON_SCHEDULE` | none — no poll loop, uploads happen on `STOR` |

### New required values

- `ftp.accounts[0].username` / `.password` (or `.passwordExistingSecret`) — at least one account; username must be letters/digits/underscore only.
- `ftp.pasv.publicHost` — IP/host advertised to scanners for PASV data connections. No default.
- `paperless.username` / `.password` (or `.passwordExistingSecret`) — same shape as before, now required at render time.

### New optional values

- `ftp.allowedExtensions` (default `[".pdf"]`) — extensions accepted by `STOR`.
- `ftp.pasv.portMin` / `portMax` (default `50000`-`50019`) — PASV port range.
- `ftp.tls.existingSecret.name` — your own `kubernetes.io/tls` Secret instead of the self-signed default; picked up on renewal without a restart.
- `service.loadBalancerIP` / `service.annotations` — passthrough to the Service.

### New Kubernetes behavior

- **`service.type: LoadBalancer`** (hardcoded) exposes the control port (`21` -> `2121`) and the full PASV range. Needs a platform that can satisfy a `LoadBalancer` (cloud LB, MetalLB, k3s klipper-lb, ...) — `NodePort` can't carry ~20 ports.
- **`replicas: 1`** (hardcoded) — PASV session affinity can't survive multiple replicas behind an L4 LoadBalancer.
- **`securityContext`** is fully locked down (non-root, dropped capabilities, read-only rootfs, seccomp). Only relevant if your cluster's pod security policies conflict with these.
- **`readinessProbe`** (`/readyz`) added alongside `livenessProbe` (`/healthz`, now reflecting only the FTP listener). `/readyz` also tracks paperless-ngx reachability, taking the pod out of rotation (not restarting it) during an outage.

### Example

```yaml
# before
ftp:
  host: "ftp.example.org:21"
  user: "scanner1"
  password: "secret"
  path: "."
interval: "*/5 * * * *"

# after
ftp:
  accounts:
    - username: "scanner1"
      password: "secret"
  pasv:
    publicHost: "203.0.113.10" # your cluster's external IP/host
```
