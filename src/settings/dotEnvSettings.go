package settings

import (
	"os"

	"github.com/joho/godotenv"
)

type dotEnvSettings struct{}

func NewDotEnvSettings() dotEnvSettings {
	return dotEnvSettings{}
}

func (sts dotEnvSettings) Load() Settings {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	return Settings{
		Port:          os.Getenv("port"),
		Gmail:         os.Getenv("gmail"),
		GmailPassword: os.Getenv("gmailPassword"),
		GmailServer:   os.Getenv("gmailServer"),
	}
}
