package main

import (
	log "log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/taskmedia/paperlessngx-ftp-bridge/internal/ftpserver"
	"github.com/taskmedia/paperlessngx-ftp-bridge/internal/paperless"
)

// ftpAccountEnvPrefix is the env var prefix an operator's Helm-rendered FTP
// account is delivered under, e.g. FTP_ACCOUNT_SCANNER1=some-password. The
// chart validates the username portion is env-var-safe (letters, digits,
// underscore) before rendering it, so no further sanitization happens here:
// the username is recoverable directly from the var name.
const ftpAccountEnvPrefix = "FTP_ACCOUNT_"

// Config holds the environment-derived settings for the embedded FTP
// server and the paperless-ngx client it uploads to.
type Config struct {
	ftpListenAddr          string
	ftpAccounts            []ftpserver.Account
	ftpAllowedExtensions   []string
	ftpPublicHost          string
	ftpPASVPortMin         int
	ftpPASVPortMax         int
	paperlessURL           string
	paperlessUser          string
	paperlessPassword      string
	paperlessCheckInterval time.Duration
}

func main() {
	setLogLevel()

	log.Info("Starting FTP-Paperless bridge...")
	config := loadConfig()

	paperlessClient := paperless.NewClient(config.paperlessURL, config.paperlessUser, config.paperlessPassword)

	srv, err := ftpserver.NewServer(ftpserver.Config{
		ListenAddr:        config.ftpListenAddr,
		Accounts:          config.ftpAccounts,
		AllowedExtensions: config.ftpAllowedExtensions,
		PublicHost:        config.ftpPublicHost,
		PASVPortMin:       config.ftpPASVPortMin,
		PASVPortMax:       config.ftpPASVPortMax,
		Uploader:          paperlessClient,
	})
	if err != nil {
		log.Error("Failed to prepare embedded FTP server", "error", err)
		os.Exit(1)
	}

	checker := newReadinessChecker(srv, paperlessClient, config.paperlessCheckInterval)
	go checker.run()
	go startHealthCheckServer(checker)

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
		ftpListenAddr:          os.Getenv("FTP_LISTEN_ADDR"),
		ftpAccounts:            loadFTPAccounts(os.Environ()),
		ftpAllowedExtensions:   allowedExtensions,
		ftpPublicHost:          os.Getenv("FTP_PASV_PUBLIC_HOST"),
		ftpPASVPortMin:         atoiOrZero(os.Getenv("FTP_PASV_PORT_MIN")),
		ftpPASVPortMax:         atoiOrZero(os.Getenv("FTP_PASV_PORT_MAX")),
		paperlessURL:           os.Getenv("PAPERLESS_URL"),
		paperlessUser:          os.Getenv("PAPERLESS_USER"),
		paperlessPassword:      os.Getenv("PAPERLESS_PASSWORD"),
		paperlessCheckInterval: paperlessCheckInterval(),
	}

	if len(config.ftpAccounts) == 0 || config.paperlessURL == "" || config.paperlessUser == "" || config.paperlessPassword == "" {
		log.Error("One or more required environment variables are missing")
		os.Exit(1)
	}

	return config
}

// loadFTPAccounts extracts one Account per FTP_ACCOUNT_<USERNAME> entry in
// env (the "NAME=VALUE" shape of os.Environ()), sorted by username for
// deterministic ordering.
func loadFTPAccounts(env []string) []ftpserver.Account {
	var accounts []ftpserver.Account

	for _, entry := range env {
		name, value, ok := strings.Cut(entry, "=")
		if !ok || !strings.HasPrefix(name, ftpAccountEnvPrefix) {
			continue
		}

		username := strings.TrimPrefix(name, ftpAccountEnvPrefix)
		if username == "" {
			continue
		}

		accounts = append(accounts, ftpserver.Account{Username: username, Password: value})
	}

	sort.Slice(accounts, func(i, j int) bool { return accounts[i].Username < accounts[j].Username })

	return accounts
}

func atoiOrZero(raw string) int {
	if raw == "" {
		return 0
	}

	value, err := strconv.Atoi(raw)
	if err != nil {
		return 0
	}

	return value
}

// paperlessCheckInterval reads PAPERLESS_CHECK_INTERVAL_SECONDS, falling
// back to DefaultPaperlessCheckInterval when unset or invalid.
func paperlessCheckInterval() time.Duration {
	raw := os.Getenv("PAPERLESS_CHECK_INTERVAL_SECONDS")
	if raw == "" {
		return DefaultPaperlessCheckInterval
	}

	seconds, err := strconv.Atoi(raw)
	if err != nil || seconds <= 0 {
		log.Warn("Invalid PAPERLESS_CHECK_INTERVAL_SECONDS, using default", "value", raw, "default", DefaultPaperlessCheckInterval)

		return DefaultPaperlessCheckInterval
	}

	return time.Duration(seconds) * time.Second
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
