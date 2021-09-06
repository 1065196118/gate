package helpers

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

func EmailUsersWelcomeMessage(email string) {
	from := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	to := []string{email}
	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("emails/users/welcome.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "Puneet Singh",
		Message: "This is a test message in a HTML template",
	})

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
}
