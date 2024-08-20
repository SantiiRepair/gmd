package main

import (
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
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

func reverseSlice(s interface{}) interface{} {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return nil
	}

	reversed := reflect.MakeSlice(val.Type(), val.Len(), val.Cap())
	for i := 0; i < val.Len(); i++ {
		reversed.Index(i).Set(val.Index(val.Len() - 1 - i))
	}

	return reversed.Interface()
}

// getEmojiFor returns an emoji based on the provided URL string.
func getEmojiFor(s string) react.Reaction {
	parsedUrl, err := url.Parse(s)
	if err != nil {
		return getRandomEmoji()
	}

	for _, data := range sources {
		urls := strings.Split(data, "\n")
		for _, u := range urls {
			if hostname(parsedUrl.Hostname()) == hostname(u) {
				return react.EvilFace
			}
		}
	}

	return getRandomEmoji()
}

// getRandomEmoji returns a random emoji from a predefined list.
func getRandomEmoji() react.Reaction {
	options := []react.Reaction{react.Eyes, react.Sunglasses}
	return options[rand.Intn(len(options))]
}
