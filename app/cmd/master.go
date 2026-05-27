package cmd

import (
	"fmt"
	"passg/app"

	"github.com/spf13/cobra"
)

func init() {
	masterCmd := &cobra.Command{
		Use:   "master",
		Short: "Set a new master password to access the storage",
		Long:  `Set a new master password to access the storage`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			master()
		},
	}
	app.Root.AddCommand(masterCmd)
}

func master() {
	fmt.Println("Hello")
}
