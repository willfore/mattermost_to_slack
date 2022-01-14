package types

type Teams []Team

type Team struct {
	Type string `json:"type"`
	Team struct {
		Name            string `json:"name"`
		DisplayName     string `json:"display_name"`
		Type            string `json:"type"`
		Description     string `json:"description"`
		AllowOpenInvite bool   `json:"allow_open_invite"`
	} `json:"team"`
}
