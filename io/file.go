package io

import (
	"os"
	"path/filepath"

	"github.com/notelho/voisynt/error"
)

func CreateDir(dir string) string {
	path := filepath.Join(dir, TmpDir)
	err := os.MkdirAll(path, 0755)
	if err != nil {
		error.Throw(error.Exceptions.TempDirFailed)
	}
	return path
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
	return filepath.Join(path, name)
}
