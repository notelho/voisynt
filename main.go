package main

import (
	"fmt"

	"github.com/notelho/voisynt/audio"
	"github.com/notelho/voisynt/cli"
	"github.com/notelho/voisynt/error"
	"github.com/notelho/voisynt/io"
)

func main() {
	args := cli.CreateArguments()
	audioMessage := args.Message
	outputDir := args.Output
	tmpDir := args.Tmp

	if !io.FileExists(outputDir) {
		error.Throw(error.Exceptions.PathNotFound)
	}

	fileName := io.AudioFileName(audioMessage)
	filePath := io.Path(outputDir, fileName)
	tempPath := io.Path(outputDir, tmpDir)
	error.DeleteOnError(tempPath)

	if !io.FileExists(filePath) {
		tmpDir := io.CreateDir(outputDir)
		audioConfig := audio.CreateAudioConfig(audioMessage, "en")
		audioName := io.CreateAudioName(audioMessage)
		dirName := io.CreateAudioDir(outputDir, tmpDir)
		audio.VoiceDownload(audioName, dirName, audioConfig)
		audio.VoiceSynth(audioName, dirName)
		io.RemoveDir(tmpDir)
	}

	fmt.Print(io.Path(filePath, fileName))
}
