package settings

import (
	"fmt"
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
		fmt.Printf("Warning: can not load dot env: %v\n", err)
	}

	return Settings{
		Port:          os.Getenv("port"),
		Gmail:         os.Getenv("gmail"),
		GmailPassword: os.Getenv("gmailPassword"),
		GmailServer:   os.Getenv("gmailServer"),
		RabbitMQUrl:   os.Getenv("rabbitUrl"),
	}
}
