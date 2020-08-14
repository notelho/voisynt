package io

type DirType = string

type AudioDir struct {
	Output DirType
	Tmp    DirType
}

func CreateAudioDir(outputDir string, tmpDir string) AudioDir {
	return AudioDir{
		Output: outputDir + "\\",
		Tmp:    tmpDir + "\\",
	}
}
