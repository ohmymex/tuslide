// Example custom demonstrates how to create fully customized slider styles.
//
// Run with: go run ./examples/custom
//
// Controls:
//   - Up/Down or J/K: Navigate between sliders
//   - Left/Right or H/L: Adjust slider values
//   - Tab: Cycle through color themes
//   - S: Cycle through symbol sets
//   - q or Ctrl+C: Quit
package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ohmymex/tuslide"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("255")).
			Background(lipgloss.Color("99")).
			Padding(0, 2)

	sectionStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("213")).
			MarginTop(1)

	labelStyle = lipgloss.NewStyle().
			Width(25).
			Foreground(lipgloss.Color("252"))

	focusedLabelStyle = lipgloss.NewStyle().
				Width(25).
				Bold(true).
				Foreground(lipgloss.Color("229"))

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("243"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

// Color themes
type colorTheme struct {
	name    string
	filled  lipgloss.Color
	empty   lipgloss.Color
	handle  lipgloss.Color
}

var colorThemes = []colorTheme{
	{"Cyberpunk", lipgloss.Color("#FF00FF"), lipgloss.Color("#330033"), lipgloss.Color("#00FFFF")},
	{"Sunset", lipgloss.Color("#FF6B35"), lipgloss.Color("#3D2914"), lipgloss.Color("#FFD700")},
	{"Ocean", lipgloss.Color("#00CED1"), lipgloss.Color("#1A3A3A"), lipgloss.Color("#E0FFFF")},
	{"Forest", lipgloss.Color("#32CD32"), lipgloss.Color("#1A331A"), lipgloss.Color("#90EE90")},
	{"Lava", lipgloss.Color("#FF4500"), lipgloss.Color("#331100"), lipgloss.Color("#FFFF00")},
	{"Arctic", lipgloss.Color("#87CEEB"), lipgloss.Color("#2F4F4F"), lipgloss.Color("#FFFFFF")},
	{"Neon Pink", lipgloss.Color("#FF1493"), lipgloss.Color("#2B0B1A"), lipgloss.Color("#FF69B4")},
	{"Matrix", lipgloss.Color("#00FF00"), lipgloss.Color("#001100"), lipgloss.Color("#ADFF2F")},
}

// Custom symbol sets for demonstration
var customSymbolSets = []struct {
	name    string
	symbols tuslide.Symbols
}{
	{"Blocks", tuslide.Symbols{Filled: "█", Empty: "░", Handle: "▓"}},
	{"Lines", tuslide.Symbols{Filled: "━", Empty: "─", Handle: "●"}},
	{"Dots", tuslide.Symbols{Filled: "●", Empty: "○", Handle: "◉"}},
	{"Squares", tuslide.Symbols{Filled: "■", Empty: "□", Handle: "▣"}},
	{"Stars", tuslide.Symbols{Filled: "★", Empty: "☆", Handle: "✦"}},
	{"Diamonds", tuslide.Symbols{Filled: "◆", Empty: "◇", Handle: "◈"}},
	{"Arrows", tuslide.Symbols{Filled: "▶", Empty: "▷", Handle: "►"}},
	{"Waves", tuslide.Symbols{Filled: "≈", Empty: "~", Handle: "≋"}},
	{"Hearts", tuslide.Symbols{Filled: "♥", Empty: "♡", Handle: "❤"}},
	{"Music", tuslide.Symbols{Filled: "♪", Empty: "♩", Handle: "♫"}},
	{"Emoji", tuslide.Symbols{Filled: "🟢", Empty: "⚪", Handle: "🔵"}},
	{"Fire", tuslide.Symbols{Filled: "🔥", Empty: "💨", Handle: "⚡"}},
}

type sliderConfig struct {
	name       string
	state      *tuslide.SliderState
	themeIdx   int
	symbolIdx  int
	segmented  bool
	showHandle bool
}

type model struct {
	sliders  []sliderConfig
	focused  int
	quitting bool
}

func initialModel() model {
	return model{
		sliders: []sliderConfig{
			{
				name:       "Standard Slider",
				state:      tuslide.NewState(tuslide.WithValue(50), tuslide.WithStep(5)),
				themeIdx:   0,
				symbolIdx:  0,
				showHandle: true,
			},
			{
				name:       "Progress Bar (no handle)",
				state:      tuslide.NewState(tuslide.WithValue(75), tuslide.WithStep(5)),
				themeIdx:   1,
				symbolIdx:  1,
				showHandle: false,
			},
			{
				name:       "Segmented Style",
				state:      tuslide.NewState(tuslide.WithValue(60), tuslide.WithStep(10)),
				themeIdx:   2,
				symbolIdx:  2,
				segmented:  true,
				showHandle: true,
			},
			{
				name:       "Star Rating",
				state:      tuslide.NewState(tuslide.WithValue(40), tuslide.WithStep(10)),
				themeIdx:   3,
				symbolIdx:  4,
				segmented:  true,
				showHandle: false,
			},
			{
				name:       "Emoji Slider",
				state:      tuslide.NewState(tuslide.WithValue(30), tuslide.WithStep(5)),
				themeIdx:   4,
				symbolIdx:  10,
				showHandle: true,
			},
		},
		focused: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "up", "k":
			m.focused--
			if m.focused < 0 {
				m.focused = len(m.sliders) - 1
			}

		case "down", "j":
			m.focused = (m.focused + 1) % len(m.sliders)

		case "left", "h":
			m.sliders[m.focused].state.Decrement()

		case "right", "l":
			m.sliders[m.focused].state.Increment()

		case "tab":
			// Cycle color theme for focused slider
			m.sliders[m.focused].themeIdx = (m.sliders[m.focused].themeIdx + 1) % len(colorThemes)

		case "s":
			// Cycle symbol set for focused slider
			m.sliders[m.focused].symbolIdx = (m.sliders[m.focused].symbolIdx + 1) % len(customSymbolSets)

		case "g":
			// Toggle segmented mode
			m.sliders[m.focused].segmented = !m.sliders[m.focused].segmented

		case "n":
			// Toggle handle visibility
			m.sliders[m.focused].showHandle = !m.sliders[m.focused].showHandle
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "Thanks for customizing!\n"
	}

	var b strings.Builder

	// Title
	b.WriteString("\n  ")
	b.WriteString(titleStyle.Render("Custom Styles Demo"))
	b.WriteString("\n")

	// Instructions
	b.WriteString(sectionStyle.Render("\n  Customize each slider with different themes and symbols:\n"))

	// Render each slider
	for i, cfg := range m.sliders {
		theme := colorThemes[cfg.themeIdx]
		symbols := customSymbolSets[cfg.symbolIdx]

		// Focus indicator
		indicator := "  "
		lStyle := labelStyle
		if i == m.focused {
			indicator = "▸ "
			lStyle = focusedLabelStyle
		}

		// Create slider with custom style
		opts := []tuslide.SliderOption{
			tuslide.WithWidth(30),
			tuslide.WithSymbols(symbols.symbols),
			tuslide.WithFilledStyle(lipgloss.NewStyle().Foreground(theme.filled)),
			tuslide.WithEmptyStyle(lipgloss.NewStyle().Foreground(theme.empty)),
			tuslide.WithHandleStyle(lipgloss.NewStyle().Foreground(theme.handle)),
			tuslide.WithHandle(cfg.showHandle),
			tuslide.WithShowValue(true),
		}

		if cfg.segmented {
			opts = append(opts,
				tuslide.WithSegmented(true),
				tuslide.WithSegmentCount(10),
			)
		}

		slider := tuslide.New(cfg.state, opts...)

		b.WriteString("\n")
		b.WriteString(indicator)
		b.WriteString(lStyle.Render(cfg.name))
		b.WriteString("\n   ")
		b.WriteString(slider.View())
		b.WriteString("\n   ")
		b.WriteString(infoStyle.Render(fmt.Sprintf("Theme: %s | Symbols: %s", theme.name, symbols.name)))
		if cfg.segmented {
			b.WriteString(infoStyle.Render(" | Segmented"))
		}
		if !cfg.showHandle {
			b.WriteString(infoStyle.Render(" | No Handle"))
		}
		b.WriteString("\n")
	}

	// Available themes and symbols
	b.WriteString(sectionStyle.Render("\n  Available Color Themes:"))
	b.WriteString("\n  ")
	for i, theme := range colorThemes {
		style := lipgloss.NewStyle().Foreground(theme.filled)
		if i > 0 {
			b.WriteString(" | ")
		}
		b.WriteString(style.Render(theme.name))
	}
	b.WriteString("\n")

	b.WriteString(sectionStyle.Render("\n  Available Symbol Sets:"))
	b.WriteString("\n  ")
	count := 0
	for _, sym := range customSymbolSets {
		if count > 0 {
			b.WriteString(" | ")
		}
		b.WriteString(fmt.Sprintf("%s %s%s%s", sym.name, sym.symbols.Filled, sym.symbols.Handle, sym.symbols.Empty))
		count++
		if count%4 == 0 {
			b.WriteString("\n  ")
			count = 0
		}
	}
	b.WriteString("\n")

	// Help
	b.WriteString(helpStyle.Render("\n  [↑↓] Navigate  [←→] Adjust  [Tab] Theme  [S] Symbols  [G] Segmented  [N] Handle  [Q] Quit\n"))

	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
