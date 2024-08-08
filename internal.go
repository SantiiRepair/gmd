package main

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	tele "gopkg.in/telebot.v3"

	"os/exec"
)

var sources = map[string]string{}

func getMediaSource(url string) (tele.Inputtable, error) {
	this, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	UUID := uuid.New().String()
	mediaFilePath := filepath.Join(this, "temp", UUID+".mp4")
	downloadCmd := exec.Command("yt-dlp", "--merge-output-format", "mp4", "-f", "bestvideo+bestaudio[ext=m4a]/best", "-o", mediaFilePath, url)

	var downloadOut bytes.Buffer
	downloadCmd.Stdout = &downloadOut
	downloadCmd.Stderr = &downloadOut

	if err := downloadCmd.Run(); err != nil {
		return nil, err
	}

	fileContent, err := os.ReadFile(mediaFilePath)
	if err != nil {
		return nil, err
	}

	contentType := http.DetectContentType(fileContent)
	if strings.Contains(contentType, "video") {
		v := &tele.Video{File: tele.FromDisk(mediaFilePath)}
		return v, nil
	} else if strings.Contains(contentType, "audio") {
		a := &tele.Audio{File: tele.FromDisk(mediaFilePath)}
		return a, nil
	} else {
		return nil, errors.New("")
	}
}
