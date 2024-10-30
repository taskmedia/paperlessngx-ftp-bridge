package main

import (
	"bytes"
	"crypto/tls"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jlaffaye/ftp"
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
	interval          time.Duration
}

func main() {
	config := loadConfig()

	ticker := time.NewTicker(config.interval)
	defer ticker.Stop()

	for range ticker.C {
		handle(config)
	}
}

func handle(config Config) {
	// Establish FTP connection with explicit SSL/TLS
	conn, err := ftp.Dial(
		config.ftpHost,
		ftp.DialWithTimeout(5*time.Second),
		ftp.DialWithExplicitTLS(&tls.Config{
			InsecureSkipVerify: true,
		}))
	if err != nil {
		log.Printf("Failed to connect to FTP server: %v\n", err)
		return
	}
	defer func() {
		if err := conn.Quit(); err != nil {
			log.Printf("Failed to close FTP connection: %v\n", err)
		}
	}()

	// Login to FTP server
	err = conn.Login(config.ftpUsername, config.ftpPassword)
	if err != nil {
		log.Printf("Failed to login to FTP server: %v\n", err)
		return
	}

	// List files in the FTP server root directory
	entries, err := conn.List(config.ftpPath)
	if err != nil {
		log.Printf("Failed to list files on FTP server: %v\n", err)
		return
	}

	// Iterate over the files and process .pdf files
	for _, entry := range entries {
		processFile(conn, entry, config)
	}

	log.Println("All files processed. Exiting.")
}

func processFile(conn *ftp.ServerConn, entry *ftp.Entry, config Config) {
	if entry.Type != ftp.EntryTypeFile || filepath.Ext(entry.Name) != ".pdf" {
		log.Printf("Skipping file: %s", entry.Name)
		return
	}
	log.Printf("Detected PDF file: %s", entry.Name)

	// Download the file from FTP server
	resp, err := conn.Retr(entry.Name)
	if err != nil {
		log.Printf("Failed to retrieve file %s: %v", entry.Name, err)
		return
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp)
	if err != nil {
		log.Printf("Failed to read file %s: %v", entry.Name, err)
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
		log.Printf("Failed to upload file %s to API: %v", entry.Name, err)
		return
	}

	if apiResp.IsError() {
		log.Printf("API returned an error for file %s: %s", entry.Name, apiResp.Status())
		return
	}

	log.Printf("Successfully uploaded file %s to API", entry.Name)

	// Delete the file from FTP server
	err = conn.Delete(entry.Name)
	if err != nil {
		log.Printf("Failed to delete file %s from FTP server: %v", entry.Name, err)
		return
	}

	log.Printf("Successfully deleted file %s from FTP server", entry.Name)
}

func loadConfig() Config {
	intervalStr := os.Getenv("INTERVAL_SECONDS")
	interval := 5 * time.Minute
	if intervalStr != "" {
		var err error
		intervalInt, err := strconv.Atoi(intervalStr)
		if err != nil {
			log.Fatalf("Invalid INTERVAL_SECONDS value: %v", err)
		}
		interval = time.Duration(intervalInt) * time.Second
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
		log.Fatalf("One or more required environment variables are missing")
	}

	return config
}
