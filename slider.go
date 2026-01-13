package tuslide

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

// Orientation defines whether the slider is horizontal or vertical.
type Orientation int

const (
	// Horizontal renders the slider left-to-right.
	Horizontal Orientation = iota
	// Vertical renders the slider bottom-to-top.
	Vertical
)

// LabelPosition defines where labels appear relative to the slider.
type LabelPosition int

const (
	// LabelNone hides the label.
	LabelNone LabelPosition = iota
	// LabelLeft places the label to the left.
	LabelLeft
	// LabelRight places the label to the right.
	LabelRight
	// LabelTop places the label above.
	LabelTop
	// LabelBottom places the label below.
	LabelBottom
)

// ValuePosition defines where the numeric value appears.
type ValuePosition int

const (
	// ValueRight places the value to the right of the track (default for horizontal).
	ValueRight ValuePosition = iota
	// ValueLeft places the value to the left of the track.
	ValueLeft
	// ValueTop places the value above the track.
	ValueTop
	// ValueBottom places the value below the track (default for vertical).
	ValueBottom
	// ValueInline places the value inside the track (centered).
	ValueInline
)

// HorizontalBarAlignment defines vertical positioning of bar in horizontal sliders.
type HorizontalBarAlignment int

const (
	// BarCenter positions the bar in the center (default).
	BarCenter HorizontalBarAlignment = iota
	// BarTop positions the bar at the top.
	BarTop
	// BarBottom positions the bar at the bottom.
	BarBottom
)

// TitleAlignment defines how the title/label is aligned.
type TitleAlignment int

const (
	// TitleAlignLeft aligns title to the left.
	TitleAlignLeft TitleAlignment = iota
	// TitleAlignCenter aligns title to the center.
	TitleAlignCenter
	// TitleAlignRight aligns title to the right.
	TitleAlignRight
)

// BorderStyle defines the border appearance.
type BorderStyle int

const (
	// BorderNone shows no border.
	BorderNone BorderStyle = iota
	// BorderRounded shows a rounded border.
	BorderRounded
	// BorderNormal shows a normal border.
	BorderNormal
	// BorderThick shows a thick border.
	BorderThick
	// BorderDouble shows a double-line border.
	BorderDouble
)

// VerticalValueAlignment defines horizontal alignment of value in vertical sliders.
type VerticalValueAlignment int

const (
	// VValueCenter aligns value to center (default).
	VValueCenter VerticalValueAlignment = iota
	// VValueLeft aligns value to left.
	VValueLeft
	// VValueRight aligns value to right.
	VValueRight
)

// VerticalLabelPosition defines where labels appear in vertical sliders.
type VerticalLabelPosition int

const (
	// VLabelTop places the label above the vertical slider (default).
	VLabelTop VerticalLabelPosition = iota
	// VLabelBottom places the label below the vertical slider.
	VLabelBottom
)

// VerticalValuePosition defines where the value appears in vertical sliders.
type VerticalValuePosition int

const (
	// VValuePosBottom places the value below the vertical slider (default).
	VValuePosBottom VerticalValuePosition = iota
	// VValuePosTop places the value above the vertical slider.
	VValuePosTop
	// VValuePosMiddle places the value at the middle of the vertical slider.
	VValuePosMiddle
)

// ValueAlignment defines horizontal alignment of value text.
type ValueAlignment int

const (
	// AlignLeft aligns value to the left.
	AlignLeft ValueAlignment = iota
	// AlignCenter aligns value to the center.
	AlignCenter
	// AlignRight aligns value to the right.
	AlignRight
)

// Symbols defines the characters used to render the slider.
type Symbols struct {
	Filled string // Character for the filled portion
	Empty  string // Character for the empty portion
	Handle string // Character for the handle (current position)
}

// DefaultSymbols returns the default slider symbols.
func DefaultSymbols() Symbols {
	return Symbols{
		Filled: "█",
		Empty:  "░",
		Handle: "●",
	}
}

// ASCIISymbols returns ASCII-only symbols for compatibility.
func ASCIISymbols() Symbols {
	return Symbols{
		Filled: "=",
		Empty:  "-",
		Handle: "O",
	}
}

// BlockSymbols returns block-style symbols.
func BlockSymbols() Symbols {
	return Symbols{
		Filled: "█",
		Empty:  "▒",
		Handle: "█",
	}
}

// Slider is a TUI slider widget that renders based on a SliderState.
type Slider struct {
	state          *SliderState
	width          int
	height         int
	orientation    Orientation
	symbols        Symbols
	showHandle     bool
	label          string
	labelPosition  LabelPosition
	showValue      bool
	valuePosition  ValuePosition
	valueFormat    string // Format string for value display (e.g., "%.1f", "%d%%")
	collisionCheck bool   // Enable label/value collision detection

	// Additional positioning options
	horizontalBarAlignment  HorizontalBarAlignment
	titleAlignment          TitleAlignment
	verticalValueAlignment  VerticalValueAlignment
	verticalLabelPosition   VerticalLabelPosition
	verticalValuePosition   VerticalValuePosition
	valueAlignment          ValueAlignment

	// Border options
	borderStyle BorderStyle
	borderTitle string
	borderColor lipgloss.Color

	// Segmented mode
	segmented    bool
	segmentCount int // Number of segments (0 = auto based on width)
	segmentGap   int // Gap between segments in characters

	// Styles
	filledStyle lipgloss.Style
	emptyStyle  lipgloss.Style
	handleStyle lipgloss.Style
	labelStyle  lipgloss.Style
	valueStyle  lipgloss.Style
	borderStyle_ lipgloss.Style
}

// SliderOption is a functional option for configuring a Slider.
type SliderOption func(*Slider)

// New creates a new Slider with the given state and options.
// If state is nil, a default state is created.
func New(state *SliderState, opts ...SliderOption) *Slider {
	if state == nil {
		state = NewState()
	}

	s := &Slider{
		state:          state,
		width:          20,
		height:         10,
		orientation:    Horizontal,
		symbols:        DefaultSymbols(),
		showHandle:     true,
		label:          "",
		labelPosition:  LabelNone,
		showValue:      false,
		valuePosition:  ValueRight,
		valueFormat:    "",
		collisionCheck: true,
		// New positioning options
		horizontalBarAlignment:  BarCenter,
		titleAlignment:          TitleAlignLeft,
		verticalValueAlignment:  VValueCenter,
		verticalLabelPosition:   VLabelTop,
		verticalValuePosition:   VValuePosBottom,
		valueAlignment:          AlignRight,
		// Border options
		borderStyle: BorderNone,
		borderTitle: "",
		borderColor: lipgloss.Color("240"),
		// Segmented mode
		segmented:    false,
		segmentCount: 0,
		segmentGap:   1,
		// Styles
		filledStyle:  lipgloss.NewStyle(),
		emptyStyle:   lipgloss.NewStyle(),
		handleStyle:  lipgloss.NewStyle(),
		labelStyle:   lipgloss.NewStyle(),
		valueStyle:   lipgloss.NewStyle(),
		borderStyle_: lipgloss.NewStyle(),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// WithWidth sets the width of the slider track (for horizontal orientation).
func WithWidth(width int) SliderOption {
	return func(s *Slider) {
		if width > 0 {
			s.width = width
		}
	}
}

// WithHeight sets the height of the slider track (for vertical orientation).
func WithHeight(height int) SliderOption {
	return func(s *Slider) {
		if height > 0 {
			s.height = height
		}
	}
}

// WithOrientation sets the slider orientation.
func WithOrientation(o Orientation) SliderOption {
	return func(s *Slider) {
		s.orientation = o
	}
}

// WithSymbols sets custom symbols for rendering.
func WithSymbols(symbols Symbols) SliderOption {
	return func(s *Slider) {
		s.symbols = symbols
	}
}

// WithHandle controls whether the handle is visible.
// Set to false for progress bar mode.
func WithHandle(show bool) SliderOption {
	return func(s *Slider) {
		s.showHandle = show
	}
}

// WithLabel sets the slider label.
func WithLabel(label string) SliderOption {
	return func(s *Slider) {
		s.label = label
	}
}

// WithLabelPosition sets where the label appears.
func WithLabelPosition(pos LabelPosition) SliderOption {
	return func(s *Slider) {
		s.labelPosition = pos
	}
}

// WithShowValue controls whether the numeric value is displayed.
func WithShowValue(show bool) SliderOption {
	return func(s *Slider) {
		s.showValue = show
	}
}

// WithValuePosition sets where the numeric value appears.
func WithValuePosition(pos ValuePosition) SliderOption {
	return func(s *Slider) {
		s.valuePosition = pos
	}
}

// WithValueFormat sets a custom format string for the value display.
// Use Go fmt-style format strings like "%.1f" or "%d%%".
// If empty, a sensible default is used.
func WithValueFormat(format string) SliderOption {
	return func(s *Slider) {
		s.valueFormat = format
	}
}

// WithCollisionCheck enables or disables label/value collision detection.
// When enabled, overlapping text is automatically repositioned.
func WithCollisionCheck(enabled bool) SliderOption {
	return func(s *Slider) {
		s.collisionCheck = enabled
	}
}

// WithFilledStyle sets the style for the filled portion.
func WithFilledStyle(style lipgloss.Style) SliderOption {
	return func(s *Slider) {
		s.filledStyle = style
	}
}

// WithEmptyStyle sets the style for the empty portion.
func WithEmptyStyle(style lipgloss.Style) SliderOption {
	return func(s *Slider) {
		s.emptyStyle = style
	}
}

// WithHandleStyle sets the style for the handle.
func WithHandleStyle(style lipgloss.Style) SliderOption {
	return func(s *Slider) {
		s.handleStyle = style
	}
}

// WithLabelStyle sets the style for the label text.
func WithLabelStyle(style lipgloss.Style) SliderOption {
	return func(s *Slider) {
		s.labelStyle = style
	}
}

// WithValueStyle sets the style for the value display.
func WithValueStyle(style lipgloss.Style) SliderOption {
	return func(s *Slider) {
		s.valueStyle = style
	}
}

// WithHorizontalBarAlignment sets vertical positioning of bar in horizontal sliders.
func WithHorizontalBarAlignment(align HorizontalBarAlignment) SliderOption {
	return func(s *Slider) {
		s.horizontalBarAlignment = align
	}
}

// WithTitleAlignment sets alignment for the title/label text.
func WithTitleAlignment(align TitleAlignment) SliderOption {
	return func(s *Slider) {
		s.titleAlignment = align
	}
}

// WithVerticalValueAlignment sets horizontal alignment of value in vertical sliders.
func WithVerticalValueAlignment(align VerticalValueAlignment) SliderOption {
	return func(s *Slider) {
		s.verticalValueAlignment = align
	}
}

// WithVerticalLabelPosition sets where labels appear in vertical sliders.
func WithVerticalLabelPosition(pos VerticalLabelPosition) SliderOption {
	return func(s *Slider) {
		s.verticalLabelPosition = pos
	}
}

// WithVerticalValuePosition sets where the value appears in vertical sliders.
func WithVerticalValuePosition(pos VerticalValuePosition) SliderOption {
	return func(s *Slider) {
		s.verticalValuePosition = pos
	}
}

// WithValueAlignment sets horizontal alignment of the value text.
// This is particularly useful for horizontal sliders where you want
// to control where the value appears (left, center, or right aligned).
func WithValueAlignment(align ValueAlignment) SliderOption {
	return func(s *Slider) {
		s.valueAlignment = align
	}
}

// WithBorder sets the border style for the slider.
func WithBorder(style BorderStyle) SliderOption {
	return func(s *Slider) {
		s.borderStyle = style
	}
}

// WithBorderTitle sets a title for the border.
func WithBorderTitle(title string) SliderOption {
	return func(s *Slider) {
		s.borderTitle = title
	}
}

// WithBorderColor sets the color of the border.
func WithBorderColor(color lipgloss.Color) SliderOption {
	return func(s *Slider) {
		s.borderColor = color
	}
}

// WithSegmented enables segmented slider mode.
func WithSegmented(enabled bool) SliderOption {
	return func(s *Slider) {
		s.segmented = enabled
	}
}

// WithSegmentCount sets the number of segments (0 = auto).
func WithSegmentCount(count int) SliderOption {
	return func(s *Slider) {
		if count >= 0 {
			s.segmentCount = count
		}
	}
}

// WithSegmentGap sets the gap between segments in characters.
func WithSegmentGap(gap int) SliderOption {
	return func(s *Slider) {
		if gap >= 0 {
			s.segmentGap = gap
		}
	}
}

// WithStyle applies a predefined SliderStyle to the slider.
func WithStyle(style SliderStyle) SliderOption {
	return func(s *Slider) {
		s.symbols = style.Symbols
		s.filledStyle = style.FilledStyle
		s.emptyStyle = style.EmptyStyle
		s.handleStyle = style.HandleStyle
		s.labelStyle = style.LabelStyle
		s.valueStyle = style.ValueStyle
		s.segmented = style.Segmented
	}
}

// WithSymbolSet applies a predefined SymbolSet to the slider.
func WithSymbolSet(set SymbolSet) SliderOption {
	return func(s *Slider) {
		s.symbols = set.ToSymbols()
	}
}

// State returns the slider's state.
func (s *Slider) State() *SliderState {
	return s.state
}

// SetState updates the slider's state.
func (s *Slider) SetState(state *SliderState) {
	s.state = state
}

// View renders the slider and returns the string representation.
// This is compatible with Bubble Tea's View method pattern.
func (s *Slider) View() string {
	var content string
	switch s.orientation {
	case Vertical:
		content = s.renderVertical()
	default:
		content = s.renderHorizontal()
	}

	// Apply border if configured
	if s.borderStyle != BorderNone {
		content = s.applyBorder(content)
	}

	return content
}

// applyBorder wraps the content with a border.
func (s *Slider) applyBorder(content string) string {
	var border lipgloss.Border
	switch s.borderStyle {
	case BorderRounded:
		border = lipgloss.RoundedBorder()
	case BorderNormal:
		border = lipgloss.NormalBorder()
	case BorderThick:
		border = lipgloss.ThickBorder()
	case BorderDouble:
		border = lipgloss.DoubleBorder()
	default:
		return content
	}

	style := lipgloss.NewStyle().
		Border(border).
		BorderForeground(s.borderColor)

	if s.borderTitle != "" {
		// Add title above the border
		titleStyle := lipgloss.NewStyle().Foreground(s.borderColor)
		return titleStyle.Render(s.borderTitle) + "\n" + style.Render(content)
	}

	return style.Render(content)
}

// String implements fmt.Stringer.
func (s *Slider) String() string {
	return s.View()
}

// renderHorizontal renders a horizontal slider.
func (s *Slider) renderHorizontal() string {
	track := s.buildHorizontalTrack()
	label := ""
	value := ""

	if s.label != "" {
		label = s.labelStyle.Render(s.label)
	}
	if s.showValue {
		value = s.valueStyle.Render(s.formatValue())
	}

	// Resolve collisions if enabled
	labelPos := s.labelPosition
	valuePos := s.valuePosition
	if s.collisionCheck && label != "" && value != "" {
		labelPos, valuePos = s.resolveCollision(labelPos, valuePos)
	}

	var topLine, midLine, bottomLine strings.Builder
	leftPad := ""
	rightPad := ""

	// Build left side
	if labelPos == LabelLeft {
		leftPad = label + " "
	}
	if s.showValue && valuePos == ValueLeft {
		if leftPad != "" {
			leftPad = value + " " + leftPad
		} else {
			leftPad = value + " "
		}
	}

	// Build right side
	if labelPos == LabelRight {
		rightPad = " " + label
	}
	if s.showValue && valuePos == ValueRight {
		if rightPad != "" {
			rightPad = rightPad + " " + value
		} else {
			rightPad = " " + value
		}
	}

	// Build middle line (track with left/right elements)
	midLine.WriteString(leftPad)
	midLine.WriteString(track)
	midLine.WriteString(rightPad)

	// Build top line
	if labelPos == LabelTop {
		topLine.WriteString(strings.Repeat(" ", runewidth.StringWidth(leftPad)))
		topLine.WriteString(label)
	}
	if s.showValue && valuePos == ValueTop {
		if topLine.Len() > 0 {
			topLine.WriteString(" ")
		} else {
			topLine.WriteString(strings.Repeat(" ", runewidth.StringWidth(leftPad)))
		}
		topLine.WriteString(value)
	}

	// Build bottom line
	if labelPos == LabelBottom {
		bottomLine.WriteString(strings.Repeat(" ", runewidth.StringWidth(leftPad)))
		bottomLine.WriteString(label)
	}
	if s.showValue && valuePos == ValueBottom {
		if bottomLine.Len() > 0 {
			bottomLine.WriteString(" ")
		} else {
			bottomLine.WriteString(strings.Repeat(" ", runewidth.StringWidth(leftPad)))
		}
		bottomLine.WriteString(value)
	}

	// Combine lines
	var result strings.Builder
	if topLine.Len() > 0 {
		result.WriteString(topLine.String())
		result.WriteString("\n")
	}
	result.WriteString(midLine.String())
	if bottomLine.Len() > 0 {
		result.WriteString("\n")
		result.WriteString(bottomLine.String())
	}

	return result.String()
}

// resolveCollision adjusts positions when label and value would overlap.
func (s *Slider) resolveCollision(labelPos LabelPosition, valuePos ValuePosition) (LabelPosition, ValuePosition) {
	// Map label positions to conflict groups
	// Left/Right are horizontal, Top/Bottom are vertical
	labelHorizontal := labelPos == LabelLeft || labelPos == LabelRight
	labelVertical := labelPos == LabelTop || labelPos == LabelBottom
	valueHorizontal := valuePos == ValueLeft || valuePos == ValueRight
	valueVertical := valuePos == ValueTop || valuePos == ValueBottom

	// Check for same-side collision
	if labelPos == LabelLeft && valuePos == ValueLeft {
		// Move value to right
		return labelPos, ValueRight
	}
	if labelPos == LabelRight && valuePos == ValueRight {
		// Move value to left
		return labelPos, ValueLeft
	}
	if labelPos == LabelTop && valuePos == ValueTop {
		// Move value to bottom
		return labelPos, ValueBottom
	}
	if labelPos == LabelBottom && valuePos == ValueBottom {
		// Move value to top
		return labelPos, ValueTop
	}

	// If both want horizontal space and we have room, keep them
	if labelHorizontal && valueHorizontal {
		return labelPos, valuePos
	}

	// If both want vertical space and we have room, keep them
	if labelVertical && valueVertical {
		return labelPos, valuePos
	}

	return labelPos, valuePos
}

// buildHorizontalTrack builds just the horizontal slider track.
func (s *Slider) buildHorizontalTrack() string {
	if s.segmented {
		return s.buildSegmentedHorizontalTrack()
	}

	pct := s.state.Percentage()
	trackWidth := s.width

	// Calculate handle width using runewidth for Unicode accuracy
	handleWidth := 0
	if s.showHandle {
		handleWidth = runewidth.StringWidth(s.symbols.Handle)
	}

	// Available width for filled + empty (excluding handle)
	availableWidth := trackWidth - handleWidth

	if availableWidth < 0 {
		availableWidth = 0
	}

	// Calculate filled cells
	filledCells := int(float64(availableWidth) * pct)
	if filledCells > availableWidth {
		filledCells = availableWidth
	}

	emptyCells := availableWidth - filledCells

	var track strings.Builder

	// Build filled portion
	filledSymbolWidth := runewidth.StringWidth(s.symbols.Filled)
	for i := 0; i < filledCells; {
		track.WriteString(s.filledStyle.Render(s.symbols.Filled))
		i += filledSymbolWidth
		if i > filledCells {
			break
		}
	}

	// Render handle
	if s.showHandle {
		track.WriteString(s.handleStyle.Render(s.symbols.Handle))
	}

	// Build empty portion
	emptySymbolWidth := runewidth.StringWidth(s.symbols.Empty)
	for i := 0; i < emptyCells; {
		track.WriteString(s.emptyStyle.Render(s.symbols.Empty))
		i += emptySymbolWidth
		if i > emptyCells {
			break
		}
	}

	return track.String()
}

// buildSegmentedHorizontalTrack builds a segmented horizontal slider track.
func (s *Slider) buildSegmentedHorizontalTrack() string {
	pct := s.state.Percentage()

	// Determine segment count
	segmentCount := s.segmentCount
	if segmentCount <= 0 {
		// Auto-calculate: approximately one segment per 2-3 characters
		segmentCount = s.width / 3
		if segmentCount < 5 {
			segmentCount = 5
		}
		if segmentCount > 20 {
			segmentCount = 20
		}
	}

	// Calculate how many segments should be filled
	filledSegments := int(float64(segmentCount) * pct)
	if filledSegments > segmentCount {
		filledSegments = segmentCount
	}

	// Handle position (at the filled/empty boundary)
	handlePos := filledSegments
	if handlePos >= segmentCount {
		handlePos = segmentCount - 1
	}

	var track strings.Builder
	gap := strings.Repeat(" ", s.segmentGap)

	for i := 0; i < segmentCount; i++ {
		if i > 0 {
			track.WriteString(gap)
		}

		if s.showHandle && i == handlePos {
			track.WriteString(s.handleStyle.Render(s.symbols.Handle))
		} else if i < filledSegments {
			track.WriteString(s.filledStyle.Render(s.symbols.Filled))
		} else {
			track.WriteString(s.emptyStyle.Render(s.symbols.Empty))
		}
	}

	return track.String()
}

// renderVertical renders a vertical slider.
func (s *Slider) renderVertical() string {
	trackLines := s.buildVerticalTrack()
	label := ""
	value := ""

	if s.label != "" {
		label = s.labelStyle.Render(s.label)
	}
	if s.showValue {
		value = s.valueStyle.Render(s.formatValue())
	}

	// Resolve collisions if enabled
	labelPos := s.labelPosition
	valuePos := s.valuePosition
	if s.collisionCheck && label != "" && value != "" {
		labelPos, valuePos = s.resolveCollision(labelPos, valuePos)
	}

	// For vertical sliders, default value position to bottom
	if s.showValue && valuePos == ValueRight {
		valuePos = ValueBottom
	}

	var result strings.Builder
	labelWidth := runewidth.StringWidth(label)
	valueWidth := runewidth.StringWidth(value)

	// Unused but kept for potential future alignment features
	_ = runewidth.StringWidth(label)

	// Build top section
	if labelPos == LabelTop && label != "" {
		result.WriteString(label)
		result.WriteString("\n")
	}
	if s.showValue && valuePos == ValueTop {
		result.WriteString(value)
		result.WriteString("\n")
	}

	// Build track with optional left/right labels
	midRow := len(trackLines) / 2
	for i, line := range trackLines {
		// Left side
		if labelPos == LabelLeft && label != "" {
			if i == midRow {
				result.WriteString(label + " ")
			} else {
				result.WriteString(strings.Repeat(" ", labelWidth+1))
			}
		}
		if s.showValue && valuePos == ValueLeft {
			if i == midRow && labelPos != LabelLeft {
				result.WriteString(value + " ")
			} else if labelPos != LabelLeft {
				result.WriteString(strings.Repeat(" ", valueWidth+1))
			}
		}

		// Track line
		result.WriteString(line)

		// Right side
		if labelPos == LabelRight && label != "" {
			if i == midRow {
				result.WriteString(" " + label)
			}
		}
		if s.showValue && valuePos == ValueRight {
			if i == midRow && labelPos != LabelRight {
				result.WriteString(" " + value)
			}
		}

		if i < len(trackLines)-1 {
			result.WriteString("\n")
		}
	}

	// Build bottom section
	if labelPos == LabelBottom && label != "" {
		result.WriteString("\n")
		result.WriteString(label)
	}
	if s.showValue && valuePos == ValueBottom {
		result.WriteString("\n")
		result.WriteString(value)
	}

	return result.String()
}

// buildVerticalTrack builds the vertical slider track lines.
func (s *Slider) buildVerticalTrack() []string {
	pct := s.state.Percentage()
	trackHeight := s.height

	// Calculate filled rows (from bottom)
	filledRows := int(float64(trackHeight) * pct)
	if filledRows > trackHeight {
		filledRows = trackHeight
	}
	emptyRows := trackHeight - filledRows

	// Handle position calculation
	handleRow := -1
	if s.showHandle {
		handleRow = emptyRows
		if handleRow >= trackHeight {
			handleRow = trackHeight - 1
		}
		if handleRow < 0 {
			handleRow = 0
		}
	}

	var lines []string

	// Build from top to bottom
	for i := 0; i < trackHeight; i++ {
		if s.showHandle && i == handleRow {
			lines = append(lines, s.handleStyle.Render(s.symbols.Handle))
		} else if i < emptyRows {
			lines = append(lines, s.emptyStyle.Render(s.symbols.Empty))
		} else {
			lines = append(lines, s.filledStyle.Render(s.symbols.Filled))
		}
	}

	return lines
}

// formatValue formats the current value for display.
func (s *Slider) formatValue() string {
	v := s.state.Value()

	// If custom format is specified, use it
	if s.valueFormat != "" {
		return fmt.Sprintf(s.valueFormat, v)
	}

	// Check if value is effectively an integer (no fractional part)
	if v == math.Trunc(v) {
		return fmt.Sprintf("%.0f", v)
	}

	// For decimal values, show one decimal place
	return fmt.Sprintf("%.1f", v)
}
