package version_command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Brazil Gov Data",
	Long:  `Print the version number of Brazil Gov Data`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Brazil Gov Data v0.0.0")
	},
}
