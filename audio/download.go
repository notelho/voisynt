package audio

import (
	"net/url"

	"github.com/notelho/voisynt/cli"
	"github.com/notelho/voisynt/io"
)

func CreateAudioLink(message string, language string) string {
	query := url.QueryEscape(message)
	target := "http://translate.google.com/translate_tts"
	params := "?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=" + query + "&tl=" + language
	return target + params
}

func VoiceDownload(audioName io.AudioName, dirName io.AudioDir, audioConfig AudioConfig) {
	audioDir := dirName.Tmp
	audioMessage := audioConfig.Message
	audioLanguage := audioConfig.Language
	audioDownloaded := audioName.Downloaded
	audioPath := io.Path(audioDir, audioDownloaded)
	audioLink := CreateAudioLink(audioMessage, audioLanguage)
	cli.CopyFromStreaming(audioLink, audioPath)
}
