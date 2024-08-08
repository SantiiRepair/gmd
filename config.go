package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("No .env file found")
	}
}

func botConfig() *BotConfig {

	botApi := os.Getenv("BOT_API")
	botToken := os.Getenv("BOT_TOKEN")

	return &BotConfig{
		BotAPI:   botApi,
		BotToken: botToken,
	}
}
