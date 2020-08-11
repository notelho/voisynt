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
	var targetMessage = args.Message
	var outputDir = args.Output

	fmt.Println("targetMessage: " + targetMessage)
	fmt.Println("outputDir: " + outputDir)

	if io.FileExists(outputDir) {

		var fileName = audio.AudioName(targetMessage)
		var filePath = io.FilePath(outputDir, fileName)

		if !io.FileExists(filePath) {
			tempDir := io.TempDir(outputDir)
			audio.DownloadVoice(targetMessage, fileName, tempDir)
			audio.SynthVoice(fileName, tempDir)
			audio.CreateVoice(fileName, outputDir)
			io.RemoveDir(tempDir)
		}

		fmt.Print(io.FilePath(filePath, fileName))

	} else {
		error.ThrowExit("Cant resolve provided output path", 1)
	}

}
