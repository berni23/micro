package main

import (
	"bytes"
	"html/template"
	"time"

	"github.com/vanng822/go-premailer/premailer"

	mail "github.com/xhit/go-simple-mail/v2"
)

type Mail struct {
	Domain      string
	Host        string
	port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        any
	DataMap     map[string]any
}

//third party packages for making mail service easier

func (m *Mail) SendSMTPMessage(msg Message) error {

	if msg.From == "" {
		msg.From = m.FromAddress
	}

	if msg.From == "" {
		msg.From = m.FromName
	}

	data := map[string]any{

		"message": msg.Data,
	}

	msg.DataMap = data

	formattedMessage, err := m.buildHTMLMessage(msg)

	if err != nil {
		return err
	}

	plainMessage, err := m.buildPlainTextMessage(msg)

	if err != nil {
		return err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.getEncryption(m.Encryption)

	server.KeepAlive = false // we do not want to keep mail server connection alive

	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	// call mail template and pass data to those templates ( html template for mail creation)

	smtpClient, err := server.Connect()

	if err != nil {

		return err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)
	email.SetBody(mail.TextPlain, plainMessage)
	email.AddAlternative(mail.TextHTML, formattedMessage)


	if len(msg.Attachments)>0 {


		//we have some attatchments

		for _,x:=range msg.Attachments {
			//each entry is the full path to whatever i want to attatch 
			email.AddAttachment(x)
		}
	}

	// last step, actually send the email


	err = email.Send(smtpClient)
	return err 

}

func (m *Mail) buildHTMLMessage(msg Message) (string, error) {

	//point to a specific template

	tempalteToRender := "./templates/mail.html.gohtml"

	t, err := template.New("email-html").ParseFiles(tempalteToRender)

	if err != nil {

		return _, err
	}

	var tpl bytes.Buffer

	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {

		return "", err
	}
	formattedMessage := tpl.String
	formattedMessage, err = m.inlineCSS(formattedMessage)

	if err != nil {

		return "", err
	}

	return formattedMessage, nil

}

func (m *Mail) buildPlainTextMessage(msg Message) (string, error) {

	//point to a specific template

	tempalteToRender := "./templates/mail.plain.gohtml"

	t, err := template.New("email-plain").ParseFiles(tempalteToRender)

	if err != nil {

		return _, err
	}

	var tpl bytes.Buffer

	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {

		return "", err
	}
	plainMessage := tpl.String()

	return plainMessage, nil

}

func (m *Mail) inlineCSS(s string) (string, error) {

	options := premailer.Options{

		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(s, &options)

	if err != nil {

		return "", err
	}

	html, err := prem.Transform()

	if err != nil {

		return "", err
	}

	return html, nil
}

func (m *Mail) getEncryption(s string) mail.Encryption {

	switch s {
	case "tls":

		return mail.EncryptionSTARTTLS

	case "ssl":

		return mail.EncryptionSSLTLS

	case "none", "":

		return mail.EncryptionNone

	default:

		return mail.EncryptionSTARTTLS

	}

}
