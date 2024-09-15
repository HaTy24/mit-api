package mail

import (
	"log"
	"net/smtp"
)

func Send(subject string, body string) {
	from := "hathienty2000@gmail.com"
	pass := "mvfo ayga psdn pyhe"
	to := "hathienty24@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Println("Successfully sended to " + to)
}
