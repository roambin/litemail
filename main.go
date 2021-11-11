package main

import (
	"github.com/emersion/go-smtp"
	"github.com/roambin/litemail/web"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

// The Backend implements SMTP server methods.
type Backend struct{}

func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	log.Println("Login:", username, password)
	return &Session{}, nil
}

func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	log.Println("AnonymousLogin")
	return &Session{}, nil
}

// A Session is returned after EHLO.
type Session struct{
	toAddr string
}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	log.Println("Mail from:", from)
	return nil
}

func (s *Session) Rcpt(to string) error {
	log.Println("Rcpt to:", to)
	s.toAddr = to
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		log.Println("Data err:", err)
		return err
	} else {
		log.Println("Data:", string(b))
		web.DealMail(s.toAddr, string(b))
	}
	return nil
}

func (s *Session) Reset() {
	log.Println("Reset")
}

func (s *Session) Logout() error {
	log.Println("Logout")
	return nil
}

func main() {
	if len(os.Args) == 4 && os.Args[1] == "post" {
		n, _ := strconv.Atoi(os.Args[3])
		for i := 1; i <= n; i++ {
			time.Sleep(time.Second)
			web.PostMail(os.Args[2], i)
		}
	} else {
		start()
	}
}

func start() {
	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = ":25"
	s.Domain = "localhost"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}