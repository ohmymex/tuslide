package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ohmymex/tuslide"
)

// Style definitions
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			MarginBottom(1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(1, 2)

	focusedBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("205")).
			Padding(1, 2)

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	valueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86")).
			Bold(true)
)

type model struct {
	// Slider states
	redState   *tuslide.SliderState
	greenState *tuslide.SliderState
	blueState  *tuslide.SliderState

	// Mouse states for each slider
	redMouse   *tuslide.MouseState
	greenMouse *tuslide.MouseState
	blueMouse  *tuslide.MouseState

	// Which slider is focused (for keyboard control)
	focused int

	// Track cursor position for display
	cursorX, cursorY int
}

func newModel() model {
	return model{
		redState:   tuslide.NewState(tuslide.WithValue(128), tuslide.WithMin(0), tuslide.WithMax(255)),
		greenState: tuslide.NewState(tuslide.WithValue(128), tuslide.WithMin(0), tuslide.WithMax(255)),
		blueState:  tuslide.NewState(tuslide.WithValue(128), tuslide.WithMin(0), tuslide.WithMax(255)),
		redMouse:   tuslide.NewMouseState(),
		greenMouse: tuslide.NewMouseState(),
		blueMouse:  tuslide.NewMouseState(),
		focused:    0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			m.focused = (m.focused + 1) % 3
		case "shift+tab":
			m.focused = (m.focused - 1 + 3) % 3
		case "left", "h":
			m.currentState().Decrement()
		case "right", "l":
			m.currentState().Increment()
		case "0":
			m.currentState().SetValue(0)
		case "1":
			m.currentState().SetFromPercentage(0.25)
		case "2":
			m.currentState().SetFromPercentage(0.5)
		case "3":
			m.currentState().SetFromPercentage(0.75)
		case "4":
			m.currentState().SetValue(255)
		}

	case tea.MouseMsg:
		m.cursorX = msg.X
		m.cursorY = msg.Y

		// Handle mouse for each slider
		// The bounds are set in View() - we need to track them
		if m.redMouse.HandleMouse(msg, m.createSlider(m.redState, "205")) {
			m.focused = 0
		} else if m.greenMouse.HandleMouse(msg, m.createSlider(m.greenState, "82")) {
			m.focused = 1
		} else if m.blueMouse.HandleMouse(msg, m.createSlider(m.blueState, "39")) {
			m.focused = 2
		}
	}

	return m, nil
}

func (m model) currentState() *tuslide.SliderState {
	switch m.focused {
	case 0:
		return m.redState
	case 1:
		return m.greenState
	case 2:
		return m.blueState
	}
	return m.redState
}

func (m model) createSlider(state *tuslide.SliderState, color string) *tuslide.Slider {
	return tuslide.New(state,
		tuslide.WithWidth(30),
		tuslide.WithShowValue(true),
		tuslide.WithValueFormat("%.0f"),
		tuslide.WithFilledStyle(lipgloss.NewStyle().Foreground(lipgloss.Color(color))),
		tuslide.WithEmptyStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("238"))),
		tuslide.WithHandleStyle(lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Bold(true)),
	)
}

func (m model) View() string {
	var b strings.Builder

	// Title
	b.WriteString(titleStyle.Render("🖱️  Mouse-Enabled RGB Color Mixer"))
	b.WriteString("\n\n")

	// Instructions
	b.WriteString(helpStyle.Render("Click and drag sliders to change values • Tab to switch • Arrow keys to adjust"))
	b.WriteString("\n\n")

	// Calculate current color
	r := int(m.redState.Value())
	g := int(m.greenState.Value())
	bl := int(m.blueState.Value())

	// Color preview
	colorHex := fmt.Sprintf("#%02X%02X%02X", r, g, bl)
	colorPreview := lipgloss.NewStyle().
		Width(20).
		Height(3).
		Background(lipgloss.Color(colorHex)).
		Align(lipgloss.Center, lipgloss.Center).
		Render(colorHex)

	// Build sliders
	// Red slider
	redSlider := tuslide.New(m.redState,
		tuslide.WithWidth(30),
		tuslide.WithShowValue(true),
		tuslide.WithValueFormat("%.0f"),
		tuslide.WithLabel("Red"),
		tuslide.WithLabelPosition(tuslide.LabelLeft),
		tuslide.WithFilledStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("205"))),
		tuslide.WithEmptyStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("238"))),
		tuslide.WithHandleStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)),
		tuslide.WithLabelStyle(labelStyle),
		tuslide.WithValueStyle(valueStyle),
	)

	// Green slider
	greenSlider := tuslide.New(m.greenState,
		tuslide.WithWidth(30),
		tuslide.WithShowValue(true),
		tuslide.WithValueFormat("%.0f"),
		tuslide.WithLabel("Green"),
		tuslide.WithLabelPosition(tuslide.LabelLeft),
		tuslide.WithFilledStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("82"))),
		tuslide.WithEmptyStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("238"))),
		tuslide.WithHandleStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)),
		tuslide.WithLabelStyle(labelStyle),
		tuslide.WithValueStyle(valueStyle),
	)

	// Blue slider
	blueSlider := tuslide.New(m.blueState,
		tuslide.WithWidth(30),
		tuslide.WithShowValue(true),
		tuslide.WithValueFormat("%.0f"),
		tuslide.WithLabel("Blue"),
		tuslide.WithLabelPosition(tuslide.LabelLeft),
		tuslide.WithFilledStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("39"))),
		tuslide.WithEmptyStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("238"))),
		tuslide.WithHandleStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)),
		tuslide.WithLabelStyle(labelStyle),
		tuslide.WithValueStyle(valueStyle),
	)

	// Render sliders with focus indication
	redBox := boxStyle
	greenBox := boxStyle
	blueBox := boxStyle

	switch m.focused {
	case 0:
		redBox = focusedBoxStyle
	case 1:
		greenBox = focusedBoxStyle
	case 2:
		blueBox = focusedBoxStyle
	}

	// Update mouse bounds based on approximate positions
	// In a real app, you'd calculate these more precisely
	m.redMouse.SetBounds(10, 5, 30, 1)
	m.greenMouse.SetBounds(10, 8, 30, 1)
	m.blueMouse.SetBounds(10, 11, 30, 1)

	// Layout
	sliders := lipgloss.JoinVertical(lipgloss.Left,
		redBox.Render(redSlider.View()),
		greenBox.Render(greenSlider.View()),
		blueBox.Render(blueSlider.View()),
	)

	// Side by side: sliders and color preview
	content := lipgloss.JoinHorizontal(lipgloss.Top,
		sliders,
		"  ",
		lipgloss.NewStyle().MarginTop(2).Render(colorPreview),
	)

	b.WriteString(content)
	b.WriteString("\n\n")

	// Debug info
	b.WriteString(helpStyle.Render(fmt.Sprintf("Mouse: (%d, %d) • Focused: %s",
		m.cursorX, m.cursorY,
		[]string{"Red", "Green", "Blue"}[m.focused])))
	b.WriteString("\n")
	b.WriteString(helpStyle.Render("Press 'q' to quit"))

	return b.String()
}

func main() {
	// Enable mouse support
	p := tea.NewProgram(
		newModel(),
		tuslide.EnableMouse(), // Enable mouse tracking
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
