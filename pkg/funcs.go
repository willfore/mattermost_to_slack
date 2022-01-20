package pkg

import (
	"sort"

	"github.com/willfore/mattermost_to_slack/types"
)

func FindPostUser(username string, mmUsers types.Users) (result bool, slackUserID string) {
	result = false
	for _, user := range mmUsers {
		if user.User.Username == username {
			slackUserID = user.User.SlackID
			result = true
			break
		}
	}
	return result, slackUserID
}

func FindPostUserEmail(username string, mmUsers types.Users) (email string) {
	for _, user := range mmUsers {
		if user.User.Username == username {
			email = user.User.Email
			break
		}
	}
	return email
}

func CheckIgnoredChannels(channel string, ignoredChannels []string) bool {
	for _, ignoredChannel := range ignoredChannels {
		if channel == ignoredChannel {
			return true
		}
	}
	return false
}

func FindChannelPosts(channel string, mmPosts types.ChannelPosts) types.ChannelPosts {
	var channelPosts types.ChannelPosts
	for _, post := range mmPosts {
		if post.Post.Channel == channel {
			channelPosts = append(channelPosts, post)
		}
	}
	sort.SliceStable(channelPosts, func(a, b int) bool {
		return channelPosts[a].Post.CreateAt < channelPosts[b].Post.CreateAt
	})
	return channelPosts
}
