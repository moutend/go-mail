// Package mail provides SMTP utility.
package mail

import (
	"os"
	"strconv"
)

var (
	defaultSMTP *smtp
)

// GetDefaultUsername returns default username.
func GetDefaultUsername() string {
	return defaultSMTP.Username
}

// Send sends an email with default configuration.
func Send(opt Option) error {
	return defaultSMTP.Send(opt)
}

func parseInt(s string) int {
	i64, _ := strconv.ParseInt(s, 10, 64)

	return int(i64)
}

func init() {
	defaultSMTP = &smtp{
		Host:     os.Getenv("GMAIL_SMTP_HOST"),
		Port:     parseInt(os.Getenv("GMAIL_SMTP_PORT")),
		Username: os.Getenv("GMAIL_SMTP_USERNAME"),
		Password: os.Getenv("GMAIL_SMTP_PASSWORD"),
	}
}
