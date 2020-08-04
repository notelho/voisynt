package helper

import (
	"flag"

	errors "github.com/enbot/voisynt/errors"
)

func ArgumentsCreate() Arguments {
	messagePtr := flag.String("message", "", "The message you want to get an audio file")
	pathPtr := flag.String("path", "", "The audio folder path to save and check the cache")
	flag.Parse()
	if *messagePtr == "" {
		errors.ThrowExit(errors.ErrorAttributeMessage)
	} else if *pathPtr == "" {
		errors.ThrowExit(errors.ErrorAttributePath)
	}
	return Arguments{Message: *messagePtr, Path: *pathPtr}
}

type Arguments struct {
	Message string
	Path    string
}
