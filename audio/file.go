package audio

import (
	"regexp"
	"strings"
)

var extension string = ".mp3"

func normalizeMessage(message string) string {
	message = strings.ReplaceAll(message, " ", "_")
	reg, _ := regexp.Compile("[^a-zA-Z0-9_]+")
	message = reg.ReplaceAllString(message, "")
	message = strings.ToUpper(message)
	return message
}

func AudioName(message string) string {
	return normalizeMessage(message) + extension
}

func AudioNamePrefix(name string, prefix string) string {
	return prefix + "@" + name
}

func CreateAudioInfo(message string) AudioInfo {
	var output = AudioName(message)
	return AudioInfo{
		Message:             message,
		Output:              output,
		Downloaded:          AudioNamePrefix(output, "downloaded"),
		Thin:                AudioNamePrefix(output, "thin"),
		Thick:               AudioNamePrefix(output, "thick"),
		Robot:               AudioNamePrefix(output, "robot"),
		ThinRobot:           AudioNamePrefix(output, "thin-robot"),
		ThickRobot:          AudioNamePrefix(output, "thick-robot"),
		MergeThickThin:      AudioNamePrefix(output, "merge-thick-thin"),
		MergeThickThinRobot: AudioNamePrefix(output, "merge-thick-thin-robot"),
		MonoOutput:          AudioNamePrefix(output, "mono-output-file"),
		MonoOutputThick:     AudioNamePrefix(output, "mono-output-thick"),
	}
}

type AudioInfo struct {
	Message             string
	Output              string
	Downloaded          string
	Thin                string
	Thick               string
	Robot               string
	ThinRobot           string
	ThickRobot          string
	MergeThickThin      string
	MergeThickThinRobot string
	MonoOutput          string
	MonoOutputThick     string
}
