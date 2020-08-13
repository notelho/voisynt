package cli

import (
	"os/exec"
	"strings"

	"github.com/enbot/voisynt/error"
)

func ffmpeg(args ...string) {
	err := exec.Command("ffmpeg", args...).Run()
	if err != nil {
		error.ThrowExit("Failed to run a ffmpeg command: "+strings.Join(args, " "), 1)
	}
}
