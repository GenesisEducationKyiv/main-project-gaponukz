package settings

import (
	"os"

	"github.com/joho/godotenv"
)

type DotEnvSettings struct{}

func NewDotEnvSettings() *DotEnvSettings {
	return &DotEnvSettings{}
}

func (sts DotEnvSettings) Load() Settings {
	godotenv.Load()

	return Settings{
		Port:          os.Getenv("port"),
		Gmail:         os.Getenv("gmail"),
		GmailPassword: os.Getenv("gmailPassword"),
	}
}
