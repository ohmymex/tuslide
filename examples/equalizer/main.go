// Example equalizer demonstrates vertical sliders in an audio equalizer layout.
//
// Run with: go run ./examples/equalizer
//
// Controls:
//   - Left/Right or H/L: Navigate between bands
//   - Up/Down or J/K: Adjust band level
//   - R: Reset all bands to 50%
//   - P: Toggle preset cycling
//   - 1-8: Jump to specific band
//   - q or Ctrl+C: Quit
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ohmymex/tuslide"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Padding(0, 2)

	bandLabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("250")).
			Width(6).
			Align(lipgloss.Center)

	focusedBandStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("229")).
				Bold(true).
				Width(6).
				Align(lipgloss.Center)

	dbStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("243")).
		Width(6).
		Align(lipgloss.Center)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	presetStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("213")).
			Bold(true)
)

// Band frequencies
var bandFreqs = []string{"32Hz", "64Hz", "125Hz", "250Hz", "500Hz", "1kHz", "2kHz", "4kHz"}

// Preset configurations
var presets = map[string][]float64{
	"Flat":      {50, 50, 50, 50, 50, 50, 50, 50},
	"Bass":      {80, 75, 65, 50, 45, 45, 50, 50},
	"Treble":    {45, 45, 50, 50, 55, 65, 75, 80},
	"V-Shape":   {75, 65, 50, 40, 40, 50, 65, 75},
	"Vocal":     {40, 45, 60, 70, 70, 60, 45, 40},
	"Electronic": {70, 75, 55, 50, 55, 65, 70, 65},
}

var presetNames = []string{"Flat", "Bass", "Treble", "V-Shape", "Vocal", "Electronic"}

type tickMsg time.Time

type model struct {
	bands        []*tuslide.SliderState
	focusedBand  int
	presetIndex  int
	cycling      bool
	quitting     bool
}

func initialModel() model {
	bands := make([]*tuslide.SliderState, 8)
	for i := range bands {
		bands[i] = tuslide.NewState(
			tuslide.WithMin(0),
			tuslide.WithMax(100),
			tuslide.WithValue(50),
			tuslide.WithStep(5),
		)
	}

	return model{
		bands:       bands,
		focusedBand: 0,
		presetIndex: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if m.cycling {
			m.presetIndex = (m.presetIndex + 1) % len(presetNames)
			m.applyPreset(presetNames[m.presetIndex])
			return m, tickCmd()
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "left", "h":
			m.focusedBand--
			if m.focusedBand < 0 {
				m.focusedBand = len(m.bands) - 1
			}

		case "right", "l":
			m.focusedBand = (m.focusedBand + 1) % len(m.bands)

		case "up", "k":
			m.bands[m.focusedBand].Increment()

		case "down", "j":
			m.bands[m.focusedBand].Decrement()

		case "r":
			// Reset all to 50%
			for _, band := range m.bands {
				band.SetValue(50)
			}

		case "p":
			m.cycling = !m.cycling
			if m.cycling {
				return m, tickCmd()
			}

		case "1", "2", "3", "4", "5", "6", "7", "8":
			idx := int(msg.String()[0] - '1')
			if idx < len(m.bands) {
				m.focusedBand = idx
			}

		case "[":
			m.presetIndex--
			if m.presetIndex < 0 {
				m.presetIndex = len(presetNames) - 1
			}
			m.applyPreset(presetNames[m.presetIndex])

		case "]":
			m.presetIndex = (m.presetIndex + 1) % len(presetNames)
			m.applyPreset(presetNames[m.presetIndex])
		}
	}

	return m, nil
}

func (m *model) applyPreset(name string) {
	if values, ok := presets[name]; ok {
		for i, v := range values {
			if i < len(m.bands) {
				m.bands[i].SetValue(v)
			}
		}
	}
}

func (m model) View() string {
	if m.quitting {
		return "Thanks for using the equalizer!\n"
	}

	var b strings.Builder

	// Title
	b.WriteString(titleStyle.Render("  Audio Equalizer  "))
	b.WriteString("\n\n")

	// Current preset
	cycleStatus := ""
	if m.cycling {
		cycleStatus = " (auto-cycling)"
	}
	b.WriteString(fmt.Sprintf("  Preset: %s%s\n\n", presetStyle.Render(presetNames[m.presetIndex]), cycleStatus))

	// Render all equalizer bands
	sliderHeight := 12

	// Create sliders for each band
	sliders := make([]*tuslide.Slider, len(m.bands))
	for i, band := range m.bands {
		// Choose color based on frequency (low=red, mid=yellow, high=cyan)
		var filledColor lipgloss.Color
		switch {
		case i < 2:
			filledColor = lipgloss.Color("196") // Red for bass
		case i < 4:
			filledColor = lipgloss.Color("208") // Orange for low-mid
		case i < 6:
			filledColor = lipgloss.Color("226") // Yellow for mid
		default:
			filledColor = lipgloss.Color("87") // Cyan for treble
		}

		sliders[i] = tuslide.New(band,
			tuslide.WithHeight(sliderHeight),
			tuslide.WithOrientation(tuslide.Vertical),
			tuslide.WithFilledStyle(lipgloss.NewStyle().Foreground(filledColor)),
			tuslide.WithEmptyStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("238"))),
			tuslide.WithHandleStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("255"))),
		)
	}

	// Render sliders side by side
	views := make([][]string, len(sliders))
	maxLines := 0
	for i, s := range sliders {
		views[i] = strings.Split(s.View(), "\n")
		if len(views[i]) > maxLines {
			maxLines = len(views[i])
		}
	}

	// dB scale on the left
	dbLabels := []string{"+12", "+6", "0", "-6", "-12"}
	dbStep := maxLines / (len(dbLabels) - 1)

	for line := 0; line < maxLines; line++ {
		// dB label
		dbIdx := line / dbStep
		if line%dbStep == 0 && dbIdx < len(dbLabels) {
			b.WriteString(dbStyle.Render(dbLabels[dbIdx]))
		} else {
			b.WriteString(strings.Repeat(" ", 6))
		}

		// Each band
		for i := range sliders {
			b.WriteString(" ")
			if line < len(views[i]) {
				// Center the slider character in a 6-char width
				content := views[i][line]
				padding := (6 - len(content)) / 2
				b.WriteString(strings.Repeat(" ", padding))
				b.WriteString(content)
				b.WriteString(strings.Repeat(" ", 6-padding-len(content)))
			} else {
				b.WriteString(strings.Repeat(" ", 6))
			}
		}
		b.WriteString("\n")
	}

	// Band labels
	b.WriteString(strings.Repeat(" ", 6)) // Align with dB column
	for i, freq := range bandFreqs {
		labelStyle := bandLabelStyle
		if i == m.focusedBand {
			labelStyle = focusedBandStyle
		}
		b.WriteString(" ")
		b.WriteString(labelStyle.Render(freq))
	}
	b.WriteString("\n")

	// Focus indicator
	b.WriteString(strings.Repeat(" ", 6))
	for i := range m.bands {
		if i == m.focusedBand {
			b.WriteString("   ")
			b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("229")).Render("▲"))
			b.WriteString("  ")
		} else {
			b.WriteString(strings.Repeat(" ", 7))
		}
	}
	b.WriteString("\n")

	// Value display
	b.WriteString(strings.Repeat(" ", 6))
	for i, band := range m.bands {
		val := fmt.Sprintf("%.0f%%", band.Value())
		style := dbStyle
		if i == m.focusedBand {
			style = focusedBandStyle
		}
		b.WriteString(" ")
		b.WriteString(style.Render(val))
	}
	b.WriteString("\n")

	// Help
	b.WriteString(helpStyle.Render("\n  [←→] Select band  [↑↓] Adjust  [R] Reset  [P] Auto-cycle  [/]] Preset  [Q] Quit"))

	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
