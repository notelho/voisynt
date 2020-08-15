package error

import "os"

var deletionPaths = make([]string, 0)

func DeleteOnError(path string) {
	deletionPaths = append(deletionPaths, path)
}

func TriggerDeletion() {
	for _, path := range deletionPaths {
		os.RemoveAll(path)
	}
}
