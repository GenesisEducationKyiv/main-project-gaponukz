package settings

type Settings struct {
	Port          string `json:"port"`
	Gmail         string `json:"gmail"`
	GmailPassword string `json:"gmailPassword"`
}

type ISettingsExporter interface {
	load() Settings
}
