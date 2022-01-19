package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/willfore/mattermost_to_slack/types"
)

func GetPosts() *cobra.Command {
	var command = &cobra.Command{
		Use:          "get_posts",
		Short:        "Print the found posts",
		Long:         "Print the found posts from specified json export",
		Example:      ` mm2slack get_posts --export-file <path_to_file> --team-name <team_name> --slack-team-id <team_id> --export-from <unix_timestamp_milliseconds>`,
		SilenceUsage: false,
	}

	command.Flags().String("export-file", "~/Downloads/bulk.json", "Provide the path to the export .json file")
	command.MarkFlagRequired("export-file")
	command.Flags().String("team-name", "my-team", "Provide the team name")
	command.MarkFlagRequired("team-name")
	command.Flags().String("slack-team-id", "my_team_id", "Provide the team id")
	command.MarkFlagRequired("slack-team-id")
	command.Flags().String("export-from", "1642634522645", "Provide the unix timestamp of the earliest post to import")
	command.MarkFlagRequired("export-from")

	command.PreRunE = func(command *cobra.Command, args []string) error {
		return nil
	}

	command.RunE = func(cmd *cobra.Command, args []string) error {
		fmt.Println("This will take a while... Grab a coffee or something")
		exportFile, _ := command.Flags().GetString("export-file")
		jsonFile, err := os.Open(exportFile)
		if err != nil {
			fmt.Errorf("could not open json file %s", err)
		}

		defer jsonFile.Close()
		fileReader := bufio.NewScanner(jsonFile)
		fileReader.Split(bufio.ScanLines)

		var jsonLines []string
		var ChannelPosts types.ChannelPosts

		for fileReader.Scan() {
			jsonLines = append(jsonLines, fileReader.Text())
		}

		for _, line := range jsonLines {
			var post types.Post
			teamName, _ := command.Flags().GetString("team-name")
			json.Unmarshal([]byte(line), &post)
			if post.Type == "post" && post.Post.Team == teamName {
				var channelPost types.ChannelPost
				json.Unmarshal([]byte(line), &channelPost)
				ChannelPosts = append(ChannelPosts, channelPost)
			}
		}

		sort.SliceStable(ChannelPosts, func(a, b int) bool {
			return ChannelPosts[a].Post.CreateAt < ChannelPosts[b].Post.CreateAt
		})

		mmUsersFile, err := os.Open("mm_users.json")
		if err != nil {
			fmt.Errorf("could not open mm_users.json file %s", err)
		}
		var mmUsers types.Users
		defer mmUsersFile.Close()
		byteValue, _ := ioutil.ReadAll(mmUsersFile)
		json.Unmarshal(byteValue, &mmUsers)

		fmt.Println("Found", len(ChannelPosts), "channel posts")
		var slackChannelPosts types.SlackChannelPosts
		for _, channelPost := range ChannelPosts {
			exportFrom, _ := command.Flags().GetString("export-from")
			exportFromInt, _ := strconv.ParseInt(exportFrom, 10, 64)
			if channelPost.Post.CreateAt >= exportFromInt {
				result, slackID := FindPostUser(channelPost.Post.User, mmUsers)
				teamId, _ := command.Flags().GetString("slack-team-id")
				var slackChannelPost types.SlackChannelPost
				os.Mkdir(dirName()+"/"+channelPost.Post.Channel, 0777)
				slackChannelPost.Type = "message"
				slackChannelPost.Text = channelPost.Post.Message
				slackChannelPost.Ts = strconv.FormatInt(int64(channelPost.Post.CreateAt), 10)
				slackChannelPost.Team = teamId
				slackChannelPost.UserTeam = teamId
				slackChannelPost.SourceTeam = teamId
				if result {
					slackChannelPost.User = slackID
				}
				slackChannelPosts = append(slackChannelPosts, slackChannelPost)
				file, _ := json.MarshalIndent(slackChannelPosts, "", " ")
				_ = ioutil.WriteFile(dirName()+"/"+channelPost.Post.Channel+"/posts.json", file, 0644)
				fmt.Printf("Adding Channel Post from: %s -> %s\n", channelPost.Post.User, channelPost.Post.Channel)
			} else {
				fmt.Println("Skipping Channel Post from:", channelPost.Post.CreateAt)
			}
		}
		fmt.Println("Done creating poasts.json file")

		return nil
	}
	return command
}

func GetDirectPosts() *cobra.Command {
	var command = &cobra.Command{
		Use:          "get_direct_posts",
		Short:        "Print the found direct posts",
		Long:         "Print the found direct posts from specified json export",
		Example:      ` mm2slack get_direct_posts --export-file <path_to_file>`,
		SilenceUsage: false,
	}

	command.Flags().String("export-file", "~/Downloads/bulk.json", "Provide the path to the export .json file")
	command.MarkFlagRequired("export-file")

	command.PreRunE = func(command *cobra.Command, args []string) error {
		_, err := command.Flags().GetString("export-file")
		if err != nil {
			return fmt.Errorf("error with --export-file usage: %s", err)
		}

		return nil
	}
	command.RunE = func(cmd *cobra.Command, args []string) error {
		exportFile, _ := command.Flags().GetString("export-file")
		jsonFile, err := os.Open(exportFile)
		if err != nil {
			fmt.Errorf("could not open json file %s", err)
		}

		defer jsonFile.Close()
		fileReader := bufio.NewScanner(jsonFile)
		fileReader.Split(bufio.ScanLines)

		var jsonLines []string
		var DirectPosts types.DirectPosts

		for fileReader.Scan() {
			jsonLines = append(jsonLines, fileReader.Text())
		}

		for _, line := range jsonLines {
			var post types.Post
			json.Unmarshal([]byte(line), &post)
			if post.Type == "direct_post" {
				var directPost types.DirectPost
				json.Unmarshal([]byte(line), &directPost)
				DirectPosts = append(DirectPosts, directPost)
			}
		}

		sort.SliceStable(DirectPosts, func(a, b int) bool {
			return DirectPosts[a].DirectPost.CreateAt < DirectPosts[b].DirectPost.CreateAt
		})

		fmt.Println("Found", len(DirectPosts), "direct posts")
		prompt := promptui.Select{
			Label: "Would you like to import these direct posts into slack?",
			Items: []string{"Yes", "No"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
		}

		if result == "Yes" {
			for _, directPost := range DirectPosts {
				fmt.Printf("Adding Direct Post From: %s -> %s\n", directPost.DirectPost.User, directPost.DirectPost.ChannelMembers[len(directPost.DirectPost.ChannelMembers)-1])
			}
		} else {
			fmt.Println("Exiting...")
		}
		return nil
	}
	return command
}

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
