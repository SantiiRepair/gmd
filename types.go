package main

const (
	_ int = iota
	Video
	Audio
	Unknown
)

type BotConfig struct {
	BotToken string
	BotAPI   string
}

type MediaInfo struct {
	Title     string   `json:"title"`
	Filename  string   `json:"filename"`
	Duration  float64  `json:"duration"`
	Thumbnail string   `json:"thumbnail"`
	Formats   []Format `json:"formats"`
	Format    string   `json:"format"`
	Width     int      `json:"width"`
	Height    int      `json:"height"`
	Ext       string   `json:"ext"`
}

type Format struct {
	Ext      string `json:"ext"`
	FormatId string `json:"format_id"`
	Acodec   string `json:"acodec"`
	Vcodec   string `json:"vcodec"`
}
