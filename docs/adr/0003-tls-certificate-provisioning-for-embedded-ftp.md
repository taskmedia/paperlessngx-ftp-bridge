# TLS certificate provisioning for the embedded FTP server

The embedded FTP server always offers explicit FTPS (`AUTH TLS`) alongside plain FTP — there is no `ftp.tls.enabled` switch, since a cert is cheap to have ready and one fewer flag means one fewer invalid combination to validate. By default the app generates a self-signed cert (ECDSA P-256, ~1yr validity, SAN set to `ftp.pasv.publicHost`) fresh in memory on every startup; nothing is persisted, matching the no-local-disk precedent already set for uploads. An operator can instead point `ftp.tls.existingSecret.name` at a Secret holding standard `tls.crt`/`tls.key` keys (the `kubernetes.io/tls` shape, which is also cert-manager's default `Certificate` output) — mirroring the existing `passwordExistingSecret` pattern. Which path is used is presence-based: setting `existingSecret.name` is the toggle, there's no separate mode enum. The chart mounts that Secret at a fixed, hardcoded container path (not configurable — the chart controls both the mount and the app binary in the same release).

Rotation is handled in-app rather than via a chart-side restart. `ftpserverlib`'s `GetTLSConfig` driver hook is invoked per-connection, so the app re-stats the mounted cert file's mtime on each call and reloads it when it changes — giving zero-downtime pickup of cert-manager's renewal cycle for free, instead of needing a Helm checksum annotation to force a rolling restart on every Secret update.

## Consequences

The self-signed cert changes on every pod restart; clients connecting via explicit FTPS to a self-signed cert have no stable identity to pin to anyway, so this is not a regression. Usernames and passwords aside, no TLS material survives a restart unless the operator supplies their own Secret.
