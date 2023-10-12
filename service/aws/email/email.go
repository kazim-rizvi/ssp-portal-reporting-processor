package email

import (
	// "fmt"
	// "net/smtp"
	"ssp-portal-reporting-processor/config"
)

type EmailSender struct {
	cfg config.EmailConfig
}

func NewEmailSender(cfg config.EmailConfig) *EmailSender {
	return &EmailSender{cfg: cfg}
}

func (es *EmailSender) SendEmail() error {
	// subject := "Reporting Service Notification"
	// message := "Hello, reporting is complete!"

	// // auth := smtp.PlainAuth("", es.cfg.SMTPUsername, es.cfg.SMTPPassword, es.cfg.SMTPServer)
	// // to := es.cfg.RecipientEmails
	// // msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, message))

	// err := smtp.SendMail(fmt.Sprintf("%s:%d", es.cfg.SMTPServer, es.cfg.SMTPPort), auth, es.cfg.SenderEmail, to, msg)
	// if err != nil {
	// 	return err
	// }

	return nil
}
