package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const CliName = "goBusyBox"

var (
	rootCmd = &cobra.Command{
		Use:   CliName,
		Short: "A busybox-like tool built with Cobra",
		Long:  `goBusyBox is a multi-functional tool that behaves differently based on the command name used to invoke it.`,
		Run: func(cmd *cobra.Command, args []string) {
			// 获取程序被调用时的名称
			runCmdName := filepath.Base(os.Args[0])
			fmt.Printf("Hello from MyBox! You invoked me as '%s'\n", runCmdName)
			fmt.Println("Use one of the available commands: hello, echo, info")
		},
	}
)

func Execute() {
	// 获取运行的指令名称
	runCmdName := filepath.Base(os.Args[0])

	if runCmdName != CliName {
		// 是否找到指令
		var found bool

		// 检查是否存在与调用名称相同的子命令
		for _, cmd := range rootCmd.Commands() {
			if cmd.Name() == runCmdName {
				os.Args = append([]string{CliName, runCmdName}, os.Args[1:]...)
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
