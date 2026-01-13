// Example showcase demonstrates all TuSlide features in a comprehensive interactive demo.
//
// Run with: go run ./examples/showcase
//
// Controls:
//   - Tab/Shift+Tab: Navigate between pages
//   - Up/Down or J/K: Navigate between sliders
//   - Left/Right or H/L: Adjust slider values
//   - 1-9: Jump to specific slider on current page
//   - Space: Toggle animation (on progress page)
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

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			Background(lipgloss.Color("235")).
			Padding(0, 2)

	pageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("39")).
			Bold(true)

	sectionStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("243")).
			MarginTop(1)

	focusedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			MarginTop(1)

	dimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))
)

type tickMsg time.Time

type page int

const (
	pageHorizontal page = iota
	pageVertical
	pageStyles
	pageSegmented
	pageBorders
	pageProgress
	totalPages
)

func (p page) String() string {
	switch p {
	case pageHorizontal:
		return "Horizontal Sliders"
	case pageVertical:
		return "Vertical Sliders"
	case pageStyles:
		return "Style Presets"
	case pageSegmented:
		return "Segmented Styles"
	case pageBorders:
		return "Border Styles"
	case pageProgress:
		return "Progress Bars"
	default:
		return "Unknown"
	}
}

// model holds the application state.
type model struct {
	currentPage page
	focusIndex  int
	animating   bool
	quitting    bool

	// Shared states for sliders
	states []*tuslide.SliderState
}

func initialModel() model {
	// Create states for all demos
	// Total needed: Horizontal(4) + Vertical(3) + Styles(5) + Segmented(4) + Borders(4) + Progress(5) = 25
	states := make([]*tuslide.SliderState, 25)
	for i := range states {
		states[i] = tuslide.NewState(
			tuslide.WithMin(0),
			tuslide.WithMax(100),
			tuslide.WithValue(float64(30+(i%10)*5)),
			tuslide.WithStep(5),
		)
	}

	return model{
		currentPage: pageHorizontal,
		focusIndex:  0,
		states:      states,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*80, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if m.animating && m.currentPage == pageProgress {
			// Animate progress bars (states 20-24)
			for i := 20; i < 25; i++ {
				if i < len(m.states) {
					current := m.states[i].Value()
					if current >= 100 {
						m.states[i].SetValue(0)
					} else {
						m.states[i].SetValue(current + 2)
					}
				}
			}
			return m, tickCmd()
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "tab":
			m.currentPage = (m.currentPage + 1) % totalPages
			m.focusIndex = 0

		case "shift+tab":
			m.currentPage--
			if m.currentPage < 0 {
				m.currentPage = totalPages - 1
			}
			m.focusIndex = 0

		case "up", "k":
			m.focusIndex--
			if m.focusIndex < 0 {
				m.focusIndex = m.getPageSliderCount() - 1
			}

		case "down", "j":
			m.focusIndex = (m.focusIndex + 1) % m.getPageSliderCount()

		case "left", "h":
			m.adjustFocused(-1)

		case "right", "l":
			m.adjustFocused(1)

		case "space":
			if m.currentPage == pageProgress {
				m.animating = !m.animating
				if m.animating {
					return m, tickCmd()
				}
			}

		case "home":
			m.getActiveState().SetFromPercentage(0)

		case "end":
			m.getActiveState().SetFromPercentage(1)

		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			idx := int(msg.String()[0] - '1')
			if idx < m.getPageSliderCount() {
				m.focusIndex = idx
			}
		}
	}

	return m, nil
}

func (m model) getPageSliderCount() int {
	switch m.currentPage {
	case pageHorizontal:
		return 4
	case pageVertical:
		return 3
	case pageStyles:
		return 5
	case pageSegmented:
		return 4
	case pageBorders:
		return 4
	case pageProgress:
		return 5
	default:
		return 1
	}
}

// getPageBaseIndex returns the starting state index for each page.
// Each page has its own dedicated range of states to avoid conflicts.
func (m model) getPageBaseIndex() int {
	switch m.currentPage {
	case pageHorizontal:
		return 0 // states 0-3 (4 sliders)
	case pageVertical:
		return 4 // states 4-6 (3 sliders)
	case pageStyles:
		return 7 // states 7-11 (5 sliders)
	case pageSegmented:
		return 12 // states 12-15 (4 sliders)
	case pageBorders:
		return 16 // states 16-19 (4 sliders)
	case pageProgress:
		return 20 // states 20-24 (5 sliders)
	default:
		return 0
	}
}

func (m model) getActiveState() *tuslide.SliderState {
	baseIdx := m.getPageBaseIndex()
	idx := baseIdx + m.focusIndex
	if idx < len(m.states) {
		return m.states[idx]
	}
	return m.states[0]
}

func (m *model) adjustFocused(direction int) {
	state := m.getActiveState()
	if direction > 0 {
		state.Increment()
	} else {
		state.Decrement()
	}
}

func (m model) View() string {
	if m.quitting {
		return "Thanks for exploring TuSlide!\n"
	}

	var b strings.Builder

	// Title bar
	b.WriteString(titleStyle.Render("🎚️  TuSlide Feature Showcase"))
	b.WriteString("\n\n")

	// Page indicator
	var pages []string
	for i := page(0); i < totalPages; i++ {
		if i == m.currentPage {
			pages = append(pages, pageStyle.Render(fmt.Sprintf("[%s]", i.String())))
		} else {
			pages = append(pages, dimStyle.Render(i.String()))
		}
	}
	b.WriteString(strings.Join(pages, " • "))
	b.WriteString("\n\n")

	// Render current page
	switch m.currentPage {
	case pageHorizontal:
		b.WriteString(m.renderHorizontalPage())
	case pageVertical:
		b.WriteString(m.renderVerticalPage())
	case pageStyles:
		b.WriteString(m.renderStylesPage())
	case pageSegmented:
		b.WriteString(m.renderSegmentedPage())
	case pageBorders:
		b.WriteString(m.renderBordersPage())
	case pageProgress:
		b.WriteString(m.renderProgressPage())
	}

	// Help
	help := "[Tab] Page • [↑↓] Select • [←→] Adjust • [Home/End] Min/Max"
	if m.currentPage == pageProgress {
		help += " • [Space] Animate"
	}
	help += " • [q] Quit"
	b.WriteString(helpStyle.Render("\n" + help))

	return b.String()
}

func (m model) renderHorizontalPage() string {
	var b strings.Builder
	b.WriteString(sectionStyle.Render("Horizontal sliders with various label positions and symbol sets"))
	b.WriteString("\n\n")

	sliders := []struct {
		name   string
		slider *tuslide.Slider
	}{
		{
			"Default Style (Label Left)",
			tuslide.New(m.states[0],
				tuslide.WithWidth(30),
				tuslide.WithLabel("Volume"),
				tuslide.WithLabelPosition(tuslide.LabelLeft),
				tuslide.WithShowValue(true),
				tuslide.WithStyle(tuslide.StyleDefault()),
			),
		},
		{
			"Block Style (Label Top)",
			tuslide.New(m.states[1],
				tuslide.WithWidth(30),
				tuslide.WithLabel("Brightness"),
				tuslide.WithLabelPosition(tuslide.LabelTop),
				tuslide.WithShowValue(true),
				tuslide.WithValuePosition(tuslide.ValueBottom),
				tuslide.WithStyle(tuslide.StyleBlock()),
			),
		},
		{
			"Progress Style (Label Right)",
			tuslide.New(m.states[2],
				tuslide.WithWidth(30),
				tuslide.WithLabel("Bass"),
				tuslide.WithLabelPosition(tuslide.LabelRight),
				tuslide.WithShowValue(true),
				tuslide.WithValuePosition(tuslide.ValueLeft),
				tuslide.WithStyle(tuslide.StyleProgress()),
			),
		},
		{
			"Wave Style (Label Bottom)",
			tuslide.New(m.states[3],
				tuslide.WithWidth(30),
				tuslide.WithLabel("Treble"),
				tuslide.WithLabelPosition(tuslide.LabelBottom),
				tuslide.WithShowValue(true),
				tuslide.WithStyle(tuslide.StyleWave()),
			),
		},
	}

	for i, s := range sliders {
		indicator := "  "
		if i == m.focusIndex {
			indicator = focusedStyle.Render("▸ ")
		}
		b.WriteString(dimStyle.Render(fmt.Sprintf("%d. %s", i+1, s.name)))
		b.WriteString("\n")
		b.WriteString(indicator)
		b.WriteString(s.slider.View())
		b.WriteString("\n\n")
	}

	return b.String()
}

func (m model) renderVerticalPage() string {
	var b strings.Builder
	b.WriteString(sectionStyle.Render("Vertical sliders - perfect for equalizers and level meters"))
	b.WriteString("\n\n")

	// Vertical page uses states 4-6
	// Create vertical sliders
	sliders := []*tuslide.Slider{
		tuslide.New(m.states[4],
			tuslide.WithHeight(8),
			tuslide.WithOrientation(tuslide.Vertical),
			tuslide.WithLabel("Bass"),
			tuslide.WithLabelPosition(tuslide.LabelBottom),
			tuslide.WithShowValue(true),
			tuslide.WithStyle(tuslide.StyleNeon()),
		),
		tuslide.New(m.states[5],
			tuslide.WithHeight(8),
			tuslide.WithOrientation(tuslide.Vertical),
			tuslide.WithLabel("Mid"),
			tuslide.WithLabelPosition(tuslide.LabelBottom),
			tuslide.WithShowValue(true),
			tuslide.WithStyle(tuslide.StyleOcean()),
		),
		tuslide.New(m.states[6],
			tuslide.WithHeight(8),
			tuslide.WithOrientation(tuslide.Vertical),
			tuslide.WithLabel("Treble"),
			tuslide.WithLabelPosition(tuslide.LabelBottom),
			tuslide.WithShowValue(true),
			tuslide.WithStyle(tuslide.StyleSunset()),
		),
	}

	// Render side by side
	views := make([][]string, len(sliders))
	maxLines := 0
	for i, s := range sliders {
		views[i] = strings.Split(s.View(), "\n")
		if len(views[i]) > maxLines {
			maxLines = len(views[i])
		}
	}

	for line := 0; line < maxLines; line++ {
		for i := range sliders {
			// Focus indicator
			if line == maxLines/2 {
				if i == m.focusIndex {
					b.WriteString(focusedStyle.Render("▸"))
				} else {
					b.WriteString(" ")
				}
			} else {
				b.WriteString(" ")
			}

			if line < len(views[i]) {
				b.WriteString(fmt.Sprintf("%-15s", views[i][line]))
			} else {
				b.WriteString(strings.Repeat(" ", 15))
			}
		}
		b.WriteString("\n")
	}

	b.WriteString(dimStyle.Render("\nUse ↑/↓ to increase/decrease vertical slider values"))

	return b.String()
}

func (m model) renderStylesPage() string {
	var b strings.Builder
	b.WriteString(sectionStyle.Render("Predefined style presets - apply complete visual themes"))
	b.WriteString("\n\n")

	// Styles page uses states 7-11 (5 sliders)
	styles := []struct {
		name  string
		style tuslide.SliderStyle
	}{
		{"Ocean Theme", tuslide.StyleOcean()},
		{"Forest Theme", tuslide.StyleForest()},
		{"Sunset Theme", tuslide.StyleSunset()},
		{"Neon Theme", tuslide.StyleNeon()},
		{"Monochrome Theme", tuslide.StyleMonochrome()},
	}

	for i, s := range styles {
		stateIdx := 7 + i // states 7-11

		slider := tuslide.New(m.states[stateIdx],
			tuslide.WithWidth(35),
			tuslide.WithLabel(s.name),
			tuslide.WithLabelPosition(tuslide.LabelLeft),
			tuslide.WithShowValue(true),
			tuslide.WithStyle(s.style),
		)

		indicator := "  "
		if i == m.focusIndex {
			indicator = focusedStyle.Render("▸ ")
		}
		b.WriteString(indicator)
		b.WriteString(slider.View())
		b.WriteString("\n\n")
	}

	return b.String()
}

func (m model) renderSegmentedPage() string {
	var b strings.Builder
	b.WriteString(sectionStyle.Render("Segmented sliders - discrete segments with gaps"))
	b.WriteString("\n\n")

	// Segmented page uses states 12-15 (4 sliders)
	segmentedStyles := []struct {
		name  string
		style tuslide.SliderStyle
	}{
		{"Segmented Dots", tuslide.StyleSegmentedDots()},
		{"Segmented Stars", tuslide.StyleSegmentedStars()},
		{"Segmented Squares", tuslide.StyleSegmentedSquares()},
		{"Segmented Diamonds", tuslide.StyleSegmentedDiamonds()},
	}

	for i, s := range segmentedStyles {
		stateIdx := 12 + i // states 12-15

		slider := tuslide.New(m.states[stateIdx],
			tuslide.WithWidth(40),
			tuslide.WithLabel(s.name),
			tuslide.WithLabelPosition(tuslide.LabelLeft),
			tuslide.WithShowValue(true),
			tuslide.WithStyle(s.style),
			tuslide.WithSegmentCount(10),
		)

		indicator := "  "
		if i == m.focusIndex {
			indicator = focusedStyle.Render("▸ ")
		}
		b.WriteString(indicator)
		b.WriteString(slider.View())
		b.WriteString("\n\n")
	}

	return b.String()
}

func (m model) renderBordersPage() string {
	var b strings.Builder
	b.WriteString(sectionStyle.Render("Border styles - wrap sliders with decorative borders"))
	b.WriteString("\n\n")

	// Borders page uses states 16-19 (4 sliders)
	borders := []struct {
		name   string
		border tuslide.BorderStyle
		color  lipgloss.Color
	}{
		{"Rounded Border", tuslide.BorderRounded, lipgloss.Color("86")},
		{"Normal Border", tuslide.BorderNormal, lipgloss.Color("213")},
		{"Thick Border", tuslide.BorderThick, lipgloss.Color("220")},
		{"Double Border", tuslide.BorderDouble, lipgloss.Color("39")},
	}

	for i, b_ := range borders {
		stateIdx := 16 + i // states 16-19

		slider := tuslide.New(m.states[stateIdx],
			tuslide.WithWidth(25),
			tuslide.WithShowValue(true),
			tuslide.WithStyle(tuslide.StyleDefault()),
			tuslide.WithBorder(b_.border),
			tuslide.WithBorderColor(b_.color),
		)

		// For bordered sliders, add indicator to the label line instead
		indicator := "  "
		if i == m.focusIndex {
			indicator = focusedStyle.Render("▸ ")
		}
		b.WriteString(indicator)
		b.WriteString(dimStyle.Render(fmt.Sprintf("%d. %s", i+1, b_.name)))
		b.WriteString("\n")
		// Add proper indentation to align with indicator
		sliderView := slider.View()
		lines := strings.Split(sliderView, "\n")
		for j, line := range lines {
			b.WriteString("  ") // Indent to align with indicator space
			b.WriteString(line)
			if j < len(lines)-1 {
				b.WriteString("\n")
			}
		}
		b.WriteString("\n\n")
	}

	return b.String()
}

func (m model) renderProgressPage() string {
	var b strings.Builder
	b.WriteString(sectionStyle.Render("Progress bars - sliders without handles for status indication"))
	b.WriteString("\n\n")

	// Progress page uses states 20-24 (5 sliders)
	progressStyles := []struct {
		name  string
		style tuslide.SliderStyle
	}{
		{"Download Progress", tuslide.StyleProgressDownload()},
		{"Upload Progress", tuslide.StyleProgressUpload()},
		{"Health Bar", tuslide.StyleHealth()},
		{"Mana Bar", tuslide.StyleMana()},
		{"Experience Bar", tuslide.StyleExperience()},
	}

	for i, s := range progressStyles {
		stateIdx := 20 + i // states 20-24

		slider := tuslide.New(m.states[stateIdx],
			tuslide.WithWidth(35),
			tuslide.WithLabel(s.name),
			tuslide.WithLabelPosition(tuslide.LabelLeft),
			tuslide.WithShowValue(true),
			tuslide.WithHandle(false), // No handle for progress bars
			tuslide.WithStyle(s.style),
		)

		indicator := "  "
		if i == m.focusIndex {
			indicator = focusedStyle.Render("▸ ")
		}
		b.WriteString(indicator)
		b.WriteString(slider.View())
		b.WriteString("\n")
	}

	status := "paused"
	if m.animating {
		status = "running"
	}
	b.WriteString(fmt.Sprintf("\n%s [%s]", dimStyle.Render("Animation:"), status))

	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
