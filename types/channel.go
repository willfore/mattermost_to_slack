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

type SlackChannels []SlackChannel
type SlackChannel struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	IsChannel          bool          `json:"is_channel"`
	IsGroup            bool          `json:"is_group"`
	IsIm               bool          `json:"is_im"`
	Created            int           `json:"created"`
	Creator            string        `json:"creator"`
	IsArchived         bool          `json:"is_archived"`
	IsGeneral          bool          `json:"is_general"`
	Unlinked           int           `json:"unlinked"`
	NameNormalized     string        `json:"name_normalized"`
	IsShared           bool          `json:"is_shared"`
	IsExtShared        bool          `json:"is_ext_shared"`
	IsOrgShared        bool          `json:"is_org_shared"`
	PendingShared      []interface{} `json:"pending_shared"`
	IsPendingExtShared bool          `json:"is_pending_ext_shared"`
	IsMember           bool          `json:"is_member"`
	IsPrivate          bool          `json:"is_private"`
	IsMpim             bool          `json:"is_mpim"`
	LastRead           string        `json:"last_read"`
	Latest             interface{}   `json:"latest"`
	UnreadCount        int           `json:"unread_count"`
	UnreadCountDisplay int           `json:"unread_count_display"`
	Members            []string      `json:"members"`
	Topic              struct {
		Value   string `json:"value"`
		Creator string `json:"creator"`
		LastSet int    `json:"last_set"`
	} `json:"topic"`
	Purpose struct {
		Value   string `json:"value"`
		Creator string `json:"creator"`
		LastSet int    `json:"last_set"`
	} `json:"purpose"`
	PreviousNames []interface{} `json:"previous_names"`
	Priority      int           `json:"priority"`
}
