# Migration notes

Breaking-change upgrade guides for this chart, newest first. Each entry is written once, when the breaking change ships, and referenced from the corresponding GitHub Release notes.

## Embedded FTP server (no more external FTP server)

Spec: [#81](https://github.com/taskmedia/paperlessngx-ftp-bridge/issues/81). Implemented across PRs #88-#93.

### What changed

`ftp-paperless-bridge` used to be an FTP *client*: it polled an external FTP server on a schedule, pulled new files, pushed them into paperless-ngx, and deleted them from the server. It is now the FTP *server* itself — scanners connect directly to the bridge (`scanner -> ftp-paperless-bridge -> paperless-ngx`), and every upload is forwarded to paperless-ngx the moment it finishes, with no polling and no external FTP server to run or keep in sync.

This is a full replacement, not a dual-mode rollout — there is no flag that keeps the old polling behavior alive.

### Action required

1. **Point your scanner(s) at the bridge directly.** Use the chart's new Service external IP (or `service.loadBalancerIP` if you pinned one) instead of your old external FTP server's address. Explicit FTPS (`AUTH TLS`) is offered automatically; see `ftp.tls.existingSecret.name` below if you want your own certificate instead of the self-signed default.
2. **Rewrite your `values.yaml`** per the table below — old keys are removed outright, not deprecated-and-ignored. A chart render fails loudly on any new required value left unset.
3. **Confirm your cluster/platform can satisfy a `LoadBalancer` Service.** The chart now hardcodes `service.type: LoadBalancer` to expose both the FTP control port and the full PASV port range (see "New Kubernetes resources" below).
4. **Decommission the old external FTP server** once scanners are repointed — the bridge no longer talks to it.

### Removed values / env vars

| Old `values.yaml` key | Old env var | Replacement |
|---|---|---|
| `ftp.host` | `FTP_HOST` | none — scanners connect to the chart's own Service instead |
| `ftp.user` | `FTP_USERNAME` | `ftp.accounts[].username` |
| `ftp.password` / `ftp.passwordExistingSecret` | `FTP_PASSWORD` | `ftp.accounts[].password` / `ftp.accounts[].passwordExistingSecret` |
| `ftp.path` | `FTP_PATH` | none — there is no remote directory to target; uploads are forwarded in-flight |
| `interval` (top-level) | `CRON_SCHEDULE` | none — there is no poll loop; uploads happen the instant a scanner finishes `STOR` |

### New required values

- `ftp.accounts[0].username` / `ftp.accounts[0].password` (or `passwordExistingSecret`) — at least one FTP account. The chart fails to render with no accounts configured, or with a username containing anything other than letters/digits/underscore.
- `ftp.pasv.publicHost` — the IP/host advertised to scanners for PASV data connections. No default; required for scanners behind NAT to reach the data channel at all.
- `paperless.username` / `paperless.password` (or `passwordExistingSecret`) — unchanged in shape, but now also required at render time if left empty.

### New optional values

- `ftp.allowedExtensions` (default `[".pdf"]`) — extensions accepted by `STOR`; anything else is rejected before any bytes are forwarded.
- `ftp.pasv.portMin` / `ftp.pasv.portMax` (default `50000`-`50019`, 20 ports) — the PASV port range the Service and container both expose.
- `ftp.tls.existingSecret.name` — point at your own `kubernetes.io/tls` Secret (e.g. from cert-manager) instead of the self-signed certificate generated in memory at startup. Picked up on renewal without a pod restart.
- `service.loadBalancerIP` / `service.annotations` — optional passthrough to the now-hardcoded `LoadBalancer` Service.

### New Kubernetes resources / behavior

- **Service is now `type: LoadBalancer`** (hardcoded, not a value), exposing the control port (external `21` -> container `2121`) and every PASV port as individual `spec.ports` entries. Your cluster/platform needs to be able to satisfy a `LoadBalancer` Service (cloud LB, MetalLB, klipper-lb on k3s, etc.) — `NodePort` was deliberately not used since it can't sanely carry ~20 ports.
- **`replicas: 1` is hardcoded** in the Deployment, not a `.Values` knob. PASV session affinity can't survive more than one replica behind an L4 passthrough LoadBalancer.
- **Container `securityContext` is fully locked down** (`runAsNonRoot`, fixed non-root UID/GID, dropped capabilities, read-only root filesystem, default seccomp profile). Shouldn't require any action unless your cluster enforces stricter pod security policies that conflict with these specific settings.
- **`readinessProbe` added** (`/readyz`), alongside the existing `livenessProbe` (`/healthz`, semantics narrowed to reflect only the FTP listener's bound state). `/readyz` additionally reflects a background paperless-ngx reachability check — a paperless-ngx outage takes the pod out of Service rotation without restarting it.

### Example

Before:

```yaml
ftp:
  host: "ftp.example.org:21"
  user: "scanner1"
  password: "secret"
  path: "."

interval: "*/5 * * * *"
```

After:

```yaml
ftp:
  accounts:
    - username: "scanner1"
      password: "secret"
  pasv:
    publicHost: "203.0.113.10" # your cluster's external IP/host
```
