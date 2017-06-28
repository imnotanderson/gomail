package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func main() {
	err := SendToMail("user", "123456", "smtp.gmail.com:587", "xxx@xx.com", "hello", `nothing here <a href="http://google.com">xxxxxxx</a>`, "html")
	fmt.Print(err)
}

func SendToMail(authuser, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", authuser, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + authuser + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, authuser, send_to, msg)
	return err
}
