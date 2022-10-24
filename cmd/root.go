/*
Copyright © 2022 Mohammed Imran <mohammedimran86992@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gvm",
	Short:   "GitLab variable Manager",
	Long:    "A CLI tool to easily manage your GitLab CI/CD variable.\nQuickly fetch, debug and update CI/CD variables directly from your terminal.",
	Example: "gvm get --host https://gitlab.com -p 123 -t XXX",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
