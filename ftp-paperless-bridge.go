package main

import (
	"bytes"
	"crypto/tls"
	log "log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jlaffaye/ftp"
	"github.com/robfig/cron/v3"
)

type Config struct {
	ftpHost           string
	ftpUsername       string
	ftpPassword       string
	ftpPath           string
	paperlessURL      string
	paperlessUser     string
	paperlessPassword string
	paperlessApiURL   string
	interval          string
}

func main() {
	setLogLevel()

	log.Info("Starting FTP-Paperless bridge...")
	config := loadConfig()

	// Start health check server
	go startHealthCheckServer()

	// check readiness
	if !readinessProbe(config) {
		log.Error("Initial readiness probe failed")
		os.Exit(1)
	}

	c := cron.New()

	_, err := c.AddFunc(config.interval, func() {
		success := handle(config)
		updateLastResults(success)
	})
	if err != nil {
		log.Error("Failed to schedule job", "error", err)
		os.Exit(1)
	}

	c.Start()

	// Run the first job immediately
	success := handle(config)
	updateLastResults(success)

	// Wait forever
	select {}
}

func handle(config Config) bool {
	log.Debug("Starting file processing...")

	// Establish FTP connection with explicit SSL/TLS
	conn, err := ftp.Dial(
		config.ftpHost,
		ftp.DialWithTimeout(5*time.Second),
		ftp.DialWithExplicitTLS(&tls.Config{
			InsecureSkipVerify: true,
		}))
	if err != nil {
		log.Warn("Failed to connect to FTP server", "error", err)
		return false
	}
	defer quitFTPConnection(conn)

	// Login to FTP server
	err = conn.Login(config.ftpUsername, config.ftpPassword)
	if err != nil {
		log.Warn("Failed to login to FTP server", "error", err)
		return false
	}

	// List files in the FTP server root directory
	entries, err := conn.List(config.ftpPath)
	if err != nil {
		log.Warn("Failed to list files on FTP server", "error", err)
		return false
	}

	// Iterate over the files and process .pdf files
	for _, entry := range entries {
		processFile(conn, entry, config)
	}

	log.Debug("All files processed. Exiting.")
	return true
}

func processFile(conn *ftp.ServerConn, entry *ftp.Entry, config Config) {
	if entry.Type != ftp.EntryTypeFile || filepath.Ext(entry.Name) != ".pdf" {
		log.Debug("Skipping file", "fileName", entry.Name)
		return
	}
	log.Debug("Detected PDF file", "fileName", entry.Name)

	// Download the file from FTP server
	resp, err := conn.Retr(entry.Name)
	if err != nil {
		log.Warn("Failed to retrieve file", "fileName", entry.Name, "error", err)
		return
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp)
	if err != nil {
		log.Warn("Failed to read file", "fileName", entry.Name, "error", err)
		return
	}
	resp.Close()

	// Create a Resty client
	client := resty.New()

	// Upload the file to the Paperless-ngx API
	apiResp, err := client.R().
		SetBasicAuth(config.paperlessUser, config.paperlessPassword).
		SetFileReader("document", entry.Name, buf).
		Post(config.paperlessApiURL)

	if err != nil {
		log.Warn("Failed to upload file to API", "fileName", entry.Name, "error", err)
		return
	}

	if apiResp.IsError() {
		log.Warn("API returned an error for file", "fileName", entry.Name, "status", apiResp.Status())
		return
	}
	log.Debug("Successfully uploaded file to API", "fileName", entry.Name)

	// Delete the file from FTP server
	err = conn.Delete(entry.Name)
	if err != nil {
		log.Error("Failed to delete file from FTP server", "fileName", entry.Name, "error", err)
		return
	}
	log.Debug("Successfully deleted file from FTP server", "fileName", entry.Name)

	log.Info("Successfully processed file", "fileName", entry.Name)
}

func loadConfig() Config {
	interval := os.Getenv("CRON_SCHEDULE")
	if interval == "" {
		interval = "*/5 7-20 * * *" // Default to every 5 minutes from 7 AM to 8 PM
	}

	config := Config{
		ftpHost:           os.Getenv("FTP_HOST"),
		ftpUsername:       os.Getenv("FTP_USERNAME"),
		ftpPassword:       os.Getenv("FTP_PASSWORD"),
		ftpPath:           os.Getenv("FTP_PATH"),
		paperlessURL:      os.Getenv("PAPERLESS_URL"),
		paperlessUser:     os.Getenv("PAPERLESS_USER"),
		paperlessPassword: os.Getenv("PAPERLESS_PASSWORD"),
		paperlessApiURL:   os.Getenv("PAPERLESS_URL") + "/api/documents/post_document/",
		interval:          interval,
	}

	if config.ftpHost == "" || config.ftpUsername == "" || config.ftpPassword == "" || config.paperlessURL == "" || config.paperlessUser == "" || config.paperlessPassword == "" {
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

func quitFTPConnection(conn *ftp.ServerConn) {
	if err := conn.Quit(); err != nil {
		log.Warn("Failed to close FTP connection", "error", err)
	}
}
