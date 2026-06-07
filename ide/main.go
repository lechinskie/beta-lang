package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ── Focus panels ─────────────────────────────────────────────────────────────

const (
	panelEditor = iota
	panelOutput
)

// ── Compile modes ─────────────────────────────────────────────────────────────

const (
	compileNormal = iota
	compileSymTab
	compileAssembly
)

// ── Palette ───────────────────────────────────────────────────────────────────

const (
	colBg       = "#1C1C2E"
	colBgPanel  = "#16213E"
	colBgActive = "#0F3460"
	colBgAccent = "#E94560"
	colBorder   = "#414868"
	colBorderHi = "#7AA2F7"
	colText     = "#C0CAF5"
	colTextDim  = "#565F89"
	colTextHint = "#9AA5CE"
	colGreen    = "#9ECE6A"
	colRed      = "#F7768E"
	colYellow   = "#E0AF68"
	colCyan     = "#7DCFFF"
	colMagenta  = "#BB9AF7"
)

// ── Styles ────────────────────────────────────────────────────────────────────

var (
	styleTitleLeft = lipgloss.NewStyle().
			Background(lipgloss.Color(colBgAccent)).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true).
			Padding(0, 2)

	styleTitleFile = lipgloss.NewStyle().
			Background(lipgloss.Color(colBgActive)).
			Foreground(lipgloss.Color(colCyan)).
			Padding(0, 2)

	styleTitleFill = lipgloss.NewStyle().
			Background(lipgloss.Color(colBgPanel)).
			Foreground(lipgloss.Color(colTextDim))

	styleTitleMode = lipgloss.NewStyle().
			Background(lipgloss.Color(colMagenta)).
			Foreground(lipgloss.Color("#1C1C2E")).
			Bold(true).
			Padding(0, 2)

	stylePanelActive = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(colBorderHi)).
				Background(lipgloss.Color(colBg))

	stylePanelInactive = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(colBorder)).
				Background(lipgloss.Color(colBg))

	styleTabOK = lipgloss.NewStyle().
			Background(lipgloss.Color(colGreen)).
			Foreground(lipgloss.Color("#1C1C2E")).
			Bold(true).
			Padding(0, 1)

	styleTabErr = lipgloss.NewStyle().
			Background(lipgloss.Color(colRed)).
			Foreground(lipgloss.Color("#1C1C2E")).
			Bold(true).
			Padding(0, 1)

	styleTabIdle = lipgloss.NewStyle().
			Background(lipgloss.Color(colBgPanel)).
			Foreground(lipgloss.Color(colTextDim)).
			Padding(0, 1)

	styleOutputOK = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colGreen)).
			Bold(true)

	styleOutputErr = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colRed))

	styleOutputDim = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colTextDim)).
			Italic(true)

	styleCompiling = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colYellow)).
			Bold(true).
			Italic(true)

	styleHintKey = lipgloss.NewStyle().
			Background(lipgloss.Color(colBgActive)).
			Foreground(lipgloss.Color(colCyan)).
			Bold(true).
			Padding(0, 1)

	styleHintLabel = lipgloss.NewStyle().
			Background(lipgloss.Color(colBgPanel)).
			Foreground(lipgloss.Color(colTextHint)).
			Padding(0, 1)

	styleHintsFill = lipgloss.NewStyle().
			Background(lipgloss.Color(colBgPanel)).
			Foreground(lipgloss.Color(colTextDim))
)

// ── Tea messages ──────────────────────────────────────────────────────────────

type compileResultMsg struct {
	success     bool
	output      string
	compileMode int
}

// ── Model ─────────────────────────────────────────────────────────────────────

type model struct {
	editor       textarea.Model
	output       string
	compiled     bool
	success      bool
	compiling    bool
	focus        int
	width        int
	height       int
	filename     string
	outputScroll int
	compileMode  int
}

const (
	outputPanelH = 10
	titleBarH    = 1
	hintsBarH    = 1
	borderV      = 2
)

func newModel(filename string) model {
	ta := textarea.New()
	ta.Placeholder = "// Write your B code here..."
	ta.ShowLineNumbers = true
	ta.CharLimit = 0
	ta.SetWidth(80)
	ta.SetHeight(20)
	ta.Focus()

	base := lipgloss.NewStyle().
		Background(lipgloss.Color(colBg)).
		Foreground(lipgloss.Color(colText))

	lineNum := lipgloss.NewStyle().
		Background(lipgloss.Color(colBg)).
		Foreground(lipgloss.Color(colTextDim)).
		PaddingRight(1)

	cursorLine := lipgloss.NewStyle().
		Background(lipgloss.Color(colBgPanel))

	blurredBase := lipgloss.NewStyle().
		Background(lipgloss.Color(colBg)).
		Foreground(lipgloss.Color(colTextDim))

	blurredLineNum := lipgloss.NewStyle().
		Background(lipgloss.Color(colBg)).
		Foreground(lipgloss.Color(colBorder)).
		PaddingRight(1)

	ta.FocusedStyle.Base = base
	ta.FocusedStyle.LineNumber = lineNum
	ta.FocusedStyle.CursorLine = cursorLine
	ta.BlurredStyle.Base = blurredBase
	ta.BlurredStyle.LineNumber = blurredLineNum
	ta.BlurredStyle.CursorLine = lipgloss.NewStyle()

	content := ""
	if filename != "" {
		if data, err := os.ReadFile(filename); err == nil {
			content = string(data)
		}
	}
	if content != "" {
		ta.SetValue(content)
	}

	return model{
		editor:   ta,
		filename: filename,
		focus:    panelEditor,
	}
}

// ── Init ──────────────────────────────────────────────────────────────────────

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

// ── Compile command ───────────────────────────────────────────────────────────

func runCompile(code string, mode int, flags ...string) tea.Cmd {
	return func() tea.Msg {
		tmpFile, err := os.CreateTemp("", "beta-*.b")
		if err != nil {
			return compileResultMsg{false, "failed to create temp file: " + err.Error(), mode}
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(code); err != nil {
			tmpFile.Close()
			return compileResultMsg{false, "failed to write source: " + err.Error(), mode}
		}
		tmpFile.Close()

		compiler := "./beta"
		if _, statErr := os.Stat(compiler); os.IsNotExist(statErr) {
			if p, lookErr := exec.LookPath("beta"); lookErr == nil {
				compiler = p
			}
		} else {
			if abs, absErr := filepath.Abs(compiler); absErr == nil {
				compiler = abs
			}
		}

		args := append([]string{}, flags...)
		args = append(args, tmpFile.Name())
		cmd := exec.Command(compiler, args...)
		out, err := cmd.CombinedOutput()
		output := strings.TrimSpace(string(out))

		if err != nil {
			if output == "" {
				output = err.Error()
			}
			return compileResultMsg{false, output, mode}
		}
		if output == "" {
			output = "Compilation successful — no output."
		}
		return compileResultMsg{true, output, mode}
	}
}

// ── Update ────────────────────────────────────────────────────────────────────

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.relayout()

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+q":
			return m, tea.Quit

		case "ctrl+b":
			if !m.compiling {
				m.compiling = true
				m.compiled = false
				m.compileMode = compileNormal
				return m, runCompile(m.editor.Value(), compileNormal)
			}

		case "ctrl+t":
			if !m.compiling {
				m.compiling = true
				m.compiled = false
				m.compileMode = compileSymTab
				return m, runCompile(m.editor.Value(), compileSymTab, "--exp-st")
			}

		case "ctrl+s":
			if m.filename != "" {
				_ = os.WriteFile(m.filename, []byte(m.editor.Value()), 0644)
			}
			if !m.compiling {
				m.compiling = true
				m.compiled = false
				m.compileMode = compileAssembly
				return m, runCompile(m.editor.Value(), compileAssembly, "--S")
			}

		case "tab":
			if m.focus == panelEditor {
				m.focus = panelOutput
				m.editor.Blur()
			} else {
				m.focus = panelEditor
				m.editor.Focus()
			}
			return m, nil

		case "up":
			if m.focus == panelOutput && m.outputScroll > 0 {
				m.outputScroll--
				return m, nil
			}

		case "down":
			if m.focus == panelOutput {
				m.outputScroll++
				return m, nil
			}
		}

	case compileResultMsg:
		m.compiling = false
		m.compiled = true
		m.success = msg.success
		m.output = msg.output
		m.outputScroll = 0
		if msg.success && msg.compileMode == compileAssembly {
			_ = clipboard.WriteAll(msg.output)
		}
		if !msg.success {
			m.focus = panelOutput
			m.editor.Blur()
		}
		return m, nil
	}

	if m.focus == panelEditor {
		var cmd tea.Cmd
		m.editor, cmd = m.editor.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *model) relayout() {
	innerW := m.width - 2
	if innerW < 10 {
		innerW = 10
	}

	editorH := m.height - titleBarH - hintsBarH - (outputPanelH + borderV) - borderV
	if editorH < 3 {
		editorH = 3
	}

	m.editor.SetWidth(innerW)
	m.editor.SetHeight(editorH)
}

// ── View ──────────────────────────────────────────────────────────────────────

func (m model) View() string {
	if m.width == 0 {
		return "Initialising…\n"
	}
	return lipgloss.JoinVertical(lipgloss.Left,
		m.viewTitleBar(),
		m.viewEditorPanel(),
		m.viewOutputPanel(),
		m.viewHintsBar(),
	)
}

func (m model) viewTitleBar() string {
	left := styleTitleLeft.Render(" B(eta) IDE ")

	label := m.filename
	if label == "" {
		label = "[unsaved]"
	}
	file := styleTitleFile.Render(" " + label + " ")

	mode := styleTitleMode.Render(" INSERT ")

	leftSection := lipgloss.JoinHorizontal(lipgloss.Top, left, file)
	usedW := lipgloss.Width(leftSection) + lipgloss.Width(mode)
	fillW := clampMin(m.width-usedW, 0)
	fill := styleTitleFill.Render(strings.Repeat(" ", fillW))

	return lipgloss.JoinHorizontal(lipgloss.Top, leftSection, fill, mode)
}

func (m model) viewEditorPanel() string {
	panelW := m.width - 2

	var panel lipgloss.Style
	if m.focus == panelEditor {
		panel = stylePanelActive.Width(panelW)
	} else {
		panel = stylePanelInactive.Width(panelW)
	}

	return panel.Render(m.editor.View())
}

func (m model) viewOutputPanel() string {
	panelW := m.width - 2

	var tabLabel string
	switch {
	case m.compiling:
		switch m.compileMode {
		case compileSymTab:
			tabLabel = styleCompiling.Render(" ⟳  Symbol Table… ")
		case compileAssembly:
			tabLabel = styleCompiling.Render(" ⟳  Assembly… ")
		default:
			tabLabel = styleCompiling.Render(" ⟳  Compiling… ")
		}
	case !m.compiled:
		tabLabel = styleTabIdle.Render(" OUTPUT ")
	case m.compileMode == compileSymTab && m.success:
		tabLabel = styleTabOK.Render(" ◇  SYMTAB ")
	case m.compileMode == compileSymTab && !m.success:
		tabLabel = styleTabErr.Render(" ✗  SYMTAB ERR ")
	case m.compileMode == compileAssembly && m.success:
		tabLabel = styleTabOK.Render(" ◆  ASSEMBLY ")
	case m.compileMode == compileAssembly && !m.success:
		tabLabel = styleTabErr.Render(" ✗  ASM ERR ")
	case m.success:
		tabLabel = styleTabOK.Render(" ✓  OK ")
	default:
		tabLabel = styleTabErr.Render(" ✗  ERRORS ")
	}

	used := lipgloss.Width(tabLabel) + 1
	fillLen := clampMin(panelW-used, 0)
	dim := lipgloss.NewStyle().Foreground(lipgloss.Color(colBorder))
	header := dim.Render("─") + tabLabel + dim.Render(strings.Repeat("─", fillLen))

	var body string
	switch {
	case m.compiling:
		body = styleCompiling.Render("  Compiling, please wait…")
	case !m.compiled:
		body = styleOutputDim.Render("  Press Ctrl+B to compile.")
	default:
		lines := strings.Split(m.output, "\n")
		maxScroll := clampMin(len(lines)-1, 0)
		if m.outputScroll > maxScroll {
			m.outputScroll = maxScroll
		}
		visible := lines[m.outputScroll:]
		if len(visible) > outputPanelH {
			visible = visible[:outputPanelH]
		}
		var sb strings.Builder
		lineStyle := styleOutputOK
		if !m.success {
			lineStyle = styleOutputErr
		}
		for _, l := range visible {
			sb.WriteString(lineStyle.Render("  "+l) + "\n")
		}
		body = strings.TrimRight(sb.String(), "\n")
	}

	bodyLines := strings.Split(body, "\n")
	for len(bodyLines) < outputPanelH {
		bodyLines = append(bodyLines, "")
	}
	body = strings.Join(bodyLines, "\n")

	borderColor := colBorder
	if m.focus == panelOutput {
		borderColor = colBorderHi
	}

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(borderColor)).
		Width(panelW).
		Background(lipgloss.Color(colBg)).
		Render(header + "\n" + body)

	return box
}

func (m model) viewHintsBar() string {
	hints := []struct{ key, label string }{
		{"^B", "compile"},
		{"^T", "sym table"},
		{"^S", "save+asm"},
		{"Tab", "switch panel"},
		{"↑↓", "scroll output"},
		{"^Q", "quit"},
	}

	var parts []string
	for _, h := range hints {
		parts = append(parts,
			styleHintKey.Render(h.key),
			styleHintLabel.Render(h.label),
		)
	}

	bar := lipgloss.JoinHorizontal(lipgloss.Top, parts...)
	usedW := lipgloss.Width(bar)
	fillW := clampMin(m.width-usedW, 0)
	fill := styleHintsFill.Render(strings.Repeat(" ", fillW))

	return lipgloss.JoinHorizontal(lipgloss.Top, bar, fill)
}

// ── Helpers ───────────────────────────────────────────────────────────────────

func clampMin(v, min int) int {
	if v < min {
		return min
	}
	return v
}

// ── Entry point ───────────────────────────────────────────────────────────────

func main() {
	filename := ""
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	p := tea.NewProgram(
		newModel(filename),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
