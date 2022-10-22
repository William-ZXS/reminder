package main

import (
	"gopkg.in/gomail.v2"
)

func (r Mail) SendMail(subject string, body string) error {
	m := gomail.NewMessage()
	// 这种方式可以添加别名，即 nickname， 也可以直接用<code>m.SetHeader("From", MAIL_USER)</code>
	nickname := "家庭助手:"
	m.SetHeader("From", nickname+"<"+r.User+">")
	// 发送给多个用户
	m.SetHeader("To", r.Receiver...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(r.Host, r.Port, r.User, r.Pwd)
	// 发送邮件
	err := d.DialAndSend(m)
	return err
}
