package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/s19835/stencil/internal/model"
	"github.com/s19835/stencil/internal/storage"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new snippet",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Title: ")
		title, _ := reader.ReadString('\n')
		title = strings.TrimSpace(title)

		fmt.Print("Language (e.g., go, python, bash): ")
		lang, _ := reader.ReadString('\n')
		lang = strings.TrimSpace(lang)

		fmt.Print("Tags (comma-separated): ")
		tagsInput, _ := reader.ReadString('\n')
		tagsInput = strings.TrimSpace(tagsInput)
		tags := []string{}

		if tagsInput != "" {
			tags = strings.Split(tagsInput, ",")
			for i := range tags {
				tags[i] = strings.TrimSpace(tags[i])
			}
		}

		fmt.Print("Description: ")
		desc, _ := reader.ReadString('\n')
		desc = strings.TrimSpace(desc)

		fmt.Println("Content (end with a single dot '.' on a line):")
		var contentLines []string
		for {
			line, _ := reader.ReadString('\n')
			line = strings.TrimRight(line, "\n")
			if line == "." {
				break
			}
			contentLines = append(contentLines, line)
		}
		content := strings.Join(contentLines, "\n")

		snippet := &model.Snippet{
			ID:          uuid.NewString(),
			Title:       title,
			Content:     content,
			Language:    lang,
			Tags:        tags,
			Description: desc,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err := storage.SaveSnippet(snippet)
		if err != nil {
			fmt.Println("Failed to save snippet:", err)
			os.Exit(1)
		}

		fmt.Println("Snippet saved successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
