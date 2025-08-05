/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/s19835/stencil/internal/storage"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view <id>",
	Short: "View a snippet by ID",
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		snippet, err := storage.GetSnippetByID(id)
		if err != nil {
			fmt.Println("❌", err)
			os.Exit(1)
		}

		// Print metadata
		fmt.Printf("\n📄 %s\n", snippet.Title)
		fmt.Printf("📝 Description: %s\n", snippet.Description)
		fmt.Printf("🏷️  Tags: %s\n", strings.Join(snippet.Tags, ", "))
		fmt.Printf("💻 Language: %s\n", snippet.Language)
		fmt.Printf("📅 Created: %s\n", snippet.CreatedAt.Format("2006-01-02 15:04"))
		fmt.Println(strings.Repeat("-", 50))

		// Print highlighted code
		fmt.Println("🔽 Content:")
		if snippet.Language != "" {
			// try syntax highlighting
			err = quick.Highlight(os.Stdout, snippet.Content, snippet.Language, "terminal16m", "monokai")
			if err != nil {
				// fallback plain
				fmt.Println(snippet.Content)
			}
		} else {
			// no language provided
			fmt.Println(snippet.Content)
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
