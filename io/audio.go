package io

import (
	"io"
	"os"
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

func AudioCreate(path string, name string, stream io.ReadCloser) *os.File {
	audio := AudioPath(path, name)
	dir, _ := os.Open(path)
	file, _ := os.Create(audio)
	_, _ = io.Copy(file, stream)
	defer dir.Close()
	defer stream.Close()
	return file
}
