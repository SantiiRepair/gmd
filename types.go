package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

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

func (m *MediaInfo) GetThumbnail() (string, error) {
	if m.Thumbnail == "" {
		return "", errors.New("no thumbnail URL provided")
	}

	thumbId := uuid.New().String()

	filename := fmt.Sprintf(tempDir, "%s.jpg", thumbId)
	thumbnailPath := filepath.Join(tempDir, filename) 

	resp, err := http.Get(m.Thumbnail)
	if err != nil {
		return "", err
	}
	
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download thumbnail: %s", resp.Status)
	}

	out, err := os.Create(thumbnailPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return thumbnailPath, nil
}