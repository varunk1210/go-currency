package mailer

import (
	"crypto/tls"
	"net/smtp"
)

func SendMailViaTLS(toEmail, fromEmail, subject, smtpServer, smtpUser, smtpPassword string, body []byte) error {
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpServer)

	msg := body

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer,
	}

	conn, err := tls.Dial("tcp", smtpServer+":465", tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, smtpServer)
	if err != nil {
		return err
	}

	if err = c.Auth(auth); err != nil {
		return err
	}

	if err = c.Mail(fromEmail); err != nil {
		return err
	}

	if err = c.Rcpt(toEmail); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()

	return nil
}
