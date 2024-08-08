package main

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"os"
	"path/filepath"

	"os/exec"
)

var sources = map[string]string{}

func getMediaSource(url string) (string, error) {
	this, err := os.Getwd()
	if err != nil {
		return "", err
	}

	UUID := uuid.New().String()
	mediaFilePath := filepath.Join("temp", UUID+".mp4")
	downloadCmd := exec.Command("yt-dlp", "--merge-output-format", "mp4", "-f", "bestvideo+bestaudio[ext=m4a]/best", "-o", mediaFilePath, url)

	var downloadOut bytes.Buffer
	downloadCmd.Stdout = &downloadOut
	downloadCmd.Stderr = &downloadOut

	if err := downloadCmd.Run(); err != nil {
		return "", err
	}

	if _, err := os.Stat(mediaFilePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file path not found for the downloaded media")
	}

	return filepath.Join(this, mediaFilePath), nil
}
