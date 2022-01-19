package types

type Post struct {
	Type string `json:"type"`
	Post struct {
		Team string `json:"team"`
	} `json:"post"`
}

type ChannelPosts []ChannelPost
type ChannelPost struct {
	Type string `json:"type"`
	Post struct {
		Team    string `json:"team"`
		Channel string `json:"channel"`
		User    string `json:"user"`
		Message string `json:"message"`
		Props   struct {
		} `json:"props"`
		CreateAt  int64         `json:"create_at"`
		Reactions []interface{} `json:"reactions"`
		Replies   interface{}   `json:"replies"`
	} `json:"post"`
}

type DirectPosts []DirectPost
type DirectPost struct {
	Type       string `json:"type"`
	DirectPost struct {
		ChannelMembers []string `json:"channel_members"`
		User           string   `json:"user"`
		Message        string   `json:"message"`
		Props          struct {
		} `json:"props"`
		CreateAt    int64       `json:"create_at"`
		FlaggedBy   interface{} `json:"flagged_by"`
		Reactions   interface{} `json:"reactions"`
		Replies     interface{} `json:"replies"`
		Attachments interface{} `json:"attachments"`
	} `json:"direct_post"`
}

type SlackChannelPosts []SlackChannelPost
type SlackChannelPost struct {
	ClientMsgID string `json:"client_msg_id"`
	Type        string `json:"type"`
	Text        string `json:"text"`
	User        string `json:"user"`
	Ts          string `json:"ts"`
	Team        string `json:"team"`
	UserTeam    string `json:"user_team"`
	SourceTeam  string `json:"source_team"`
	UserProfile struct {
		AvatarHash        string `json:"avatar_hash"`
		Image72           string `json:"image_72"`
		FirstName         string `json:"first_name"`
		RealName          string `json:"real_name"`
		DisplayName       string `json:"display_name"`
		Team              string `json:"team"`
		Name              string `json:"name"`
		IsRestricted      bool   `json:"is_restricted"`
		IsUltraRestricted bool   `json:"is_ultra_restricted"`
	} `json:"user_profile"`
	Blocks []struct {
		Type     string `json:"type"`
		BlockID  string `json:"block_id"`
		Elements []struct {
			Type     string `json:"type"`
			Elements []struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"elements"`
		} `json:"elements"`
	} `json:"blocks"`
}
