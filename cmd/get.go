/*
Copyright © 2022 Mohammed Imran <mohammedimran86992@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var host string
var project int
var token string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Example: "gvm get --host https://gitlab.com",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(host, project, token)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&host, "host", "s", "https://gitlab.com", "GitLab host example: https://gitlab.com")
	getCmd.Flags().IntVarP(&project, "project", "p", -1, "GitLab host example: https://gitlab.com")
	getCmd.Flags().StringVarP(&token, "token", "t", "", "GitLab host example: https://gitlab.com")
}
