package audio

func CreateBaseFiles(fileName string, outputDir string) string {
	// > downloaded > thin e thick
	return ""
}

func CreateRoboticFiles(fileName string, outputDir string) string {
	// > thick-robot , downloaded-robot e thin-robot
	return ""
}

func CreateMergedFiles(fileName string, outputDir string) string {
	// > merge thick e thin > merge downlaoded e thick-thin
	return ""
}

func CreateFinalFile(fileName string, outputDir string) string {
	// > mono > thick again
	return ""
}

// downloadFileName := io.DownloadAudioName(targetMessage)
// downloadFilePath := io.FilePath(tempDir, downloadFileName)
// downloadStream := io.DownloadAudio(targetMessage, "en")
// downloadFile := io.FileFromStream(downloadFilePath, downloadStream)

// fmt.Println("===========================================================")

// var distortFileName = io.DistortedAudioName(targetMessage) // "distort@HELLO_WORLD.mp3"
// var robotFileName = io.RoboticAudioName(targetMessage)     // "robot@HELLO_WORLD.mp3"

// fmt.Println(downloadFileName)
// fmt.Println(distortFileName)
// fmt.Println(robotFileName)
// fmt.Println(downloadFile)

// fmt.Println("===========================================================")
