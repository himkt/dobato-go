/*
Copyright © 2022 himkt <himkt@klis.tsukuba.ac.jp>

*/
package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/spf13/cobra"
)

type Payload struct {
	Content string `json:"content"`
}

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post text to discord server",
	Long:  `Post text to discord server using HTTP.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := os.UserHomeDir()
		configPath := path.Join(home, ".config/dobato/webhook")
		if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
			fmt.Println("No config available, run dobato setup first.")
			return
		}

		f, err := os.Open(configPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
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
			fmt.Fprintln(os.Stderr)
			return
		}

		payload := new(Payload)
		payload.Content = text

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postCmd.PersistentFlags().String("foo", "", "A help for foo")
	postCmd.PersistentFlags().String("text", "test", "Text")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
