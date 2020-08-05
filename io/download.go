package io

import (
	"io"
	"net/http"
	"net/url"

	"github.com/enbot/voisynt/error"
)

const language string = "en"

func DownloadAudio(message string, language string) io.ReadCloser {
	// func DownloadAudio(message string, language string) io.Reader {
	query := url.QueryEscape(message)
	target := "http://translate.google.com/translate_tts"
	params := "?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=" + query + "&tl=" + language
	link := target + params
	res, err := http.Get(link)
	if err != nil {
		error.ThrowExit("Failed to download tts file", 1)
	}
	// defer res.Body.Close()
	return res.Body
}
