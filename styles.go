package tuslide

import "github.com/charmbracelet/lipgloss"

// SliderStyle defines a complete visual style for a slider.
// It combines symbols with colors for a cohesive appearance.
type SliderStyle struct {
	Name         string
	Symbols      Symbols
	FilledStyle  lipgloss.Style
	EmptyStyle   lipgloss.Style
	HandleStyle  lipgloss.Style
	LabelStyle   lipgloss.Style
	ValueStyle   lipgloss.Style
	Segmented    bool // Whether to render as discrete segments
}

// Apply applies this style to a slider via functional options.
func (s SliderStyle) Apply() []SliderOption {
	return []SliderOption{
		WithSymbols(s.Symbols),
		WithFilledStyle(s.FilledStyle),
		WithEmptyStyle(s.EmptyStyle),
		WithHandleStyle(s.HandleStyle),
		WithLabelStyle(s.LabelStyle),
		WithValueStyle(s.ValueStyle),
	}
}

// ============================================================================
// PREDEFINED STYLES
// ============================================================================

// StyleDefault is the default clean and professional style.
func StyleDefault() SliderStyle {
	return SliderStyle{
		Name:        "Default",
		Symbols:     SymbolSetDefault.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("86")),  // Cyan
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleBlock is a bold and solid block style.
func StyleBlock() SliderStyle {
	return SliderStyle{
		Name:        "Block",
		Symbols:     SymbolSetBlock.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("46")),  // Green
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleDots uses braille patterns with yellow fills.
func StyleDots() SliderStyle {
	return SliderStyle{
		Name:        "Dots",
		Symbols:     SymbolSetDots.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("226")), // Yellow
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleMinimal is clean and subtle with blue tones.
func StyleMinimal() SliderStyle {
	return SliderStyle{
		Name:        "Minimal",
		Symbols:     SymbolSetMinimal.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("33")),  // Blue
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("87")),  // Cyan
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleDoubleLine has a formal appearance with red fills.
func StyleDoubleLine() SliderStyle {
	return SliderStyle{
		Name:        "Double Line",
		Symbols:     SymbolSetDoubleLine.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("196")), // Red
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleWave has a fluid cyan appearance.
func StyleWave() SliderStyle {
	return SliderStyle{
		Name:        "Wave",
		Symbols:     SymbolSetWave.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("87")),  // Cyan
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleProgress is a progress bar style with green fills.
func StyleProgress() SliderStyle {
	return SliderStyle{
		Name:        "Progress",
		Symbols:     SymbolSetProgress.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("46")),  // Green
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("226")), // Yellow
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleThick is bold with magenta fills.
func StyleThick() SliderStyle {
	return SliderStyle{
		Name:        "Thick",
		Symbols:     SymbolSetThick.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("201")), // Magenta
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("238")), // Darker gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleGradient has a shaded blue effect.
func StyleGradient() SliderStyle {
	return SliderStyle{
		Name:        "Gradient",
		Symbols:     SymbolSetGradient.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("33")),  // Blue
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("87")),  // Cyan
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleRetro is old-school ASCII with green fills.
func StyleRetro() SliderStyle {
	return SliderStyle{
		Name:        "Retro",
		Symbols:     SymbolSetRetro.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("46")),  // Green
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// ============================================================================
// PROGRESS BAR STYLES (No handle)
// ============================================================================

// StyleProgressDownload is a download progress bar style.
func StyleProgressDownload() SliderStyle {
	return SliderStyle{
		Name: "Download",
		Symbols: Symbols{
			Filled: FilledBlock,
			Empty:  EmptyLightShade,
			Handle: "",
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("46")),  // Green
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle(),
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleProgressUpload is an upload progress bar style.
func StyleProgressUpload() SliderStyle {
	return SliderStyle{
		Name: "Upload",
		Symbols: Symbols{
			Filled: FilledProgress,
			Empty:  EmptyProgress,
			Handle: "",
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("33")),  // Blue
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle(),
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleHealth is a health bar style for gaming UIs.
func StyleHealth() SliderStyle {
	return SliderStyle{
		Name: "Health",
		Symbols: Symbols{
			Filled: FilledMediumShade,
			Empty:  EmptyLightShade,
			Handle: "",
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("196")), // Red
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("236")), // Very dark gray
		HandleStyle: lipgloss.NewStyle(),
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleMana is a mana bar style for gaming UIs.
func StyleMana() SliderStyle {
	return SliderStyle{
		Name: "Mana",
		Symbols: Symbols{
			Filled: FilledMediumShade,
			Empty:  EmptyLightShade,
			Handle: "",
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("87")),  // Cyan
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("236")), // Very dark gray
		HandleStyle: lipgloss.NewStyle(),
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleExperience is an experience bar style for gaming UIs.
func StyleExperience() SliderStyle {
	return SliderStyle{
		Name: "Experience",
		Symbols: Symbols{
			Filled: FilledThickLine,
			Empty:  EmptyThinLine,
			Handle: "",
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("226")), // Yellow
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle(),
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// ============================================================================
// SEGMENTED STYLES
// ============================================================================

// StyleSegmented is a discrete segments style.
func StyleSegmented() SliderStyle {
	return SliderStyle{
		Name: "Segmented",
		Symbols: Symbols{
			Filled: "─",
			Empty:  "─",
			Handle: HandleCircle,
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("196")), // Red
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
		Segmented:   true,
	}
}

// StyleSegmentedBlocks uses block segments.
func StyleSegmentedBlocks() SliderStyle {
	return SliderStyle{
		Name: "Segmented Blocks",
		Symbols: Symbols{
			Filled: FilledBlock,
			Empty:  EmptyLightShade,
			Handle: HandleSquare,
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("46")),  // Green
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
		Segmented:   true,
	}
}

// StyleSegmentedDots uses dot segments.
func StyleSegmentedDots() SliderStyle {
	return SliderStyle{
		Name: "Segmented Dots",
		Symbols: Symbols{
			Filled: FilledCircle,
			Empty:  EmptyCircle,
			Handle: HandleDiamond,
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("87")),  // Cyan
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("226")), // Yellow
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
		Segmented:   true,
	}
}

// StyleSegmentedStars uses star segments.
func StyleSegmentedStars() SliderStyle {
	return SliderStyle{
		Name: "Segmented Stars",
		Symbols: Symbols{
			Filled: FilledStar,
			Empty:  EmptyStar,
			Handle: HandleSparkle,
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("226")), // Yellow
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("87")),  // Cyan
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
		Segmented:   true,
	}
}

// StyleSegmentedSquares uses square segments.
func StyleSegmentedSquares() SliderStyle {
	return SliderStyle{
		Name: "Segmented Squares",
		Symbols: Symbols{
			Filled: FilledSquare,
			Empty:  EmptySquare,
			Handle: HandleDoubleCircle,
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("33")),  // Blue
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("87")),  // Cyan
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
		Segmented:   true,
	}
}

// StyleSegmentedDiamonds uses diamond segments.
func StyleSegmentedDiamonds() SliderStyle {
	return SliderStyle{
		Name: "Segmented Diamonds",
		Symbols: Symbols{
			Filled: FilledDiamond,
			Empty:  EmptyDiamond,
			Handle: HandleHexagon,
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("226")), // Yellow
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
		Segmented:   true,
	}
}

// ============================================================================
// COLOR THEME STYLES
// ============================================================================

// StyleOcean is a blue ocean theme.
func StyleOcean() SliderStyle {
	return SliderStyle{
		Name:        "Ocean",
		Symbols:     SymbolSetDefault.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("39")),  // Deep sky blue
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("24")),  // Dark blue
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("123")), // Light blue
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("159")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("117")),
	}
}

// StyleForest is a green forest theme.
func StyleForest() SliderStyle {
	return SliderStyle{
		Name:        "Forest",
		Symbols:     SymbolSetDefault.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("34")),  // Forest green
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("22")),  // Dark green
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("118")), // Light green
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("157")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("120")),
	}
}

// StyleSunset is a warm orange/red theme.
func StyleSunset() SliderStyle {
	return SliderStyle{
		Name:        "Sunset",
		Symbols:     SymbolSetDefault.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("208")), // Orange
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("94")),  // Dark orange
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("220")), // Gold
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("223")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("215")),
	}
}

// StyleNeon is a vibrant neon theme.
func StyleNeon() SliderStyle {
	return SliderStyle{
		Name:        "Neon",
		Symbols:     SymbolSetDefault.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("201")), // Magenta
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("53")),  // Dark magenta
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("51")),  // Cyan
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("207")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("213")),
	}
}

// StyleMonochrome is a grayscale theme.
func StyleMonochrome() SliderStyle {
	return SliderStyle{
		Name:        "Monochrome",
		Symbols:     SymbolSetDefault.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("252")), // Light gray
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("238")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("250")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("246")),
	}
}

// ============================================================================
// ADDITIONAL PROGRESS BAR STYLES
// ============================================================================

// StyleProgressLoading is a loading progress bar style.
func StyleProgressLoading() SliderStyle {
	return SliderStyle{
		Name: "Loading",
		Symbols: Symbols{
			Filled: FilledDoubleLine,
			Empty:  EmptyThinLine,
			Handle: "",
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("201")), // Magenta
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle(),
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleProgressInstallation is an installation progress bar style.
func StyleProgressInstallation() SliderStyle {
	return SliderStyle{
		Name: "Installation",
		Symbols: Symbols{
			Filled: FilledBar,
			Empty:  EmptyBarOutline,
			Handle: "",
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("118")), // Light green
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle(),
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleProgressBattery is a battery level bar style.
func StyleProgressBattery() SliderStyle {
	return SliderStyle{
		Name: "Battery",
		Symbols: Symbols{
			Filled: FilledSquare,
			Empty:  EmptySquare,
			Handle: "",
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("227")), // Light yellow
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle(),
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// ============================================================================
// ADDITIONAL SEGMENTED STYLES
// ============================================================================

// StyleSegmentedBars uses vertical bar segments.
func StyleSegmentedBars() SliderStyle {
	return SliderStyle{
		Name: "Segmented Bars",
		Symbols: Symbols{
			Filled: "│",
			Empty:  "┆",
			Handle: HandleTriangleRight,
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("201")), // Magenta
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
		Segmented:   true,
	}
}

// StyleSegmentedArrows uses arrow segments.
func StyleSegmentedArrows() SliderStyle {
	return SliderStyle{
		Name: "Segmented Arrows",
		Symbols: Symbols{
			Filled: "▶",
			Empty:  "▷",
			Handle: HandleTriangleRight,
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("196")), // Red
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
		Segmented:   true,
	}
}

// StyleSegmentedThick uses thick line segments.
func StyleSegmentedThick() SliderStyle {
	return SliderStyle{
		Name: "Segmented Thick",
		Symbols: Symbols{
			Filled: "━",
			Empty:  "╌",
			Handle: HandleLargeCircle,
		},
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("87")),  // Cyan
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
		Segmented:   true,
	}
}

// ============================================================================
// HORIZONTAL SLIDER STYLES
// ============================================================================

// StyleHorizontal is a clean horizontal line style.
func StyleHorizontal() SliderStyle {
	return SliderStyle{
		Name:        "Horizontal",
		Symbols:     SymbolSetHorizontal.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("87")),  // Cyan
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleHorizontalThick uses thick lines.
func StyleHorizontalThick() SliderStyle {
	return SliderStyle{
		Name:        "Horizontal Thick",
		Symbols:     SymbolSetHorizontalThick.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("87")),  // Cyan
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleHorizontalBlocks uses bold blocks.
func StyleHorizontalBlocks() SliderStyle {
	return SliderStyle{
		Name:        "Horizontal Blocks",
		Symbols:     SymbolSetHorizontalBlocks.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("46")),  // Green
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleHorizontalGradient uses shaded gradient.
func StyleHorizontalGradient() SliderStyle {
	return SliderStyle{
		Name:        "Horizontal Gradient",
		Symbols:     SymbolSetHorizontalGradient.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("201")), // Magenta
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleHorizontalDots uses dots/circles.
func StyleHorizontalDots() SliderStyle {
	return SliderStyle{
		Name:        "Horizontal Dots",
		Symbols:     SymbolSetHorizontalDots.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("226")), // Yellow
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleHorizontalSquares uses squares.
func StyleHorizontalSquares() SliderStyle {
	return SliderStyle{
		Name:        "Horizontal Squares",
		Symbols:     SymbolSetHorizontalSquares.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("33")),  // Blue
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleHorizontalDouble uses double lines.
func StyleHorizontalDouble() SliderStyle {
	return SliderStyle{
		Name:        "Horizontal Double",
		Symbols:     SymbolSetHorizontalDouble.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("123")), // Light cyan
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// ============================================================================
// VERTICAL SLIDER STYLES
// ============================================================================

// StyleVertical is a clean vertical line style.
func StyleVertical() SliderStyle {
	return SliderStyle{
		Name:        "Vertical",
		Symbols:     SymbolSetVertical.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("87")),  // Cyan
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleVerticalBlocks uses bold blocks.
func StyleVerticalBlocks() SliderStyle {
	return SliderStyle{
		Name:        "Vertical Blocks",
		Symbols:     SymbolSetVerticalBlocks.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("46")),  // Green
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleVerticalGradient uses shaded gradient.
func StyleVerticalGradient() SliderStyle {
	return SliderStyle{
		Name:        "Vertical Gradient",
		Symbols:     SymbolSetVerticalGradient.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("201")), // Magenta
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleVerticalDots uses dots/circles.
func StyleVerticalDots() SliderStyle {
	return SliderStyle{
		Name:        "Vertical Dots",
		Symbols:     SymbolSetVerticalDots.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("226")), // Yellow
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleVerticalSquares uses squares.
func StyleVerticalSquares() SliderStyle {
	return SliderStyle{
		Name:        "Vertical Squares",
		Symbols:     SymbolSetVerticalSquares.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("33")),  // Blue
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleVerticalEqualizer uses equalizer bars.
func StyleVerticalEqualizer() SliderStyle {
	return SliderStyle{
		Name:        "Equalizer",
		Symbols:     SymbolSetVerticalEqualizer.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("118")), // Light green
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// StyleRounded has a soft appearance.
func StyleRounded() SliderStyle {
	return SliderStyle{
		Name:        "Rounded",
		Symbols:     SymbolSetRounded.ToSymbols(),
		FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("226")), // Yellow
		EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // Dark gray
		HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // White
		LabelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		ValueStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("249")),
	}
}

// AllStyles returns all predefined styles.
func AllStyles() []SliderStyle {
	return []SliderStyle{
		// Basic styles
		StyleDefault(),
		StyleBlock(),
		StyleDots(),
		StyleMinimal(),
		StyleDoubleLine(),
		StyleWave(),
		StyleProgress(),
		StyleThick(),
		StyleGradient(),
		StyleRetro(),
		StyleRounded(),
		// Progress bar styles
		StyleProgressDownload(),
		StyleProgressUpload(),
		StyleHealth(),
		StyleMana(),
		StyleExperience(),
		StyleProgressLoading(),
		StyleProgressInstallation(),
		StyleProgressBattery(),
		// Segmented styles
		StyleSegmented(),
		StyleSegmentedBlocks(),
		StyleSegmentedDots(),
		StyleSegmentedStars(),
		StyleSegmentedSquares(),
		StyleSegmentedDiamonds(),
		StyleSegmentedBars(),
		StyleSegmentedArrows(),
		StyleSegmentedThick(),
		// Color themes
		StyleOcean(),
		StyleForest(),
		StyleSunset(),
		StyleNeon(),
		StyleMonochrome(),
		// Horizontal styles
		StyleHorizontal(),
		StyleHorizontalThick(),
		StyleHorizontalBlocks(),
		StyleHorizontalGradient(),
		StyleHorizontalDots(),
		StyleHorizontalSquares(),
		StyleHorizontalDouble(),
		// Vertical styles
		StyleVertical(),
		StyleVerticalBlocks(),
		StyleVerticalGradient(),
		StyleVerticalDots(),
		StyleVerticalSquares(),
		StyleVerticalEqualizer(),
	}
}
