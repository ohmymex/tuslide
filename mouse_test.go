package tuslide

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewMouseState(t *testing.T) {
	ms := NewMouseState()
	if ms == nil {
		t.Fatal("NewMouseState returned nil")
	}
	if ms.Dragging {
		t.Error("New mouse state should not be dragging")
	}
	if ms.Focused {
		t.Error("New mouse state should not be focused")
	}
}

func TestMouseState_SetBounds(t *testing.T) {
	ms := NewMouseState()
	ms.SetBounds(10, 20, 100, 50)

	if ms.X != 10 {
		t.Errorf("Expected X=10, got %d", ms.X)
	}
	if ms.Y != 20 {
		t.Errorf("Expected Y=20, got %d", ms.Y)
	}
	if ms.Width != 100 {
		t.Errorf("Expected Width=100, got %d", ms.Width)
	}
	if ms.Height != 50 {
		t.Errorf("Expected Height=50, got %d", ms.Height)
	}
}

func TestMouseState_Contains(t *testing.T) {
	ms := NewMouseState()
	ms.SetBounds(10, 20, 100, 50)

	tests := []struct {
		name     string
		x, y     int
		expected bool
	}{
		{"inside", 50, 40, true},
		{"top-left corner", 10, 20, true},
		{"bottom-right edge", 109, 69, true},
		{"outside left", 5, 40, false},
		{"outside right", 115, 40, false},
		{"outside top", 50, 15, false},
		{"outside bottom", 50, 75, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ms.Contains(tt.x, tt.y)
			if result != tt.expected {
				t.Errorf("Contains(%d, %d) = %v, expected %v", tt.x, tt.y, result, tt.expected)
			}
		})
	}
}

func TestMouseState_HandleMouse_Click(t *testing.T) {
	ms := NewMouseState()
	ms.SetBounds(0, 0, 100, 10)

	state := NewState(WithMin(0), WithMax(100), WithValue(50))
	slider := New(state, WithWidth(100))

	// Click in the middle of the slider
	msg := tea.MouseMsg{
		X:      50,
		Y:      5,
		Button: tea.MouseButtonLeft,
		Action: tea.MouseActionPress,
	}

	handled := ms.HandleMouse(msg, slider)

	if !handled {
		t.Error("Click should have been handled")
	}
	if !ms.Dragging {
		t.Error("Should be dragging after click")
	}
	if !ms.Focused {
		t.Error("Should be focused after click")
	}

	// Value should be approximately 50% (50 in 0-100 range)
	pct := state.Percentage()
	if pct < 0.45 || pct > 0.55 {
		t.Errorf("Expected percentage around 0.5, got %f", pct)
	}
}

func TestMouseState_HandleMouse_Drag(t *testing.T) {
	ms := NewMouseState()
	ms.SetBounds(0, 0, 100, 10)

	state := NewState(WithMin(0), WithMax(100), WithValue(0))
	slider := New(state, WithWidth(100))

	// Start drag
	ms.Dragging = true

	// Drag to 75%
	msg := tea.MouseMsg{
		X:      75,
		Y:      5,
		Action: tea.MouseActionMotion,
	}

	handled := ms.HandleMouse(msg, slider)

	if !handled {
		t.Error("Drag should have been handled")
	}

	pct := state.Percentage()
	if pct < 0.70 || pct > 0.80 {
		t.Errorf("Expected percentage around 0.75, got %f", pct)
	}
}

func TestMouseState_HandleMouse_Release(t *testing.T) {
	ms := NewMouseState()
	ms.SetBounds(0, 0, 100, 10)
	ms.Dragging = true

	state := NewState(WithMin(0), WithMax(100), WithValue(0))
	slider := New(state, WithWidth(100))

	msg := tea.MouseMsg{
		X:      25,
		Y:      5,
		Action: tea.MouseActionRelease,
	}

	handled := ms.HandleMouse(msg, slider)

	if !handled {
		t.Error("Release should have been handled")
	}
	if ms.Dragging {
		t.Error("Should not be dragging after release")
	}
}

func TestMouseState_HandleMouse_ClickOutside(t *testing.T) {
	ms := NewMouseState()
	ms.SetBounds(0, 0, 100, 10)

	state := NewState(WithMin(0), WithMax(100), WithValue(50))
	slider := New(state, WithWidth(100))

	// Click outside bounds
	msg := tea.MouseMsg{
		X:      150,
		Y:      5,
		Button: tea.MouseButtonLeft,
		Action: tea.MouseActionPress,
	}

	handled := ms.HandleMouse(msg, slider)

	if handled {
		t.Error("Click outside should not be handled")
	}
	if ms.Dragging {
		t.Error("Should not be dragging after click outside")
	}
}

func TestMouseState_HandleMouse_VerticalSlider(t *testing.T) {
	ms := NewMouseState()
	ms.SetBounds(0, 0, 10, 100)

	state := NewState(WithMin(0), WithMax(100), WithValue(50))
	slider := New(state, WithWidth(10), WithHeight(100), WithOrientation(Vertical))

	// Click at top (should be 100%)
	msg := tea.MouseMsg{
		X:      5,
		Y:      0,
		Button: tea.MouseButtonLeft,
		Action: tea.MouseActionPress,
	}

	ms.HandleMouse(msg, slider)

	pct := state.Percentage()
	if pct < 0.95 {
		t.Errorf("Click at top should give ~100%%, got %f", pct)
	}

	// Click at bottom (should be 0%)
	// Note: Y=99 is the last valid row inside bounds (0-99 for height 100)
	msg.Y = 99
	ms.HandleMouse(msg, slider)

	pct = state.Percentage()
	if pct > 0.05 {
		t.Errorf("Click at bottom should give ~0%%, got %f", pct)
	}
}

func TestMouseState_HandleMouse_NilSlider(t *testing.T) {
	ms := NewMouseState()
	ms.SetBounds(0, 0, 100, 10)

	msg := tea.MouseMsg{
		X:      50,
		Y:      5,
		Button: tea.MouseButtonLeft,
		Action: tea.MouseActionPress,
	}

	// Should not panic with nil slider
	handled := ms.HandleMouse(msg, nil)
	if handled {
		t.Error("Should not handle mouse with nil slider")
	}
}

func TestMouseState_HandleMouse_NilState(t *testing.T) {
	ms := NewMouseState()
	ms.SetBounds(0, 0, 100, 10)

	slider := &Slider{state: nil}

	msg := tea.MouseMsg{
		X:      50,
		Y:      5,
		Button: tea.MouseButtonLeft,
		Action: tea.MouseActionPress,
	}

	// Should not panic with nil state
	handled := ms.HandleMouse(msg, slider)
	if handled {
		t.Error("Should not handle mouse with nil state")
	}
}

func TestSliderGroup_Add(t *testing.T) {
	group := NewSliderGroup()

	state1 := NewState()
	slider1 := New(state1)
	idx1 := group.Add(slider1)

	state2 := NewState()
	slider2 := New(state2)
	idx2 := group.Add(slider2)

	if idx1 != 0 {
		t.Errorf("First slider should have index 0, got %d", idx1)
	}
	if idx2 != 1 {
		t.Errorf("Second slider should have index 1, got %d", idx2)
	}
	if group.Count() != 2 {
		t.Errorf("Expected count 2, got %d", group.Count())
	}
}

func TestSliderGroup_Get(t *testing.T) {
	group := NewSliderGroup()

	state := NewState()
	slider := New(state)
	idx := group.Add(slider)

	got := group.Get(idx)
	if got != slider {
		t.Error("Get should return the same slider that was added")
	}

	// Out of bounds
	if group.Get(-1) != nil {
		t.Error("Get(-1) should return nil")
	}
	if group.Get(100) != nil {
		t.Error("Get(100) should return nil")
	}
}

func TestSliderGroup_GetMouseState(t *testing.T) {
	group := NewSliderGroup()

	state := NewState()
	slider := New(state)
	idx := group.Add(slider)

	ms := group.GetMouseState(idx)
	if ms == nil {
		t.Error("GetMouseState should return a mouse state")
	}

	// Out of bounds
	if group.GetMouseState(-1) != nil {
		t.Error("GetMouseState(-1) should return nil")
	}
	if group.GetMouseState(100) != nil {
		t.Error("GetMouseState(100) should return nil")
	}
}

func TestSliderGroup_Focus(t *testing.T) {
	group := NewSliderGroup()

	if group.Focused() != -1 {
		t.Error("Initial focus should be -1")
	}

	state := NewState()
	slider := New(state)
	group.Add(slider)

	group.SetFocused(0)
	if group.Focused() != 0 {
		t.Errorf("Expected focused 0, got %d", group.Focused())
	}

	// Out of bounds should be ignored
	group.SetFocused(100)
	if group.Focused() != 0 {
		t.Error("Out of bounds focus should be ignored")
	}

	group.SetFocused(-1)
	if group.Focused() != -1 {
		t.Error("Should be able to set focus to -1")
	}
}

func TestSliderGroup_SetBounds(t *testing.T) {
	group := NewSliderGroup()

	state := NewState()
	slider := New(state)
	idx := group.Add(slider)

	group.SetBounds(idx, 10, 20, 100, 50)

	ms := group.GetMouseState(idx)
	if ms.X != 10 || ms.Y != 20 || ms.Width != 100 || ms.Height != 50 {
		t.Error("SetBounds did not set bounds correctly")
	}

	// Out of bounds should not panic
	group.SetBounds(100, 0, 0, 0, 0)
}

func TestSliderGroup_HandleMouse(t *testing.T) {
	group := NewSliderGroup()

	state1 := NewState(WithMin(0), WithMax(100))
	slider1 := New(state1, WithWidth(100))
	group.Add(slider1)
	group.SetBounds(0, 0, 0, 100, 10)

	state2 := NewState(WithMin(0), WithMax(100))
	slider2 := New(state2, WithWidth(100))
	group.Add(slider2)
	group.SetBounds(1, 0, 20, 100, 10)

	// Click on first slider
	msg := tea.MouseMsg{
		X:      50,
		Y:      5,
		Button: tea.MouseButtonLeft,
		Action: tea.MouseActionPress,
	}

	handled := group.HandleMouse(msg)
	if !handled {
		t.Error("Click on slider should be handled")
	}
	if group.Focused() != 0 {
		t.Errorf("Expected focus on slider 0, got %d", group.Focused())
	}

	// Release
	msg.Action = tea.MouseActionRelease
	group.HandleMouse(msg)

	// Click on second slider
	msg.Y = 25
	msg.Action = tea.MouseActionPress
	handled = group.HandleMouse(msg)
	if !handled {
		t.Error("Click on second slider should be handled")
	}
	if group.Focused() != 1 {
		t.Errorf("Expected focus on slider 1, got %d", group.Focused())
	}
}

func TestEnableMouse(t *testing.T) {
	opt := EnableMouse()
	if opt == nil {
		t.Error("EnableMouse should return a non-nil option")
	}
}

func TestEnableMouseAllMotion(t *testing.T) {
	opt := EnableMouseAllMotion()
	if opt == nil {
		t.Error("EnableMouseAllMotion should return a non-nil option")
	}
}
