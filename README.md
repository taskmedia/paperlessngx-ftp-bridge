[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/taskmedia)](https://artifacthub.io/packages/helm/taskmedia/paperlessngx-ftp-bridge)

# paperless-ngx FTP bridge

An embedded FTP/FTPS server that forwards every uploaded file straight to the paperless-ngx API.

Your scanner connects directly to this bridge over FTP; there's no separate FTP server to run or keep in sync, and uploads reach paperless-ngx the moment the transfer finishes.

## Helm chart

The Helm chart for this application is maintained in [taskmedia/helm](https://github.com/taskmedia/helm), not in this repository:

- Source: [`charts/paperlessngx-ftp-bridge`](https://github.com/taskmedia/helm/tree/main/charts/paperlessngx-ftp-bridge)
- Upgrade notes: [`docs/MIGRATION.md`](https://github.com/taskmedia/helm/blob/main/charts/paperlessngx-ftp-bridge/docs/MIGRATION.md)
- Issues and pull requests for the chart belong in [taskmedia/helm](https://github.com/taskmedia/helm/issues)

## Installation

```bash
$ helm repo add taskmedia https://helm.task.media
$ helm repo update

$ helm show values taskmedia/paperlessngx-ftp-bridge > ./my-values.yaml
$ vi ./my-values.yaml

$ helm upgrade --install paperlessngx-ftp-bridge taskmedia/paperlessngx-ftp-bridge --values ./my-values.yaml
```

You can also use OCI Helm charts from [ghcr.io](https://ghcr.io/):

```bash
$ helm upgrade --install paperlessngx-ftp-bridge oci://ghcr.io/taskmedia/paperlessngx-ftp-bridge
```
