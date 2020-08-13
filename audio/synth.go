package audio

import (
	"os"

	"github.com/enbot/voisynt/cli"
	"github.com/enbot/voisynt/io"
)

func DownloadVoice(audioInfo AudioInfo, tempDir string) *os.File {
	downloadMessage := audioInfo.Message
	downloadStream := io.DownloadAudio(downloadMessage, "en")
	downloadPath := io.FilePath(tempDir, audioInfo.Downloaded)
	downloadFile := io.FileFromStream(downloadPath, downloadStream)
	return downloadFile
}

func VoiceSynth(audioInfo AudioInfo, tempDir string, outputDir string) {
	thinPath := io.FilePath(tempDir, audioInfo.Thin)
	thickPath := io.FilePath(tempDir, audioInfo.Thick)
	robotPath := io.FilePath(tempDir, audioInfo.Robot)
	downloadPath := io.FilePath(tempDir, audioInfo.Downloaded)

	cli.SynthThinVoice(downloadPath, thinPath)
	cli.SynthThickVoice(downloadPath, thickPath)
	cli.SynthRobotVoice(downloadPath, robotPath)

	thinRobotPath := io.FilePath(tempDir, audioInfo.ThinRobot)
	thickRobotPath := io.FilePath(tempDir, audioInfo.ThickRobot)

	cli.SynthRobotVoice(thinPath, thinRobotPath)
	cli.SynthRobotVoice(thickPath, thickRobotPath)

	mergeThickThinPath := io.FilePath(tempDir, audioInfo.MergeThickThin)
	mergeThickThinRobotPath := io.FilePath(tempDir, audioInfo.MergeThickThinRobot)
	outputPath := io.FilePath(outputDir, audioInfo.Output)

	cli.MergeVoiceFiles(thickRobotPath, 1, thinRobotPath, 1, mergeThickThinPath)
	cli.MergeVoiceFiles(robotPath, 2, mergeThickThinPath, 0.5, mergeThickThinRobotPath)
	cli.SynthThickVoice(mergeThickThinRobotPath, outputPath)
}
