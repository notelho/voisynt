package main

import (
	"fmt"

	errors "github.com/enbot/voisynt/errors"
	helpers "github.com/enbot/voisynt/helpers"
)

func main() {

	var args = helpers.ArgumentsCreate()
	var message = args.Message
	var path = args.Path

	fmt.Print("message: " + args.Message)
	fmt.Print("path: " + args.Path)

	if helpers.PathExists(path) {

		var name = helpers.FixName(message)

		if !helpers.PathExists(path + name) {
			file = audio.fetch(message)
			audio = audio.distortion(file)
			helpers.FileCreate(audio)
		}

		fmt.Print(path + name)

	} else {
		errors.ThrowExit("Provided path not found", 1)
	}

}
