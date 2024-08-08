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

func hostname(s string) string {
	parts := strings.Split(s, ".")
	if len(parts) >= 2 {
		return parts[len(parts)-2]
	}

	return ""
}

func getEmojiFor(s string) react.Reaction {
	var emoji react.Reaction
	parsedUrl, err := url.Parse(s)
	if errors.Is(err, nil) {
		for _, data := range sources {
			urls := strings.Split(data, "\n")
			for _, u := range urls {
				url := hostname(parsedUrl.Hostname())
				if url == hostname(u) {
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
