package main

import (
	"fmt"

	errors "github.com/enbot/voisynt/errors"
	helper "github.com/enbot/voisynt/helper"
	language "github.com/enbot/voisynt/language"
)

func main() {

	var args = helper.ArgumentsCreate()
	var message = args.Message
	var path = args.Path

	fmt.Print("message: " + args.Message)
	fmt.Print("path: " + args.Path)

	if helper.PathExists(path) {

		var name = language.FixName(message)

		if !helper.PathExists(path + name) {
			file = audio.fetch(message)
			audio = audio.distortion(file)
			helper.FileCreate(audio)
		}

		fmt.Print(path + name)

	} else {
		errors.ThrowExit(errors.ErrorPathNotFound)
	}

}
