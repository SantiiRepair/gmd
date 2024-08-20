package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/react"
)

func init() {
	var err error
	this, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	tempDir = filepath.Join(this, "temp")
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		panic(err)
	}

	for _, url := range []string{pornBlackList} {
		resp, err := http.Get(url)
		if errors.Is(err, nil) {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if errors.Is(err, nil) {
				sources[url] = string(body)
			}
		}
	}

}

func main() {
	bot, err := tele.NewBot(tele.Settings{
		URL:    botConfig().BotAPI,
		Token:  botConfig().BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, ctx tele.Context) {
			fmt.Println(err.Error())
		},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle(tele.OnText, func(c tele.Context) error {
		messageText := c.Text()

		if isURL(messageText) && urlExists(messageText) {
			emoji := getEmojiFor(messageText)

			errChan := make(chan error)

			go func() {
				errChan <- c.Bot().React(c.Recipient(), c.Message(), react.React(emoji))
			}()

			if err := <-errChan; err != nil {
				return fmt.Errorf("error reacting to message: %w", err)
			}

			media, err := getMediaSource(messageText)
			if err != nil {
				return fmt.Errorf("error getting media source: %w", err)
			}

			return c.SendAlbum(tele.Album{media})
		}

		return nil
	})

	bot.Start()
}
