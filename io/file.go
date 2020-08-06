package io

import (
	"io/ioutil"
	"os"

	"github.com/enbot/voisynt/error"
)

func TempDir(path string) string {
	name, err := ioutil.TempDir(path, "temp")
	if err != nil {
		error.ThrowExit("Failed to create audio file from download", 1)
	}
	return name
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func FilePath(path string, name string) string {
	return path + "\\" + name
}
