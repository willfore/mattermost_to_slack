package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Version string
	GitCommit string
)

func MakeVersion() *cobra.Command {
	var command = &cobra.Command{
		Use: "version",
		Short: "Print the version",
		Example: ` mm2slack version`,
		Aliases: []string{"v"},
		SilenceUsage: false,
	}
	
	command.Run = func(cmd *cobra.Command, args []string) {
		if len(Version) == 0 {
			fmt.Println("Version: dev")
		} else {
			fmt.Println("Version:", Version)
		}
		fmt.Println("Git Commit:", GitCommit)
		
	}
	return command
}