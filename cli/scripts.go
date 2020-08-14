package cli

import "strconv"

func CopyFromStreaming(link string, output string) {
	ffmpeg(
		"-i",
		link,
		"-c",
		"copy",
		output,
	)
}

func SynthThinVoice(input string, output string) {
	ffmpeg(
		"-i",
		input,
		"-af",
		"asetrate=44100*1.07,atempo=0.5",
		output,
	)
}

func SynthThickVoice(input string, output string) {
	ffmpeg(
		"-i",
		input,
		"-af",
		"asetrate=44100*0.25,aresample=44100,atempo=2.17",
		output,
	)
}

func SynthRobotVoice(input string, output string) {
	ffmpeg(
		"-i",
		input,
		"-filter_complex",
		"afftfilt=real='hypot(re,im)*sin(0)':imag='hypot(re,im)*cos(0)':win_size=512:overlap=0.75",
		output,
	)
}

func MergeVoiceFiles(file1 string, volume1 float64, file2 string, volume2 float64, output string) {
	ffmpeg(
		"-i",
		file1,
		"-i",
		file2,
		"-filter_complex",
		"[0:a]volume="+strconv.FormatFloat(volume1, 'f', 2, 64)+"[a1];[1:a]volume="+strconv.FormatFloat(volume2, 'f', 2, 64)+"[a2];[a1][a2]amerge",
		"-c:v",
		"copy",
		"-shortest",
		output,
	)
}
