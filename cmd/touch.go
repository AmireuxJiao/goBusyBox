package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var touchCommand = &cobra.Command{
	Use:   "touch [file]",
	Short: "Create Files and Modify timestamp",
	Run:   runTouch,
}

func init() {
	rootCmd.AddCommand(touchCommand)
}

func runTouch(cmd *cobra.Command, args []string) {
	now := time.Now()

	for _, filename := range args {
		if err := touchFile(filename, now); err != nil {
			fmt.Printf("touch %s 失败 %s", filename, err)
		}
	}
}

func touchFile(filename string, t time.Time) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)

	if err != nil {
		// file not exist
		if os.IsNotExist(err) {
			file, err := os.Create(filename)
			if err != nil {
				return err
			}
			defer file.Close()
			return nil
		} else {
			// other error
			return err
		}
	}
	defer file.Close()

	return os.Chtimes(filename, t, t) // 更新访问时间和修改时间
}
