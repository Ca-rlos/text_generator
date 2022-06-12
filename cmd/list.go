/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List files in the current directory",
	Long:  `List files in the current directory, as a test.`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := os.ReadDir(".")
		if err != nil {
			fmt.Println("Error reading directory:", err)
			os.Exit(1)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
