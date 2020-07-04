package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/moutend/go-mail/pkg/mail"
)

// MailCC contains mail CC values.
type MailCC struct {
	CC []string
}

// Set implements flag.Value.
func (v *MailCC) Set(s string) error {
	v.CC = append(v.CC, s)

	return nil
}

// String implements flag.Value.
func (v MailCC) String() string {
	return strings.Join(v.CC, ",")
}

// Strings returns mail cc values.
func (v MailCC) Strings() []string {
	return v.CC
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("error: ")

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		mailSubject, mailTo, mailFromName, mailFromAddress string
		mailCC                                             MailCC
		htmlBody, textBody                                 []byte
	)

	flagVerbose := flag.Bool("verbose", false, "enable verbose output")
	flagSubject := flag.String("subject", "", "mail subject")
	flagTo := flag.String("to", "", "mail to")
	flagFromName := flag.String("from-name", "", "mail from name")
	flagFromAddress := flag.String("from-address", "", "mail from address")
	flagHTMLPath := flag.String("html", "", "Path to HTML file (optional)")
	flagTextPath := flag.String("text", "", "Path to Text file (optional)")
	flag.Var(&mailCC, "cc", "mail CC")

	flag.Parse()

	if flagSubject != nil {
		mailSubject = *flagSubject
	}
	if flagTo != nil {
		mailTo = *flagTo
	}
	if flagFromName != nil && *flagFromName != "" {
		mailFromName = *flagFromName
	}
	if flagFromAddress != nil && *flagFromAddress != "" {
		mailFromAddress = *flagFromAddress
	} else {
		mailFromAddress = mail.GetDefaultUsername()
	}
	if flagHTMLPath != nil && *flagHTMLPath != "" {
		htmlBody, _ = ioutil.ReadFile(*flagHTMLPath)
	}
	if flagTextPath != nil && *flagTextPath != "" {
		textBody, _ = ioutil.ReadFile(*flagTextPath)
	}

	option := mail.Option{
		To:          mailTo,
		CC:          mailCC.Strings(),
		Subject:     mailSubject,
		FromName:    mailFromName,
		FromAddress: mailFromAddress,
		HTMLBody:    string(htmlBody),
		TextBody:    string(textBody),
	}

	if flagVerbose != nil && *flagVerbose {
		fmt.Println("You are attempt to send an email ...\n")
		fmt.Printf("subject: %q\n", option.Subject)
		fmt.Printf("to: %q\n", option.To)
		fmt.Printf("from: \"%s <%s>\"\n", option.FromName, option.FromAddress)
		fmt.Printf("cc: %v\n", option.CC)
		fmt.Printf("body (HTML): %q\n", option.HTMLBody)
		fmt.Printf("body (Text): %q\n", option.TextBody)
	}

	return mail.Send(option)
}
