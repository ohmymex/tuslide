package tuslide

// Symbol constants for slider customization.
// These provide a rich library of Unicode characters for creating
// visually distinctive sliders.

// ============================================================================
// FILLED SYMBOLS - Used for the filled portion of the slider
// ============================================================================

const (
	// FilledThickLine is the default filled symbol - thick horizontal line.
	FilledThickLine = "━"
	// FilledThinLine is a thin horizontal line.
	FilledThinLine = "─"
	// FilledDoubleLine is a double horizontal line.
	FilledDoubleLine = "═"
	// FilledBlock is a full block character.
	FilledBlock = "█"
	// FilledDarkShade is a dark shade block.
	FilledDarkShade = "▓"
	// FilledMediumShade is a medium shade block.
	FilledMediumShade = "▒"
	// FilledLightShade is a light shade block.
	FilledLightShade = "░"
	// FilledBar is a horizontal bar.
	FilledBar = "▬"
	// FilledProgress is a progress bar filled segment.
	FilledProgress = "▰"
	// FilledBraille is a braille full pattern.
	FilledBraille = "⣿"
	// FilledWave is a wave character.
	FilledWave = "≈"
	// FilledDiamond is a filled diamond.
	FilledDiamond = "◆"
	// FilledHash is a hash/number sign.
	FilledHash = "#"
	// FilledEquals is an equals sign.
	FilledEquals = "="
	// FilledLowerBar is a lower bar.
	FilledLowerBar = "▂"
	// FilledStar is a filled star.
	FilledStar = "★"
	// FilledSquare is a filled small square.
	FilledSquare = "■"
	// FilledCircle is a filled circle.
	FilledCircle = "●"
	// FilledVerticalBar is a vertical bar.
	FilledVerticalBar = "│"
	// FilledVerticalLine is for vertical sliders.
	FilledVerticalLine = "│"
)

// ============================================================================
// EMPTY SYMBOLS - Used for the unfilled portion of the slider
// ============================================================================

const (
	// EmptyThinLine is the default empty symbol - thin horizontal line.
	EmptyThinLine = "─"
	// EmptySpace is an invisible space.
	EmptySpace = " "
	// EmptyLightShade is a light shade block.
	EmptyLightShade = "░"
	// EmptyDotted is a dotted line.
	EmptyDotted = "┄"
	// EmptyDashed is a dashed line.
	EmptyDashed = "╌"
	// EmptyProgress is a progress bar empty segment.
	EmptyProgress = "▱"
	// EmptyBraille is braille lower dots.
	EmptyBraille = "⣀"
	// EmptyWave is a tilde/wave.
	EmptyWave = "˜"
	// EmptyDiamond is a white diamond.
	EmptyDiamond = "◇"
	// EmptyDot is a period/dot.
	EmptyDot = "."
	// EmptyHyphen is a hyphen.
	EmptyHyphen = "-"
	// EmptyUnderscore is an underscore.
	EmptyUnderscore = "_"
	// EmptyLowerBar is a lower bar.
	EmptyLowerBar = "▁"
	// EmptyStar is a white star.
	EmptyStar = "☆"
	// EmptyBarOutline is a horizontal bar outline.
	EmptyBarOutline = "▭"
	// EmptySquare is a small white square.
	EmptySquare = "□"
	// EmptyCircle is a white circle.
	EmptyCircle = "○"
	// EmptyVerticalBar is a vertical bar.
	EmptyVerticalBar = "│"
	// EmptyColon is a colon.
	EmptyColon = ":"
)

// ============================================================================
// HANDLE SYMBOLS - Used for the slider handle/thumb
// ============================================================================

const (
	// HandleCircle is the default handle symbol - filled circle.
	HandleCircle = "●"
	// HandleWhiteCircle is a white circle.
	HandleWhiteCircle = "○"
	// HandleDoubleCircle is a double circle.
	HandleDoubleCircle = "◉"
	// HandleLargeCircle is a large circle.
	HandleLargeCircle = "◯"
	// HandleBullseye is a bullseye.
	HandleBullseye = "◎"
	// HandleBlackCircle is a black circle.
	HandleBlackCircle = "⬤"
	// HandleSquare is a filled square.
	HandleSquare = "■"
	// HandleWhiteSquare is a white square.
	HandleWhiteSquare = "□"
	// HandleSmallSquare is a small square.
	HandleSmallSquare = "▪"
	// HandleMediumBlock is a medium shade block.
	HandleMediumBlock = "▓"
	// HandleDiamond is a filled diamond.
	HandleDiamond = "◆"
	// HandleWhiteDiamond is a white diamond.
	HandleWhiteDiamond = "◇"
	// HandleDoubleDiamond is a double diamond.
	HandleDoubleDiamond = "◈"
	// HandleTriangleRight is a right-pointing triangle.
	HandleTriangleRight = "▶"
	// HandleTriangleLeft is a left-pointing triangle.
	HandleTriangleLeft = "◀"
	// HandleTriangleUp is an up-pointing triangle.
	HandleTriangleUp = "▲"
	// HandleTriangleDown is a down-pointing triangle.
	HandleTriangleDown = "▼"
	// HandleVerticalBar is a vertical bar.
	HandleVerticalBar = "│"
	// HandlePipe is a pipe character.
	HandlePipe = "|"
	// HandleAt is an at sign.
	HandleAt = "@"
	// HandleStar is a star.
	HandleStar = "✦"
	// HandleSparkle is a sparkle.
	HandleSparkle = "✨"
	// HandleWhiteStar is a white star.
	HandleWhiteStar = "☆"
	// HandleFilledStar is a filled star.
	HandleFilledStar = "★"
	// HandleHexagon is a hexagon.
	HandleHexagon = "⬢"
	// HandleOctagon is an octagon.
	HandleOctagon = "⬣"
	// HandleHorizontalLine is for vertical sliders.
	HandleHorizontalLine = "━"
	// HandleThickLine is a thick horizontal line.
	HandleThickLine = "═"
	// HandleLowerBar is a lower bar handle.
	HandleLowerBar = "▃"
	// HandleArrowUp is an up arrow.
	HandleArrowUp = "↑"
	// HandleArrowDown is a down arrow.
	HandleArrowDown = "↓"
	// HandleArrowLeft is a left arrow.
	HandleArrowLeft = "←"
	// HandleArrowRight is a right arrow.
	HandleArrowRight = "→"
	// HandleVerticalLine is for horizontal sliders.
	HandleVerticalLine = "│"
)

// ============================================================================
// ADDITIONAL FILLED SYMBOLS
// ============================================================================

const (
	// FilledPlus is a plus sign.
	FilledPlus = "+"
	// FilledAsterisk is an asterisk.
	FilledAsterisk = "*"
	// FilledVerticalRect is a vertical rectangle.
	FilledVerticalRect = "▮"
	// FilledHorizontalLine is for horizontal sliders.
	FilledHorizontalLine = "─"
	// FilledArrowRight is a right arrow.
	FilledArrowRight = "▶"
	// FilledArrowLeft is a left arrow.
	FilledArrowLeft = "◀"
)

// ============================================================================
// PREDEFINED SYMBOL SETS
// ============================================================================

// SymbolSet is a complete symbol set for a slider style.
type SymbolSet struct {
	Name   string
	Filled string
	Empty  string
	Handle string
}

// Predefined symbol sets for common use cases.
var (
	// SymbolSetDefault is the default clean and professional style.
	SymbolSetDefault = SymbolSet{
		Name:   "Default",
		Filled: FilledThickLine,
		Empty:  EmptyThinLine,
		Handle: HandleCircle,
	}

	// SymbolSetBlock is a bold and solid block style.
	SymbolSetBlock = SymbolSet{
		Name:   "Block",
		Filled: FilledBlock,
		Empty:  FilledLightShade,
		Handle: FilledDarkShade,
	}

	// SymbolSetDots uses braille patterns.
	SymbolSetDots = SymbolSet{
		Name:   "Dots",
		Filled: FilledBraille,
		Empty:  EmptyBraille,
		Handle: HandleBlackCircle,
	}

	// SymbolSetMinimal is clean and subtle.
	SymbolSetMinimal = SymbolSet{
		Name:   "Minimal",
		Filled: FilledThinLine,
		Empty:  EmptySpace,
		Handle: HandleVerticalBar,
	}

	// SymbolSetDoubleLine has a formal appearance.
	SymbolSetDoubleLine = SymbolSet{
		Name:   "Double Line",
		Filled: FilledDoubleLine,
		Empty:  EmptyThinLine,
		Handle: HandleDoubleCircle,
	}

	// SymbolSetWave has a fluid appearance.
	SymbolSetWave = SymbolSet{
		Name:   "Wave",
		Filled: FilledWave,
		Empty:  EmptyWave,
		Handle: HandleDoubleDiamond,
	}

	// SymbolSetProgress is a progress bar style.
	SymbolSetProgress = SymbolSet{
		Name:   "Progress",
		Filled: FilledProgress,
		Empty:  EmptyProgress,
		Handle: HandleTriangleRight,
	}

	// SymbolSetThick is bold.
	SymbolSetThick = SymbolSet{
		Name:   "Thick",
		Filled: FilledBar,
		Empty:  FilledBar,
		Handle: HandleSquare,
	}

	// SymbolSetGradient has a shaded effect.
	SymbolSetGradient = SymbolSet{
		Name:   "Gradient",
		Filled: FilledDarkShade,
		Empty:  FilledLightShade,
		Handle: HandleCircle,
	}

	// SymbolSetRounded has a soft appearance.
	SymbolSetRounded = SymbolSet{
		Name:   "Rounded",
		Filled: FilledThinLine,
		Empty:  EmptyDashed,
		Handle: HandleLargeCircle,
	}

	// SymbolSetRetro is old-school ASCII.
	SymbolSetRetro = SymbolSet{
		Name:   "Retro",
		Filled: FilledHash,
		Empty:  EmptyDot,
		Handle: HandleAt,
	}

	// SymbolSetASCII uses only ASCII characters.
	SymbolSetASCII = SymbolSet{
		Name:   "ASCII",
		Filled: FilledEquals,
		Empty:  EmptyHyphen,
		Handle: "O",
	}

	// SymbolSetStars uses star characters.
	SymbolSetStars = SymbolSet{
		Name:   "Stars",
		Filled: FilledStar,
		Empty:  EmptyStar,
		Handle: HandleSparkle,
	}

	// SymbolSetSquares uses square characters.
	SymbolSetSquares = SymbolSet{
		Name:   "Squares",
		Filled: FilledSquare,
		Empty:  EmptySquare,
		Handle: HandleSquare,
	}

	// SymbolSetCircles uses circle characters.
	SymbolSetCircles = SymbolSet{
		Name:   "Circles",
		Filled: FilledCircle,
		Empty:  EmptyCircle,
		Handle: HandleDoubleCircle,
	}

	// SymbolSetDiamonds uses diamond characters.
	SymbolSetDiamonds = SymbolSet{
		Name:   "Diamonds",
		Filled: FilledDiamond,
		Empty:  EmptyDiamond,
		Handle: HandleDoubleDiamond,
	}

	// SymbolSetVertical is optimized for vertical sliders.
	SymbolSetVertical = SymbolSet{
		Name:   "Vertical",
		Filled: FilledBlock,
		Empty:  EmptyLightShade,
		Handle: HandleHorizontalLine,
	}

	// SymbolSetNeon has a modern neon look.
	SymbolSetNeon = SymbolSet{
		Name:   "Neon",
		Filled: FilledLowerBar,
		Empty:  EmptyLowerBar,
		Handle: HandleLowerBar,
	}

	// SymbolSetArrow has a directional look.
	SymbolSetArrow = SymbolSet{
		Name:   "Arrow",
		Filled: FilledBar,
		Empty:  EmptyBarOutline,
		Handle: HandleDiamond,
	}

	// SymbolSetSegmented is for discrete segments.
	SymbolSetSegmented = SymbolSet{
		Name:   "Segmented",
		Filled: FilledThinLine,
		Empty:  EmptySpace,
		Handle: HandleCircle,
	}

	// SymbolSetSegmentedBlocks uses vertical bars.
	SymbolSetSegmentedBlocks = SymbolSet{
		Name:   "Segmented Blocks",
		Filled: FilledVerticalBar,
		Empty:  EmptyVerticalBar,
		Handle: HandleCircle,
	}

	// SymbolSetSegmentedDots uses circles.
	SymbolSetSegmentedDots = SymbolSet{
		Name:   "Segmented Dots",
		Filled: FilledCircle,
		Empty:  EmptyCircle,
		Handle: HandleCircle,
	}

	// SymbolSetSegmentedSquares uses squares.
	SymbolSetSegmentedSquares = SymbolSet{
		Name:   "Segmented Squares",
		Filled: FilledSquare,
		Empty:  EmptySquare,
		Handle: HandleCircle,
	}

	// SymbolSetHorizontal is for horizontal sliders.
	SymbolSetHorizontal = SymbolSet{
		Name:   "Horizontal",
		Filled: FilledHorizontalLine,
		Empty:  EmptyThinLine,
		Handle: HandleVerticalLine,
	}

	// SymbolSetHorizontalThick uses thick lines.
	SymbolSetHorizontalThick = SymbolSet{
		Name:   "Horizontal Thick",
		Filled: FilledThickLine,
		Empty:  EmptyThinLine,
		Handle: HandleCircle,
	}

	// SymbolSetHorizontalBlocks uses blocks.
	SymbolSetHorizontalBlocks = SymbolSet{
		Name:   "Horizontal Blocks",
		Filled: FilledBlock,
		Empty:  FilledLightShade,
		Handle: HandleCircle,
	}

	// SymbolSetHorizontalGradient uses shading.
	SymbolSetHorizontalGradient = SymbolSet{
		Name:   "Horizontal Gradient",
		Filled: FilledDarkShade,
		Empty:  FilledLightShade,
		Handle: HandleCircle,
	}

	// SymbolSetHorizontalDots uses circles.
	SymbolSetHorizontalDots = SymbolSet{
		Name:   "Horizontal Dots",
		Filled: FilledCircle,
		Empty:  EmptyCircle,
		Handle: HandleCircle,
	}

	// SymbolSetHorizontalSquares uses squares.
	SymbolSetHorizontalSquares = SymbolSet{
		Name:   "Horizontal Squares",
		Filled: FilledSquare,
		Empty:  EmptySquare,
		Handle: HandleCircle,
	}

	// SymbolSetHorizontalDouble uses double lines.
	SymbolSetHorizontalDouble = SymbolSet{
		Name:   "Horizontal Double",
		Filled: FilledDoubleLine,
		Empty:  EmptyThinLine,
		Handle: HandleDoubleCircle,
	}

	// SymbolSetVerticalBlocks uses blocks for vertical sliders.
	SymbolSetVerticalBlocks = SymbolSet{
		Name:   "Vertical Blocks",
		Filled: FilledBlock,
		Empty:  EmptyVerticalBar,
		Handle: HandleHorizontalLine,
	}

	// SymbolSetVerticalGradient uses shading for vertical sliders.
	SymbolSetVerticalGradient = SymbolSet{
		Name:   "Vertical Gradient",
		Filled: FilledDarkShade,
		Empty:  FilledLightShade,
		Handle: HandleHorizontalLine,
	}

	// SymbolSetVerticalDots uses circles for vertical sliders.
	SymbolSetVerticalDots = SymbolSet{
		Name:   "Vertical Dots",
		Filled: FilledCircle,
		Empty:  EmptyCircle,
		Handle: HandleHorizontalLine,
	}

	// SymbolSetVerticalSquares uses squares for vertical sliders.
	SymbolSetVerticalSquares = SymbolSet{
		Name:   "Vertical Squares",
		Filled: FilledSquare,
		Empty:  EmptySquare,
		Handle: HandleHorizontalLine,
	}

	// SymbolSetVerticalEqualizer uses bars like an audio equalizer.
	SymbolSetVerticalEqualizer = SymbolSet{
		Name:   "Equalizer",
		Filled: FilledVerticalBar,
		Empty:  EmptyVerticalBar,
		Handle: HandleHorizontalLine,
	}
)

// ToSymbols converts a SymbolSet to the Symbols struct used by Slider.
func (ss SymbolSet) ToSymbols() Symbols {
	return Symbols{
		Filled: ss.Filled,
		Empty:  ss.Empty,
		Handle: ss.Handle,
	}
}

// AllSymbolSets returns all predefined symbol sets.
func AllSymbolSets() []SymbolSet {
	return []SymbolSet{
		SymbolSetDefault,
		SymbolSetBlock,
		SymbolSetDots,
		SymbolSetMinimal,
		SymbolSetDoubleLine,
		SymbolSetWave,
		SymbolSetProgress,
		SymbolSetThick,
		SymbolSetGradient,
		SymbolSetRounded,
		SymbolSetRetro,
		SymbolSetASCII,
		SymbolSetStars,
		SymbolSetSquares,
		SymbolSetCircles,
		SymbolSetDiamonds,
		SymbolSetVertical,
		SymbolSetNeon,
		SymbolSetArrow,
		SymbolSetSegmented,
		SymbolSetSegmentedBlocks,
		SymbolSetSegmentedDots,
		SymbolSetSegmentedSquares,
		SymbolSetHorizontal,
		SymbolSetHorizontalThick,
		SymbolSetHorizontalBlocks,
		SymbolSetHorizontalGradient,
		SymbolSetHorizontalDots,
		SymbolSetHorizontalSquares,
		SymbolSetHorizontalDouble,
		SymbolSetVerticalBlocks,
		SymbolSetVerticalGradient,
		SymbolSetVerticalDots,
		SymbolSetVerticalSquares,
		SymbolSetVerticalEqualizer,
	}
}
