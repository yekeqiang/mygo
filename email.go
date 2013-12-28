package main
import (
	"net/smtp"
	"log"
)

type EmailUser struct {
  Username    string
  Password    string
  EmailServer string
  Port        int
}

emailUser := &EmailUser{'yekeqiang', 'ykq19870530', 'smtp.gmail.com', 587}

auth := smtp.PlainAuth("",
  emailUser.Username,
  emailUser.Password,
  emailUser.EmailServer
)

auth := smtp.PlainAuth("",
	emailUser.Username,
	emailUser.Password,
	emailUser.EmailServer
)


type SmtpTemplateData struct {
  From    string
  To      string
  Subject string
  Body    string
}

const emailTemplate = `From: {{.From}}
To: {{.To}}}
Subject: {{.Subject}}

{{.Body}}

Sincerely,

{{.From}}
`
var err error
var doc bytes.Buffer

context := &SmtpTemplateData{
  "yekeqiang@umpay.com",
  "yekeqiang@umpay.com",
  "This is the e-mail subject line!",
  "Hello, this is a test e-mail body."
}
t := template.New("emailTemplate")
t, err = t.Parse(emailTemplate)
if err != nil {
  log.Print("error trying to parse mail template")
}
err = t.Execute(&doc, context)
if err != nil {
  log.Print("error trying to execute mail template")
}

err = smtp.SendMail(emailUser.EmailServer+":"+strconv.Itoa(emailUser.Port), // in our case, "smtp.google.com:587"
  auth,
  emailUser.Username,
  []string{"yekeqiang@gmail.com"},
  doc.Bytes())
if err != nil {
  log.Print("ERROR: attempting to send a mail ", err)
}