/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// answerCmd represents the answer command
var responseCmd = &cobra.Command{
	Use:   "response",
	Short: "Get a response from a web server.",
	Long:  `Get a response from a web server, as a test.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get("https://www.google.com")
		if err != nil {
			fmt.Printf("Error making the HTTP request: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Request: GET https://www.google.com\n")
		fmt.Printf("Result: Success\n")
		fmt.Printf("Status Code: %d\n", res.StatusCode)
	},
}

func init() {
	rootCmd.AddCommand(responseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// answerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// answerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
