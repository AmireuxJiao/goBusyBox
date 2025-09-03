package cmd

import "github.com/spf13/cobra"

var parCommand = &cobra.Command{
	Use:   "par [flag]",
	Short: "并发任务执行工具",
	Run:   runPar,
}

func runPar(cmd *cobra.Command, args []string) {

}

func init() {
	rootCmd.AddCommand(parCommand)
}
