[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/taskmedia)](https://artifacthub.io/packages/helm/taskmedia/paperlessngx-ftp-bridge)

# Helm chart: paperless-ngx FTP bridge

Kubernetes [Helm](https://helm.sh) chart to automatically upload PDF files from a FTP server to paperless-ngx.

This application will automatically search for PDF files on your FTP server and upload them to the paperless-ngx API.
The application will run as a cronjob and will be executed every 5 minutes (can be changed).

You can use this application e.g. if your document scanner can only upload files to a FTP server.
With this bridge your scan device will be able to upload the documents directly with the FTP as file storage inbetween.

## Configuration

The configuration of the application will be set in the [`values.yaml`](./values.yaml)-file.
Everything is pretty straight forward and should be self-explanatory.
If you think more information should be provided or need help, feel free to open an issue.

## Installation

To deploy the Helm chart first copy the [`values.yaml`](./values.yaml)-file and customize your deployment.
After it was modified you can deploy the chart with the following command.

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
