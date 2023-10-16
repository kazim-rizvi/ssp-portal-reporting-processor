package email

import (

	// "net/smtp"
	"ssp-portal-reporting-processor/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
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

func SendEmailWithSES() error {
	// Initialize an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("your-aws-region"),
	})
	if err != nil {
		return err
	}

	// Create an SES client
	svc := ses.New(sess)

	// Specify the sender's email address
	sender := "your-sender@example.com"

	// Specify the recipient email addresses
	to := []string{"recipient1@example.com"}
	cc := []string{"cc@example.com"}
	bcc := []string{"bcc@example.com"}

	toAddresses := make([]*string, len(to))
	for i, address := range to {
		toAddresses[i] = aws.String(address)
	}

	ccAddresses := make([]*string, len(cc))
	for i, address := range cc {
		ccAddresses[i] = aws.String(address)
	}

	bccAddresses := make([]*string, len(bcc))
	for i, address := range bcc {
		bccAddresses[i] = aws.String(address)
	}

	// Create the recipient list
	recipients := &ses.Destination{
		ToAddresses:  toAddresses,
		CcAddresses:  ccAddresses,
		BccAddresses: bccAddresses,
	}
	// Create the message subject and body
	subject := "Your Email Subject"
	htmlBody := "<html><body><p>This is an HTML email content.</p></body></html>"

	// Compose the message
	message := &ses.Message{
		Subject: &ses.Content{
			Data: aws.String(subject),
		},
		Body: &ses.Body{
			Html: &ses.Content{
				Data: aws.String(htmlBody),
			},
		},
	}

	// Specify the email input parameters
	input := &ses.SendEmailInput{
		Destination: recipients,
		Message:     message,
		Source:      aws.String(sender),
	}

	// Send the email
	_, err = svc.SendEmail(input)
	if err != nil {
		return err
	}

	return nil
}
