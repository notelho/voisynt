package io

import (
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/enbot/voisynt/error"
)

const language string = "en"

func DownloadAudio(message string, language string) io.ReadCloser {
	query := url.QueryEscape(message)
	target := "http://translate.google.com/translate_tts"
	params := "?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=" + query + "&tl=" + language
	link := target + params
	res, err := http.Get(link)
	if err != nil {
		error.ThrowExit("Failed to download tts file", 1)
	}
	return res.Body
}

func FileFromStream(path string, stream io.ReadCloser) *os.File {
	file, createErr := os.Create(path)
	_, copyErr := io.Copy(file, stream)
	defer stream.Close()
	if createErr != nil || copyErr != nil {
		error.ThrowExit("Failed to create audio file from download", 1)
	}
	return file
}
