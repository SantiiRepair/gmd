package main

var (
	this    string
	tempDir string
	sources = make(map[string]string)
	tools   = []string{"yt-dlp", "ffmpeg"}
)

const (
	pornBlackList = "https://raw.githubusercontent.com/SantiiRepair/gmd/main/.github/blacklist/porn.txt"
)
