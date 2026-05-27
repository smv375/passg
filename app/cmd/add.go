package cmd

import (
	"fmt"
	"passg/app"

	"github.com/spf13/cobra"
)

func init() {
	addCmd := &cobra.Command{
		Use: "add [your login] [resourse]",
		Example: `add username http://example.passg.com
add username ftp://example1.passg.com`,
		Short: "Add a credential to the storage",
		Long: `Add a credential to the storage. 
The password will be asked after entering this command`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			add()
		},
	}
	app.This.AddCommand(addCmd)
}

func add() {
	fmt.Println("Hello")
}
