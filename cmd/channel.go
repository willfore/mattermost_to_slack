package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/willfore/mattermost_to_slack/types"
)

func GetChannels() *cobra.Command {
	var command = &cobra.Command{
		Use:          "get_channels",
		Short:        "Print the found channels",
		Long:         "Print the found channels from specified json export",
		Example:      ` mm2slack get_channels --export-file <path_to_file> --team-name <team_name>`,
		SilenceUsage: false,
	}

	command.Flags().String("export-file", "~/Downloads/bulk.json", "Provide the path to the export .json file")
	command.Flags().String("team-name", "my-team", "Provide the team name")

	command.PreRunE = func(command *cobra.Command, args []string) error {
		_, err := command.Flags().GetString("export-file")
		if err != nil {
			return fmt.Errorf("error with --export-file usage: %s", err)
		}

		_, teamErr := command.Flags().GetString("team-name")
		if teamErr != nil {
			return fmt.Errorf("error with --team-name usage: %s", err)
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
		var Channels types.Channels

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

		fmt.Println("Found", len(Channels), "channels")
		prompt := promptui.Select{
			Label: "Would you like to import these channels into slack?",
			Items: []string{"Yes", "No"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
		}

		if result == "Yes" {
			for _, channel := range Channels {
				fmt.Printf("Adding Channel: %s - %s\n", channel.Channel.Name, channel.Channel.Type)
			}
		} else {
			fmt.Println("Exiting...")
		}
		return nil
	}

	return command
}
