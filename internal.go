package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	tele "gopkg.in/telebot.v3"

	"os/exec"
)

// getMediaSource retrieves media from the given URL and returns a Telegram inputtable object.
func getMediaSource(url string) (tele.Inputtable, error) {
	fileId := uuid.New().String()

	mediaInfo, err := fetchMediaInfo(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch media info: %w", err)
	}

	output, err := downloadMedia(url, mediaInfo, fileId)
	if err != nil {
		return nil, fmt.Errorf("failed to download media: %w", err)
	}

	return createTelegramMedia(mediaInfo, output)
}

// fetchMediaInfo retrieves media information using yt-dlp.
func fetchMediaInfo(url string) (MediaInfo, error) {
	formatCmd := exec.Command("yt-dlp", "--get-format", "-j", url)
	var formatOut bytes.Buffer
	formatCmd.Stdout = &formatOut
	formatCmd.Stderr = &formatOut

	if err := formatCmd.Run(); err != nil {
		return MediaInfo{}, err
	}

	result := formatOut.String()
	lines := strings.Split(result, "\n")
	if len(lines) > 1 {
		result = strings.Join(lines[1:], "\n")
	}

	var mediaInfo MediaInfo
	if err := json.Unmarshal([]byte(result), &mediaInfo); err != nil {
		return MediaInfo{}, fmt.Errorf("failed to unmarshal media info: %w", err)
	}

	return mediaInfo, nil
}

// downloadMedia downloads the media file based on the provided media information.
func downloadMedia(url string, mediaInfo MediaInfo, fileId string) (string, error) {
	formats := reverseSlice(mediaInfo.Formats)
	bestFormat := formats.([]Format)[0]
	formatId := bestFormat.FormatId

	fileName := fmt.Sprintf("%s.%s", fileId, mediaInfo.Ext)
	output := filepath.Join(tempDir, fileName)

	cmd := exec.Command("yt-dlp", "-f", formatId, "--no-overwrites", "-o", output, url)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return output, nil
}

// createTelegramMedia creates a Telegram media object based on the media type.
func createTelegramMedia(mediaInfo MediaInfo, output string) (tele.Inputtable, error) {
	mediaType := whichFormat(mediaInfo.Format)
	if mediaType == Video {
		return &tele.Video{
			File:      tele.FromDisk(output),
			FileName:  mediaInfo.Filename,
			Width:     mediaInfo.Width,
			Height:    mediaInfo.Height,
			Duration:  int(mediaInfo.Duration),
			Thumbnail: &tele.Photo{File: tele.FromURL(mediaInfo.Thumbnail)},
		}, nil
	} else if mediaType == Audio {
		return &tele.Audio{
			File:     tele.FromDisk(output),
			Title:    mediaInfo.Title,
			FileName: mediaInfo.Filename,
		}, nil
	}

	return nil, errors.New("could not determine content type")
}
