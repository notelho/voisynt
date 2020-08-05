package main

import (
	"fmt"

	"github.com/enbot/voisynt/helpers"
)

func main() {

	var args = helpers.ArgumentsCreate()
	var message = args.Message
	var path = args.Path

	fmt.Println("message: " + args.Message)
	fmt.Println("path: " + args.Path)

	if helpers.FileExists(path) {

		var name = helpers.FileName(message)

		if !helpers.FileExists(path + name) {

			fmt.Println("file not exists")

			// 	file = audio.fetch(message)
			// 	audio = audio.distortion(file)
			// 	helpers.FileCreate(audio)
		}

		// fmt.Print(path + name)

	} else {
		helpers.ThrowExit("Provided path not found", 1)
	}

}
