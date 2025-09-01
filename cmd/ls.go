package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var lsCommand = &cobra.Command{
	Use:   "ls [path]",
	Short: "List directory contents",
	Args:  cobra.MaximumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}
		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ls: %v\n", err)
			os.Exit(1)
		}
		for _, f := range files {
			fmt.Println(f.Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCommand)
}
