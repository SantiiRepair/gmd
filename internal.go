package main

import (
	"bytes"
	"encoding/json"
	"errors"

	"io"
	"os/exec"
	"strings"
)

var sources = map[string]string{}

func getMediaSource(s string) (string, error) {
	meta, err := getMetadata(s)
	if err != nil {
		return "", err
	}

	if meta.VideoURL != "" {
		ok, err := hasAudio(meta.VideoURL)
		if errors.Is(err, nil) && ok {
			return meta.VideoURL, nil
		}
	}

	if meta.AudioURL != "" && meta.VideoURL != "" {
		_, err := mergeMedia(meta.VideoURL, meta.AudioURL)
		if !errors.Is(err, nil) {
			return "", nil
		}
	}

	return meta.VideoURL, nil
}

func getMetadata(s string) (MediaURLs, error) {
	cmd := exec.Command("yt-dlp", "-f", "b", "-j", "--skip-download", s)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return MediaURLs{}, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(out.Bytes(), &result); err != nil {
		return MediaURLs{}, err
	}

	videoURL, audioURL := "", ""
	if formats, ok := result["formats"].([]interface{}); ok {
		for _, f := range formats {
			format := f.(map[string]interface{})
			if format["vcodec"] != "none" && videoURL == "" {
				videoURL = format["url"].(string)
			}
			if format["acodec"] != "none" && audioURL == "" {
				audioURL = format["url"].(string)
			}
			if videoURL != "" && audioURL != "" {
				break
			}
		}
	}

	return MediaURLs{VideoURL: videoURL, AudioURL: audioURL}, nil
}

func hasAudio(s string) (bool, error) {
	cmd := exec.Command("ffmpeg", "-i", s, "-hide_banner")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if !strings.Contains(stderr.String(), "Stream #0") {
			return false, err
		}
	}

	output := stderr.String()
	return strings.Contains(output, "Audio:"), nil
}

func mergeMedia(videoURL, audioURL string) (io.Reader, error) {
	cmd := exec.Command("ffmpeg", "-i", videoURL, "-i", audioURL, "-c:v", "copy", "-c:a", "aac", "-f", "matroska", "pipe:1")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	go func() {
		cmd.Wait()
		stdout.Close()
	}()

	return stdout, nil
}
