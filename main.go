package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/react"
)

func init() {
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
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle(tele.OnText, func(c tele.Context) error {
		s := c.Text()
		if isURL(s) && urlExists(s) {
			emoji := getEmojiFor(s)
			go c.Bot().React(c.Recipient(), c.Message(), react.React(emoji))
			media, err := getMediaSource(s)
			if errors.Is(err, nil) {
				return c.SendAlbum(tele.Album{media})
			}
		}

		return nil
	})

	bot.Start()
}
