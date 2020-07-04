package mail

import (
	"fmt"
	"net/mail"

	"github.com/pkg/errors"
	gomail "gopkg.in/mail.v2"
)

// Option corresponds to mail header.
type Option struct {
	To          string
	CC          []string
	FromName    string
	FromAddress string
	Subject     string
	HTMLBody    string
	TextBody    string
}

// smtp holds values SMTP parameters.
type smtp struct {
	Host     string
	Port     int
	Username string
	Password string
}

// Send sends an email.
func (v *smtp) Send(opt Option) error {
	if v == nil {
		return fmt.Errorf("mail: failed to send an email")
	}
	if opt.To == "" {
		return errors.New("mail: mail \"to\" is required")
	}
	m := gomail.NewMessage()

	m.SetHeader(`From`, (&mail.Address{
		Name:    opt.FromName,
		Address: opt.FromAddress,
	}).String())
	m.SetHeader(`To`, opt.To)
	m.SetHeader(`Subject`, opt.Subject)
	m.SetBody(`text/plain`, opt.TextBody)
	m.AddAlternative(`text/html`, opt.HTMLBody)

	if len(opt.CC) > 0 {
		m.SetHeader(`Cc`, opt.CC...)
	}

	d := &gomail.Dialer{
		LocalName: v.Host,
		Host:      v.Host,
		Port:      v.Port,
		Username:  v.Username,
		Password:  v.Password,
	}

	if err := d.DialAndSend(m); err != nil {
		return errors.Wrap(err, "mail: failed to send an email")
	}

	return nil
}
