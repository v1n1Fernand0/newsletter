package email

import (
	"log"

	"gopkg.in/gomail.v2"
)

func SendEmail(smtpServer string, smtpPort int, username, password, from, to, subject, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(smtpServer, smtpPort, username, password)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send email: %v", err)
	}
}
