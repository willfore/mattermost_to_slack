package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tushar2708/altcsv"
	"github.com/willfore/mattermost_to_slack/pkg"
	"github.com/willfore/mattermost_to_slack/types"
)

func GetPostsCSV() *cobra.Command {
	var command = &cobra.Command{
		Use:          "csv_posts",
		Short:        "Print the found posts",
		Long:         "Print the found posts from specified json export",
		Example:      ` mm2slack csv_posts --export-file <path_to_file> --team-name <team_name> --slack-team-id <team_id> --export-from <unix_timestamp_milliseconds> --ignore-channels <comma_separated_channel_names>`,
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
	command.Flags().StringSlice("ignore-channels", []string{"test-channel-1", "test-channel-2"}, "Provide the comma separated list of channel names to ignore")

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

		channelsFile, err := os.Open(dirName() + "/channels.json")
		if err != nil {
			fmt.Errorf("could not open channels.json file %s", err)
		}
		var slackChannels types.SlackChannels
		defer channelsFile.Close()
		byteValue, _ = ioutil.ReadAll(channelsFile)
		json.Unmarshal(byteValue, &slackChannels)

		for _, slackChannel := range slackChannels {
			exportFrom, _ := command.Flags().GetString("export-from")
			exportFromInt, _ := strconv.ParseInt(exportFrom, 10, 64)
			os.Mkdir(dirName()+"/"+slackChannel.Name, 0777)
			csvFile, err := os.Create(dirName() + "/" + slackChannel.Name + "/" + slackChannel.Name + ".csv")
			if err != nil {
				fmt.Errorf("could not create csv file %s", err)
			}
			writer := altcsv.NewWriter(csvFile)
			writer.AllQuotes = true
			mmPosts := pkg.FindChannelPosts(slackChannel.Name, ChannelPosts)
			for _, post := range mmPosts {
				if post.Post.CreateAt >= exportFromInt {
					var row []string
					row = append(row, strconv.FormatInt(int64(post.Post.CreateAt/1000), 10))
					row = append(row, post.Post.Channel)
					row = append(row, post.Post.User)
					row = append(row, post.Post.Message)
					writer.Write(row)
				} else {
					fmt.Println("Skipping Channel Post from:", post.Post.CreateAt)
				}
			}
			writer.Flush()
			csvFile.Close()
			fmt.Println("Channel: " + slackChannel.Name)
			fmt.Println("Wrote " + fmt.Sprintf("%d", len(mmPosts)) + " posts to " + dirName() + "/" + slackChannel.Name + "/" + slackChannel.Name + ".csv")
		}
		return nil
	}
	return command
}
