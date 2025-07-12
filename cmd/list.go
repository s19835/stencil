package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snippets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing snippets... (to be implemented)")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
