package smtp_infra

import (
	"net/smtp"
	"os"
)

func SendMail(toEmail string, sub string, bo string) {
	from := "mdfilipiaki@gmail.com"
	password := os.Getenv("GOOGLE_PASSWORD_EMAIL")

	toEmailAddress := toEmail
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := sub
	body := bo
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}
}
