package tuslide

import (
	"github.com/charmbracelet/bubbletea"
)

// MouseState tracks mouse interaction state for a slider.
type MouseState struct {
	// Bounding box of the slider track (for hit testing)
	X, Y          int
	Width, Height int

	// Interaction state
	Dragging bool
	Focused  bool
}

// NewMouseState creates a new mouse state.
func NewMouseState() *MouseState {
	return &MouseState{}
}

// SetBounds sets the bounding box for mouse hit testing.
// Call this after rendering to update the clickable area.
func (m *MouseState) SetBounds(x, y, width, height int) {
	m.X = x
	m.Y = y
	m.Width = width
	m.Height = height
}

// Contains checks if a point is within the slider bounds.
func (m *MouseState) Contains(x, y int) bool {
	return x >= m.X && x < m.X+m.Width &&
		y >= m.Y && y < m.Y+m.Height
}

// HandleMouse processes a mouse event and returns true if the slider was interacted with.
// It also updates the slider state value based on click/drag position.
func (m *MouseState) HandleMouse(msg tea.MouseMsg, slider *Slider) bool {
	if slider == nil || slider.state == nil {
		return false
	}

	switch msg.Action {
	case tea.MouseActionPress:
		if msg.Button == tea.MouseButtonLeft && m.Contains(msg.X, msg.Y) {
			m.Dragging = true
			m.Focused = true
			m.updateValue(msg.X, msg.Y, slider)
			return true
		}

	case tea.MouseActionMotion:
		if m.Dragging {
			m.updateValue(msg.X, msg.Y, slider)
			return true
		}

	case tea.MouseActionRelease:
		if m.Dragging {
			m.Dragging = false
			m.updateValue(msg.X, msg.Y, slider)
			return true
		}
	}

	return false
}

// updateValue updates the slider value based on mouse position.
func (m *MouseState) updateValue(mouseX, mouseY int, slider *Slider) {
	var percentage float64

	switch slider.orientation {
	case Horizontal:
		// Calculate percentage based on X position within track
		if m.Width > 0 {
			relX := mouseX - m.X
			if relX < 0 {
				relX = 0
			}
			if relX > m.Width {
				relX = m.Width
			}
			percentage = float64(relX) / float64(m.Width)
		}

	case Vertical:
		// Calculate percentage based on Y position within track
		// Note: Y increases downward, but slider fills from bottom
		if m.Height > 0 {
			relY := mouseY - m.Y
			if relY < 0 {
				relY = 0
			}
			if relY > m.Height {
				relY = m.Height
			}
			// Invert because Y=0 is top, but 100% should be at top
			percentage = 1.0 - (float64(relY) / float64(m.Height))
		}
	}

	// Clamp percentage
	if percentage < 0 {
		percentage = 0
	}
	if percentage > 1 {
		percentage = 1
	}

	// Update slider state
	slider.state.SetFromPercentage(percentage)
}

// MouseHandler is a helper interface for components that want
// mouse-aware sliders. Implement this in your Bubble Tea model.
type MouseHandler interface {
	// HandleSliderMouse processes mouse events for sliders.
	// Returns the updated model and whether any slider was interacted with.
	HandleSliderMouse(msg tea.MouseMsg) (tea.Model, bool)
}

// SliderGroup manages multiple sliders with mouse support.
// This is useful when you have multiple sliders in a single view.
type SliderGroup struct {
	sliders    []*Slider
	mouseState []*MouseState
	focused    int // Currently focused slider index (-1 if none)
}

// NewSliderGroup creates a new slider group.
func NewSliderGroup() *SliderGroup {
	return &SliderGroup{
		focused: -1,
	}
}

// Add adds a slider to the group and returns its index.
func (g *SliderGroup) Add(slider *Slider) int {
	idx := len(g.sliders)
	g.sliders = append(g.sliders, slider)
	g.mouseState = append(g.mouseState, NewMouseState())
	return idx
}

// Get returns the slider at the given index.
func (g *SliderGroup) Get(idx int) *Slider {
	if idx >= 0 && idx < len(g.sliders) {
		return g.sliders[idx]
	}
	return nil
}

// GetMouseState returns the mouse state for the slider at the given index.
func (g *SliderGroup) GetMouseState(idx int) *MouseState {
	if idx >= 0 && idx < len(g.mouseState) {
		return g.mouseState[idx]
	}
	return nil
}

// Count returns the number of sliders in the group.
func (g *SliderGroup) Count() int {
	return len(g.sliders)
}

// Focused returns the index of the currently focused slider, or -1 if none.
func (g *SliderGroup) Focused() int {
	return g.focused
}

// SetFocused sets the focused slider index.
func (g *SliderGroup) SetFocused(idx int) {
	if idx >= -1 && idx < len(g.sliders) {
		g.focused = idx
	}
}

// SetBounds sets the bounds for a specific slider.
func (g *SliderGroup) SetBounds(idx, x, y, width, height int) {
	if idx >= 0 && idx < len(g.mouseState) {
		g.mouseState[idx].SetBounds(x, y, width, height)
	}
}

// HandleMouse processes a mouse event for all sliders in the group.
// Returns true if any slider was interacted with.
func (g *SliderGroup) HandleMouse(msg tea.MouseMsg) bool {
	// Check if any slider is being dragged
	for i, ms := range g.mouseState {
		if ms.Dragging {
			if ms.HandleMouse(msg, g.sliders[i]) {
				g.focused = i
				return true
			}
		}
	}

	// Check for new clicks on any slider
	for i, ms := range g.mouseState {
		if ms.HandleMouse(msg, g.sliders[i]) {
			g.focused = i
			// Clear other slider focus
			for j := range g.mouseState {
				if j != i {
					g.mouseState[j].Focused = false
				}
			}
			return true
		}
	}

	return false
}

// EnableMouse returns a tea.ProgramOption that enables mouse support.
// Use this when creating your Bubble Tea program:
//
//	p := tea.NewProgram(model, tuslide.EnableMouse())
func EnableMouse() tea.ProgramOption {
	return tea.WithMouseCellMotion()
}

// EnableMouseAllMotion returns a tea.ProgramOption that enables mouse support
// with all motion events (more detailed tracking).
// Use this when you need continuous mouse position updates:
//
//	p := tea.NewProgram(model, tuslide.EnableMouseAllMotion())
func EnableMouseAllMotion() tea.ProgramOption {
	return tea.WithMouseAllMotion()
}
