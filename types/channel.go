package types

type Channels []Channel

type Channel struct {
	Type    string `json:"type"`
	Channel struct {
		Team        string `json:"team"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Type        string `json:"type"`
		Header      string `json:"header"`
		Purpose     string `json:"purpose"`
	} `json:"channel"`
}
