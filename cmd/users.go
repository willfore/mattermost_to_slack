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

func GetUsers() *cobra.Command {
	var command = &cobra.Command{
		Use:          "get_users",
		Short:        "Print the found users",
		Long:         "Print the found users from specified json export",
		Example:      ` mm2slack get_users --export-file <path_to_file>`,
		SilenceUsage: false,
	}

	command.Flags().String("export-file", "~/Downloads/bulk.json", "Provide the path to the export .json file")

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
		prompt := promptui.Select{
			Label: "Do you want to create these slack uers?",
			Items: []string{"Yes", "No"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
		if result == "Yes" {
			for _, user := range Users {
				fmt.Println(user.User.Username)
			}
		}
		return nil
	}

	return command
}
