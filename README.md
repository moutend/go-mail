# Sending an Email via SMTP Server in Go

This repository shows how to send an email via SMTP server in Go. It assumes that you have a Gmail account, create an account before getting started.

## Setup

1. Visit https://myaccount.google.com
2. Open Security, App Passwords and then generate an app password.
3. For security reasons, delete the issued password if no longer used.

## Configure environment variables

Run the following commands to set the environment variables:

```console
export GMAIL_SMTP_HOST="smtp.gmail.com"
export GMAIL_SMTP_PORT=587
export GMAIL_SMTP_USERNAME="you@gmail.com"
export GMAIL_SMTP_PASSWORD="REPLACE ME!"
```

## Test

Install `send` command at first.

```console
go get -u github.com/moutend/go-mail/cmd/send
```


Now, you are able to send an email by the following command.

```console
send \
  --subject "Hello, World!" \
  --to recipient@example.com \
  -from-name "Your name" \
  --from-address "Your email address" \
```

For more information, run `send --help`.

## Caution

- You are NOT able to send 500 emails per day.
- You might encounter an security alerts when youd sent many emails in short period.

## See also

1. [How To Use Google's SMTP Server | DigitalOcean](https://www.digitalocean.com/community/tutorials/how-to-use-google-s-smtp-server)
2. [Limits for sending & getting mail - Gmail Help](https://support.google.com/mail/answer/22839?hl=en)
