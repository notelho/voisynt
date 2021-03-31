package audio

import (
	"github.com/notelho/voisynt/cli"
	"github.com/notelho/voisynt/io"
)

func VoiceSynth(audioName io.AudioName, dirName io.AudioDir) {
	thinPath := io.Path(dirName.Tmp, audioName.Thin)
	thickPath := io.Path(dirName.Tmp, audioName.Thick)
	robotPath := io.Path(dirName.Tmp, audioName.Robot)
	downloadPath := io.Path(dirName.Tmp, audioName.Downloaded)

	cli.SynthThinVoice(downloadPath, thinPath)
	cli.SynthThickVoice(downloadPath, thickPath)
	cli.SynthRobotVoice(downloadPath, robotPath)

	thinRobotPath := io.Path(dirName.Tmp, audioName.ThinRobot)
	thickRobotPath := io.Path(dirName.Tmp, audioName.ThickRobot)

	cli.SynthRobotVoice(thinPath, thinRobotPath)
	cli.SynthRobotVoice(thickPath, thickRobotPath)

	mergeThickThinPath := io.Path(dirName.Tmp, audioName.MergeThickThin)
	mergeThickThinRobotPath := io.Path(dirName.Tmp, audioName.MergeThickThinRobot)
	outputPath := io.Path(dirName.Output, audioName.Output)

	cli.MergeVoiceFiles(thickRobotPath, 1, thinRobotPath, 1, mergeThickThinPath)
	cli.MergeVoiceFiles(robotPath, 2, mergeThickThinPath, 0.5, mergeThickThinRobotPath)
	cli.SynthThickVoice(mergeThickThinRobotPath, outputPath)
}
