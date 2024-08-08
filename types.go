package main

type BotConfig struct {
	BotToken string
}

type VideoInfo struct {
	Formats []struct {
		Ext    string `json:"ext"`
		Acodec string `json:"acodec"`
		Vcodec string `json:"vcodec"`
	} `json:"formats"`
}

type MediaURLs struct {
	VideoURL string
	AudioURL string
}
