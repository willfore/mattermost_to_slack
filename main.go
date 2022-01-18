// Copyright (c) mattermost_to_slack author(s) 2020. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/willfore/mattermost_to_slack/cmd"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "mm2slack",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(cmd.MakeAuth())
	rootCmd.AddCommand(cmd.GetUsers())
	rootCmd.AddCommand(cmd.GetChannels())
	rootCmd.AddCommand(cmd.GetPosts())
	rootCmd.AddCommand(cmd.GetDirectPosts())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
