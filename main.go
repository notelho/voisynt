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
	var targetMessage = args.Message // "hello world"
	var outputDir = args.Output      // "C:\Users\NathanBotelho\www\enbot\encore\audio"

	fmt.Println("targetMessage: " + targetMessage)
	fmt.Println("outputDir: " + outputDir)

	fmt.Println("===========================================================")

	if io.FileExists(outputDir) { // "C:\Users\NathanBotelho\www\enbot\encore\audio"

		var fileName = audio.DefaultAudioName(targetMessage) // "HELLO_WORLD.mp3"
		var filePath = io.FilePath(outputDir, fileName)      // "C:\Users\NathanBotelho\www\enbot\encore\audio\HELLO_WORLD.mp3"

		if !io.FileExists(filePath) {

			fmt.Println("===========================================================")

			tempDir := io.TempDir(outputDir) // "C:\Users\NathanBotelho\www\enbot\encore\audio\temp"

			audio.CreateBaseFiles(fileName, tempDir)
			audio.CreateRoboticFiles(fileName, tempDir)
			audio.CreateMergedFiles(fileName, tempDir)
			audio.CreateFinalFile(fileName, outputDir)

			// io.MoveFile( final audio , output dir )
			// io.RenameFile ()

			// // 	audio = audio.AudioDistort(file)

			// io.DeleteTempFolder(tempDir)

		}

		// fmt.Print(io.AudioPath(path, name))

	} else {
		error.ThrowExit("Cant resolve provided output path", 1)
	}

}
