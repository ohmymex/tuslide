package tuslide

import (
	"strings"
	"testing"
)

func TestNewAccessibleSlider(t *testing.T) {
	state := NewState(WithValue(50), WithMax(100))
	slider := New(state, WithLabel("Volume"))

	accessible := NewAccessibleSlider(slider)

	if accessible.Slider() != slider {
		t.Error("Slider() should return the underlying slider")
	}
	if accessible.State() != state {
		t.Error("State() should return the slider's state")
	}
}

func TestAccessibleSlider_HighContrastMode(t *testing.T) {
	state := NewState(WithValue(50))
	slider := New(state)

	accessible := NewAccessibleSlider(slider,
		WithAccessibilityMode(AccessibilityHighContrast),
	)

	// High contrast should apply bold to handle
	_ = accessible.View()
	// Just verify it doesn't panic and renders
}

func TestAccessibleSlider_ASCIIMode(t *testing.T) {
	state := NewState(WithValue(50))
	slider := New(state)

	accessible := NewAccessibleSlider(slider,
		WithAccessibilityMode(AccessibilityASCII),
	)

	view := accessible.View()

	// Should not contain Unicode characters (basic check)
	if strings.Contains(view, "█") || strings.Contains(view, "░") || strings.Contains(view, "●") {
		t.Error("ASCII mode should not use Unicode characters")
	}
}

func TestAccessibleSlider_ScreenReaderMode(t *testing.T) {
	state := NewState(WithValue(50), WithMax(100))
	slider := New(state, WithShowValue(false))

	accessible := NewAccessibleSlider(slider,
		WithAccessibilityMode(AccessibilityScreenReader),
	)

	// Screen reader mode should force value display
	if !accessible.Slider().showValue {
		t.Error("Screen reader mode should enable value display")
	}
}

func TestAccessibleSlider_Description(t *testing.T) {
	state := NewState(WithValue(50), WithMax(100))
	slider := New(state, WithLabel("Volume"))

	accessible := NewAccessibleSlider(slider,
		WithDescription("Custom volume control"),
	)

	desc := accessible.GetDescription()
	if desc != "Custom volume control" {
		t.Errorf("Expected custom description, got: %s", desc)
	}
}

func TestAccessibleSlider_DefaultDescription(t *testing.T) {
	state := NewState(WithValue(50), WithMax(100))
	slider := New(state, WithLabel("Volume"))

	accessible := NewAccessibleSlider(slider)

	desc := accessible.GetDescription()
	if !strings.Contains(desc, "Volume") {
		t.Errorf("Description should contain label, got: %s", desc)
	}
	if !strings.Contains(desc, "50") {
		t.Errorf("Description should contain value, got: %s", desc)
	}
	if !strings.Contains(desc, "50%") {
		t.Errorf("Description should contain percentage, got: %s", desc)
	}
}

func TestAccessibleSlider_Announcer(t *testing.T) {
	state := NewState(WithValue(50), WithMax(100), WithStep(10))
	slider := New(state)

	var announced []string
	accessible := NewAccessibleSlider(slider,
		WithAnnouncer(func(msg string) {
			announced = append(announced, msg)
		}),
	)

	// Focus should announce
	accessible.SetFocused(true)
	if len(announced) != 1 {
		t.Errorf("Expected 1 announcement on focus, got %d", len(announced))
	}

	// Increment should announce
	accessible.Increment()
	if len(announced) != 2 {
		t.Errorf("Expected 2 announcements after increment, got %d", len(announced))
	}

	// Decrement should announce
	accessible.Decrement()
	if len(announced) != 3 {
		t.Errorf("Expected 3 announcements after decrement, got %d", len(announced))
	}

	// SetValue should announce
	accessible.SetValue(75)
	if len(announced) != 4 {
		t.Errorf("Expected 4 announcements after SetValue, got %d", len(announced))
	}
}

func TestAccessibleSlider_FocusState(t *testing.T) {
	state := NewState()
	slider := New(state)
	accessible := NewAccessibleSlider(slider)

	if accessible.IsFocused() {
		t.Error("Should not be focused initially")
	}

	accessible.SetFocused(true)
	if !accessible.IsFocused() {
		t.Error("Should be focused after SetFocused(true)")
	}

	accessible.SetFocused(false)
	if accessible.IsFocused() {
		t.Error("Should not be focused after SetFocused(false)")
	}
}

func TestAccessibleSlider_ValueAnnouncement(t *testing.T) {
	state := NewState(WithValue(75), WithMax(100))
	slider := New(state)
	accessible := NewAccessibleSlider(slider)

	announcement := accessible.GetValueAnnouncement()
	if !strings.Contains(announcement, "75") {
		t.Errorf("Announcement should contain value, got: %s", announcement)
	}
	if !strings.Contains(announcement, "75 percent") {
		t.Errorf("Announcement should contain percentage, got: %s", announcement)
	}
}

func TestAccessibleSlider_FocusAnnouncement(t *testing.T) {
	state := NewState(WithValue(50), WithMax(100))
	slider := New(state, WithLabel("Volume"))
	accessible := NewAccessibleSlider(slider)

	announcement := accessible.GetFocusAnnouncement()
	if !strings.Contains(announcement, "Volume") {
		t.Errorf("Focus announcement should contain label, got: %s", announcement)
	}
	if !strings.Contains(announcement, "arrow keys") {
		t.Errorf("Focus announcement should contain usage hint, got: %s", announcement)
	}
}

func TestAccessibleSlider_BoundaryAnnouncement(t *testing.T) {
	state := NewState(WithValue(0), WithMax(100))
	slider := New(state)
	accessible := NewAccessibleSlider(slider)

	announcement := accessible.GetBoundaryAnnouncement()
	if !strings.Contains(announcement, "Minimum") {
		t.Errorf("Expected minimum boundary message, got: %s", announcement)
	}

	state.SetValue(100)
	announcement = accessible.GetBoundaryAnnouncement()
	if !strings.Contains(announcement, "Maximum") {
		t.Errorf("Expected maximum boundary message, got: %s", announcement)
	}

	state.SetValue(50)
	announcement = accessible.GetBoundaryAnnouncement()
	if announcement != "" {
		t.Error("No boundary announcement expected when not at boundary")
	}
}

func TestAccessibleSlider_FocusIndicatorInView(t *testing.T) {
	state := NewState(WithValue(50))
	slider := New(state)
	accessible := NewAccessibleSlider(slider)

	accessible.SetFocused(true)
	view := accessible.View()

	if !strings.Contains(view, "▸") {
		t.Error("Focused view should contain focus indicator")
	}

	accessible.SetFocused(false)
	view = accessible.View()

	if strings.Contains(view, "▸") {
		t.Error("Unfocused view should not contain focus indicator")
	}
}

func TestFocusIndicator(t *testing.T) {
	fi := NewFocusIndicator()

	// Test Render
	if fi.Render(false) != "" {
		t.Error("Should return empty string when not focused")
	}
	if fi.Render(true) == "" {
		t.Error("Should return indicator when focused")
	}

	// Test Wrap
	wrapped := fi.Wrap("content", true)
	if !strings.Contains(wrapped, "content") {
		t.Error("Wrap should contain the content")
	}
	if !strings.Contains(wrapped, "▸") {
		t.Error("Wrap should contain the indicator when focused")
	}

	notWrapped := fi.Wrap("content", false)
	if notWrapped != "content" {
		t.Error("Wrap should return just content when not focused")
	}
}

func TestFocusIndicator_Positions(t *testing.T) {
	fi := NewFocusIndicator()

	// Left position (default)
	wrapped := fi.Wrap("content", true)
	if !strings.Contains(wrapped, "▸") || !strings.Contains(wrapped, "content") {
		t.Error("Left position should contain indicator and content")
	}
	// Indicator should come before content
	idx := strings.Index(wrapped, "content")
	if idx == 0 {
		t.Error("Left position should have indicator before content")
	}

	// Right position
	fi.WithPosition("right")
	wrapped = fi.Wrap("content", true)
	if !strings.Contains(wrapped, "▸") || !strings.Contains(wrapped, "content") {
		t.Error("Right position should contain indicator and content")
	}

	// Both positions
	fi.WithPosition("both")
	wrapped = fi.Wrap("content", true)
	// Should contain content and multiple indicators
	if !strings.Contains(wrapped, "content") {
		t.Error("Both position should contain content")
	}
	// Count indicators (the char appears twice)
	if strings.Count(wrapped, "▸") < 2 {
		t.Error("Both position should have indicators on both sides")
	}
}

func TestFocusIndicator_CustomChar(t *testing.T) {
	fi := NewFocusIndicator().WithChar(">>")

	indicator := fi.Render(true)
	if !strings.Contains(indicator, ">>") {
		t.Error("Should use custom character")
	}
}

func TestHighContrastPalettes(t *testing.T) {
	// Just verify they return non-empty palettes
	palettes := []HighContrastPalette{
		DefaultHighContrastPalette(),
		DarkHighContrastPalette(),
		LightHighContrastPalette(),
	}

	for i, p := range palettes {
		if p.Foreground == "" {
			t.Errorf("Palette %d has empty foreground", i)
		}
		if p.Background == "" {
			t.Errorf("Palette %d has empty background", i)
		}
		if p.Accent == "" {
			t.Errorf("Palette %d has empty accent", i)
		}
	}
}

func TestApplyPalette(t *testing.T) {
	state := NewState()
	slider := New(state)
	palette := DefaultHighContrastPalette()

	ApplyPalette(slider, palette)

	// Just verify it doesn't panic
	_ = slider.View()
}

func TestKeyboardHints(t *testing.T) {
	hints := DefaultKeyboardHints()

	full := hints.Render()
	if !strings.Contains(full, "Increase") {
		t.Error("Full render should contain 'Increase'")
	}
	if !strings.Contains(full, "Decrease") {
		t.Error("Full render should contain 'Decrease'")
	}

	compact := hints.RenderCompact()
	if len(compact) > 30 {
		t.Error("Compact render should be short")
	}
}

func TestProgressAnnouncer(t *testing.T) {
	var announcements []string
	announcer := NewProgressAnnouncer(25, func(msg string) {
		announcements = append(announcements, msg)
	})

	state := NewState(WithValue(0), WithMax(100))

	// 0% - no announcement
	announcer.Update(state)
	if len(announcements) != 0 {
		t.Errorf("Expected 0 announcements at 0%%, got %d", len(announcements))
	}

	// 25%
	state.SetValue(25)
	announcer.Update(state)
	if len(announcements) != 1 {
		t.Errorf("Expected 1 announcement at 25%%, got %d", len(announcements))
	}

	// 30% - no new announcement (interval is 25)
	state.SetValue(30)
	announcer.Update(state)
	if len(announcements) != 1 {
		t.Errorf("Expected still 1 announcement at 30%%, got %d", len(announcements))
	}

	// 50%
	state.SetValue(50)
	announcer.Update(state)
	if len(announcements) != 2 {
		t.Errorf("Expected 2 announcements at 50%%, got %d", len(announcements))
	}

	// 75%
	state.SetValue(75)
	announcer.Update(state)
	if len(announcements) != 3 {
		t.Errorf("Expected 3 announcements at 75%%, got %d", len(announcements))
	}

	// 100% - completion (100% threshold + "Complete" message)
	state.SetValue(100)
	announcer.Update(state)
	// Should get both "100 percent complete" and "Complete"
	if len(announcements) < 4 {
		t.Errorf("Expected at least 4 announcements at 100%%, got %d", len(announcements))
	}
	
	// Check that completion was announced
	hasComplete := false
	for _, a := range announcements {
		if strings.Contains(a, "Complete") {
			hasComplete = true
			break
		}
	}
	if !hasComplete {
		t.Error("Should have a completion announcement")
	}
}

func TestProgressAnnouncer_Reset(t *testing.T) {
	var count int
	announcer := NewProgressAnnouncer(50, func(msg string) {
		count++
	})

	state := NewState(WithValue(50), WithMax(100))
	announcer.Update(state)
	initialCount := count

	// Reset
	announcer.Reset()

	// Same value should announce again
	announcer.Update(state)
	if count != initialCount+1 {
		t.Error("Should announce again after reset")
	}
}

func TestProgressAnnouncer_NilAnnouncer(t *testing.T) {
	announcer := NewProgressAnnouncer(10, nil)
	state := NewState(WithValue(50), WithMax(100))

	// Should not panic
	announcer.Update(state)
}

func TestProgressAnnouncer_InvalidInterval(t *testing.T) {
	announcer := NewProgressAnnouncer(0, func(msg string) {})

	// Should default to 10
	if announcer.interval != 10 {
		t.Errorf("Expected interval 10 for invalid input, got %d", announcer.interval)
	}
}

func TestContrastChecker(t *testing.T) {
	checker := NewContrastChecker(4.5)

	// Same colors should fail
	ok, _ := checker.CheckColors("white", "white")
	if ok {
		t.Error("Same colors should fail contrast check")
	}

	// Obviously different should pass
	ok, _ = checker.CheckColors("white", "black")
	if !ok {
		t.Error("White on black should pass contrast check")
	}
}

func TestContrastChecker_SuggestHighContrast(t *testing.T) {
	checker := NewContrastChecker(4.5)

	tests := []struct {
		bg       string
		expected string
	}{
		{"black", "white"},
		{"dark", "white"},
		{"darkblue", "white"},
		{"white", "black"},
		{"light", "black"},
		{"lightyellow", "black"},
	}

	for _, tt := range tests {
		suggestion := checker.SuggestHighContrast(tt.bg)
		if suggestion != tt.expected {
			t.Errorf("For bg=%s, expected %s, got %s", tt.bg, tt.expected, suggestion)
		}
	}
}
