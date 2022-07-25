package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/spf13/cobra"
)

type Payload struct {
	Content string `json:"content"`
}

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post text to discord server",
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := os.UserHomeDir()
		configPath := path.Join(home, ".config/dobato/webhook")

		_, err := os.Stat(configPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		f, err := os.Open(configPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		defer f.Close()
		scanner := bufio.NewScanner(f)

		var webhookUrl string
		if scanner.Scan() {
			webhookUrl = scanner.Text()
		} else {
			panic("Failed to read webhook")
		}

		text, err := cmd.Flags().GetString("text")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		payload := Payload{
			Content: text,
		}

		json, err := json.Marshal(payload)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		fmt.Println("web: ", webhookUrl)
		res, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(json))
		if err != nil {
			fmt.Println("errrr")
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer res.Body.Close()
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.PersistentFlags().String("text", "test", "Text")
}
