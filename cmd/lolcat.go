package cmd

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/spf13/cobra"
)

var lolcatCommand = &cobra.Command{
	Use:   "lolcat",
	Short: "like-lolcat: rainbow output from pipe",
	Run:   normal,
}

// generateRGBCode 用来生成RGB颜色数据
func generateRGBCode(freq float32, i int) (int, int, int) {
	return int(math.Sin(float64(freq)*float64(i)+0)*127 + 128),
		int(math.Sin(float64(freq)*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(float64(freq)*float64(i)+4*math.Pi/3)*127 + 128)
}

// 根据颜色打印所有字符串的rune(rune是UTF-8的单位，即一次输出一个UTF-8单元)
func print(output_string []rune) {
	for j := range output_string {
		r, g, b := generateRGBCode(0.1, j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output_string[j])
	}
}

// 代码核心，从管道中读取数据，循环打印。
func normal(cmd *cobra.Command, args []string) {
	info, _ := os.Stdin.Stat()
	var output_string []rune

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gorainbow")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		output_string = append(output_string, input)
	}
	print(output_string)
}

func init() {
	rootCmd.AddCommand(lolcatCommand)
}
