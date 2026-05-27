package app

import (
	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use:   "passg",
	Short: "A simple password manager.",
	Long:  `passg is a CLI utility for encrypting and managing passwords`}

func Execute() error {
	Root.CompletionOptions.DisableDefaultCmd = true
	return Root.Execute()
}
