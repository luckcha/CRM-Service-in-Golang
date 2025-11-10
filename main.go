package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"crm-exotel/db"
	"crm-exotel/exotel"
	"crm-exotel/webhook"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	db.InitDB()
	go webhook.StartWebhookServer()

	agent := "+917302588760"
	lead := "+917302652380"
	exophone := os.Getenv("EXOPHONE_NUMBER")
	webhookURL := "https://YOUR_NGROK_URL/webhook"

	err = exotel.MakeCall(agent, lead, exophone, webhookURL)
	if err != nil {
		log.Println("❌ Call failed:", err)
	} else {
		fmt.Println("✅ Call initiated successfully")
	}

	select {}
}
