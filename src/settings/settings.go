package settings

type Settings struct {
	Port          string `json:"port"`
	Gmail         string `json:"gmail"`
	GmailPassword string `json:"gmailPassword"`
	GmailServer   string `json:"gmailServer"`
	RabbitMQUrl   string `json:"rabbitUrl"`
}
