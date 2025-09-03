package cmd

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// flags variable
var (
	modifyAccessTime bool
	modifyModifyTime bool
)
var touchCommand = &cobra.Command{
	Use:   "touch [file]",
	Short: "Create Files and Modify timestamp",
	Run:   runTouch,
}

func init() {
	touchCommand.Flags().BoolVarP(&modifyAccessTime, "access", "a", false, "仅更新文件访问时间")
	touchCommand.Flags().BoolVarP(&modifyModifyTime, "modify", "m", false, "仅更新文件修改时间")
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

	// modify exist file
	info, err := os.Stat(filename)
	if err != nil {
		return err
	}

	newAtime, newMtime := t, t
	switch {
	case modifyAccessTime && !modifyModifyTime:
		newMtime = info.ModTime()
	case modifyModifyTime && !modifyAccessTime:
		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			newAtime = time.Unix(stat.Atim.Sec, stat.Atim.Nsec)
		} else {
			newAtime = info.ModTime()
		}
	}
	return os.Chtimes(filename, newAtime, newMtime)
}
