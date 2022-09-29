package controller

import "gopkg.in/gomail.v2"

func SendMail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "xxx@qq.com")
	m.SetHeader("To", "xxxx@outlook.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	d := gomail.NewDialer("smtp.qq.com", 587, "xxxx@qq.com", "ispwarsqnbluhaec")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
