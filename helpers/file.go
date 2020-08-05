package helpers

import "os"

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func FixName(name string) string {
	return "ak"
}
