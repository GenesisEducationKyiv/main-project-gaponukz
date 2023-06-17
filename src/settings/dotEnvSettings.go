package settings

import (
	"os"

	"github.com/joho/godotenv"
)

type DotEnvSettings struct{}

func (sts DotEnvSettings) Load() Settings {
	godotenv.Load()

	return Settings{
		Gmail:         os.Getenv("gmail"),
		GmailPassword: os.Getenv("gmailPassword"),
	}
}
