package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	delCmd := &cobra.Command{
		Use: "del [your login] [resourse]",
		Example: `del username http://example2.passsg.org
del username+ftp://example14.passg.org`,
		Short: "Delete a credential from the storage",
		Long:  "Delete a credential from the storage",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			del()
		},
	}
	rootCmd.AddCommand(delCmd)
}

func del() {
	fmt.Println("Hello")
}
