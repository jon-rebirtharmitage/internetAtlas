package main

import (
 	"fmt"
	"net/smtp"
)



func SendMail(to []string, subject string, msg string) error {
	var (
		username string = "jon@rebirtharmitage.com"
		password string = "starLight7"
		host     string = "smtp.gmail.com"
		port     string = "587"
	)
	
	auth := smtp.PlainAuth(
		"",
		username,
		password,
		host,
	)

	address := fmt.Sprintf("%v:%v", host, port)

	//	build our message
	body := []byte("Subject: " + subject + "\r\n\r\n" + msg)

	err := smtp.SendMail(
		address,
		auth,
		username,
		to,
		body,
	)
	if err != nil {
		return err
	}

	return nil
}

func main(){
	b := []string{"jon@rebirtharmitage.com"}
	a := SendMail(b, "Subject", "This is my message.")
	fmt.Println(a)
}