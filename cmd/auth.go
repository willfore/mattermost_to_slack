package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func MakeAuth() *cobra.Command {
	var command = &cobra.Command{
		Use:          "set_auth",
		Short:        "Set the auth token",
		Long:         "set the auth token for slack api",
		Example:      ` mm2slack set_auth --auth-token <token>`,
		SilenceUsage: false,
	}

	command.Flags().String("auth-token", "123456", "Provide the slack auth token")

	command.PreRunE = func(command *cobra.Command, args []string) error {
		_, err := command.Flags().GetString("auth-token")
		if err != nil {
			return fmt.Errorf("error with --auth-token usage: %s", err)
		}
		return nil
	}

	command.RunE = func(cmd *cobra.Command, args []string) error {
		authToken, _ := command.Flags().GetString("auth-token")
		os.Setenv("SLACK_AUTH_TOKEN", authToken)
		fmt.Println("Auth token set")
		return nil
	}
	return command
}
