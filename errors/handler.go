package errors

import (
	"log"
	"os"
)

type ExitError string

const (
	ErrorAttributeMessage ExitError = "The --message attribute is required"
	ErrorAttributePath    ExitError = "The --path attribute is required"
	ErrorPathNotFound     ExitError = "Provided path not found"
)

func ThrowExit(message string) {
	log.Fatal(message)
	os.Exit(1)
}
