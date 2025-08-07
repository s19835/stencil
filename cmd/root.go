package cmd

import (
	"fmt"
	"os"

	"github.com/s19835/stencil/internal/tui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stencil",
	Short: "Stencil - Your personal code snippet vault",
	Long:  `Stencil is a CLI tool for managing your reusable code snippets with search, tags, and beautiful previews.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to stencil! Use 'stencil --help' to see commands.")

		app := tui.NewApp()
		if _, err := app.Run(); err != nil {
			fmt.Println("TUI error:", err)
			os.Exit(1)
		}
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
