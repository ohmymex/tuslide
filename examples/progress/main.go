// Example progress demonstrates various progress bar styles and animations.
//
// Run with: go run ./examples/progress
//
// Controls:
//   - Space: Start/pause all animations
//   - R: Reset all progress
//   - 1-6: Toggle individual progress bars
//   - +/-: Adjust animation speed
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
			Foreground(lipgloss.Color("255")).
			Background(lipgloss.Color("62")).
			Padding(0, 2)

	labelStyle = lipgloss.NewStyle().
			Width(20).
			Foreground(lipgloss.Color("252"))

	percentStyle = lipgloss.NewStyle().
			Width(6).
			Align(lipgloss.Right).
			Foreground(lipgloss.Color("250"))

	completeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("46"))

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("243"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	speedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("213"))
)

type progressBar struct {
	name    string
	state   *tuslide.SliderState
	style   tuslide.SliderStyle
	active  bool
	speed   float64 // increment per tick
}

type tickMsg time.Time

type model struct {
	bars      []progressBar
	running   bool
	tickSpeed time.Duration
	quitting  bool
}

func initialModel() model {
	return model{
		bars: []progressBar{
			{
				name:   "Downloading...",
				state:  tuslide.NewState(tuslide.WithMax(100), tuslide.WithValue(0)),
				style:  tuslide.StyleProgressDownload(),
				active: true,
				speed:  2.5,
			},
			{
				name:   "Uploading...",
				state:  tuslide.NewState(tuslide.WithMax(100), tuslide.WithValue(0)),
				style:  tuslide.StyleProgressUpload(),
				active: true,
				speed:  1.8,
			},
			{
				name:   "Installing...",
				state:  tuslide.NewState(tuslide.WithMax(100), tuslide.WithValue(0)),
				style:  tuslide.StyleProgressInstallation(),
				active: true,
				speed:  1.2,
			},
			{
				name:   "Loading assets...",
				state:  tuslide.NewState(tuslide.WithMax(100), tuslide.WithValue(0)),
				style:  tuslide.StyleProgressLoading(),
				active: true,
				speed:  3.0,
			},
			{
				name:   "Compiling...",
				state:  tuslide.NewState(tuslide.WithMax(100), tuslide.WithValue(0)),
				style: tuslide.SliderStyle{
					Name: "Compile",
					Symbols: tuslide.Symbols{
						Filled: "━",
						Empty:  "─",
						Handle: "",
					},
					FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("208")),
					EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
				},
				active: true,
				speed:  0.8,
			},
			{
				name:   "Battery charging...",
				state:  tuslide.NewState(tuslide.WithMax(100), tuslide.WithValue(0)),
				style:  tuslide.StyleProgressBattery(),
				active: true,
				speed:  1.5,
			},
		},
		running:   false,
		tickSpeed: time.Millisecond * 100,
	}
}

func (m model) Init() tea.Cmd {
	return tickCmd(m.tickSpeed)
}

func tickCmd(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if m.running {
			for i := range m.bars {
				if m.bars[i].active && m.bars[i].state.Value() < 100 {
					newVal := m.bars[i].state.Value() + m.bars[i].speed
					if newVal > 100 {
						newVal = 100
					}
					m.bars[i].state.SetValue(newVal)
				}
			}
		}
		return m, tickCmd(m.tickSpeed)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case " ":
			m.running = !m.running

		case "r":
			for i := range m.bars {
				m.bars[i].state.SetValue(0)
			}
			m.running = false

		case "1", "2", "3", "4", "5", "6":
			idx := int(msg.String()[0] - '1')
			if idx < len(m.bars) {
				m.bars[idx].active = !m.bars[idx].active
			}

		case "+", "=":
			if m.tickSpeed > time.Millisecond*20 {
				m.tickSpeed -= time.Millisecond * 20
			}

		case "-", "_":
			if m.tickSpeed < time.Millisecond*500 {
				m.tickSpeed += time.Millisecond * 20
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}

	var b strings.Builder

	// Title
	b.WriteString("\n  ")
	b.WriteString(titleStyle.Render("Progress Bar Demo"))
	b.WriteString("\n\n")

	// Status
	status := "PAUSED"
	if m.running {
		status = "RUNNING"
	}
	speedMs := m.tickSpeed.Milliseconds()
	b.WriteString(fmt.Sprintf("  Status: %s  |  Speed: %s\n\n",
		statusStyle.Render(status),
		speedStyle.Render(fmt.Sprintf("%dms/tick", speedMs))))

	// Progress bars
	for i, bar := range m.bars {
		// Number indicator
		numStyle := statusStyle
		if bar.active {
			numStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("46"))
		}
		b.WriteString(fmt.Sprintf("  %s ", numStyle.Render(fmt.Sprintf("[%d]", i+1))))

		// Label
		label := bar.name
		if bar.state.Value() >= 100 {
			label = strings.Replace(label, "...", " Complete!", 1)
		}
		
		lStyle := labelStyle
		if !bar.active {
			lStyle = lStyle.Foreground(lipgloss.Color("240"))
		}
		b.WriteString(lStyle.Render(label))

		// Progress bar
		slider := tuslide.New(bar.state,
			tuslide.WithWidth(35),
			tuslide.WithHandle(false),
			tuslide.WithStyle(bar.style),
		)
		
		if !bar.active {
			// Dim the bar if inactive
			slider = tuslide.New(bar.state,
				tuslide.WithWidth(35),
				tuslide.WithHandle(false),
				tuslide.WithSymbols(bar.style.Symbols),
				tuslide.WithFilledStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("240"))),
				tuslide.WithEmptyStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("236"))),
			)
		}
		b.WriteString(slider.View())

		// Percentage
		pct := bar.state.Value()
		pctText := fmt.Sprintf("%.0f%%", pct)
		if pct >= 100 {
			b.WriteString(" ")
			b.WriteString(completeStyle.Render(pctText))
		} else {
			b.WriteString(" ")
			b.WriteString(percentStyle.Render(pctText))
		}
		b.WriteString("\n")
	}

	// Summary
	completed := 0
	for _, bar := range m.bars {
		if bar.state.Value() >= 100 {
			completed++
		}
	}
	b.WriteString(fmt.Sprintf("\n  Completed: %d/%d", completed, len(m.bars)))
	if completed == len(m.bars) {
		b.WriteString(completeStyle.Render("  All tasks complete!"))
	}
	b.WriteString("\n")

	// Help
	b.WriteString(helpStyle.Render("\n  [Space] Start/Pause  [R] Reset  [1-6] Toggle bar  [+/-] Speed  [Q] Quit\n"))

	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
