package main

import (
	log "log/slog"
	"os"
	"strings"

	"github.com/taskmedia/paperlessngx-ftp-bridge/internal/ftpserver"
	"github.com/taskmedia/paperlessngx-ftp-bridge/internal/paperless"
)

// Config holds the environment-derived settings for the embedded FTP
// server and the paperless-ngx client it uploads to. Account and
// allowed-extension configuration is intentionally minimal here; a
// follow-on ticket widens it to a multi-account, Helm-driven schema.
type Config struct {
	ftpListenAddr        string
	ftpUsername          string
	ftpPassword          string
	ftpAllowedExtensions []string
	ftpPublicHost        string
	paperlessURL         string
	paperlessUser        string
	paperlessPassword    string
}

func main() {
	setLogLevel()

	log.Info("Starting FTP-Paperless bridge...")
	config := loadConfig()

	// Start health check server
	go startHealthCheckServer()

	paperlessClient := paperless.NewClient(config.paperlessURL, config.paperlessUser, config.paperlessPassword)

	srv, err := ftpserver.NewServer(ftpserver.Config{
		ListenAddr:        config.ftpListenAddr,
		Username:          config.ftpUsername,
		Password:          config.ftpPassword,
		AllowedExtensions: config.ftpAllowedExtensions,
		PublicHost:        config.ftpPublicHost,
		Uploader:          paperlessClient,
	})
	if err != nil {
		log.Error("Failed to prepare embedded FTP server", "error", err)
		os.Exit(1)
	}

	log.Info("Starting embedded FTP server...")
	if err := srv.ListenAndServe(); err != nil {
		log.Error("FTP server stopped", "error", err)
		os.Exit(1)
	}
}

func loadConfig() Config {
	allowedExtensions := ftpserver.DefaultAllowedExtensions
	if raw := os.Getenv("FTP_ALLOWED_EXTENSIONS"); raw != "" {
		allowedExtensions = strings.Split(raw, ",")
	}

	config := Config{
		ftpListenAddr:        os.Getenv("FTP_LISTEN_ADDR"),
		ftpUsername:          os.Getenv("FTP_USERNAME"),
		ftpPassword:          os.Getenv("FTP_PASSWORD"),
		ftpAllowedExtensions: allowedExtensions,
		ftpPublicHost:        os.Getenv("FTP_PASV_PUBLIC_HOST"),
		paperlessURL:         os.Getenv("PAPERLESS_URL"),
		paperlessUser:        os.Getenv("PAPERLESS_USER"),
		paperlessPassword:    os.Getenv("PAPERLESS_PASSWORD"),
	}

	if config.ftpUsername == "" || config.ftpPassword == "" || config.paperlessURL == "" || config.paperlessUser == "" || config.paperlessPassword == "" {
		log.Error("One or more required environment variables are missing")
		os.Exit(1)
	}

	return config
}

func setLogLevel() {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "ERROR"
	}
	logLevel = strings.ToUpper(logLevel)

	switch logLevel {
	case "DEBUG":
		log.SetLogLoggerLevel(log.LevelDebug)
	case "INFO":
		log.SetLogLoggerLevel(log.LevelInfo)
	case "WARN":
		log.SetLogLoggerLevel(log.LevelWarn)
	case "ERROR":
		log.SetLogLoggerLevel(log.LevelError)
	default:
		log.SetLogLoggerLevel(log.LevelInfo)
	}
}
