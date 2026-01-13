package tuslide

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// AccessibilityMode defines the accessibility profile.
type AccessibilityMode int

const (
	// AccessibilityDefault uses standard rendering.
	AccessibilityDefault AccessibilityMode = iota
	// AccessibilityHighContrast uses high contrast colors.
	AccessibilityHighContrast
	// AccessibilityASCII uses ASCII-only characters.
	AccessibilityASCII
	// AccessibilityScreenReader optimizes for screen readers.
	AccessibilityScreenReader
)

// AccessibleSlider wraps a Slider with accessibility features.
type AccessibleSlider struct {
	slider      *Slider
	mode        AccessibilityMode
	description string
	announcer   func(string) // Callback for announcements
	focused     bool
}

// AccessibleOption configures an AccessibleSlider.
type AccessibleOption func(*AccessibleSlider)

// NewAccessibleSlider creates an accessibility-enhanced slider.
func NewAccessibleSlider(slider *Slider, opts ...AccessibleOption) *AccessibleSlider {
	a := &AccessibleSlider{
		slider:      slider,
		mode:        AccessibilityDefault,
		description: "",
		focused:     false,
	}

	for _, opt := range opts {
		opt(a)
	}

	// Apply mode-specific settings
	a.applyMode()

	return a
}

// WithAccessibilityMode sets the accessibility mode.
func WithAccessibilityMode(mode AccessibilityMode) AccessibleOption {
	return func(a *AccessibleSlider) {
		a.mode = mode
	}
}

// WithDescription sets a screen-reader friendly description.
func WithDescription(desc string) AccessibleOption {
	return func(a *AccessibleSlider) {
		a.description = desc
	}
}

// WithAnnouncer sets a callback for value change announcements.
func WithAnnouncer(f func(string)) AccessibleOption {
	return func(a *AccessibleSlider) {
		a.announcer = f
	}
}

// applyMode applies mode-specific slider settings.
func (a *AccessibleSlider) applyMode() {
	switch a.mode {
	case AccessibilityHighContrast:
		a.applyHighContrast()
	case AccessibilityASCII:
		a.applyASCII()
	case AccessibilityScreenReader:
		a.applyScreenReader()
	}
}

// applyHighContrast applies high contrast styling.
func (a *AccessibleSlider) applyHighContrast() {
	// Use maximum contrast colors
	a.slider.filledStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")). // Bright white
		Background(lipgloss.Color("0"))   // Black
	a.slider.emptyStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("8")).  // Dark gray
		Background(lipgloss.Color("0"))   // Black
	a.slider.handleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("11")). // Bright yellow
		Bold(true)
	a.slider.labelStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")).
		Bold(true)
	a.slider.valueStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("14")). // Bright cyan
		Bold(true)
}

// applyASCII uses ASCII-only characters.
func (a *AccessibleSlider) applyASCII() {
	a.slider.symbols = ASCIISymbols()
}

// applyScreenReader optimizes for screen reader usage.
func (a *AccessibleSlider) applyScreenReader() {
	// Use simple ASCII and ensure value is always shown
	a.slider.symbols = ASCIISymbols()
	a.slider.showValue = true
}

// Slider returns the underlying slider.
func (a *AccessibleSlider) Slider() *Slider {
	return a.slider
}

// State returns the slider's state.
func (a *AccessibleSlider) State() *SliderState {
	return a.slider.state
}

// SetFocused sets the focus state and announces if changed.
func (a *AccessibleSlider) SetFocused(focused bool) {
	if a.focused != focused {
		a.focused = focused
		if focused && a.announcer != nil {
			a.announcer(a.GetFocusAnnouncement())
		}
	}
}

// IsFocused returns the focus state.
func (a *AccessibleSlider) IsFocused() bool {
	return a.focused
}

// Increment increases the value and announces the change.
func (a *AccessibleSlider) Increment() {
	oldValue := a.slider.state.Value()
	a.slider.state.Increment()
	newValue := a.slider.state.Value()

	if oldValue != newValue && a.announcer != nil {
		a.announcer(a.GetValueAnnouncement())
	}
}

// Decrement decreases the value and announces the change.
func (a *AccessibleSlider) Decrement() {
	oldValue := a.slider.state.Value()
	a.slider.state.Decrement()
	newValue := a.slider.state.Value()

	if oldValue != newValue && a.announcer != nil {
		a.announcer(a.GetValueAnnouncement())
	}
}

// SetValue sets the value and announces the change.
func (a *AccessibleSlider) SetValue(value float64) {
	oldValue := a.slider.state.Value()
	a.slider.state.SetValue(value)
	newValue := a.slider.state.Value()

	if oldValue != newValue && a.announcer != nil {
		a.announcer(a.GetValueAnnouncement())
	}
}

// View renders the slider with accessibility enhancements.
func (a *AccessibleSlider) View() string {
	view := a.slider.View()

	// Add focus indicator if focused
	if a.focused && a.mode != AccessibilityScreenReader {
		focusIndicator := lipgloss.NewStyle().
			Foreground(lipgloss.Color("11")).
			Render("▸ ")
		view = focusIndicator + view
	}

	return view
}

// GetDescription returns a screen-reader friendly description.
func (a *AccessibleSlider) GetDescription() string {
	if a.description != "" {
		return a.description
	}

	state := a.slider.state
	label := a.slider.label
	if label == "" {
		label = "Slider"
	}

	return fmt.Sprintf("%s: %.0f of %.0f (%.0f%%)",
		label,
		state.Value(),
		state.Max(),
		state.Percentage()*100,
	)
}

// GetValueAnnouncement returns an announcement for the current value.
func (a *AccessibleSlider) GetValueAnnouncement() string {
	state := a.slider.state
	return fmt.Sprintf("%.0f, %.0f percent", state.Value(), state.Percentage()*100)
}

// GetFocusAnnouncement returns an announcement when focused.
func (a *AccessibleSlider) GetFocusAnnouncement() string {
	return a.GetDescription() + ". Use arrow keys to adjust."
}

// GetBoundaryAnnouncement returns an announcement for boundary conditions.
func (a *AccessibleSlider) GetBoundaryAnnouncement() string {
	state := a.slider.state
	if state.Value() <= state.Min() {
		return "Minimum value reached"
	}
	if state.Value() >= state.Max() {
		return "Maximum value reached"
	}
	return ""
}

// FocusIndicator provides a styled focus indicator.
type FocusIndicator struct {
	style    lipgloss.Style
	char     string
	position string // "left", "right", "both"
}

// NewFocusIndicator creates a new focus indicator.
func NewFocusIndicator() *FocusIndicator {
	return &FocusIndicator{
		style: lipgloss.NewStyle().
			Foreground(lipgloss.Color("11")).
			Bold(true),
		char:     "▸",
		position: "left",
	}
}

// WithStyle sets the indicator style.
func (f *FocusIndicator) WithStyle(style lipgloss.Style) *FocusIndicator {
	f.style = style
	return f
}

// WithChar sets the indicator character.
func (f *FocusIndicator) WithChar(char string) *FocusIndicator {
	f.char = char
	return f
}

// WithPosition sets the indicator position ("left", "right", "both").
func (f *FocusIndicator) WithPosition(pos string) *FocusIndicator {
	f.position = pos
	return f
}

// Render returns the focus indicator if focused.
func (f *FocusIndicator) Render(focused bool) string {
	if !focused {
		return ""
	}
	return f.style.Render(f.char)
}

// Wrap wraps content with focus indicators.
func (f *FocusIndicator) Wrap(content string, focused bool) string {
	if !focused {
		return content
	}

	indicator := f.style.Render(f.char)

	switch f.position {
	case "right":
		return content + " " + indicator
	case "both":
		return indicator + " " + content + " " + indicator
	default: // "left"
		return indicator + " " + content
	}
}

// HighContrastPalette provides high contrast color options.
type HighContrastPalette struct {
	Foreground lipgloss.Color
	Background lipgloss.Color
	Accent     lipgloss.Color
	Muted      lipgloss.Color
	Warning    lipgloss.Color
	Error      lipgloss.Color
	Success    lipgloss.Color
}

// DefaultHighContrastPalette returns the default high contrast palette.
func DefaultHighContrastPalette() HighContrastPalette {
	return HighContrastPalette{
		Foreground: lipgloss.Color("15"), // Bright white
		Background: lipgloss.Color("0"),  // Black
		Accent:     lipgloss.Color("11"), // Bright yellow
		Muted:      lipgloss.Color("8"),  // Dark gray
		Warning:    lipgloss.Color("11"), // Bright yellow
		Error:      lipgloss.Color("9"),  // Bright red
		Success:    lipgloss.Color("10"), // Bright green
	}
}

// DarkHighContrastPalette returns a dark high contrast palette.
func DarkHighContrastPalette() HighContrastPalette {
	return HighContrastPalette{
		Foreground: lipgloss.Color("255"), // Pure white
		Background: lipgloss.Color("16"),  // Pure black
		Accent:     lipgloss.Color("226"), // Bright yellow
		Muted:      lipgloss.Color("240"), // Gray
		Warning:    lipgloss.Color("214"), // Orange
		Error:      lipgloss.Color("196"), // Red
		Success:    lipgloss.Color("46"),  // Green
	}
}

// LightHighContrastPalette returns a light high contrast palette.
func LightHighContrastPalette() HighContrastPalette {
	return HighContrastPalette{
		Foreground: lipgloss.Color("16"),  // Pure black
		Background: lipgloss.Color("255"), // Pure white
		Accent:     lipgloss.Color("21"),  // Blue
		Muted:      lipgloss.Color("240"), // Gray
		Warning:    lipgloss.Color("202"), // Orange
		Error:      lipgloss.Color("160"), // Dark red
		Success:    lipgloss.Color("28"),  // Dark green
	}
}

// ApplyPalette applies a high contrast palette to a slider.
func ApplyPalette(slider *Slider, palette HighContrastPalette) {
	slider.filledStyle = lipgloss.NewStyle().Foreground(palette.Foreground)
	slider.emptyStyle = lipgloss.NewStyle().Foreground(palette.Muted)
	slider.handleStyle = lipgloss.NewStyle().Foreground(palette.Accent).Bold(true)
	slider.labelStyle = lipgloss.NewStyle().Foreground(palette.Foreground).Bold(true)
	slider.valueStyle = lipgloss.NewStyle().Foreground(palette.Accent)
}

// KeyboardHints provides keyboard shortcut hints.
type KeyboardHints struct {
	IncrementKey string
	DecrementKey string
	MinKey       string
	MaxKey       string
	StepSize     string
}

// DefaultKeyboardHints returns default keyboard hints.
func DefaultKeyboardHints() KeyboardHints {
	return KeyboardHints{
		IncrementKey: "→ or l",
		DecrementKey: "← or h",
		MinKey:       "Home or 0",
		MaxKey:       "End or $",
		StepSize:     "Step",
	}
}

// Render returns a formatted keyboard hints string.
func (k KeyboardHints) Render() string {
	return fmt.Sprintf(
		"Increase: %s | Decrease: %s | Min: %s | Max: %s",
		k.IncrementKey, k.DecrementKey, k.MinKey, k.MaxKey,
	)
}

// RenderCompact returns a compact keyboard hints string.
func (k KeyboardHints) RenderCompact() string {
	return fmt.Sprintf("[%s/%s]", k.IncrementKey, k.DecrementKey)
}

// ProgressAnnouncer provides progress announcements for long operations.
type ProgressAnnouncer struct {
	lastAnnounced int // Last announced percentage (to avoid spam)
	interval      int // Announce every N percent
	announcer     func(string)
}

// NewProgressAnnouncer creates a new progress announcer.
func NewProgressAnnouncer(interval int, announcer func(string)) *ProgressAnnouncer {
	if interval < 1 {
		interval = 10
	}
	return &ProgressAnnouncer{
		lastAnnounced: -1,
		interval:      interval,
		announcer:     announcer,
	}
}

// Update checks if a new announcement should be made.
func (p *ProgressAnnouncer) Update(state *SliderState) {
	if p.announcer == nil {
		return
	}

	pct := int(state.Percentage() * 100)
	threshold := (pct / p.interval) * p.interval

	// Check for completion first (special case)
	if pct >= 100 && p.lastAnnounced < 100 {
		p.lastAnnounced = 100
		p.announcer("Complete")
		return
	}

	// Regular interval announcements
	if threshold != p.lastAnnounced && threshold > 0 && threshold < 100 {
		p.lastAnnounced = threshold
		p.announcer(fmt.Sprintf("%d percent complete", threshold))
	}
}

// Reset resets the announcer for a new operation.
func (p *ProgressAnnouncer) Reset() {
	p.lastAnnounced = -1
}

// ContrastChecker helps verify color contrast ratios.
type ContrastChecker struct {
	minRatio float64
}

// NewContrastChecker creates a contrast checker with minimum ratio.
// WCAG AA requires 4.5:1 for normal text, 3:1 for large text.
func NewContrastChecker(minRatio float64) *ContrastChecker {
	return &ContrastChecker{minRatio: minRatio}
}

// CheckColors verifies if two colors have sufficient contrast.
// Returns a suggestion if contrast is insufficient.
func (c *ContrastChecker) CheckColors(fg, bg string) (bool, string) {
	// Simplified check - in production you'd calculate actual luminance
	// This is a basic check for obviously low-contrast combinations
	if fg == bg {
		return false, "Foreground and background colors are identical"
	}

	// Check for common low-contrast combinations
	lowContrast := map[string][]string{
		"gray":    {"gray", "darkgray", "lightgray"},
		"blue":    {"darkblue", "navy"},
		"green":   {"darkgreen"},
		"red":     {"darkred", "maroon"},
		"yellow":  {"lightyellow", "white"},
		"white":   {"lightyellow", "lightgray"},
		"black":   {"darkgray", "navy", "darkgreen", "darkred"},
	}

	fgLower := strings.ToLower(fg)
	bgLower := strings.ToLower(bg)

	if similar, ok := lowContrast[fgLower]; ok {
		for _, s := range similar {
			if bgLower == s {
				return false, fmt.Sprintf("Low contrast between %s and %s", fg, bg)
			}
		}
	}

	return true, ""
}

// SuggestHighContrast suggests a high contrast alternative.
func (c *ContrastChecker) SuggestHighContrast(bg string) string {
	bgLower := strings.ToLower(bg)

	// Suggest contrasting foreground based on background
	switch {
	case strings.Contains(bgLower, "dark") || bgLower == "black" || bgLower == "navy":
		return "white"
	case strings.Contains(bgLower, "light") || bgLower == "white" || bgLower == "yellow":
		return "black"
	default:
		return "white" // Default to white for safety
	}
}
