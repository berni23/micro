package main

import (
	"net/http"
	"github.com/dgrijalva/jwt-go/request"
)

// from, to , subject, message
func (app *Config) SendMail( w http.ResponseWriter, r *http.Request){


	type mailMesssage struct {

		From string `json:"from"`
		To string `json:"to"`
		Subject string `json"subject"`
		Message string `json:message"`
	}


	var requestPayload mailMesssage
	err:= app.readJSON(w,r,&requesxtPayload)

	if err!=nil {
		app.errorJSON(w,err)
		return 
	}

	msg:= Message {

		From : requestPayload.From,
		To: requestPayload.To,
		Subject: requestPayload.Subject,
		Data: requestPayload.Message
	}

	err =app.Mailer.SendSMTPMessage(msg)

	if err!=nil{

		app.errorJSON(w,err)

		return 
	}

	payload = jsonResponse{
		Error: false,
		Message: "sent to " + requestPayload.To,
	}

	app.writeJSON(w, http.StatusAccepted,payload)
}

