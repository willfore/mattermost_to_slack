package types

type Users []User

type User struct {
	Type string `json:"type"`
	User struct {
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
