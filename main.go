package main

import (
	"fmt"

	"github.com/enbot/voisynt/audio"
	"github.com/enbot/voisynt/cli"
	"github.com/enbot/voisynt/error"
	"github.com/enbot/voisynt/io"
)

func main() {
	var args = cli.CreateArguments()
	var audioMessage = args.Message
	var outputDir = args.Output

	fmt.Println("audioMessage: " + audioMessage)
	fmt.Println("outputDir: " + outputDir)

	if io.FileExists(outputDir) {

		var fileName = io.AudioFileName(audioMessage)
		var filePath = io.Path(outputDir, fileName)

		if !io.FileExists(filePath) {
			tempDir := io.TempDir(outputDir)
			audioConfig := audio.CreateAudioConfig(audioMessage, "en")
			audioName := io.CreateAudioName(audioMessage)
			dirName := io.CreateAudioDir(outputDir, tempDir)
			audio.VoiceDownload(audioName, dirName, audioConfig)
			audio.VoiceSynth(audioName, dirName)
			io.RemoveDir(tempDir)
		}

		fmt.Print(io.Path(filePath, fileName))

	} else {
		error.ThrowExit("Cant resolve provided output path", 1)
	}
}
