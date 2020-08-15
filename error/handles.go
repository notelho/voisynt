package error

import (
	"log"
	"os"
)

func Throw(err *StatusError) {
	TriggerDeletion()
	Fatal(err)
	Exit(err)
}

func Exit(err *StatusError) {
	os.Exit(err.Code)
}

func Fatal(err *StatusError) {
	log.Fatal(err.Message)
}
