# ghcr.io/taskmedia/paperlessngx-ftp-bridge-image
#
# The builder always runs on the build host's native platform (BUILDPLATFORM)
# and cross-compiles for the requested target via GOOS/GOARCH. This avoids
# running the Go toolchain itself under QEMU emulation for non-native target
# platforms, which is by far the slowest part of a multi-platform build.
FROM --platform=$BUILDPLATFORM golang:1.27-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY *.go .
COPY internal ./internal
ARG TARGETOS
ARG TARGETARCH
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o ftp-paperless-bridge .

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
