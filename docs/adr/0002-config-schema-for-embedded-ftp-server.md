# Config schema for the embedded FTP server

The old client-mode config (`FTP_HOST`/`FTP_USERNAME`/`FTP_PASSWORD`/`FTP_PATH`) is replaced wholesale, and the new shape must let a second or third FTP account be added later without a breaking change. We model `ftp.accounts` as a Helm list (one entry by default), each with `username`/`password`/`passwordExistingSecret` — the same per-field secret-ref pattern already used by `paperless.password`.

Delivering that list to the container as env vars rules out a single combined value: a JSON blob or `user:pass` CSV string would require Helm to bake every account's plaintext password into one rendered value at template time, which breaks `passwordExistingSecret` (the whole point of that pattern is that Helm never sees the plaintext). Kubernetes' `valueFrom.secretKeyRef` only composes at the level of one discrete env var, so each account needs its own var. We also rejected index-based naming (`FTP_ACCOUNTS_0_USERNAME`/`_PASSWORD`) in favor of username-keyed vars — `FTP_ACCOUNT_<SANITIZED_USERNAME>=<password>` — so the username is recoverable directly from the var name instead of needing a parallel list to resolve index to username.

## Consequences

Usernames must be valid env-var-safe tokens (letters/digits/underscore); the chart validates this per entry with `required`/regex and fails fast rather than silently mangling an unsafe username. Adding an account later is a pure list append — no new var name shape, no migration.

Addendum (ADR-0006): the default `ftp.pasv.portMin`/`portMax` shipped by the chart was revised from `50000`/`50100` (101 ports) down to `50000`/`50019` (20 ports) — the knob and its shape are unchanged, only the default size, which this ADR never reasoned about in the first place.
