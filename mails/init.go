package mails

import (
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"messagewith-server/env"
	"strconv"
)

var (
	smtpClient *mail.SMTPClient
	Service    *service
)

func InitService() {
	server := mail.NewSMTPClient()
	server.Host = env.SmtpHost
	port, err := strconv.ParseUint(env.SmtpPort, 10, 32)
	if err != nil {
		panic(err)
	}
	server.Port = int(port)
	server.Username = env.SmtpUsername
	server.Password = env.SmtpPassword

	c, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	smtpClient = c
	Service = GetService(&Client{})
}
