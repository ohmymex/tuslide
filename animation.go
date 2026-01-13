package tuslide

import (
	"math"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// EasingFunc is a function that takes a progress value (0-1) and returns
// an eased progress value (0-1).
type EasingFunc func(t float64) float64

// Standard easing functions

// Linear returns unchanged progress (no easing).
func Linear(t float64) float64 {
	return t
}

// EaseInQuad provides quadratic ease-in (accelerating from zero).
func EaseInQuad(t float64) float64 {
	return t * t
}

// EaseOutQuad provides quadratic ease-out (decelerating to zero).
func EaseOutQuad(t float64) float64 {
	return t * (2 - t)
}

// EaseInOutQuad provides quadratic ease-in-out.
func EaseInOutQuad(t float64) float64 {
	if t < 0.5 {
		return 2 * t * t
	}
	return -1 + (4-2*t)*t
}

// EaseInCubic provides cubic ease-in.
func EaseInCubic(t float64) float64 {
	return t * t * t
}

// EaseOutCubic provides cubic ease-out.
func EaseOutCubic(t float64) float64 {
	t--
	return t*t*t + 1
}

// EaseInOutCubic provides cubic ease-in-out.
func EaseInOutCubic(t float64) float64 {
	if t < 0.5 {
		return 4 * t * t * t
	}
	t = 2*t - 2
	return (t*t*t)/2 + 1
}

// EaseInExpo provides exponential ease-in.
func EaseInExpo(t float64) float64 {
	if t == 0 {
		return 0
	}
	return math.Pow(2, 10*(t-1))
}

// EaseOutExpo provides exponential ease-out.
func EaseOutExpo(t float64) float64 {
	if t == 1 {
		return 1
	}
	return 1 - math.Pow(2, -10*t)
}

// EaseInOutExpo provides exponential ease-in-out.
func EaseInOutExpo(t float64) float64 {
	if t == 0 {
		return 0
	}
	if t == 1 {
		return 1
	}
	if t < 0.5 {
		return math.Pow(2, 20*t-10) / 2
	}
	return (2 - math.Pow(2, -20*t+10)) / 2
}

// EaseOutElastic provides elastic bounce effect at the end.
func EaseOutElastic(t float64) float64 {
	if t == 0 {
		return 0
	}
	if t == 1 {
		return 1
	}
	p := 0.3
	s := p / 4
	return math.Pow(2, -10*t)*math.Sin((t-s)*(2*math.Pi)/p) + 1
}

// EaseOutBounce provides bouncing effect at the end.
func EaseOutBounce(t float64) float64 {
	if t < 1/2.75 {
		return 7.5625 * t * t
	} else if t < 2/2.75 {
		t -= 1.5 / 2.75
		return 7.5625*t*t + 0.75
	} else if t < 2.5/2.75 {
		t -= 2.25 / 2.75
		return 7.5625*t*t + 0.9375
	}
	t -= 2.625 / 2.75
	return 7.5625*t*t + 0.984375
}

// EaseInBounce provides bouncing effect at the start.
func EaseInBounce(t float64) float64 {
	return 1 - EaseOutBounce(1-t)
}

// Animation represents an in-progress value animation.
type Animation struct {
	state      *SliderState
	startValue float64
	endValue   float64
	startTime  time.Time
	duration   time.Duration
	easing     EasingFunc
	onComplete func()
}

// AnimationOption configures an Animation.
type AnimationOption func(*Animation)

// WithDuration sets the animation duration.
func WithAnimDuration(d time.Duration) AnimationOption {
	return func(a *Animation) {
		a.duration = d
	}
}

// WithEasing sets the easing function.
func WithEasing(f EasingFunc) AnimationOption {
	return func(a *Animation) {
		a.easing = f
	}
}

// WithOnComplete sets a callback to run when animation completes.
func WithOnComplete(f func()) AnimationOption {
	return func(a *Animation) {
		a.onComplete = f
	}
}

// NewAnimation creates a new animation for a slider state.
func NewAnimation(state *SliderState, targetValue float64, opts ...AnimationOption) *Animation {
	a := &Animation{
		state:      state,
		startValue: state.Value(),
		endValue:   targetValue,
		startTime:  time.Now(),
		duration:   300 * time.Millisecond,
		easing:     EaseOutQuad,
	}

	for _, opt := range opts {
		opt(a)
	}

	return a
}

// AnimationTickMsg is sent to update animation progress.
type AnimationTickMsg struct {
	ID int
}

// Update advances the animation and returns true if complete.
func (a *Animation) Update() bool {
	elapsed := time.Since(a.startTime)
	if elapsed >= a.duration {
		a.state.SetValue(a.endValue)
		if a.onComplete != nil {
			a.onComplete()
		}
		return true
	}

	// Calculate eased progress
	progress := float64(elapsed) / float64(a.duration)
	easedProgress := a.easing(progress)

	// Interpolate value
	value := a.startValue + (a.endValue-a.startValue)*easedProgress
	a.state.SetValue(value)

	return false
}

// Tick returns a command that triggers the next animation frame.
func (a *Animation) Tick() tea.Cmd {
	return tea.Tick(16*time.Millisecond, func(t time.Time) tea.Msg {
		return AnimationTickMsg{}
	})
}

// IsComplete returns true if the animation has finished.
func (a *Animation) IsComplete() bool {
	return time.Since(a.startTime) >= a.duration
}

// AnimationManager manages multiple concurrent animations.
type AnimationManager struct {
	animations map[int]*Animation
	nextID     int
}

// NewAnimationManager creates a new animation manager.
func NewAnimationManager() *AnimationManager {
	return &AnimationManager{
		animations: make(map[int]*Animation),
	}
}

// Start begins a new animation and returns its ID.
func (m *AnimationManager) Start(state *SliderState, targetValue float64, opts ...AnimationOption) int {
	id := m.nextID
	m.nextID++
	m.animations[id] = NewAnimation(state, targetValue, opts...)
	return id
}

// Update advances all animations and removes completed ones.
// Returns true if any animations are still running.
func (m *AnimationManager) Update() bool {
	for id, anim := range m.animations {
		if anim.Update() {
			delete(m.animations, id)
		}
	}
	return len(m.animations) > 0
}

// Tick returns a command that triggers the next animation frame.
func (m *AnimationManager) Tick() tea.Cmd {
	if len(m.animations) == 0 {
		return nil
	}
	return tea.Tick(16*time.Millisecond, func(t time.Time) tea.Msg {
		return AnimationTickMsg{}
	})
}

// Cancel stops an animation by ID.
func (m *AnimationManager) Cancel(id int) {
	delete(m.animations, id)
}

// CancelAll stops all animations.
func (m *AnimationManager) CancelAll() {
	m.animations = make(map[int]*Animation)
}

// IsRunning returns true if any animations are running.
func (m *AnimationManager) IsRunning() bool {
	return len(m.animations) > 0
}

// Count returns the number of running animations.
func (m *AnimationManager) Count() int {
	return len(m.animations)
}

// AnimateTo is a convenience function that returns a Bubble Tea command
// to animate a slider to a target value.
func AnimateTo(state *SliderState, targetValue float64, duration time.Duration, easing EasingFunc) tea.Cmd {
	anim := NewAnimation(state, targetValue,
		WithAnimDuration(duration),
		WithEasing(easing),
	)

	return func() tea.Msg {
		for !anim.IsComplete() {
			anim.Update()
			time.Sleep(16 * time.Millisecond)
		}
		return AnimationTickMsg{ID: -1} // Signal completion
	}
}

// Lerp performs linear interpolation between two values.
func Lerp(start, end, t float64) float64 {
	return start + (end-start)*t
}

// Clamp restricts a value to the given range.
func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// SmoothStep provides smooth Hermite interpolation between 0 and 1.
func SmoothStep(t float64) float64 {
	t = Clamp(t, 0, 1)
	return t * t * (3 - 2*t)
}

// SmootherStep provides even smoother interpolation (Ken Perlin's version).
func SmootherStep(t float64) float64 {
	t = Clamp(t, 0, 1)
	return t * t * t * (t*(t*6-15) + 10)
}

// PulseAnimation creates a pulsing effect that oscillates the value.
type PulseAnimation struct {
	state     *SliderState
	baseValue float64
	amplitude float64
	frequency float64 // cycles per second
	startTime time.Time
	duration  time.Duration // 0 for infinite
}

// NewPulseAnimation creates a new pulsing animation.
func NewPulseAnimation(state *SliderState, amplitude, frequency float64, duration time.Duration) *PulseAnimation {
	return &PulseAnimation{
		state:     state,
		baseValue: state.Value(),
		amplitude: amplitude,
		frequency: frequency,
		startTime: time.Now(),
		duration:  duration,
	}
}

// Update advances the pulse animation. Returns true if complete.
func (p *PulseAnimation) Update() bool {
	elapsed := time.Since(p.startTime)

	if p.duration > 0 && elapsed >= p.duration {
		p.state.SetValue(p.baseValue)
		return true
	}

	// Calculate sine wave offset
	t := elapsed.Seconds()
	offset := p.amplitude * math.Sin(2*math.Pi*p.frequency*t)
	p.state.SetValue(p.baseValue + offset)

	return false
}

// Tick returns a command for the next pulse frame.
func (p *PulseAnimation) Tick() tea.Cmd {
	return tea.Tick(16*time.Millisecond, func(t time.Time) tea.Msg {
		return AnimationTickMsg{}
	})
}

// SpringAnimation simulates spring physics for natural-feeling motion.
type SpringAnimation struct {
	state       *SliderState
	target      float64
	velocity    float64
	stiffness   float64 // Spring stiffness (higher = faster)
	damping     float64 // Damping ratio (1 = critical damping)
	lastUpdate  time.Time
	threshold   float64 // Velocity threshold to consider "at rest"
}

// NewSpringAnimation creates a spring-based animation.
func NewSpringAnimation(state *SliderState, target float64) *SpringAnimation {
	return &SpringAnimation{
		state:      state,
		target:     target,
		velocity:   0,
		stiffness:  180,
		damping:    12,
		lastUpdate: time.Now(),
		threshold:  0.01,
	}
}

// WithStiffness sets the spring stiffness.
func (s *SpringAnimation) WithStiffness(stiffness float64) *SpringAnimation {
	s.stiffness = stiffness
	return s
}

// WithDamping sets the damping ratio.
func (s *SpringAnimation) WithDamping(damping float64) *SpringAnimation {
	s.damping = damping
	return s
}

// SetTarget changes the animation target.
func (s *SpringAnimation) SetTarget(target float64) {
	s.target = target
}

// Update advances the spring simulation. Returns true if at rest.
func (s *SpringAnimation) Update() bool {
	now := time.Now()
	dt := now.Sub(s.lastUpdate).Seconds()
	s.lastUpdate = now

	// Clamp dt to avoid instability
	if dt > 0.1 {
		dt = 0.1
	}

	current := s.state.Value()
	displacement := current - s.target

	// Spring force: F = -kx - cv
	springForce := -s.stiffness * displacement
	dampingForce := -s.damping * s.velocity
	acceleration := springForce + dampingForce

	// Update velocity and position
	s.velocity += acceleration * dt
	newValue := current + s.velocity*dt
	s.state.SetValue(newValue)

	// Check if at rest
	if math.Abs(s.velocity) < s.threshold && math.Abs(displacement) < s.threshold {
		s.state.SetValue(s.target)
		s.velocity = 0
		return true
	}

	return false
}

// Tick returns a command for the next spring frame.
func (s *SpringAnimation) Tick() tea.Cmd {
	return tea.Tick(16*time.Millisecond, func(t time.Time) tea.Msg {
		return AnimationTickMsg{}
	})
}
