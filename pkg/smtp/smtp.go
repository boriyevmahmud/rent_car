package smtp

import (
	"backend_course/rent_car/config"
	"net/smtp"
)

func SendMail(toEmail string, msg string) error {

	// Compose the email message
	from := "akromjonotaboyev@gmail.com"
	to := []string{toEmail}
	subject := "Register for RENT_CAR"
	message := msg

	// Create the email message
	body := "To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + message

	auth := smtp.PlainAuth("", config.SmtpUsername, config.SmtpPassword, config.SmtpServer)

	// Connectin to the SMTP server
	err := smtp.SendMail(config.SmtpServer+":"+config.SmtpPort, auth, from, to, []byte(body))
	if err != nil {
		return err
	}

	return nil
}
