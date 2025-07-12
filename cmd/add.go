package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new snippet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding a new snippet... (to be implemented)")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
