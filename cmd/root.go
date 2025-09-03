package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const cliName = "goBusyBox"

var (
	logLevel string // 全局 flag：日志级别

	rootCmd = &cobra.Command{
		Use:   cliName,
		Short: "A busybox-like tool built with Cobra",
		Long:  `goBusyBox is a multi-functional tool that behaves differently based on the command name used to invoke it.`,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// 在任何一个子命令执行前，统一初始化 logrus
			level, err := logrus.ParseLevel(logLevel)
			if err != nil {
				logrus.Fatalf("invalid log level: %v", err)
			}
			logrus.SetLevel(level)
			logrus.SetFormatter(&logrus.JSONFormatter{
				TimestampFormat: "2006-01-02 15:04:05",
			})
			logrus.SetOutput(os.Stdout)
		},

		Run: func(cmd *cobra.Command, args []string) {
			// 获取程序被调用时的名称
			runCmdName := filepath.Base(os.Args[0])
			fmt.Printf("Hello from MyBox! You invoked me as '%s'\n", runCmdName)
			fmt.Println("Use one of the available commands: hello, echo, info")
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info",
		"日志级别: debug, info, warn, error, fatal, panic")
}

func Execute() {
	// 获取运行的指令名称
	runCmdName := filepath.Base(os.Args[0])

	if runCmdName != cliName {
		// 是否找到指令
		var found bool

		// 检查是否存在与调用名称相同的子命令
		for _, cmd := range rootCmd.Commands() {
			if cmd.Name() == runCmdName {
				os.Args = append([]string{cliName, runCmdName}, os.Args[1:]...)
				found = true
				break
			}
		}

		// 未找到则打印
		if !found {
			fmt.Printf("Unknown command: %s\n", runCmdName)
			os.Exit(1)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
