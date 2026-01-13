// Example basic demonstrates a simple interactive slider using TuSlide with Bubble Tea.
//
// Run with: go run ./examples/basic
//
// Controls:
//   - Left/Right arrows: Adjust the slider value
//   - Up/Down arrows: Adjust the second slider
//   - q or Ctrl+C: Quit
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ohmymex/tuslide"
)

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			MarginBottom(1)

	focusedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86"))

	unfocusedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			MarginTop(1)
)

// model holds the application state.
type model struct {
	sliders     []*tuslide.Slider
	states      []*tuslide.SliderState
	focusIndex  int
	quitting    bool
}

func initialModel() model {
	// Create states for two sliders
	volumeState := tuslide.NewState(
		tuslide.WithMin(0),
		tuslide.WithMax(100),
		tuslide.WithValue(50),
		tuslide.WithStep(5),
	)

	brightnessState := tuslide.NewState(
		tuslide.WithMin(0),
		tuslide.WithMax(100),
		tuslide.WithValue(75),
		tuslide.WithStep(1),
	)

	// Create sliders with different configurations
	volumeSlider := tuslide.New(volumeState,
		tuslide.WithWidth(30),
		tuslide.WithLabel("Volume"),
		tuslide.WithLabelPosition(tuslide.LabelLeft),
		tuslide.WithShowValue(true),
		tuslide.WithFilledStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("86"))),
		tuslide.WithEmptyStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("241"))),
		tuslide.WithHandleStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("205"))),
	)

	brightnessSlider := tuslide.New(brightnessState,
		tuslide.WithWidth(30),
		tuslide.WithLabel("Brightness"),
		tuslide.WithLabelPosition(tuslide.LabelLeft),
		tuslide.WithShowValue(true),
		tuslide.WithFilledStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("220"))),
		tuslide.WithEmptyStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("241"))),
		tuslide.WithHandleStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("205"))),
	)

	return model{
		sliders:    []*tuslide.Slider{volumeSlider, brightnessSlider},
		states:     []*tuslide.SliderState{volumeState, brightnessState},
		focusIndex: 0,
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
			// Move focus up
			m.focusIndex--
			if m.focusIndex < 0 {
				m.focusIndex = len(m.sliders) - 1
			}

		case "down", "j":
			// Move focus down
			m.focusIndex++
			if m.focusIndex >= len(m.sliders) {
				m.focusIndex = 0
			}

		case "left", "h":
			// Decrease value of focused slider
			m.states[m.focusIndex].Decrement()

		case "right", "l":
			// Increase value of focused slider
			m.states[m.focusIndex].Increment()

		case "home":
			// Set to minimum
			m.states[m.focusIndex].SetFromPercentage(0)

		case "end":
			// Set to maximum
			m.states[m.focusIndex].SetFromPercentage(1)
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	s := titleStyle.Render("🎚️  TuSlide Demo") + "\n\n"

	for i, slider := range m.sliders {
		// Apply focus styling
		indicator := "  "
		if i == m.focusIndex {
			indicator = focusedStyle.Render("▸ ")
		}

		s += indicator + slider.View() + "\n"
	}

	s += helpStyle.Render("\n↑/↓: Select • ←/→: Adjust • Home/End: Min/Max • q: Quit")

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
