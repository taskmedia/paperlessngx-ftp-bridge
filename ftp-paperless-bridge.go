package main

import (
	"bytes"
	"crypto/tls"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jlaffaye/ftp"
)

func main() {
	// Read configuration from environment variables
	ftpHost := os.Getenv("FTP_HOST")
	ftpUsername := os.Getenv("FTP_USERNAME")
	ftpPassword := os.Getenv("FTP_PASSWORD")
	ftpPath := os.Getenv("FTP_PATH")
	paperlessUrl := os.Getenv("PAPERLESS_URL")
	paperlessUser := os.Getenv("PAPERLESS_USER")
	paperlessPassword := os.Getenv("PAPERLESS_PASSWORD")
	paperlessApiUrl := paperlessUrl + "/api/documents/post_document/"

	if ftpHost == "" || ftpUsername == "" || ftpPassword == "" || paperlessUrl == "" || paperlessUser == "" || paperlessPassword == "" {
		log.Fatalf("One or more required environment variables are missing")
	}

	// Establish FTP connection with explicit SSL/TLS
	conn, err := ftp.Dial(ftpHost, ftp.DialWithTimeout(5*time.Second), ftp.DialWithExplicitTLS(&tls.Config{
		InsecureSkipVerify: true,
	}))
	if err != nil {
		log.Fatalf("Failed to connect to FTP server: %v", err)
	}
	defer func() {
		if err := conn.Quit(); err != nil {
			log.Printf("Failed to close FTP connection: %v", err)
		}
	}()

	// Login to FTP server
	err = conn.Login(ftpUsername, ftpPassword)
	if err != nil {
		log.Fatalf("Failed to login to FTP server: %v", err)
	}

	// List files in the FTP server root directory
	entries, err := conn.List(ftpPath)
	if err != nil {
		log.Fatalf("Failed to list files on FTP server: %v", err)
	}

	// Create a Resty client
	client := resty.New()

	// Iterate over the files and process .pdf files
	for _, entry := range entries {
		if entry.Type != ftp.EntryTypeFile || filepath.Ext(entry.Name) != ".pdf" {
			log.Printf("Skipping file: %s", entry.Name)
			continue
		}
		log.Printf("Detected PDF file: %s", entry.Name)

		// Download the file from FTP server
		resp, err := conn.Retr(entry.Name)
		if err != nil {
			log.Printf("Failed to retrieve file %s: %v", entry.Name, err)
			continue
		}

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp)
		if err != nil {
			log.Printf("Failed to read file %s: %v", entry.Name, err)
			continue
		}
		resp.Close()

		// Upload the file to the Paperless-ngx API
		apiResp, err := client.R().
			SetBasicAuth(paperlessUser, paperlessPassword).
			SetFileReader("document", entry.Name, buf).
			Post(paperlessApiUrl)

		if err != nil {
			log.Printf("Failed to upload file %s to API: %v", entry.Name, err)
			continue
		}

		if apiResp.IsError() {
			log.Printf("API returned an error for file %s: %s", entry.Name, apiResp.Status())
			continue
		}

		log.Printf("Successfully uploaded file %s to API", entry.Name)

		// Delete the file from FTP server
		err = conn.Delete(entry.Name)
		if err != nil {
			log.Printf("Failed to delete file %s from FTP server: %v", entry.Name, err)
			continue
		}

		log.Printf("Successfully deleted file %s from FTP server", entry.Name)
	}

	log.Println("All files processed. Exiting.")
}
