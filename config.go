package main

import (
	"os"

	env "github.com/joho/godotenv"
)

func init() {
	err := env.Load(".env")
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
