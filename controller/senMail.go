package controller

import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

func createRandCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixMicro()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}

func SendMail(email string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "xxxx@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "验证码!")
	m.SetBody("text/html", `<h1>Hello validation code is <span style="color:red;">`+createRandCode()+`</span> !</h1>`)

	d := gomail.NewDialer("smtp.qq.com", 587, "xxxx@qq.com", "xxxx")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return
}
