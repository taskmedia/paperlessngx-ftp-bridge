# ghcr.io/taskmedia/paperlessngx-ftp-bridge-image
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go .
RUN go build -o ftp-paperless-bridge .

FROM scratch

# Image annotations
# see: https://github.com/opencontainers/image-spec/blob/main/annotations.md#pre-defined-annotation-keys
LABEL org.opencontainers.image.title=paperless-ftp-bridge
LABEL org.opencontainers.image.description="uploads files to a paperless-ng instance via FTP"
LABEL org.opencontainers.image.url=https://github.com/taskmedia/paperlessngx-ftp-bridge/pkgs/container/paperless-ftp-bridge
LABEL org.opencontainers.image.source=https://github.com/taskmedia/paperlessngx-ftp-bridge/blob/main/Dockerfile
LABEL org.opencontainers.image.vendor=task.media
LABEL org.opencontainers.image.licenses=MIT

COPY --from=builder /app/ftp-paperless-bridge /ftp-paperless-bridge
CMD ["/ftp-paperless-bridge"]
