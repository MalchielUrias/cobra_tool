package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cobra tool",
	Long:  `All software has versions. This is cobra_tool's version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cobra Tool v1.0 -- HEAD")
	},
}
