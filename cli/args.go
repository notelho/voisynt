package cli

import (
	"flag"

	"github.com/enbot/voisynt/error"
	"github.com/enbot/voisynt/io"
)

type ArgumentsType = string

type Arguments struct {
	Message ArgumentsType
	Output  ArgumentsType
	Tmp     ArgumentsType
}

func CreateArguments() Arguments {
	messagePtr := flag.String("message", "", "The message you want to get an audio file")
	outputPtr := flag.String("output", "", "The audio folder output to save and check the cache")
	tmpPtr := flag.String("tmp", io.TmpDir, "The audio tmp folder to create temporary files")
	flag.Parse()
	if *messagePtr == "" {
		error.Throw(error.Exceptions.ArgsMessageMissing)
	} else if *outputPtr == "" {
		error.Throw(error.Exceptions.ArgsOutputMissing)
	}
	return Arguments{
		Message: *messagePtr,
		Output:  *outputPtr,
		Tmp:     *tmpPtr,
	}
}
