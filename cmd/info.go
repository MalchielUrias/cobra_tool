package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print information about Cobra tool",
	Long:  `You may want to know more about cobra tool and what it does, this would help`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cobra is a tool that was build by the Hugo team at Golang to help you create cli applications")
	},
}
