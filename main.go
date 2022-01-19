// Copyright (c) mattermost_to_slack author(s) 2020. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package main

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/willfore/mattermost_to_slack/cmd"
)

func main() {
	if _, err := os.Stat(dirName()); os.IsNotExist(err) {
		os.Mkdir(dirName(), 0777)
	}

	var rootCmd = &cobra.Command{
		Use: "mm2slack",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(cmd.MakeUserCsv())
	rootCmd.AddCommand(cmd.GetUsers())
	rootCmd.AddCommand(cmd.GetChannels())
	rootCmd.AddCommand(cmd.GetPosts())
	rootCmd.AddCommand(cmd.GetDirectPosts())
	rootCmd.AddCommand(cmd.DoCleanup())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func dirName() string {
	now := time.Now()
	time := now.Format("2006-01-02")
	return "./mm_export-" + time + "/"
}
