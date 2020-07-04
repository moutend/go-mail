package mail

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSend(t *testing.T) {
	t.Parallel()

	require.Error(t, Send(Option{
		To:          "",
		CC:          []string{},
		FromName:    "Testing",
		FromAddress: "test@example.com",
		Subject:     "Hello, World!",
		HTMLBody:    "<h1>Hi!</h1>",
		TextBody:    "Hey!",
	}))
}

func TestGetDefaultUsername(t *testing.T) {
	t.Parallel()

	require.Equal(t, os.Getenv("GMAIL_SMTP_USERNAME"), GetDefaultUsername())
}
