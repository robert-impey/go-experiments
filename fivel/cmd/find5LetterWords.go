/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// find5LetterWordsCmd represents the find5LetterWords command
var find5LetterWordsCmd = &cobra.Command{
	Use:   "find5LetterWords",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 5 {
				fmt.Println(line)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(find5LetterWordsCmd)
}
