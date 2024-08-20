package main

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
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
