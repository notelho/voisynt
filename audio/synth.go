package audio

import (
	"os"

	"github.com/enbot/voisynt/cli"
	"github.com/enbot/voisynt/io"
)

func DownloadVoice(audioInfo AudioInfo, outputDir string) *os.File {
	downloadMessage := audioInfo.Message
	downloadStream := io.DownloadAudio(downloadMessage, "en")
	downloadFilePath := io.FilePath(outputDir, audioInfo.Downloaded)
	downloadFile := io.FileFromStream(downloadFilePath, downloadStream)
	return downloadFile
}

func VoiceSynth(audioInfo AudioInfo, tempDir string, outputDir string) {

	// cli.SynthThinVoice(audioInfo.Downloaded, audioInfo.Thin, tempDir)
	cli.SynthThickVoice(audioInfo.Downloaded, audioInfo.Thick, tempDir)
	// SynthRobotVoice(audioInfo.Downloaded, audioInfo.Robot, tempDir)
	// SynthRobotVoice(audioInfo.Thin, audioInfo.ThinRobot, tempDir)
	// SynthRobotVoice(audioInfo.Thick, audioInfo.ThickRobot, tempDir)

	// audio.SynthRobot(audioInfo, tempDir)
	// audio.MergeVoices(audioInfo, tempDir)
	// audio.SynthFinal(audioInfo, outputDir)

	// thinFileName := AudioNamePrefix(fileName, "thin")
	// thickFileName := AudioNamePrefix(fileName, "thick")
	// SynthThinVoice(thinFileName, tempDir)
	// SynthThickVoice(thickFileName, tempDir)

	// thinRobotFileName := AudioNamePrefix(fileName, "thin-robot")
	// thickRobotFileName := AudioNamePrefix(fileName, "thick-robot")
	// downloadRobotFileName := AudioNamePrefix(fileName, "downloaded-robot")
	// SynthRobotVoice(thinRobotFileName, tempDir)
	// SynthRobotVoice(thickRobotFileName, tempDir)
	// SynthRobotVoice(downloadRobotFileName, tempDir)

	// > merge thick e thin > merge downlaoded e thick-thin

	// > mono > thick again
}
