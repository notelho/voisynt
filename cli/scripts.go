package cli

import (
	"os/exec"

	"github.com/enbot/voisynt/error"
	"github.com/enbot/voisynt/io"
)

func SynthThickVoice(inputName string, outputName string, outputDir string) {
	inputAudioPath := io.FilePath(outputDir, inputName)
	outputAudioPath := io.FilePath(outputDir, outputName)

	err := exec.Command("ffmpeg", "-i", inputAudioPath, "-af", "asetrate=44100*0.25,aresample=44100,atempo=2.17", outputAudioPath).Run()

	if err != nil {
		error.ThrowExit("Failed to create "+outputName+" with "+inputName, 1)
	}
}

// func SynthThinVoice(inputName string, outputName string, outputDir string) {
// 	cli.exec("ffmpeg -i test.mp3 -af \"asetrate=44100*1.07,atempo=0.5\" thin.mp3")
// }

// func SynthRobotVoice(inputName string, outputName string, outputDir string) {
// 	cli.exec("ffmpeg -i test.mp3 -filter_complex "afftfilt=real='hypot(re,im)*sin(0)':imag='hypot(re,im)*cos(0)':win_size=512:overlap=0.75" robot.mp3")
// }

// func MergeManyFiles(inputName string, outputName string, outputDir string) {
// 	// cli.exec("")
// }

// func MergeFilesAdjustVolumn(inputName string, outputName string, outputDir string) {
// 	// cli.exec("")
// }
