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

	botToken := os.Getenv("BOT_TOKEN")

	return &BotConfig{
		BotToken: botToken,
	}
}
