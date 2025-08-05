package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/s19835/stencil/internal/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snippets",
	Run: func(cmd *cobra.Command, args []string) {
		snippets, err := storage.LoadSnippets()
		if err != nil {
			fmt.Println("Unnable to load snippets: ", err)
			os.Exit(1)
		}

		if len(snippets) == 0 {
			fmt.Println("No snippets found. Use `stencil add` to create one.")
			return
		}

		// Print header
		fmt.Printf("%-8s  %-25s %-10s %-20s %s\n", "ID", "Title", "Lang", "Tags", "Created")
		fmt.Println(strings.Repeat("-", 80))

		for _, s := range snippets {
			// Shorten ID for table
			idShort := s.ID[:8]
			tags := strings.Join(s.Tags, ", ")
			created := s.CreatedAt.Format("2006-01-02")

			fmt.Printf("%-8s  %-25s %-10s %-20s %s\n",
				idShort, s.Title, s.Language, tags, created)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
