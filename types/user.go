package types

type Users []User

type User struct {
	Type string `json:"type"`
	User struct {
		SlackID     string
		Username    string `json:"username"`
		Email       string `json:"email"`
		AuthService string `json:"auth_service"`
		AuthData    string `json:"auth_data"`
		Nickname    string `json:"nickname"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Position    string `json:"position"`
		Roles       string `json:"roles"`
		Locale      string `json:"locale"`
		DeleteAt    int    `json:"delete_at"`
		Teams       []struct {
			Name     string `json:"name"`
			Roles    string `json:"roles"`
			Channels []struct {
				Name        string `json:"name"`
				Roles       string `json:"roles"`
				NotifyProps struct {
					Desktop    string `json:"desktop"`
					Mobile     string `json:"mobile"`
					MarkUnread string `json:"mark_unread"`
				} `json:"notify_props"`
				Favorite bool `json:"favorite"`
			} `json:"channels"`
		} `json:"teams"`
		TutorialStep string `json:"tutorial_step"`
		NotifyProps  struct {
			Desktop          string `json:"desktop"`
			DesktopSound     string `json:"desktop_sound"`
			Email            string `json:"email"`
			Mobile           string `json:"mobile"`
			MobilePushStatus string `json:"mobile_push_status"`
			Channel          string `json:"channel"`
			Comments         string `json:"comments"`
			MentionKeys      string `json:"mention_keys"`
		} `json:"notify_props"`
	} `json:"user"`
}

type SlackUsers []SlackUser
type SlackUser struct {
	ID       string `json:"id"`
	TeamID   string `json:"team_id"`
	Name     string `json:"name"`
	Deleted  bool   `json:"deleted"`
	Color    string `json:"color"`
	RealName string `json:"real_name"`
	Tz       string `json:"tz"`
	TzLabel  string `json:"tz_label"`
	TzOffset int    `json:"tz_offset"`
	Profile  struct {
		Title                  string        `json:"title"`
		Phone                  string        `json:"phone"`
		Skype                  string        `json:"skype"`
		RealName               string        `json:"real_name"`
		RealNameNormalized     string        `json:"real_name_normalized"`
		DisplayName            string        `json:"display_name"`
		DisplayNameNormalized  string        `json:"display_name_normalized"`
		Fields                 interface{}   `json:"fields"`
		StatusText             string        `json:"status_text"`
		StatusEmoji            string        `json:"status_emoji"`
		StatusEmojiDisplayInfo []interface{} `json:"status_emoji_display_info"`
		StatusExpiration       int           `json:"status_expiration"`
		AvatarHash             string        `json:"avatar_hash"`
		Email                  string        `json:"email"`
		FirstName              string        `json:"first_name"`
		LastName               string        `json:"last_name"`
		Image24                string        `json:"image_24"`
		Image32                string        `json:"image_32"`
		Image48                string        `json:"image_48"`
		Image72                string        `json:"image_72"`
		Image192               string        `json:"image_192"`
		Image512               string        `json:"image_512"`
		StatusTextCanonical    string        `json:"status_text_canonical"`
		Team                   string        `json:"team"`
	} `json:"profile"`
	IsAdmin                bool   `json:"is_admin"`
	IsOwner                bool   `json:"is_owner"`
	IsPrimaryOwner         bool   `json:"is_primary_owner"`
	IsRestricted           bool   `json:"is_restricted"`
	IsUltraRestricted      bool   `json:"is_ultra_restricted"`
	IsBot                  bool   `json:"is_bot"`
	IsAppUser              bool   `json:"is_app_user"`
	Updated                int    `json:"updated"`
	IsEmailConfirmed       bool   `json:"is_email_confirmed"`
	WhoCanShareContactCard string `json:"who_can_share_contact_card"`
}

type SlackUserList struct {
	Ok      bool `json:"ok"`
	Members []struct {
		ID       string `json:"id"`
		TeamID   string `json:"team_id"`
		Name     string `json:"name"`
		Deleted  bool   `json:"deleted"`
		Color    string `json:"color"`
		RealName string `json:"real_name"`
		Tz       string `json:"tz"`
		TzLabel  string `json:"tz_label"`
		TzOffset int    `json:"tz_offset"`
		Profile  struct {
			Title                  string        `json:"title"`
			Phone                  string        `json:"phone"`
			Skype                  string        `json:"skype"`
			RealName               string        `json:"real_name"`
			RealNameNormalized     string        `json:"real_name_normalized"`
			DisplayName            string        `json:"display_name"`
			DisplayNameNormalized  string        `json:"display_name_normalized"`
			Fields                 interface{}   `json:"fields"`
			StatusText             string        `json:"status_text"`
			StatusEmoji            string        `json:"status_emoji"`
			StatusEmojiDisplayInfo []interface{} `json:"status_emoji_display_info"`
			StatusExpiration       int           `json:"status_expiration"`
			AvatarHash             string        `json:"avatar_hash"`
			AlwaysActive           bool          `json:"always_active"`
			FirstName              string        `json:"first_name"`
			LastName               string        `json:"last_name"`
			Image24                string        `json:"image_24"`
			Image32                string        `json:"image_32"`
			Image48                string        `json:"image_48"`
			Image72                string        `json:"image_72"`
			Image192               string        `json:"image_192"`
			Image512               string        `json:"image_512"`
			StatusTextCanonical    string        `json:"status_text_canonical"`
			Team                   string        `json:"team"`
		} `json:"profile"`
		IsAdmin                bool   `json:"is_admin"`
		IsOwner                bool   `json:"is_owner"`
		IsPrimaryOwner         bool   `json:"is_primary_owner"`
		IsRestricted           bool   `json:"is_restricted"`
		IsUltraRestricted      bool   `json:"is_ultra_restricted"`
		IsBot                  bool   `json:"is_bot"`
		IsAppUser              bool   `json:"is_app_user"`
		Updated                int    `json:"updated"`
		IsEmailConfirmed       bool   `json:"is_email_confirmed"`
		WhoCanShareContactCard string `json:"who_can_share_contact_card"`
	} `json:"members"`
}
