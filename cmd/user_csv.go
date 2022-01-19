package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/tushar2708/altcsv"
	"github.com/willfore/mattermost_to_slack/types"
)

func MakeUserCsv() *cobra.Command {
	now := time.Now()
	var command = &cobra.Command{
		Use:          "make_user_csv",
		Short:        "Print the found users and write to CSV file",
		Long:         "Print the found users from specified json export and write to CSV file for import into slack",
		Example:      ` mm2slack make_user_csv --export-file <path_to_file>`,
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
		fmt.Println("These Users will have to be imported via CSV. Now creating CSV file")

		csvFile, err := os.Create("users.csv")
		if err != nil {
			log.Fatalf("failed creating csv file: %s", err)
		}
		defer csvFile.Close()
		csvWriter := altcsv.NewWriter(csvFile)
		csvWriter.AllQuotes = true

		for _, user := range Users {
			time := strconv.FormatInt(int64(now.UnixMilli()), 10)
			csvWriter.Write([]string{time, "random", user.User.Username, "Adding a user"})
			fmt.Printf("Adding user %s\n", user.User.Username)
		}
		csvWriter.Flush()
		fmt.Println("Done creating CSV file")
		return nil
	}
	// here we will need to call a func to do authentication, check users and create them if needed. Need to ask what type of user we want to create.
	return command
}
