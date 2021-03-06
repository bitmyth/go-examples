package main

import "gopkg.in/gomail.v2"

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "bob@example.com", "cora@example.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/Users/gsh/Downloads/logo.png")

	d := gomail.NewDialer("localhost", 1025, "", "")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
