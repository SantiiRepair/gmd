package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/react"
)

func main() {
	bot, err := tele.NewBot(tele.Settings{
		URL:    botConfig().BotAPI,
		Token:  botConfig().BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, ctx tele.Context) {
			fmt.Println(err.Error())
		},
	})

	if !errors.Is(err, nil) {
		panic(err)
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
				return err
			}

			return c.SendAlbum(tele.Album{media})
		}

		return nil
	})

	var wg sync.WaitGroup

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		<-sigChan
		fmt.Println("\nReceived SIGINT, exiting...")

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bot.Start()
	}()

	fmt.Println("\nListening for updates. Interrupt (Ctrl+C) to stop.")

	wg.Wait()
}
