package main

import (
	"fmt"
	visa "mycurrencynotifier/api"
	mailer "mycurrencynotifier/mails"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("config.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println("Welcome to currency notifier, we will be tracking the currency and notifying you when it reaches a certain value")

	appId := os.Getenv("APP_ID")
	rate, err := visa.GetCurrency(appId)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("The current rate is %v, and the currency is %v\n", rate.Price, rate.Currency)

	msg := []byte(fmt.Sprintf("To: vperplexity07@onemail.host\r\n"+
		"Subject: Hello!\r\n"+
		"\r\n"+
		"The current rate is %v, and the currency is %v\r\n", rate.Price, rate.Currency))

	error := mailer.SendMailViaTLS("vperplexity07@onemail.host", os.Getenv("SMTP_EMAIL"), "Hello!", os.Getenv("SMTP_SERVER"), os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"), msg)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Email Sent Successfully!")
	}
}
