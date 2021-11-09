package sgctl

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kr/pretty"
	"github.com/sendgrid/sendgrid-go"
	mailhelper "github.com/sendgrid/sendgrid-go/helpers/mail"
	log "github.com/sirupsen/logrus"
)

func setup_email_body(mails *MailsConfig) []byte {

	// Create new sendgrid email object
	mailv3 := mailhelper.NewV3Mail()

	// Loop all emails configs
	for _, mail := range mails.Mails {

		// Set email sender
		mailv3.SetFrom(
			mailhelper.NewEmail(mail.From.Name, mail.From.Email),
		)

		// Set email subject
		mailv3.Subject = mail.Subject

		// Create new email personalization
		p := mailhelper.NewPersonalization()

		// Set email recipients
		tos := []*mailhelper.Email{}
		for _, to := range mail.To {
			tos = append(tos, mailhelper.NewEmail(to.Name, to.Email))
		}
		p.AddTos(tos...)

		// Add email personalization
		mailv3.AddPersonalizations(p)

		// Create email content
		c := mailhelper.NewContent(mail.Content.Type, mail.Content.Value)
		mailv3.AddContent(c)

		// Add email attachments
		for _, attachment := range mail.Attachments {

			// Read file content from filepath
			if file, err := ioutil.ReadFile(attachment.Path); err != nil {
				log.Panicln("Error reading content from file path, ", err)
			} else {
				// Create attachment object
				att := mailhelper.NewAttachment()
				fileBase64 := base64.StdEncoding.EncodeToString(file)

				// Setup attachment content
				att.SetContent(fileBase64)
				att.SetFilename(filepath.Base(attachment.Path))
				att.SetDisposition("attachment")
				// Detect file mimetype
				if fileOpen, err := os.Open(attachment.Path); err != nil {
					defer fileOpen.Close()
					log.Panicln("Error getting file information, ", err)
					if fileType, err := GetFileContentType(fileOpen); err != nil {
						log.Panicln("Error getting file mimetype, ", err)
					} else {
						att.SetType(fileType)
					}
				}

				// Add file as email attachments
				mailv3.AddAttachment(att)
			}
		}
	}

	return mailhelper.GetRequestBody(mailv3)
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func send(mails *MailsConfig) {
	// pretty.Println(mails)

	// Setup sendgrid request object
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = setup_email_body(mails)

	// Send request to sendgrid api
	response, err := sendgrid.API(request)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(response)
	}
}

func send_dryrun(mails *MailsConfig) {
	// pretty.Println(mails)

	// Setup sendgrid request object
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = setup_email_body(mails)

	// Send request to sendgrid api
	pretty.Println(string(request.Body))
}
