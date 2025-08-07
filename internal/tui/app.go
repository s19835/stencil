package tui

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/s19835/stencil/internal/model"
	"github.com/s19835/stencil/internal/storage"
)

type modelTUI struct {
	snippets []model.Snippet
	cursor   int
	quitting bool
	height   int
	width    int
	copied   bool
}

type copiedMsg struct{}
type resetCopiedMsg struct{}

func NewApp() *tea.Program {
	snippets, _ := storage.LoadSnippets()
	return tea.NewProgram(modelTUI{
		snippets: snippets,
		cursor:   0,
	})
}

func (m modelTUI) Init() tea.Cmd {
	return nil
}

func (m modelTUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.snippets)-1 {
				m.cursor++
			}
		case "c":
			content := m.snippets[m.cursor].Content
			err := clipboard.WriteAll(content)
			if err == nil {
				m.copied = true
				return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
					return resetCopiedMsg{}
				})
			} else {
				fmt.Println("Failed to copy: ", err)
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case resetCopiedMsg:
		m.copied = false
	}

	return m, nil
}

func (m modelTUI) View() string {
	if len(m.snippets) == 0 {
		return "No snippets found. Use `stencil add` to create one.\n"
	}

	// layout width
	leftWidth := m.width / 3
	rightWidth := m.width - leftWidth - 2

	// left list
	var left strings.Builder
	for i, s := range m.snippets {
		cursor := " "
		if m.cursor == i {
			cursor = "👉"
		}
		line := fmt.Sprintf("%s %s [%s]\n", cursor, s.Title, strings.Join(s.Tags, ", "))
		left.WriteString(line)
	}

	// right preview
	selected := m.snippets[m.cursor]
	var codeBuf bytes.Buffer
	if selected.Language != "" {
		_ = quick.Highlight(&codeBuf, selected.Content, selected.Language, "terminal16m", "monokai")
	} else {
		codeBuf.WriteString(selected.Content)
	}

	preview := fmt.Sprintf(
		"📄 %s\n📝 %s\n🏷️  %s\n💻 %s\n\n%s",
		selected.Title,
		selected.Description,
		strings.Join(selected.Tags, ", "),
		selected.Language,
		codeBuf.String(),
	)

	var toast string
	if m.copied {
		toastStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true).PaddingBottom(1)
		toast = toastStyle.Render("Copied to Clipboard")
	}

	// Styles
	columnStyle := lipgloss.NewStyle().Padding(1, 2)

	leftCol := columnStyle.Width(leftWidth).Render(left.String())
	righCol := columnStyle.Width(rightWidth).Render(preview)

	layout := lipgloss.JoinHorizontal(lipgloss.Top, leftCol, righCol)

	return lipgloss.JoinVertical(lipgloss.Left, layout, toast)
}
