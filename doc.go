// Package tuslide provides a flexible and customizable TUI slider component
// for terminal applications built with the Bubble Tea framework.
//
// TuSlide is a Go port of the Rust tui-slider library, adapted to idiomatic
// Go patterns. It separates state management from rendering, allowing for
// clean integration into any Bubble Tea application.
//
// # Features
//
//   - Horizontal and vertical slider orientations
//   - Customizable symbols (filled, empty, handle)
//   - Full styling support via Lip Gloss
//   - Progress bar mode (handle hidden)
//   - Label positioning (top, bottom, left, right)
//   - Unicode-accurate rendering with go-runewidth
//
// # Quick Start
//
// Create a slider state and widget:
//
//	state := tuslide.NewState(
//	    tuslide.WithMin(0),
//	    tuslide.WithMax(100),
//	    tuslide.WithValue(50),
//	)
//
//	slider := tuslide.New(state,
//	    tuslide.WithWidth(30),
//	    tuslide.WithLabel("Volume"),
//	    tuslide.WithShowValue(true),
//	)
//
// In your Bubble Tea Update function, modify the state:
//
//	case "right":
//	    state.Increment()
//	case "left":
//	    state.Decrement()
//
// In your View function, render the slider:
//
//	return slider.View()
//
// # Architecture
//
// TuSlide uses a state/widget separation pattern:
//
//   - SliderState: Manages min, max, value, step, and percentage calculations.
//     This is the "model" that holds your data.
//
//   - Slider: The widget that renders the state. Configure it with functional
//     options like WithWidth, WithSymbols, WithFilledStyle, etc.
//
// This separation allows you to:
//   - Share state between multiple views
//   - Persist state independently of rendering
//   - Apply different visual styles to the same state
//
// # Styling
//
// TuSlide integrates with Lip Gloss for styling:
//
//	slider := tuslide.New(state,
//	    tuslide.WithFilledStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("86"))),
//	    tuslide.WithEmptyStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("241"))),
//	    tuslide.WithHandleStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("205"))),
//	)
//
// # Custom Symbols
//
// Replace the default symbols with your own:
//
//	slider := tuslide.New(state,
//	    tuslide.WithSymbols(tuslide.Symbols{
//	        Filled: "▰",
//	        Empty:  "▱",
//	        Handle: "◆",
//	    }),
//	)
//
// Built-in symbol sets are available:
//   - DefaultSymbols(): █ ░ ●
//   - ASCIISymbols(): = - O
//   - BlockSymbols(): █ ▒ █
package tuslide
