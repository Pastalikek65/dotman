package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/Pastalikek65/dotman/store"
)

var (
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("99")).Padding(0,1)
	selStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	normStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	faintStyle = lipgloss.NewStyle().Faint(true)
)

type Model struct {
	files []string
	cursor int
}

func New() Model {
	files, _ := store.List()
	return Model{files: files}
}

func (m Model) Init() tea.Cmd { return nil }
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q","ctrl+c": return m, tea.Quit
		case "j","down": if m.cursor < len(m.files)-1 { m.cursor++ }
		case "k","up": if m.cursor>0 { m.cursor-- }
		}
	}
	return m, nil
}
func (m Model) View() string {
	var b strings.Builder
	b.WriteString(titleStyle.Render(" dotman — Termux ") + "\n")
	if len(m.files)==0 { b.WriteString("  no dotfiles — dotman add ~/.termux/termux.properties\n") }
	for i,f := range m.files {
		prefix:="  "
		if i==m.cursor { prefix="> " }
		line:=f
		if i==m.cursor { b.WriteString(selStyle.Render(prefix+line)+"\n") } else { b.WriteString(normStyle.Render(prefix+line)+"\n") }
	}
	b.WriteString("\n"+faintStyle.Render("j/k:nav  q:quit  (dotman restore <file>)")+"\n")
	b.WriteString(fmt.Sprintf(" %d files\n", len(m.files)))
	return b.String()
}
