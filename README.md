# TuSlide

[![Go Reference](https://pkg.go.dev/badge/github.com/ohmymex/tuslide.svg)](https://pkg.go.dev/github.com/ohmymex/tuslide)
[![Go Report Card](https://goreportcard.com/badge/github.com/ohmymex/tuslide)](https://goreportcard.com/report/github.com/ohmymex/tuslide)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A feature-rich TUI slider component library for Go, inspired by [tui-slider](https://github.com/orhun/tui-slider) for Rust. Built for use with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and styled with [Lip Gloss](https://github.com/charmbracelet/lipgloss).

```
Volume ━━━━━━━━━━━━━●─────────────── 65
```

## Features

- **Horizontal & Vertical Sliders** - Full support for both orientations
- **46+ Style Presets** - Ocean, Neon, Forest, Gaming UI, and more
- **35+ Symbol Sets** - Blocks, dots, waves, gradients, ASCII-compatible
- **Segmented Mode** - Discrete segments with customizable gaps
- **Progress Bar Mode** - Hide the handle for progress indicators
- **Border Support** - Rounded, normal, thick, and double borders
- **Flexible Positioning** - Labels and values can be placed anywhere
- **Mouse Support** - Click and drag interaction with slider groups
- **Animation Helpers** - 13 easing functions, spring physics, pulse effects
- **Accessibility** - High contrast, ASCII-only, screen reader modes
- **Unicode Accurate** - Proper width handling with `go-runewidth`
- **Bubble Tea Compatible** - Works seamlessly with the Bubble Tea framework

## Installation

```bash
go get github.com/ohmymex/tuslide
```

## Quick Start

### Basic Slider

```go
package main

import (
    "fmt"
    "github.com/ohmymex/tuslide"
)

func main() {
    // Create state
    state := tuslide.NewState(
        tuslide.WithMin(0),
        tuslide.WithMax(100),
        tuslide.WithValue(50),
        tuslide.WithStep(5),
    )

    // Create slider
    slider := tuslide.New(state,
        tuslide.WithWidth(30),
        tuslide.WithLabel("Volume"),
        tuslide.WithLabelPosition(tuslide.LabelLeft),
        tuslide.WithShowValue(true),
    )

    fmt.Println(slider.View())
    // Output: Volume ━━━━━━━━━━━━━━●─────────────── 50
}
```

### With Bubble Tea

```go
package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/ohmymex/tuslide"
)

type model struct {
    slider *tuslide.Slider
    state  *tuslide.SliderState
}

func initialModel() model {
    state := tuslide.NewState(
        tuslide.WithMin(0),
        tuslide.WithMax(100),
        tuslide.WithValue(50),
    )
    
    slider := tuslide.New(state,
        tuslide.WithWidth(40),
        tuslide.WithShowValue(true),
        tuslide.WithStyle(tuslide.StyleNeon()),
    )
    
    return model{slider: slider, state: state}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        case "left", "h":
            m.state.Decrement()
        case "right", "l":
            m.state.Increment()
        }
    }
    return m, nil
}

func (m model) View() string {
    return fmt.Sprintf("\n  %s\n\n  Press ←/→ to adjust, q to quit\n", m.slider.View())
}

func main() {
    if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
```

## Style Presets

TuSlide comes with 46+ predefined styles. See **[STYLEPRESETS.md](STYLEPRESETS.md)** for the complete list.

Quick examples:
```go
tuslide.WithStyle(tuslide.StyleNeon())      // Vibrant magenta/cyan
tuslide.WithStyle(tuslide.StyleHealth())    // Red health bar
tuslide.WithStyle(tuslide.StyleOcean())     // Deep blue tones
tuslide.WithStyle(tuslide.StyleSegmentedStars()) // ★ ★ ★ ☆ ☆
```

## Vertical Sliders

Perfect for equalizers and level meters:

```go
slider := tuslide.New(state,
    tuslide.WithHeight(10),
    tuslide.WithOrientation(tuslide.Vertical),
    tuslide.WithLabel("Bass"),
    tuslide.WithLabelPosition(tuslide.LabelBottom),
    tuslide.WithShowValue(true),
    tuslide.WithStyle(tuslide.StyleVerticalEqualizer()),
)
```

## Borders

Wrap sliders with decorative borders:

```go
slider := tuslide.New(state,
    tuslide.WithWidth(25),
    tuslide.WithBorder(tuslide.BorderRounded),
    tuslide.WithBorderColor(lipgloss.Color("86")),
    tuslide.WithShowValue(true),
)
```

Available border styles:
- `BorderNone` - No border
- `BorderRounded` - Rounded corners (╭╮╰╯)
- `BorderNormal` - Normal box (┌┐└┘)
- `BorderThick` - Thick lines (┏┓┗┛)
- `BorderDouble` - Double lines (╔╗╚╝)

## Segmented Mode

Create discrete segment sliders:

```go
slider := tuslide.New(state,
    tuslide.WithWidth(40),
    tuslide.WithSegmented(true),
    tuslide.WithSegmentCount(10),
    tuslide.WithSegmentGap(1),
    tuslide.WithStyle(tuslide.StyleSegmentedStars()),
)
```

## Custom Symbols

Use any Unicode characters:

```go
slider := tuslide.New(state,
    tuslide.WithSymbols(tuslide.Symbols{
        Filled: "🟢",
        Empty:  "⚪",
        Handle: "🔘",
    }),
)
```

Or use predefined symbol sets:

```go
tuslide.WithSymbolSet(tuslide.SymbolSetStars)
tuslide.WithSymbolSet(tuslide.SymbolSetCircles)
tuslide.WithSymbolSet(tuslide.SymbolSetDiamonds)
```

## Custom Colors with Lip Gloss

```go
slider := tuslide.New(state,
    tuslide.WithFilledStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6B6B"))),
    tuslide.WithEmptyStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#4A4A4A"))),
    tuslide.WithHandleStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))),
)
```

## State Management

The `SliderState` manages the slider's value:

```go
state := tuslide.NewState(
    tuslide.WithMin(0),
    tuslide.WithMax(100),
    tuslide.WithValue(50),
    tuslide.WithStep(5),
)

// Get values
state.Value()       // 50
state.Percentage()  // 0.5
state.Min()         // 0
state.Max()         // 100

// Set values
state.SetValue(75)
state.SetFromPercentage(0.8)
state.Increment()   // +5
state.Decrement()   // -5
```

## Label & Value Positioning

### Label Positions
```go
tuslide.LabelNone    // Hidden
tuslide.LabelLeft    // To the left
tuslide.LabelRight   // To the right
tuslide.LabelTop     // Above
tuslide.LabelBottom  // Below
```

### Value Positions
```go
tuslide.ValueLeft    // To the left
tuslide.ValueRight   // To the right (default)
tuslide.ValueTop     // Above
tuslide.ValueBottom  // Below
tuslide.ValueInline  // Inside track
```

### Vertical Slider Positioning
```go
tuslide.WithVerticalLabelPosition(tuslide.VLabelTop)
tuslide.WithVerticalLabelPosition(tuslide.VLabelBottom)

tuslide.WithVerticalValuePosition(tuslide.VValuePosTop)
tuslide.WithVerticalValuePosition(tuslide.VValuePosMiddle)
tuslide.WithVerticalValuePosition(tuslide.VValuePosBottom)

tuslide.WithVerticalValueAlignment(tuslide.VValueLeft)
tuslide.WithVerticalValueAlignment(tuslide.VValueCenter)
tuslide.WithVerticalValueAlignment(tuslide.VValueRight)
```

## Mouse Support

TuSlide includes built-in mouse support for click-and-drag slider interaction:

### Enabling Mouse

```go
import tea "github.com/charmbracelet/bubbletea"

// In your main function:
p := tea.NewProgram(
    yourModel,
    tuslide.EnableMouse(), // Enable mouse cell motion tracking
)
```

### Using MouseState

```go
type model struct {
    state      *tuslide.SliderState
    slider     *tuslide.Slider
    mouseState *tuslide.MouseState
}

func newModel() model {
    state := tuslide.NewState()
    return model{
        state:      state,
        slider:     tuslide.New(state, tuslide.WithWidth(30)),
        mouseState: tuslide.NewMouseState(),
    }
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.MouseMsg:
        // Set bounds based on where the slider is rendered
        m.mouseState.SetBounds(10, 5, 30, 1) // x, y, width, height
        
        if m.mouseState.HandleMouse(msg, m.slider) {
            // Slider was interacted with
            // m.state.Value() is automatically updated
        }
    }
    return m, nil
}
```

### SliderGroup for Multiple Sliders

```go
func newModel() model {
    group := tuslide.NewSliderGroup()
    
    // Add sliders to the group
    redState := tuslide.NewState()
    redSlider := tuslide.New(redState)
    group.Add(redSlider)
    group.SetBounds(0, 0, 0, 30, 1)
    
    greenState := tuslide.NewState()
    greenSlider := tuslide.New(greenState)
    group.Add(greenSlider)
    group.SetBounds(1, 0, 2, 30, 1)
    
    return model{group: group}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.MouseMsg:
        if m.group.HandleMouse(msg) {
            // One of the sliders was interacted with
            focused := m.group.Focused() // Index of focused slider
        }
    }
    return m, nil
}
```

## Animation Helpers

TuSlide provides animation utilities for smooth value transitions:

### Basic Animation

```go
state := tuslide.NewState(tuslide.WithValue(0), tuslide.WithMax(100))

// Create animation to target value
anim := tuslide.NewAnimation(state, 100,
    tuslide.WithAnimDuration(500*time.Millisecond),
    tuslide.WithEasing(tuslide.EaseOutQuad),
    tuslide.WithOnComplete(func() {
        fmt.Println("Animation complete!")
    }),
)

// In your Update function:
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case tuslide.AnimationTickMsg:
        if m.anim.Update() {
            // Animation complete
            return m, nil
        }
        return m, m.anim.Tick()
    }
    return m, nil
}
```

### Easing Functions

```go
tuslide.Linear         // No easing
tuslide.EaseInQuad     // Accelerate from zero
tuslide.EaseOutQuad    // Decelerate to zero
tuslide.EaseInOutQuad  // Accelerate then decelerate
tuslide.EaseInCubic    // Cubic ease-in
tuslide.EaseOutCubic   // Cubic ease-out
tuslide.EaseInOutCubic // Cubic ease-in-out
tuslide.EaseInExpo     // Exponential ease-in
tuslide.EaseOutExpo    // Exponential ease-out
tuslide.EaseOutElastic // Elastic bounce at end
tuslide.EaseOutBounce  // Bouncing effect
```

### Animation Manager (Multiple Animations)

```go
manager := tuslide.NewAnimationManager()

// Start multiple animations
id1 := manager.Start(state1, 100, tuslide.WithAnimDuration(300*time.Millisecond))
id2 := manager.Start(state2, 50, tuslide.WithEasing(tuslide.EaseOutBounce))

// In Update:
case tuslide.AnimationTickMsg:
    if manager.Update() {
        return m, manager.Tick()
    }

// Cancel specific or all animations
manager.Cancel(id1)
manager.CancelAll()
```

### Spring Animation (Physics-Based)

```go
spring := tuslide.NewSpringAnimation(state, 100).
    WithStiffness(180).  // Higher = faster
    WithDamping(12)      // 1 = critical damping

// Change target dynamically
spring.SetTarget(50)

// In Update:
case tuslide.AnimationTickMsg:
    if spring.Update() {
        // At rest
    }
    return m, spring.Tick()
```

### Pulse Animation (Oscillating)

```go
// Create pulsing effect: amplitude=10, frequency=2Hz, duration=2s
pulse := tuslide.NewPulseAnimation(state, 10, 2, 2*time.Second)

// For infinite pulse, use duration=0
infinitePulse := tuslide.NewPulseAnimation(state, 5, 1, 0)
```

### Utility Functions

```go
tuslide.Lerp(0, 100, 0.5)     // Linear interpolation → 50
tuslide.Clamp(150, 0, 100)    // Clamp to range → 100
tuslide.SmoothStep(0.5)       // Smooth Hermite interpolation
tuslide.SmootherStep(0.5)     // Ken Perlin's smoother version
```

## Accessibility Features

TuSlide includes built-in accessibility support:

### Accessibility Modes

```go
// High contrast mode - maximum visibility
accessible := tuslide.NewAccessibleSlider(slider,
    tuslide.WithAccessibilityMode(tuslide.AccessibilityHighContrast),
)

// ASCII-only mode - for terminals without Unicode
accessible := tuslide.NewAccessibleSlider(slider,
    tuslide.WithAccessibilityMode(tuslide.AccessibilityASCII),
)

// Screen reader optimized
accessible := tuslide.NewAccessibleSlider(slider,
    tuslide.WithAccessibilityMode(tuslide.AccessibilityScreenReader),
)
```

### Value Announcements

```go
accessible := tuslide.NewAccessibleSlider(slider,
    tuslide.WithAnnouncer(func(msg string) {
        // Send to screen reader or status bar
        fmt.Println("Announcement:", msg)
    }),
    tuslide.WithDescription("Volume control, 0 to 100"),
)

// Increment announces the new value automatically
accessible.Increment()

// Get descriptions for screen readers
desc := accessible.GetDescription()          // "Volume: 50 of 100 (50%)"
announcement := accessible.GetValueAnnouncement() // "50, 50 percent"
```

### High Contrast Palettes

```go
// Apply predefined high contrast palettes
tuslide.ApplyPalette(slider, tuslide.DefaultHighContrastPalette())
tuslide.ApplyPalette(slider, tuslide.DarkHighContrastPalette())
tuslide.ApplyPalette(slider, tuslide.LightHighContrastPalette())
```

### Focus Indicators

```go
fi := tuslide.NewFocusIndicator().
    WithChar("▸").
    WithPosition("left") // "left", "right", or "both"

// Wrap content with focus indicator
output := fi.Wrap(slider.View(), isFocused)
```

### Progress Announcements

```go
// Announce progress at intervals (e.g., every 25%)
announcer := tuslide.NewProgressAnnouncer(25, func(msg string) {
    fmt.Println(msg) // "25 percent complete", "50 percent complete", etc.
})

// Call during progress updates
announcer.Update(state)
announcer.Reset() // For new operations
```

### Keyboard Hints

```go
hints := tuslide.DefaultKeyboardHints()
fmt.Println(hints.Render())       // Full keyboard hints
fmt.Println(hints.RenderCompact()) // Compact version
```

## All Options

| Option | Description |
|--------|-------------|
| `WithWidth(int)` | Set horizontal track width |
| `WithHeight(int)` | Set vertical track height |
| `WithOrientation(Orientation)` | Horizontal or Vertical |
| `WithSymbols(Symbols)` | Custom filled/empty/handle characters |
| `WithSymbolSet(SymbolSet)` | Use predefined symbol set |
| `WithStyle(SliderStyle)` | Apply complete style preset |
| `WithHandle(bool)` | Show/hide handle (false for progress bars) |
| `WithLabel(string)` | Set label text |
| `WithLabelPosition(LabelPosition)` | Label placement |
| `WithShowValue(bool)` | Show numeric value |
| `WithValuePosition(ValuePosition)` | Value placement |
| `WithValueFormat(string)` | Custom value format (e.g., "%.1f%%") |
| `WithValueAlignment(ValueAlignment)` | Left/Center/Right alignment |
| `WithCollisionCheck(bool)` | Auto-reposition overlapping text |
| `WithHorizontalBarAlignment(HorizontalBarAlignment)` | Vertical position of bar |
| `WithTitleAlignment(TitleAlignment)` | Title text alignment |
| `WithVerticalValueAlignment(VerticalValueAlignment)` | Value alignment in vertical sliders |
| `WithVerticalLabelPosition(VerticalLabelPosition)` | Label position in vertical sliders |
| `WithVerticalValuePosition(VerticalValuePosition)` | Value position in vertical sliders |
| `WithBorder(BorderStyle)` | Add border around slider |
| `WithBorderTitle(string)` | Border title text |
| `WithBorderColor(lipgloss.Color)` | Border color |
| `WithSegmented(bool)` | Enable segmented mode |
| `WithSegmentCount(int)` | Number of segments |
| `WithSegmentGap(int)` | Gap between segments |
| `WithFilledStyle(lipgloss.Style)` | Style for filled portion |
| `WithEmptyStyle(lipgloss.Style)` | Style for empty portion |
| `WithHandleStyle(lipgloss.Style)` | Style for handle |
| `WithLabelStyle(lipgloss.Style)` | Style for label |
| `WithValueStyle(lipgloss.Style)` | Style for value |

## Examples

Run the interactive showcase:

```bash
go run ./examples/showcase
```

Run the basic example:

```bash
go run ./examples/basic
```

## Credits

- Inspired by [tui-slider](https://github.com/orhun/tui-slider) for Rust
- Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- Unicode width handling by [go-runewidth](https://github.com/mattn/go-runewidth)

## License

MIT License - see [LICENSE](LICENSE) for details.
