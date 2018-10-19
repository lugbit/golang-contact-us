package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
)

// Sender email address info. Email address, password and smtp server.
const (
	SENDER_EMAIL     string = "EmailAddress@email.com"
	SENDER_EMAIL_PWD string = "PasswordHere"
	SENDER_SMTP      string = "smtp.gmail.com"
)

// formData struct to hold the submitted form.
type formData struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Subject   string
	Body      string
}

// Global variable tpl holds all the parsed templates.
var tpl *template.Template
var err error

// Parse all .gohtml templates in the templates folder.
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/send", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	// Form is submitted.
	if req.Method == http.MethodPost {
		// Collect form input and validate
		fd := &formValidate{
			FirstName: req.FormValue("frmFirstName"),
			LastName:  req.FormValue("frmLastName"),
			Email:     req.FormValue("frmEmail"),
			Phone:     req.FormValue("frmPhone"),
			Subject:   req.FormValue("frmSubject"),
			Body:      req.FormValue("frmBody"),
			SentFlag:  false,
		}
		// Check for errors.
		if fd.validate() == false {
			err := tpl.ExecuteTemplate(res, "index.gohtml", fd)
			if err != nil {
				log.Fatalln(err)
			}
			return
		}
		// Send email
		sendEmail(*fd)

		/*
			SentFlag variable will be set to true once the email sendEmail function runs.
			It is then used in the template to check if the message sent html should be displayed.
		*/
		fd.SentFlag = true
		err := tpl.ExecuteTemplate(res, "index.gohtml", fd)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// sendEmail function takes a formData object and uses the info to send an email.
func sendEmail(fd formValidate) {
	// Authentication information for the sender email address. Configurable above.
	auth := smtp.PlainAuth("", SENDER_EMAIL, SENDER_EMAIL_PWD, SENDER_SMTP)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"Marck527@gmail.com"}
	msg := []byte("To: Marck527@gmail.com\r\n" +
		"Subject:" + fd.Subject + "\r\n" +
		"\r\n" +
		"Name: " + fd.FirstName + " " + fd.LastName + "\n" +
		"Email: " + fd.Email + "\n" +
		"Mobile: " + fd.Phone + "\n" +
		"Subject: " + fd.Subject + "\n" +
		"-----------------------------------------------------------------------------" + "\n" +
		fd.Body + "\n" + "\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, SENDER_EMAIL, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email successfully sent to: ", to)
}
