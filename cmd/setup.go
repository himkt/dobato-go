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
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		configDir := path.Join(homeDir, ".config/dobato")
		if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		fmt.Print("Webhook URL: ")
		reader := bufio.NewReader(os.Stdin)
		webhookUrl, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		configPath := path.Join(homeDir, ".config/dobato/webhook")
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
