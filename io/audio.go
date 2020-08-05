package io

import (
	"regexp"
	"strings"
)

func AudioName(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	reg, _ := regexp.Compile("[^a-zA-Z0-9_]+")
	name = reg.ReplaceAllString(name, "")
	name = strings.ToUpper(name)
	return name
}

func AudioPath(path string, file string) string {
	return path + "\\" + file + ".mp3"
}
