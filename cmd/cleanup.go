package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func DoCleanup() *cobra.Command {
	var command = &cobra.Command{
		Use:          "cleanup",
		Short:        "Removes all files created by the program",
		Long:         "Removes all files created by the program",
		Example:      `mm2slack cleanup`,
		SilenceUsage: false,
	}

	command.RunE = func(cmd *cobra.Command, args []string) error {
		fmt.Println("Removing all files created by the program")
		os.RemoveAll(dirName())
		os.Remove("mm_users.json")
		os.Remove(dirName() + ".zip")
		os.RemoveAll("mm_export-*")
		return nil
	}
	return command
}
