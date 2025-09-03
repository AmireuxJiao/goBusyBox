package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var parCommand = &cobra.Command{
	Use:   "par",
	Short: "并发任务执行工具",
	Run:   runPar,
}

var (
	concurrencyNum int8
	retriesNum     int8
	filePath       string
	dependFlow     string
)

func runPar(cmd *cobra.Command, args []string) {
	logrus.Debugf("并发数：%d, 重试次数：%d\n", concurrencyNum, retriesNum)
	logrus.Debugf("读取文件的名称: %s\n", filePath)
	logrus.Debugf("依赖顺序: %s\n", dependFlow)
}

func init() {
	parCommand.Flags().BoolP("test", "t", false, "并发访问测试")
	parCommand.Flags().Int8VarP(&retriesNum, "retries", "r", 3, "重试次数")
	parCommand.Flags().Int8VarP(&concurrencyNum, "concurrent", "c", 5, "并发数")
	parCommand.Flags().StringVarP(&filePath, "file", "f", "", "从文件中读取指令")
	parCommand.Flags().StringVarP(&dependFlow, "depend", "d", "", "任务的依赖关系（格式：\"A->B,C;D->E\")")

	rootCmd.AddCommand(parCommand)
}
