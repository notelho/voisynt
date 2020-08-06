package main

import (
	"fmt"

	"github.com/enbot/voisynt/cli"
	"github.com/enbot/voisynt/error"
	"github.com/enbot/voisynt/io"
)

func main() {

	var args = cli.ArgumentsCreate()
	var targetMessage = args.Message // "hello world"
	var outputDir = args.Output      // "C:\Users\NathanBotelho\www\enbot\encore\audio"

	fmt.Println("targetMessage: " + targetMessage)
	fmt.Println("outputDir: " + outputDir)

	fmt.Println("===========================================================")

	if io.FileExists(outputDir) { // "C:\Users\NathanBotelho\www\enbot\encore\audio"

		var defaultFileName = io.DefaultAudioName(targetMessage)      // "HELLO_WORLD.mp3"
		var defaultFilePath = io.FilePath(outputDir, defaultFileName) // "C:\Users\NathanBotelho\www\enbot\encore\audio\HELLO_WORLD.mp3"

		if !io.FileExists(defaultFilePath) {

			fmt.Println("===========================================================")

			tempDir := io.TempDir(outputDir) // "C:\Users\NathanBotelho\www\enbot\encore\audio\temp"

			fmt.Println("===========================================================")

			downloadFileName := io.DownloadAudioName(targetMessage)
			downloadFilePath := io.FilePath(tempDir, downloadFileName)
			downloadStream := io.DownloadAudio(targetMessage, "en")
			downloadFile := io.FileFromStream(downloadFilePath, downloadStream)

			fmt.Println("===========================================================")

			var distortFileName = io.DistortedAudioName(targetMessage) // "distort@HELLO_WORLD.mp3"
			var robotFileName = io.RoboticAudioName(targetMessage)     // "robot@HELLO_WORLD.mp3"

			fmt.Println(downloadFileName)
			fmt.Println(distortFileName)
			fmt.Println(robotFileName)
			fmt.Println(downloadFile)

			fmt.Println("===========================================================")

			// // 	audio = audio.AudioDistort(file)
			// // 	io.FileCreate(audio)

			// fmt.Println(file)

		}

		// io.DeleteTempFolder(tempFolder)
		// fmt.Print(io.AudioPath(path, name))

	} else {
		error.ThrowExit("Cant resolve provided output path", 1)
	}

}
