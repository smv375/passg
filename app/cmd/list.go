package cmd

import (
	"fmt"
	"passg/app"

	"github.com/spf13/cobra"
)

func init() {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all the credentials from the storage",
		Long:  "List all the credentials from the storage",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			list()
		},
	}
	app.This.AddCommand(listCmd)
}

func list() {
	fmt.Println("Hello")
}
