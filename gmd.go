package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func init() {
	var err error
	this, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	tempDir = filepath.Join(this, "temp")
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		panic(err)
	}

	for _, tool := range tools {
		_, err = exec.LookPath(tool)
		if err != nil {
			fmt.Printf("%s is not installed\n", tool)
			os.Exit(1)
		}
	}

	for _, url := range []string{pornBlackList} {
		resp, err := http.Get(url)
		if errors.Is(err, nil) {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if errors.Is(err, nil) {
				sources[url] = string(body)
			}
		}
	}

	ascii := filepath.Join(this, ".github", "ascii.txt")
	file, err := os.Open(ascii)
	if errors.Is(err, nil) {
		defer file.Close()
		content, err := io.ReadAll(file)
		if errors.Is(err, nil) {
			fmt.Println(string(content))
		}
	}
}
