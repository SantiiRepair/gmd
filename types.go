package main

type BotConfig struct {
	BotToken string
}

type MediaInfo struct {
	Title   string `json:"title"`
	Formats []struct {
		Ext    string `json:"ext"`
		Acodec string `json:"acodec"`
		Vcodec string `json:"vcodec"`
	} `json:"formats"`
}
