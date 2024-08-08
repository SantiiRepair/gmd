package main

import (
	"errors"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"gopkg.in/telebot.v3/react"
)

func isURL(s string) bool {
	u, err := url.Parse(s)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func urlExists(url string) bool {
	resp, err := http.Head(url)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func getEmojiFor(s string) react.Reaction {
	var emoji react.Reaction
	parsedUrl, err := url.Parse(s)
	if errors.Is(err, nil) {
		for _, item := range sources {
			urls := strings.Split(sources[item], "\n")
			for _, u := range urls {
				if strings.Contains(u, parsedUrl.Hostname()) {
					emoji = react.EvilFace
					break
				}
			}
		}

	}

	if (emoji == react.Reaction{}) {
		options := []react.Reaction{react.Eyes, react.Sunglasses}
		emoji = options[rand.Intn(len(options))]
	}

	emoji.Type = "emoji"

	return emoji
}
