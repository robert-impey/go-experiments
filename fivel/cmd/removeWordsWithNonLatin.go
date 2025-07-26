/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

// removeWordsWithNonLatinCmd represents the removeWordsWithNonLatin command
var removeWordsWithNonLatinCmd = &cobra.Command{
	Use:   "removeWordsWithNonLatin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		regexp := regexp.MustCompile(`^[a-zA-Z]+$`)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if regexp.MatchString(line) {
				fmt.Println(line)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(removeWordsWithNonLatinCmd)
}
