package tuslide

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestNew_DefaultState(t *testing.T) {
	slider := New(nil)
	if slider.state == nil {
		t.Error("Expected default state to be created when nil is passed")
	}
	if slider.width != 20 {
		t.Errorf("Expected default width 20, got %d", slider.width)
	}
	if slider.height != 10 {
		t.Errorf("Expected default height 10, got %d", slider.height)
	}
	if slider.orientation != Horizontal {
		t.Error("Expected default orientation to be Horizontal")
	}
	if !slider.showHandle {
		t.Error("Expected showHandle to be true by default")
	}
}

func TestNew_WithState(t *testing.T) {
	state := NewState(WithValue(75), WithMin(0), WithMax(100))
	slider := New(state)
	
	if slider.state != state {
		t.Error("Expected slider to use provided state")
	}
	if slider.state.Value() != 75 {
		t.Errorf("Expected value 75, got %f", slider.state.Value())
	}
}

func TestWithWidth(t *testing.T) {
	slider := New(nil, WithWidth(50))
	if slider.width != 50 {
		t.Errorf("Expected width 50, got %d", slider.width)
	}

	// Test invalid width (should not change)
	slider = New(nil, WithWidth(-10))
	if slider.width != 20 {
		t.Errorf("Expected default width 20 for invalid input, got %d", slider.width)
	}
}

func TestWithHeight(t *testing.T) {
	slider := New(nil, WithHeight(15))
	if slider.height != 15 {
		t.Errorf("Expected height 15, got %d", slider.height)
	}

	// Test invalid height (should not change)
	slider = New(nil, WithHeight(0))
	if slider.height != 10 {
		t.Errorf("Expected default height 10 for invalid input, got %d", slider.height)
	}
}

func TestWithOrientation(t *testing.T) {
	slider := New(nil, WithOrientation(Vertical))
	if slider.orientation != Vertical {
		t.Error("Expected orientation to be Vertical")
	}
}

func TestWithSymbols(t *testing.T) {
	customSymbols := Symbols{
		Filled: "X",
		Empty:  "-",
		Handle: "O",
	}
	slider := New(nil, WithSymbols(customSymbols))
	
	if slider.symbols.Filled != "X" {
		t.Errorf("Expected filled symbol 'X', got '%s'", slider.symbols.Filled)
	}
	if slider.symbols.Empty != "-" {
		t.Errorf("Expected empty symbol '-', got '%s'", slider.symbols.Empty)
	}
	if slider.symbols.Handle != "O" {
		t.Errorf("Expected handle symbol 'O', got '%s'", slider.symbols.Handle)
	}
}

func TestWithHandle(t *testing.T) {
	slider := New(nil, WithHandle(false))
	if slider.showHandle {
		t.Error("Expected showHandle to be false")
	}
}

func TestWithLabel(t *testing.T) {
	slider := New(nil, WithLabel("Volume"))
	if slider.label != "Volume" {
		t.Errorf("Expected label 'Volume', got '%s'", slider.label)
	}
}

func TestWithLabelPosition(t *testing.T) {
	positions := []LabelPosition{LabelNone, LabelLeft, LabelRight, LabelTop, LabelBottom}
	for _, pos := range positions {
		slider := New(nil, WithLabelPosition(pos))
		if slider.labelPosition != pos {
			t.Errorf("Expected label position %v, got %v", pos, slider.labelPosition)
		}
	}
}

func TestWithShowValue(t *testing.T) {
	slider := New(nil, WithShowValue(true))
	if !slider.showValue {
		t.Error("Expected showValue to be true")
	}
}

func TestWithValuePosition(t *testing.T) {
	positions := []ValuePosition{ValueRight, ValueLeft, ValueTop, ValueBottom, ValueInline}
	for _, pos := range positions {
		slider := New(nil, WithValuePosition(pos))
		if slider.valuePosition != pos {
			t.Errorf("Expected value position %v, got %v", pos, slider.valuePosition)
		}
	}
}

func TestWithValueFormat(t *testing.T) {
	slider := New(nil, WithValueFormat("%.2f%%"))
	if slider.valueFormat != "%.2f%%" {
		t.Errorf("Expected value format '%%.2f%%%%', got '%s'", slider.valueFormat)
	}
}

func TestWithCollisionCheck(t *testing.T) {
	slider := New(nil, WithCollisionCheck(false))
	if slider.collisionCheck {
		t.Error("Expected collisionCheck to be false")
	}
}

func TestWithHorizontalBarAlignment(t *testing.T) {
	alignments := []HorizontalBarAlignment{BarCenter, BarTop, BarBottom}
	for _, align := range alignments {
		slider := New(nil, WithHorizontalBarAlignment(align))
		if slider.horizontalBarAlignment != align {
			t.Errorf("Expected horizontal bar alignment %v, got %v", align, slider.horizontalBarAlignment)
		}
	}
}

func TestWithTitleAlignment(t *testing.T) {
	alignments := []TitleAlignment{TitleAlignLeft, TitleAlignCenter, TitleAlignRight}
	for _, align := range alignments {
		slider := New(nil, WithTitleAlignment(align))
		if slider.titleAlignment != align {
			t.Errorf("Expected title alignment %v, got %v", align, slider.titleAlignment)
		}
	}
}

func TestWithVerticalValueAlignment(t *testing.T) {
	alignments := []VerticalValueAlignment{VValueCenter, VValueLeft, VValueRight}
	for _, align := range alignments {
		slider := New(nil, WithVerticalValueAlignment(align))
		if slider.verticalValueAlignment != align {
			t.Errorf("Expected vertical value alignment %v, got %v", align, slider.verticalValueAlignment)
		}
	}
}

func TestWithVerticalLabelPosition(t *testing.T) {
	positions := []VerticalLabelPosition{VLabelTop, VLabelBottom}
	for _, pos := range positions {
		slider := New(nil, WithVerticalLabelPosition(pos))
		if slider.verticalLabelPosition != pos {
			t.Errorf("Expected vertical label position %v, got %v", pos, slider.verticalLabelPosition)
		}
	}
}

func TestWithVerticalValuePosition(t *testing.T) {
	positions := []VerticalValuePosition{VValuePosBottom, VValuePosTop, VValuePosMiddle}
	for _, pos := range positions {
		slider := New(nil, WithVerticalValuePosition(pos))
		if slider.verticalValuePosition != pos {
			t.Errorf("Expected vertical value position %v, got %v", pos, slider.verticalValuePosition)
		}
	}
}

func TestWithValueAlignment(t *testing.T) {
	alignments := []ValueAlignment{AlignLeft, AlignCenter, AlignRight}
	for _, align := range alignments {
		slider := New(nil, WithValueAlignment(align))
		if slider.valueAlignment != align {
			t.Errorf("Expected value alignment %v, got %v", align, slider.valueAlignment)
		}
	}
}

func TestWithBorder(t *testing.T) {
	styles := []BorderStyle{BorderNone, BorderRounded, BorderNormal, BorderThick, BorderDouble}
	for _, style := range styles {
		slider := New(nil, WithBorder(style))
		if slider.borderStyle != style {
			t.Errorf("Expected border style %v, got %v", style, slider.borderStyle)
		}
	}
}

func TestWithBorderTitle(t *testing.T) {
	slider := New(nil, WithBorderTitle("Settings"))
	if slider.borderTitle != "Settings" {
		t.Errorf("Expected border title 'Settings', got '%s'", slider.borderTitle)
	}
}

func TestWithBorderColor(t *testing.T) {
	color := lipgloss.Color("196")
	slider := New(nil, WithBorderColor(color))
	if slider.borderColor != color {
		t.Errorf("Expected border color %v, got %v", color, slider.borderColor)
	}
}

func TestWithSegmented(t *testing.T) {
	slider := New(nil, WithSegmented(true))
	if !slider.segmented {
		t.Error("Expected segmented to be true")
	}
}

func TestWithSegmentCount(t *testing.T) {
	slider := New(nil, WithSegmentCount(15))
	if slider.segmentCount != 15 {
		t.Errorf("Expected segment count 15, got %d", slider.segmentCount)
	}

	// Test invalid count (should not change)
	slider = New(nil, WithSegmentCount(-5))
	if slider.segmentCount != 0 {
		t.Errorf("Expected default segment count 0 for invalid input, got %d", slider.segmentCount)
	}
}

func TestWithSegmentGap(t *testing.T) {
	slider := New(nil, WithSegmentGap(2))
	if slider.segmentGap != 2 {
		t.Errorf("Expected segment gap 2, got %d", slider.segmentGap)
	}

	// Test invalid gap (should not change)
	slider = New(nil, WithSegmentGap(-1))
	if slider.segmentGap != 1 {
		t.Errorf("Expected default segment gap 1 for invalid input, got %d", slider.segmentGap)
	}
}

func TestWithStyle(t *testing.T) {
	style := StyleNeon()
	slider := New(nil, WithStyle(style))
	
	if slider.symbols != style.Symbols {
		t.Error("Expected symbols to match style")
	}
}

func TestWithSymbolSet(t *testing.T) {
	slider := New(nil, WithSymbolSet(SymbolSetStars))
	
	if slider.symbols.Filled != SymbolSetStars.Filled {
		t.Errorf("Expected filled symbol '%s', got '%s'", SymbolSetStars.Filled, slider.symbols.Filled)
	}
}

func TestState(t *testing.T) {
	state := NewState(WithValue(42))
	slider := New(state)
	
	if slider.State() != state {
		t.Error("State() should return the slider's state")
	}
}

func TestSetState(t *testing.T) {
	slider := New(nil)
	newState := NewState(WithValue(99))
	slider.SetState(newState)
	
	if slider.state != newState {
		t.Error("SetState should update the slider's state")
	}
}

func TestView_HorizontalBasic(t *testing.T) {
	state := NewState(WithValue(50), WithMax(100))
	slider := New(state, WithWidth(10), WithShowValue(true))
	
	view := slider.View()
	
	// Should contain the value "50"
	if !strings.Contains(view, "50") {
		t.Errorf("Expected view to contain '50', got: %s", view)
	}
}

func TestView_VerticalBasic(t *testing.T) {
	state := NewState(WithValue(50), WithMax(100))
	slider := New(state, 
		WithHeight(5), 
		WithOrientation(Vertical),
		WithShowValue(true),
	)
	
	view := slider.View()
	
	// Should contain the value "50"
	if !strings.Contains(view, "50") {
		t.Errorf("Expected view to contain '50', got: %s", view)
	}
	
	// Should have multiple lines
	lines := strings.Split(view, "\n")
	if len(lines) < 2 {
		t.Error("Expected vertical slider to have multiple lines")
	}
}

func TestView_WithLabel(t *testing.T) {
	state := NewState(WithValue(50))
	slider := New(state, 
		WithWidth(10), 
		WithLabel("Volume"),
		WithLabelPosition(LabelLeft),
	)
	
	view := slider.View()
	
	if !strings.Contains(view, "Volume") {
		t.Errorf("Expected view to contain 'Volume', got: %s", view)
	}
}

func TestView_WithBorder(t *testing.T) {
	state := NewState(WithValue(50))
	slider := New(state, 
		WithWidth(10), 
		WithBorder(BorderRounded),
	)
	
	view := slider.View()
	
	// Should contain border characters
	if !strings.Contains(view, "╭") && !strings.Contains(view, "╮") {
		t.Errorf("Expected view to contain rounded border characters, got: %s", view)
	}
}

func TestView_Segmented(t *testing.T) {
	state := NewState(WithValue(50))
	slider := New(state, 
		WithWidth(20), 
		WithSegmented(true),
		WithSegmentCount(5),
	)
	
	view := slider.View()
	
	// Should render something
	if len(view) == 0 {
		t.Error("Expected segmented slider to produce output")
	}
}

func TestView_NoHandle(t *testing.T) {
	state := NewState(WithValue(50))
	slider := New(state, 
		WithWidth(10), 
		WithHandle(false),
		WithSymbols(Symbols{Filled: "X", Empty: "-", Handle: "O"}),
	)
	
	view := slider.View()
	
	// Should not contain the handle symbol
	if strings.Contains(view, "O") {
		t.Errorf("Expected view to not contain handle 'O' when handle is disabled, got: %s", view)
	}
}

func TestView_CustomFormat(t *testing.T) {
	state := NewState(WithValue(50.5))
	slider := New(state, 
		WithWidth(10), 
		WithShowValue(true),
		WithValueFormat("%.1f%%"),
	)
	
	view := slider.View()
	
	// Should contain formatted value
	if !strings.Contains(view, "50.5%") {
		t.Errorf("Expected view to contain '50.5%%', got: %s", view)
	}
}

func TestString(t *testing.T) {
	state := NewState(WithValue(50))
	slider := New(state, WithWidth(10))
	
	// String() should return the same as View()
	if slider.String() != slider.View() {
		t.Error("String() should return the same as View()")
	}
}

func TestFormatValue_Integer(t *testing.T) {
	state := NewState(WithValue(100))
	slider := New(state, WithShowValue(true))
	
	view := slider.View()
	
	// Should show "100" not "1" or "100.0"
	if !strings.Contains(view, "100") {
		t.Errorf("Expected view to contain '100', got: %s", view)
	}
	if strings.Contains(view, "100.0") {
		t.Errorf("Expected integer to be formatted without decimal, got: %s", view)
	}
}

func TestFormatValue_Decimal(t *testing.T) {
	state := NewState(WithValue(50.5))
	slider := New(state, WithShowValue(true))
	
	view := slider.View()
	
	// Should show "50.5"
	if !strings.Contains(view, "50.5") {
		t.Errorf("Expected view to contain '50.5', got: %s", view)
	}
}

func TestDefaultSymbols(t *testing.T) {
	symbols := DefaultSymbols()
	
	if symbols.Filled != "█" {
		t.Errorf("Expected default filled '█', got '%s'", symbols.Filled)
	}
	if symbols.Empty != "░" {
		t.Errorf("Expected default empty '░', got '%s'", symbols.Empty)
	}
	if symbols.Handle != "●" {
		t.Errorf("Expected default handle '●', got '%s'", symbols.Handle)
	}
}

func TestASCIISymbols(t *testing.T) {
	symbols := ASCIISymbols()
	
	if symbols.Filled != "=" {
		t.Errorf("Expected ASCII filled '=', got '%s'", symbols.Filled)
	}
	if symbols.Empty != "-" {
		t.Errorf("Expected ASCII empty '-', got '%s'", symbols.Empty)
	}
	if symbols.Handle != "O" {
		t.Errorf("Expected ASCII handle 'O', got '%s'", symbols.Handle)
	}
}

func TestBlockSymbols(t *testing.T) {
	symbols := BlockSymbols()
	
	if symbols.Filled != "█" {
		t.Errorf("Expected block filled '█', got '%s'", symbols.Filled)
	}
	if symbols.Empty != "▒" {
		t.Errorf("Expected block empty '▒', got '%s'", symbols.Empty)
	}
	if symbols.Handle != "█" {
		t.Errorf("Expected block handle '█', got '%s'", symbols.Handle)
	}
}

func TestMultipleOptions(t *testing.T) {
	state := NewState(WithValue(75))
	slider := New(state,
		WithWidth(30),
		WithHeight(15),
		WithOrientation(Horizontal),
		WithLabel("Test"),
		WithLabelPosition(LabelLeft),
		WithShowValue(true),
		WithValuePosition(ValueRight),
		WithHandle(true),
		WithBorder(BorderRounded),
		WithBorderColor(lipgloss.Color("86")),
		WithStyle(StyleOcean()),
	)
	
	// Verify all options were applied
	if slider.width != 30 {
		t.Errorf("Expected width 30, got %d", slider.width)
	}
	if slider.height != 15 {
		t.Errorf("Expected height 15, got %d", slider.height)
	}
	if slider.label != "Test" {
		t.Errorf("Expected label 'Test', got '%s'", slider.label)
	}
	if slider.labelPosition != LabelLeft {
		t.Error("Expected label position LabelLeft")
	}
	if !slider.showValue {
		t.Error("Expected showValue to be true")
	}
	if slider.borderStyle != BorderRounded {
		t.Error("Expected border style BorderRounded")
	}
}
