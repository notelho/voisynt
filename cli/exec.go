package cli

import (
	"os/exec"

	"github.com/notelho/voisynt/error"
)

func ffmpeg(args ...string) {
	err := exec.Command("ffmpeg", args...).Run()
	if err != nil {
		error.Throw(error.Exceptions.ExecRunFailed)
	}
}
