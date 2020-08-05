package error

import (
	"errors"
	"log"
	"os"
)

func ThrowExit(message string, status int) {
	err := errors.New(message)
	log.Fatal(err)
	os.Exit(status)
}

func ThrowLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
