package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up webhook URL",
	Long:  "Set discord webhook URL from standard input.",
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		os.MkdirAll(path.Join(home, ".config/dobato"), os.ModePerm)

		fmt.Print("Webhook URL: ")
		reader := bufio.NewReader(os.Stdin)
		webhookUrl, _ := reader.ReadString('\n')

		configPath := path.Join(home, ".config/dobato/webhook")
		f, err := os.Create(configPath)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		f.WriteString(webhookUrl)
		f.Close()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
