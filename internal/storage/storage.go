package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/s19835/stencil/internal/model"
)

func getSnippetsDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".stencil", "snippets")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return "", err
		}
	}

	return dir, nil
}

func SaveSnippet(s *model.Snippet) error {
	dir, err := getSnippetsDir()
	if err != nil {
		return err
	}

	filePath := filepath.Join(dir, fmt.Sprintf("%s.json", s.ID))
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	return enc.Encode(s)
}

func LoadSnippets() ([]model.Snippet, error) {
	dir, err := getSnippetsDir()
	if err != nil {
		return nil, err
	}

	var snippets []model.Snippet
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			path := filepath.Join(dir, file.Name())
			data, err := os.ReadFile(path)
			if err != nil {
				return nil, err
			}

			var s model.Snippet
			if err := json.Unmarshal(data, &s); err != nil {
				return nil, err
			}
			snippets = append(snippets, s)
		}
	}
	return snippets, nil
}
