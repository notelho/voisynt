package main

import (
	"fmt"

	"github.com/enbot/voisynt/cli"
	"github.com/enbot/voisynt/error"
	"github.com/enbot/voisynt/io"
)

func main() {

	var args = cli.ArgumentsCreate()
	var message = args.Message
	var path = args.Path

	fmt.Println("message: " + args.Message)
	fmt.Println("path: " + args.Path)

	if io.FileExists(path) {

		var name = io.AudioName(message)

		if !io.FileExists(path + name) {

			fmt.Println("file not exists")

			// 	file = audio.AudioDownload(message)
			// 	audio = audio.AudioDistort(file)
			// 	io.FileCreate(audio)

		}

		fmt.Print(io.AudioPath(path, name))

	} else {
		error.ThrowExit("Provided path not found", 1)
	}

}
