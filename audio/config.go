package audio

type ConfigType = string

type AudioConfig struct {
	Message  ConfigType
	Language ConfigType
}

func CreateAudioConfig(message string, language string) AudioConfig {
	return AudioConfig{
		Message:  message,
		Language: language,
	}
}
