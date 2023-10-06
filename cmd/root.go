package cmd

import (
	"fmt"
	"gov-data/cmd/version_command"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gov-data",
	Short: "Brazil Gov Data is an application to retrieve, parse, structure and store open government data.",
	Long:  "Brazil Gov Data is an application to retrieve, parse, structure and store open government data.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.AddCommand(version_command.VersionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
