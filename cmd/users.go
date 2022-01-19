package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/willfore/mattermost_to_slack/slack"
	"github.com/willfore/mattermost_to_slack/types"
)

func GetUsers() *cobra.Command {
	var command = &cobra.Command{
		Use:          "get_users",
		Short:        "Print the found users and write to json file",
		Long:         "Print the found users from specified json export and write to json file for import into slack",
		Example:      `mm2slack get_users --export-file <path_to_file> --slack-team-id <slack_team_id>`,
		SilenceUsage: false,
	}

	command.Flags().String("export-file", "~/Downloads/bulk.json", "Provide the path to the export .json file")
	command.Flags().String("slack-team-id", "", "Provide the slack team id")
	command.MarkFlagRequired("export-file")
	command.MarkFlagRequired("slack-team-id")

	command.PreRunE = func(command *cobra.Command, args []string) error {
		_, err := command.Flags().GetString("export-file")
		if err != nil {
			return fmt.Errorf("error with --export-file usage: %s", err)
		}

		_, teamErr := command.Flags().GetString("slack-team-id")
		if teamErr != nil {
			return fmt.Errorf("error with --slack-team-id usage: %s", err)
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
		var Users types.Users

		for fileReader.Scan() {
			jsonLines = append(jsonLines, fileReader.Text())
		}

		for _, line := range jsonLines {
			var user types.User
			json.Unmarshal([]byte(line), &user)
			if user.Type == "user" {
				Users = append(Users, user)
			}
		}
		fmt.Println("Found", len(Users), "users")
		fmt.Println("These Users will be added to users.json... Now creating JSON file")
		fmt.Println("Fetching Slack Users...")

		currentSlackUsers, err := slack.FetchUsers()
		if err != nil {
			fmt.Errorf("could not fetch slack users: %s", err)
		}
		fmt.Println("Fetched", len(currentSlackUsers.Members), "slack users")

		var SlackUsers types.SlackUsers
		var SlackChannelMembers types.Users
		for _, user := range Users {
			var slackUser types.SlackUser
			var slackChannelMember types.User
			result, slackUserID := FindUser(user.User.Username, currentSlackUsers)
			if result {
				slackChannelMember.User.SlackID = slackUserID
				slackChannelMember.User.Username = user.User.Username
				slackChannelMember.User.Teams = user.User.Teams
				SlackChannelMembers = append(SlackChannelMembers, slackChannelMember)
			}
			slackUser.Name = user.User.Username
			slackUser.TeamID = command.Flag("slack-team-id").Value.String()
			slackUser.RealName = user.User.FirstName + " " + user.User.LastName
			slackUser.Profile.Email = user.User.Email
			slackUser.Profile.FirstName = user.User.FirstName
			slackUser.Profile.LastName = user.User.LastName
			slackUser.IsAdmin = false
			slackUser.IsOwner = false
			slackUser.IsBot = false
			slackUser.IsEmailConfirmed = true
			fmt.Printf("Adding user %s %s\n", user.User.Username, user.User.Email)
			SlackUsers = append(SlackUsers, slackUser)
		}

		file, _ := json.MarshalIndent(SlackUsers, "", " ")
		userFile, _ := json.MarshalIndent(SlackChannelMembers, "", " ")
		_ = ioutil.WriteFile(dirName()+"/users.json", file, 0644)
		_ = ioutil.WriteFile("mm_users.json", userFile, 0644)
		fmt.Println("Done creating user.json file")
		return nil
	}
	return command
}

func dirName() string {
	now := time.Now()
	time := now.Format("2006-01-02")
	return "./mm_export-" + time + "/"
}

func FindUser(username string, slackUserList types.SlackUserList) (result bool, slackUserID string) {
	result = false
	for _, user := range slackUserList.Members {
		if user.Name == username {
			slackUserID = user.ID
			result = true
			break
		}
	}
	return result, slackUserID
}
