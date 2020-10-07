package cmd

import (
	"fmt"

	"github.com/ivanklee86/homelab-ddns/constants"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of homelab-ddns.",
	Long:  `Just the semantic version of the package!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", constants.Version)
	},
}
