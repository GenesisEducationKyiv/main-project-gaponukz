package settings

type Settings struct {
	Gmail         string `json:"gmail"`
	GmailPassword string `json:"gmailPassword"`
}

type ISettingsExporter interface {
	load() Settings
}
