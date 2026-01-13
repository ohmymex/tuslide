# Style Presets

TuSlide comes with 46+ predefined styles. Apply them using `tuslide.WithStyle()`:

```go
slider := tuslide.New(state, tuslide.WithStyle(tuslide.StyleNeon()))
```

## Basic Styles

```go
tuslide.StyleDefault()      // ━━━━━━━━━●─────────
tuslide.StyleBlock()        // ████████▓░░░░░░░░░░
tuslide.StyleProgress()     // ▰▰▰▰▰▰▰▰▶▱▱▱▱▱▱▱▱▱
tuslide.StyleWave()         // ≈≈≈≈≈≈≈≈◈˜˜˜˜˜˜˜˜˜
tuslide.StyleDots()         // ⣿⣿⣿⣿⣿⣿⬤⣀⣀⣀⣀⣀⣀⣀
tuslide.StyleMinimal()      // Simple, clean style
tuslide.StyleClassic()      // Traditional slider look
```

## Color Themes

```go
tuslide.StyleOcean()        // Deep blue tones
tuslide.StyleForest()       // Green nature theme
tuslide.StyleSunset()       // Warm orange/red
tuslide.StyleNeon()         // Vibrant magenta/cyan
tuslide.StyleMonochrome()   // Grayscale
tuslide.StyleFire()         // Red/orange fire theme
tuslide.StyleIce()          // Cool blue/cyan theme
tuslide.StyleGold()         // Luxurious gold theme
tuslide.StylePurple()       // Royal purple theme
```

## Gaming UI

Perfect for game interfaces:

```go
tuslide.StyleHealth()       // Red health bar
tuslide.StyleMana()         // Cyan mana bar
tuslide.StyleExperience()   // Yellow XP bar
tuslide.StyleStamina()      // Green stamina bar
```

## Progress Bars

Ideal for showing operation progress:

```go
tuslide.StyleProgressDownload()     // Green download bar
tuslide.StyleProgressUpload()       // Blue upload bar
tuslide.StyleProgressLoading()      // Magenta loading bar
tuslide.StyleProgressInstallation() // Installation progress
tuslide.StyleProgressBattery()      // Battery level indicator
```

## Segmented Styles

Discrete segment sliders:

```go
tuslide.StyleSegmentedDots()      // ● ● ● ● ● ○ ○ ○ ○ ○
tuslide.StyleSegmentedStars()     // ★ ★ ★ ★ ★ ☆ ☆ ☆ ☆ ☆
tuslide.StyleSegmentedSquares()   // ■ ■ ■ ■ ■ □ □ □ □ □
tuslide.StyleSegmentedDiamonds()  // ◆ ◆ ◆ ◆ ◆ ◇ ◇ ◇ ◇ ◇
tuslide.StyleSegmentedArrows()    // ▶ ▶ ▶ ▶ ▶ ▷ ▷ ▷ ▷ ▷
tuslide.StyleSegmentedBars()      // ▮ ▮ ▮ ▮ ▮ ▯ ▯ ▯ ▯ ▯
tuslide.StyleSegmentedBlocks()    // █ █ █ █ █ ░ ░ ░ ░ ░
```

## Vertical Styles

Optimized for vertical sliders:

```go
tuslide.StyleVertical()           // Standard vertical style
tuslide.StyleVerticalEqualizer()  // Audio equalizer style
tuslide.StyleVerticalMeter()      // Level meter style
tuslide.StyleVerticalThermometer()// Temperature gauge style
```

## Horizontal Styles

Optimized for horizontal sliders:

```go
tuslide.StyleHorizontal()         // Standard horizontal style
tuslide.StyleHorizontalThin()     // Thin line style
tuslide.StyleHorizontalThick()    // Thick bar style
```

## Gradient Styles

```go
tuslide.StyleGradient()           // Gradient fill effect
tuslide.StyleGradientWarm()       // Warm color gradient
tuslide.StyleGradientCool()       // Cool color gradient
tuslide.StyleGradientRainbow()    // Rainbow gradient
```

## ASCII Styles

For terminals without Unicode support:

```go
tuslide.StyleASCII()              // [====O----]
tuslide.StyleASCIISimple()        // ====>-----
tuslide.StyleASCIIClassic()       // [####|....]
```

## Custom Styles

Create your own style by combining symbols and colors:

```go
style := tuslide.SliderStyle{
    Symbols: tuslide.Symbols{
        Filled: "█",
        Empty:  "░",
        Handle: "●",
    },
    FilledStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6B6B")),
    EmptyStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#4A4A4A")),
    HandleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")),
    Segmented:   false,
}

slider := tuslide.New(state, tuslide.WithStyle(style))
```

## Symbol Sets

You can also use predefined symbol sets independently:

```go
tuslide.WithSymbolSet(tuslide.SymbolSetDefault)
tuslide.WithSymbolSet(tuslide.SymbolSetBlocks)
tuslide.WithSymbolSet(tuslide.SymbolSetDots)
tuslide.WithSymbolSet(tuslide.SymbolSetStars)
tuslide.WithSymbolSet(tuslide.SymbolSetCircles)
tuslide.WithSymbolSet(tuslide.SymbolSetSquares)
tuslide.WithSymbolSet(tuslide.SymbolSetDiamonds)
tuslide.WithSymbolSet(tuslide.SymbolSetArrows)
tuslide.WithSymbolSet(tuslide.SymbolSetWaves)
tuslide.WithSymbolSet(tuslide.SymbolSetGradient)
tuslide.WithSymbolSet(tuslide.SymbolSetASCII)
// ... and 25+ more
```

See the [showcase example](examples/showcase/main.go) for a live demo of all styles.
