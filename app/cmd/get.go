package cmd

import (
	"fmt"
	"passg/app"

	"github.com/spf13/cobra"
)

func init() {
	getCmd := &cobra.Command{
		Use: "get [your login] [resourse]",
		Example: `get username http://example.passg.com
get username+ftp://example1.passg.com`,
		Short: "Get a credential from the storage",
		Long:  "Get a credential from the storage",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			get(args[0])
		},
	}
	app.This.AddCommand(getCmd)
}

func get(s string) {
	fmt.Println(s)
}
