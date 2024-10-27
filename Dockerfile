FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go .
RUN go build -o ftp-paperless-bridge .

FROM scratch
COPY --from=builder /app/ftp-paperless-bridge /ftp-paperless-bridge
CMD ["/ftp-paperless-bridge"]
