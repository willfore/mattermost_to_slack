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
