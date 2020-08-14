package io

import (
	"regexp"
	"strings"
)

var extension string = ".mp3"

type AudioType = string

type AudioName struct {
	Message
	Output              AudioType
	Downloaded          AudioType
	Thin                AudioType
	Thick               AudioType
	Robot               AudioType
	ThinRobot           AudioType
	ThickRobot          AudioType
	MergeThickThin      AudioType
	MergeThickThinRobot AudioType
}

func CreateAudioName(message string) AudioName {
	return AudioName{
		Output:              AudioFileName(message),
		Downloaded:          AudioFileNamePrefix(message, "downloaded file"),
		Thin:                AudioFileNamePrefix(message, "thin file"),
		Thick:               AudioFileNamePrefix(message, "thick file"),
		Robot:               AudioFileNamePrefix(message, "robot file"),
		ThinRobot:           AudioFileNamePrefix(message, "thin robot file"),
		ThickRobot:          AudioFileNamePrefix(message, "thick robot file"),
		MergeThickThin:      AudioFileNamePrefix(message, "merge thick thin file"),
		MergeThickThinRobot: AudioFileNamePrefix(message, "merge thick thin robot file"),
	}
}

func NormalizeAudioName(message string) AudioType {
	message = strings.ReplaceAll(message, " ", "_")
	reg, _ := regexp.Compile("[^@a-zA-Z0-9_]+")
	message = reg.ReplaceAllString(message, "")
	message = strings.ToUpper(message)
	return message
}

func AudioFileName(message string) AudioType {
	return NormalizeAudioName(message) + extension
}

func AudioFileNamePrefix(message string, prefix string) AudioType {
	return AudioFileName(prefix + "@" + message)
}
