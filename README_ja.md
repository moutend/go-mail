# GoでSMTPサーバーを経由してメール送信


このリポジトリでは、GoでSMTPサーバーを経由してメールを送信する方法を紹介します。Gmailアカウントを持っていることを前提としていますので、始める前にアカウントを作成してください。

## 準備

1. https://myaccount.google.comに移動する
2. セキュリティ→アプリパスワード→パスワードを生成する
3. 生成したパスワードはメール送信のテストが終わったら削除しておく

## 環境変数の設定

以下のコマンドを実行して環境変数を設定します。

```console
export GMAIL_SMTP_HOST="smtp.gmail.com"
export GMAIL_SMTP_PORT=587
export GMAIL_SMTP_USERNAME="you@example.com"
export GMAIL_SMTP_PASSWORD="xxxxxxxx"
```

## テスト

まずは`send`コマンドをインストールします。

```console
go get -u github.com/moutend/go-mail/cmd/send
```

以下のコマンドでメール送信できます。

send \
  --subject "Hello, World!" \
  --to recipient@example.com \
  -from-name "Your name" \
  --from-address "Your email address" \
```

## 注意事項

- 1日に送信できるメールの件数は500件です。
- 短時間になんどもメール送信するとセキュリティアラートが発生する恐れがあります

## 参考資料

1. [How To Use Google's SMTP Server | DigitalOcean](https://www.digitalocean.com/community/tutorials/how-to-use-google-s-smtp-server)
2. [メールの送受信数に関する制限 - Gmail ヘルプ](https://support.google.com/mail/answer/22839?hl=ja)
