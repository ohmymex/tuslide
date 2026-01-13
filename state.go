// Package tuslide provides a flexible and customizable TUI slider component
// for terminal applications built with the Bubble Tea framework.
//
// TuSlide separates state management from rendering, allowing for clean
// integration into any Bubble Tea application. It supports both horizontal
// and vertical orientations, custom symbols, and full styling via Lip Gloss.
//
// Basic usage:
//
//	state := tuslide.NewState(
//	    tuslide.WithMin(0),
//	    tuslide.WithMax(100),
//	    tuslide.WithValue(50),
//	)
//	slider := tuslide.New(state)
package tuslide

// SliderState manages the value and bounds of a slider.
// It handles clamping, stepping, and percentage calculations.
type SliderState struct {
	min   float64
	max   float64
	value float64
	step  float64
}

// StateOption is a functional option for configuring SliderState.
type StateOption func(*SliderState)

// NewState creates a new SliderState with the given options.
// Default values: min=0, max=100, value=0, step=1.
func NewState(opts ...StateOption) *SliderState {
	s := &SliderState{
		min:   0,
		max:   100,
		value: 0,
		step:  1,
	}

	for _, opt := range opts {
		opt(s)
	}

	// Ensure value is clamped after initialization
	s.value = s.clamp(s.value)

	return s
}

// WithMin sets the minimum value of the slider.
func WithMin(min float64) StateOption {
	return func(s *SliderState) {
		s.min = min
	}
}

// WithMax sets the maximum value of the slider.
func WithMax(max float64) StateOption {
	return func(s *SliderState) {
		s.max = max
	}
}

// WithValue sets the initial value of the slider.
// The value will be clamped to the min/max bounds.
func WithValue(value float64) StateOption {
	return func(s *SliderState) {
		s.value = value
	}
}

// WithStep sets the step size for value changes.
// Must be positive; defaults to 1 if set to zero or negative.
func WithStep(step float64) StateOption {
	return func(s *SliderState) {
		if step <= 0 {
			step = 1
		}
		s.step = step
	}
}

// Min returns the minimum value of the slider.
func (s *SliderState) Min() float64 {
	return s.min
}

// Max returns the maximum value of the slider.
func (s *SliderState) Max() float64 {
	return s.max
}

// Value returns the current value of the slider.
func (s *SliderState) Value() float64 {
	return s.value
}

// Step returns the step size of the slider.
func (s *SliderState) Step() float64 {
	return s.step
}

// SetValue sets the slider value, clamping it to the valid range.
func (s *SliderState) SetValue(value float64) {
	s.value = s.clamp(value)
}

// SetMin sets the minimum value and re-clamps the current value.
func (s *SliderState) SetMin(min float64) {
	s.min = min
	s.value = s.clamp(s.value)
}

// SetMax sets the maximum value and re-clamps the current value.
func (s *SliderState) SetMax(max float64) {
	s.max = max
	s.value = s.clamp(s.value)
}

// SetStep sets the step size. Must be positive.
func (s *SliderState) SetStep(step float64) {
	if step > 0 {
		s.step = step
	}
}

// Increment increases the value by one step, respecting the maximum bound.
func (s *SliderState) Increment() {
	s.SetValue(s.value + s.step)
}

// Decrement decreases the value by one step, respecting the minimum bound.
func (s *SliderState) Decrement() {
	s.SetValue(s.value - s.step)
}

// Percentage returns the current value as a percentage (0.0 to 1.0)
// of the range between min and max.
// Returns 0 if min equals max (to avoid division by zero).
func (s *SliderState) Percentage() float64 {
	if s.max == s.min {
		return 0
	}
	return (s.value - s.min) / (s.max - s.min)
}

// SetFromPercentage sets the value based on a percentage (0.0 to 1.0).
// The percentage is clamped to [0, 1] before calculating the value.
func (s *SliderState) SetFromPercentage(pct float64) {
	// Clamp percentage to valid range
	if pct < 0 {
		pct = 0
	}
	if pct > 1 {
		pct = 1
	}
	s.value = s.min + pct*(s.max-s.min)
}

// Range returns the difference between max and min.
func (s *SliderState) Range() float64 {
	return s.max - s.min
}

// clamp restricts the value to the valid range [min, max].
func (s *SliderState) clamp(value float64) float64 {
	if value < s.min {
		return s.min
	}
	if value > s.max {
		return s.max
	}
	return value
}
