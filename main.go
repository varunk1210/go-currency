package main

import (
	"fmt"
	visa "mycurrencynotifier/api"
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
}
