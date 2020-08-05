package audio

import (
	"io"
	"net/http"
	"net/url"
)

const language string = "en"

func DownloadAudio(s, lang string) (io.ReadCloser, error) {
	resp, err := http.Get("http://translate.google.com/translate_tts" +
		"?ie=UTF-8&tl=" + lang + "&q=" + url.QueryEscape(s))
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
