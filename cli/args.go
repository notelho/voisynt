package cli

import (
	"flag"

	"github.com/enbot/voisynt/error"
)

func ArgumentsCreate() Arguments {
	messagePtr := flag.String("message", "", "The message you want to get an audio file")
	outputPtr := flag.String("output", "", "The audio folder output to save and check the cache")
	flag.Parse()
	if *messagePtr == "" {
		error.ThrowExit("The --message attribute is required", 1)
	} else if *outputPtr == "" {
		error.ThrowExit("The --output attribute is required", 1)
	}
	return Arguments{
		Message: *messagePtr,
		Output:  *outputPtr,
	}
}

type Arguments struct {
	Message string
	Output  string
}
