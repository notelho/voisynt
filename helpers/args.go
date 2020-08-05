package helpers

import (
	"flag"
)

func ArgumentsCreate() Arguments {
	messagePtr := flag.String("message", "", "The message you want to get an audio file")
	pathPtr := flag.String("path", "", "The audio folder path to save and check the cache")
	flag.Parse()
	if *messagePtr == "" {
		ThrowExit("The --message attribute is required", 1)
	} else if *pathPtr == "" {
		ThrowExit("The --path attribute is required", 1)
	}
	return Arguments{Message: *messagePtr, Path: *pathPtr}
}

type Arguments struct {
	Message string
	Path    string
}
