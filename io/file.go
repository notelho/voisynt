package io

import (
	"os"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// func FileCreate ( path string file file ) {}