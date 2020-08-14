package io

import (
	"io/ioutil"
	"os"

	"github.com/enbot/voisynt/error"
)

func TempDir(path string) string {
	name, err := ioutil.TempDir(path, "tmpdir")
	if err != nil {
		error.ThrowExit("Failed to create a temporary directory", 1)
	}
	return name
}

func RemoveDir(path string) {
	defer os.RemoveAll(path)
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func Path(path DirType, name AudioType) string {
	return path + name
}
