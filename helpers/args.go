package helpers

import (
	"flag"
	"log"
	"os"
)

// CreateArguments function that read an object from the CLI input and throws an exit if didn't find it
func CreateArguments() Arguments {

	messagePtr := flag.String("message", "", "The message you want to get an audio file")
	pathPtr := flag.String("path", "", "The audio folder path to save and check the cache")

	if *messagePtr == "" {
		log.Fatal("The --message attribute is required")
		os.Exit(1)
	}

	if *pathPtr == "" {
		log.Fatal("The --path attribute is required")
		os.Exit(1)
	}

	flag.Parse()

	return Arguments{Message: *messagePtr, Path: *pathPtr}

}

// Arguments struct that have message and path properties
type Arguments struct {
	Message string
	Path    string
}
