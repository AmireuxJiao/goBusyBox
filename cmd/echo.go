package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var echoCommand = &cobra.Command{
	Use:   "echo",
	Short: "Echo back input",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(echoCommand)
}
