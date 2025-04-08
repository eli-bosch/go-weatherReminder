package email

import (
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmailWithSendGrid(toEmail, subject, body string) error {
	from := mail.NewEmail("Weather Reminder", "example-email@example.com")
	to := mail.NewEmail("User", toEmail)
	message := mail.NewSingleEmail(from, subject, to, body, body)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)
	if err != nil {
		return err
	} else {
		log.Println("Email Status: ", response.StatusCode)
		log.Println("Email headers: ", response.Headers)
	}

	return nil
}
