package send

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//CreateAndSendEmail will take in data and send the email
func CreateAndSendEmail(w http.ResponseWriter, r *http.Request) {

	from := mail.NewEmail(r.FormValue("name-of-sender"), r.FormValue("email-of-sender"))
	to := mail.NewEmail(r.FormValue("name-of-recipient"), r.FormValue("email-of-recipient"))
	subject := r.FormValue("subject")
	TextContent := r.FormValue("text-content")
	htmlContent := TextContent
	message := mail.NewSingleEmail(from, subject, to, TextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	res, err := client.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(res.StatusCode)
		fmt.Println(res.Body)
		fmt.Println(res.Headers)
	}

}
