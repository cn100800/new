package common

import (
	"crypto/tls"
	"fmt"
	"time"

	//"github.com/go-ozzo/ozzo-config"
	"gopkg.in/gomail.v2"
)

type CnMail struct {
	Host     string
	Port     int
	Name     string
	Username string
	Password string
	From     string
	To       string
}

func (m *CnMail) Setup(s *mailServer) {
	m.Host = s.host
	m.Port = s.port
	m.Username = s.username
	m.Password = s.password
	m.From = s.from
	m.To = s.to
	m.Name = s.fromName
	// c := config.New()
	// c.Load("config/mail.json")
	// m.Host = c.GetString("host")
	// m.Port = c.GetString("port")
	// m.Name = c.GetString("name")
	// m.Username = c.GetString("username")
	// m.Password = c.GetString("password")
	// m.From = c.GetString("send_from")
	// m.To = c.GetString("send_to")
}

func (m *CnMail) getTitle() string {
	l := time.Now()
	y, mm, d := l.Date()
	t := fmt.Sprintf("%d-%s-%d", y, mm, d)
	return t
}

func (m *CnMail) SendMail(context string) {
	g := gomail.NewMessage()
	g.SetHeader("From", m.From)
	g.SetHeader("To", m.To)

	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	g.SetHeader("Subject", m.getTitle())
	g.SetBody("text/html", context)
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(m.Host, 587, m.Username, m.Password)
	// 解决 x509: certificate signed by unknown authority
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(g); err != nil {
		panic(err)
	}
}

func NewCnMail() *CnMail {
	return &CnMail{}
}
