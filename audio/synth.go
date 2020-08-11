package audio

import (
	"os"

	"github.com/enbot/voisynt/io"
)

func DownloadVoice(targetMessage string, fileName string, tempDir string) *os.File {
	downloadFileName := DownloadAudioName(targetMessage)
	downloadFilePath := io.FilePath(tempDir, downloadFileName)
	downloadStream := io.DownloadAudio(targetMessage, "en")
	downloadFile := io.FileFromStream(downloadFilePath, downloadStream)
	return downloadFile
}

func SynthVoice(fileName string, tempDir string) {

	// audio.CreateBaseFiles(fileName, tempDir)
	// audio.CreateRoboticFiles(fileName, tempDir)
	// audio.CreateMergedFiles(fileName, tempDir)
	// audio.CreateFinalFile(fileName, outputDir)

}

func CreateVoice(fileName string, outputDir string) {
	// > downloaded > thin e thick
	// > thick-robot , downloaded-robot e thin-robot
	// > merge thick e thin > merge downlaoded e thick-thin
	// > mono > thick again
}
