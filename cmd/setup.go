/*
Copyright © 2022 himkt <himkt@klis.tsukuba.ac.jp>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
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

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Webhook URL: ")
		webhookUrl, _ := reader.ReadString('\n')
		configPath := path.Join(home, ".config/dobato/webhook")
		f, err := os.Create(configPath)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		f.WriteString(webhookUrl)
		f.Close()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
