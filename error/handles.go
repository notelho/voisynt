package error

import (
	"log"
)

func LogFatalMessage(err string) {
	if err != "" {
		log.Fatal(err)
	}
}

func LogFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
