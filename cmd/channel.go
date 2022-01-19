package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/willfore/mattermost_to_slack/types"
)

func GetChannels() *cobra.Command {
	var command = &cobra.Command{
		Use:          "get_channels",
		Short:        "Print the found channels",
		Long:         "Print the found channels from specified json export",
		Example:      ` mm2slack get_channels --export-file <path_to_file> --team-name <team_name> --slack-team-id <slack_team_id>`,
		SilenceUsage: false,
	}

	command.Flags().String("export-file", "~/Downloads/bulk.json", "Provide the path to the export .json file")
	command.MarkFlagRequired("export-file")
	command.Flags().String("team-name", "my-team", "Provide the mattermost team name")
	command.MarkFlagRequired("team-name")
	command.Flags().String("slack-team-id", "", "Provide the slack team id")
	command.MarkFlagRequired("slack-team-id")

	command.PreRunE = func(command *cobra.Command, args []string) error {
		_, err := command.Flags().GetString("export-file")
		if err != nil {
			return fmt.Errorf("error with --export-file usage: %s", err)
		}

		_, teamErr := command.Flags().GetString("team-name")
		if teamErr != nil {
			return fmt.Errorf("error with --team-name usage: %s", err)
		}

		_, teamId := command.Flags().GetString("slack-team-id")
		if teamId != nil {
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

		mmUsersFile, err := os.Open("mm_users.json")
		if err != nil {
			fmt.Errorf("could not open mm_users.json file %s", err)
		}

		defer mmUsersFile.Close()
		defer jsonFile.Close()

		fileReader := bufio.NewScanner(jsonFile)
		fileReader.Split(bufio.ScanLines)

		var jsonLines []string
		var Channels types.Channels
		var mmUsers types.Users

		for fileReader.Scan() {
			jsonLines = append(jsonLines, fileReader.Text())
		}

		for _, line := range jsonLines {
			var channel types.Channel
			teamName, _ := command.Flags().GetString("team-name")
			json.Unmarshal([]byte(line), &channel)
			if channel.Type == "channel" && channel.Channel.Team == teamName {
				Channels = append(Channels, channel)
			}
		}

		byteValue, _ := ioutil.ReadAll(mmUsersFile)
		json.Unmarshal(byteValue, &mmUsers)

		fmt.Println("Found", len(Channels), "channels")

		var SlackChannels types.SlackChannels
		for _, channel := range Channels {
			var slackChannel types.SlackChannel
			result, slackUserIDs := matchChannels(channel.Channel.Name, mmUsers)
			if result {
				slackChannel.Members = append(slackChannel.Members, slackUserIDs...)
			}
			slackChannel.Name = channel.Channel.Name
			if channel.Channel.Type == "O" {
				slackChannel.IsPrivate = false
			} else {
				slackChannel.IsPrivate = true
			}
			slackChannel.IsChannel = true
			slackChannel.Topic.Value = channel.Channel.Header
			slackChannel.Purpose.Value = channel.Channel.Purpose
			fmt.Printf("Adding Channel: %s - %s\n", channel.Channel.Name, channel.Channel.Type)
			SlackChannels = append(SlackChannels, slackChannel)
		}
		file, _ := json.MarshalIndent(SlackChannels, "", " ")
		_ = ioutil.WriteFile(dirName()+"/channels.json", file, 0644)
		fmt.Println("Done creating channels.json files")
		return nil
	}

	return command
}

func matchChannels(channelName string, mmUsers types.Users) (result bool, slackUserID []string) {
	result = false
	var slackUserIDs []string
	for _, user := range mmUsers {
		for _, team := range user.User.Teams {
			if team.Name == "irt" {
				for _, channel := range team.Channels {
					if channel.Name == channelName {
						slackUserIDs = append(slackUserIDs, user.User.SlackID)
						result = true
					}
				}
			}
		}
	}
	return result, slackUserIDs
}
