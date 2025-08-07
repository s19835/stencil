# 🧱 stencil

**stencil** is a modern, minimalistic TUI-powered code snippet manager for terminal users. Designed for developers who live in the terminal, it helps you quickly save, browse, and copy your frequently used code snippets — with style.

> Think of it as your own personal developer scratchpad — fast, searchable, and elegant.

---

## ✨ Features

- 📋 **Add snippets** with language and content
- 🔍 **Browse** through your personal snippet library in a TUI
- 📎 **Copy** snippet content to clipboard instantly (`c` key)
- ✅ **Visual toast** confirmation on copy actions
- 🎨 **Styled UI** with Bubbletea + Lipgloss
- ⏱️ Non-blocking **timer-based fade messages**
- 📁 Snippets stored locally in `~/.stencil/snippets.json`

---

## 📦 Installation

### Prerequisites

- Go 1.20+
- Git (for cloning)
- Clipboard tool (e.g., `xclip`, `pbcopy`, `wl-copy` depending on OS)

### Clone and Build

```bash
git clone https://github.com/yourname/stencil.git
cd stencil
go build -o stencil ./cmd/stencil
```

Optionally, move the binary into your `$PATH`:

```bash
sudo mv stencil /usr/local/bin/
```

---

## 🚀 Usage

### Add a snippet

```bash
stencil add
```

You'll be prompted for:

- Snippet name
- Programming language (e.g., go, py, js)
- Multiline snippet content (end with a single `.` on a new line)

### View and copy snippets

```bash
stencil list
```

Use arrow keys or `j`/`k` to navigate. Press:

- `c` to copy content
- `q` to quit

---

## 📁 Project Structure

```bash
stencil/
├── cmd/
│   └── stencil/         # CLI commands (add, list, etc.)
├── internal/
│   └── model/           # Snippet model and file handling
├── tui/                 # Bubbletea + Lipgloss UI rendering
├── go.mod
└── main.go              # Entry point
```

---

## 📋 Backlog (Next Features)

- 🔍 Fuzzy search by title/tag/lang
- 🏷️ Tag support and filtering
- 🎭 Theme toggle (light/dark)
- ✏️ Edit and delete snippets
- ☁️ GitHub Gist or markdown export
- 🔐 Encrypted snippet storage
- 🧠 Auto language detection

See [`BACKLOG.md`](./BACKLOG.md) for full roadmap.

---

## 🛠 Built With

- [Go](https://golang.org)
- [Bubbletea](https://github.com/charmbracelet/bubbletea) – TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) – Layout & styling
- [Clipboard](https://github.com/atotto/clipboard)

---

## 🧑‍💻 Author

Built with care by ❤️ [@s19835](https://github.com/s19835).

Feel free to fork, PR, or open issues!

---

## 📄 License

MIT License
