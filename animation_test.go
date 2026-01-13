package tuslide

import (
	"math"
	"testing"
	"time"
)

func TestEasingFunctions(t *testing.T) {
	easings := []struct {
		name string
		f    EasingFunc
	}{
		{"Linear", Linear},
		{"EaseInQuad", EaseInQuad},
		{"EaseOutQuad", EaseOutQuad},
		{"EaseInOutQuad", EaseInOutQuad},
		{"EaseInCubic", EaseInCubic},
		{"EaseOutCubic", EaseOutCubic},
		{"EaseInOutCubic", EaseInOutCubic},
		{"EaseInExpo", EaseInExpo},
		{"EaseOutExpo", EaseOutExpo},
		{"EaseInOutExpo", EaseInOutExpo},
		{"EaseOutElastic", EaseOutElastic},
		{"EaseOutBounce", EaseOutBounce},
		{"EaseInBounce", EaseInBounce},
	}

	for _, e := range easings {
		t.Run(e.name, func(t *testing.T) {
			// Test boundaries
			start := e.f(0)
			end := e.f(1)

			if math.Abs(start) > 0.001 {
				t.Errorf("%s(0) = %f, expected ~0", e.name, start)
			}
			if math.Abs(end-1) > 0.001 {
				t.Errorf("%s(1) = %f, expected ~1", e.name, end)
			}

			// Test that it produces values in reasonable range
			mid := e.f(0.5)
			if mid < -0.5 || mid > 1.5 {
				t.Errorf("%s(0.5) = %f, out of reasonable range", e.name, mid)
			}
		})
	}
}

func TestLinear(t *testing.T) {
	for i := 0.0; i <= 1.0; i += 0.1 {
		result := Linear(i)
		if math.Abs(result-i) > 0.001 {
			t.Errorf("Linear(%f) = %f, expected %f", i, result, i)
		}
	}
}

func TestNewAnimation(t *testing.T) {
	state := NewState(WithValue(0), WithMax(100))
	anim := NewAnimation(state, 100)

	if anim.startValue != 0 {
		t.Errorf("Expected startValue 0, got %f", anim.startValue)
	}
	if anim.endValue != 100 {
		t.Errorf("Expected endValue 100, got %f", anim.endValue)
	}
	if anim.duration != 300*time.Millisecond {
		t.Errorf("Expected default duration 300ms, got %v", anim.duration)
	}
}

func TestAnimationWithOptions(t *testing.T) {
	state := NewState(WithValue(0), WithMax(100))
	completed := false

	anim := NewAnimation(state, 100,
		WithAnimDuration(500*time.Millisecond),
		WithEasing(EaseInCubic),
		WithOnComplete(func() { completed = true }),
	)

	if anim.duration != 500*time.Millisecond {
		t.Errorf("Expected duration 500ms, got %v", anim.duration)
	}

	// Run animation to completion
	for !anim.Update() {
		time.Sleep(20 * time.Millisecond)
	}

	if !completed {
		t.Error("OnComplete callback was not called")
	}
	if state.Value() != 100 {
		t.Errorf("Expected value 100 after animation, got %f", state.Value())
	}
}

func TestAnimationUpdate(t *testing.T) {
	state := NewState(WithValue(0), WithMax(100))
	anim := NewAnimation(state, 100, WithAnimDuration(100*time.Millisecond))

	// First update should not be complete
	if anim.Update() {
		t.Error("Animation should not be complete immediately")
	}

	// Value should have changed
	if state.Value() == 0 {
		t.Error("Value should have changed after update")
	}

	// Wait for animation to complete
	time.Sleep(150 * time.Millisecond)

	if !anim.Update() {
		t.Error("Animation should be complete after duration")
	}

	if state.Value() != 100 {
		t.Errorf("Expected final value 100, got %f", state.Value())
	}
}

func TestAnimationIsComplete(t *testing.T) {
	state := NewState(WithValue(0), WithMax(100))
	anim := NewAnimation(state, 100, WithAnimDuration(50*time.Millisecond))

	if anim.IsComplete() {
		t.Error("Animation should not be complete immediately")
	}

	time.Sleep(100 * time.Millisecond)

	if !anim.IsComplete() {
		t.Error("Animation should be complete after duration")
	}
}

func TestAnimationTick(t *testing.T) {
	state := NewState()
	anim := NewAnimation(state, 100)

	cmd := anim.Tick()
	if cmd == nil {
		t.Error("Tick should return a command")
	}
}

func TestAnimationManager(t *testing.T) {
	manager := NewAnimationManager()

	if manager.IsRunning() {
		t.Error("Manager should not be running initially")
	}
	if manager.Count() != 0 {
		t.Error("Manager should have 0 animations initially")
	}

	state1 := NewState(WithValue(0), WithMax(100))
	state2 := NewState(WithValue(0), WithMax(100))

	id1 := manager.Start(state1, 50, WithAnimDuration(50*time.Millisecond))
	id2 := manager.Start(state2, 100, WithAnimDuration(50*time.Millisecond))

	if !manager.IsRunning() {
		t.Error("Manager should be running after starting animations")
	}
	if manager.Count() != 2 {
		t.Errorf("Expected 2 animations, got %d", manager.Count())
	}

	// IDs should be different
	if id1 == id2 {
		t.Error("Animation IDs should be unique")
	}

	// Update until complete
	for manager.Update() {
		time.Sleep(20 * time.Millisecond)
	}

	if manager.IsRunning() {
		t.Error("Manager should not be running after animations complete")
	}
}

func TestAnimationManagerCancel(t *testing.T) {
	manager := NewAnimationManager()
	state := NewState(WithValue(0), WithMax(100))

	id := manager.Start(state, 100, WithAnimDuration(1*time.Second))

	if manager.Count() != 1 {
		t.Error("Should have 1 animation")
	}

	manager.Cancel(id)

	if manager.Count() != 0 {
		t.Error("Should have 0 animations after cancel")
	}
}

func TestAnimationManagerCancelAll(t *testing.T) {
	manager := NewAnimationManager()

	state1 := NewState()
	state2 := NewState()
	manager.Start(state1, 100)
	manager.Start(state2, 100)

	manager.CancelAll()

	if manager.Count() != 0 {
		t.Error("Should have 0 animations after CancelAll")
	}
}

func TestAnimationManagerTick(t *testing.T) {
	manager := NewAnimationManager()

	// No animations - should return nil
	cmd := manager.Tick()
	if cmd != nil {
		t.Error("Tick should return nil with no animations")
	}

	// With animation - should return command
	state := NewState()
	manager.Start(state, 100)
	cmd = manager.Tick()
	if cmd == nil {
		t.Error("Tick should return command with active animations")
	}
}

func TestLerp(t *testing.T) {
	tests := []struct {
		start, end, t, expected float64
	}{
		{0, 100, 0, 0},
		{0, 100, 1, 100},
		{0, 100, 0.5, 50},
		{50, 150, 0.25, 75},
		{-100, 100, 0.5, 0},
	}

	for _, tt := range tests {
		result := Lerp(tt.start, tt.end, tt.t)
		if math.Abs(result-tt.expected) > 0.001 {
			t.Errorf("Lerp(%f, %f, %f) = %f, expected %f",
				tt.start, tt.end, tt.t, result, tt.expected)
		}
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		value, min, max, expected float64
	}{
		{50, 0, 100, 50},
		{-10, 0, 100, 0},
		{150, 0, 100, 100},
		{0, 0, 100, 0},
		{100, 0, 100, 100},
	}

	for _, tt := range tests {
		result := Clamp(tt.value, tt.min, tt.max)
		if result != tt.expected {
			t.Errorf("Clamp(%f, %f, %f) = %f, expected %f",
				tt.value, tt.min, tt.max, result, tt.expected)
		}
	}
}

func TestSmoothStep(t *testing.T) {
	// Test boundaries
	if SmoothStep(0) != 0 {
		t.Error("SmoothStep(0) should be 0")
	}
	if SmoothStep(1) != 1 {
		t.Error("SmoothStep(1) should be 1")
	}

	// Test middle value
	mid := SmoothStep(0.5)
	if math.Abs(mid-0.5) > 0.001 {
		t.Errorf("SmoothStep(0.5) = %f, expected 0.5", mid)
	}

	// Test clamping
	if SmoothStep(-1) != 0 {
		t.Error("SmoothStep should clamp negative values")
	}
	if SmoothStep(2) != 1 {
		t.Error("SmoothStep should clamp values > 1")
	}
}

func TestSmootherStep(t *testing.T) {
	// Test boundaries
	if SmootherStep(0) != 0 {
		t.Error("SmootherStep(0) should be 0")
	}
	if SmootherStep(1) != 1 {
		t.Error("SmootherStep(1) should be 1")
	}

	// Test middle value
	mid := SmootherStep(0.5)
	if math.Abs(mid-0.5) > 0.001 {
		t.Errorf("SmootherStep(0.5) = %f, expected 0.5", mid)
	}
}

func TestPulseAnimation(t *testing.T) {
	state := NewState(WithValue(50), WithMin(0), WithMax(100))
	pulse := NewPulseAnimation(state, 10, 2, 100*time.Millisecond)

	// Run a few frames
	for i := 0; i < 5; i++ {
		pulse.Update()
		time.Sleep(20 * time.Millisecond)
	}

	// Value should have changed from base
	// (not checking exact values due to timing)

	// Wait for completion
	time.Sleep(100 * time.Millisecond)
	if !pulse.Update() {
		t.Error("Pulse should complete after duration")
	}

	// Should return to base value
	if state.Value() != 50 {
		t.Errorf("Expected return to base value 50, got %f", state.Value())
	}
}

func TestPulseAnimationInfinite(t *testing.T) {
	state := NewState(WithValue(50))
	pulse := NewPulseAnimation(state, 10, 1, 0) // 0 duration = infinite

	// Should not complete
	pulse.Update()
	time.Sleep(50 * time.Millisecond)

	if pulse.Update() {
		t.Error("Infinite pulse should not complete")
	}
}

func TestPulseAnimationTick(t *testing.T) {
	state := NewState()
	pulse := NewPulseAnimation(state, 10, 1, time.Second)

	cmd := pulse.Tick()
	if cmd == nil {
		t.Error("Tick should return a command")
	}
}

func TestSpringAnimation(t *testing.T) {
	state := NewState(WithValue(0), WithMin(0), WithMax(100))
	spring := NewSpringAnimation(state, 100)

	// Run animation for a bit
	for i := 0; i < 50; i++ {
		if spring.Update() {
			break
		}
		time.Sleep(16 * time.Millisecond)
	}

	// Value should have moved toward target
	if state.Value() < 50 {
		t.Errorf("Spring should have moved toward target, got %f", state.Value())
	}
}

func TestSpringAnimationSetTarget(t *testing.T) {
	state := NewState(WithValue(50))
	spring := NewSpringAnimation(state, 100)

	spring.SetTarget(0)

	// Run a few frames
	for i := 0; i < 10; i++ {
		spring.Update()
		time.Sleep(16 * time.Millisecond)
	}

	// Should be moving toward new target (0)
	if state.Value() > 50 {
		t.Error("Spring should be moving toward new target 0")
	}
}

func TestSpringAnimationWithOptions(t *testing.T) {
	state := NewState(WithValue(0), WithMax(100))
	spring := NewSpringAnimation(state, 100).
		WithStiffness(500).
		WithDamping(20)

	if spring.stiffness != 500 {
		t.Errorf("Expected stiffness 500, got %f", spring.stiffness)
	}
	if spring.damping != 20 {
		t.Errorf("Expected damping 20, got %f", spring.damping)
	}
}

func TestSpringAnimationTick(t *testing.T) {
	state := NewState()
	spring := NewSpringAnimation(state, 100)

	cmd := spring.Tick()
	if cmd == nil {
		t.Error("Tick should return a command")
	}
}
