package io

import (
	"regexp"
	"strings"
)

var extension string = ".mp3"

func normalizeName(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	reg, _ := regexp.Compile("[^a-zA-Z0-9_]+")
	name = reg.ReplaceAllString(name, "")
	name = strings.ToUpper(name)
	return name
}

func DefaultAudioName(name string) string {
	return normalizeName(name) + extension
}

func DownloadAudioName(name string) string {
	return "download@" + normalizeName(name) + extension
}

func DistortedAudioName(name string) string {
	return "distorted@" + normalizeName(name) + extension
}

func RoboticAudioName(name string) string {
	return "robotic@" + normalizeName(name) + extension
}
