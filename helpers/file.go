package helpers

import (
	"os"
	"regexp"
	"strings"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// func FileCreate ( path string file file ) {}

func FileName(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	reg, _ := regexp.Compile("[^a-zA-Z0-9_]+")
	name = reg.ReplaceAllString(name, "")
	name = strings.ToUpper(name)
	return name
}
