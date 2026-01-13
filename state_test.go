package tuslide

import (
	"math"
	"testing"
)

func TestNewState_Defaults(t *testing.T) {
	s := NewState()

	if s.Min() != 0 {
		t.Errorf("expected min=0, got %f", s.Min())
	}
	if s.Max() != 100 {
		t.Errorf("expected max=100, got %f", s.Max())
	}
	if s.Value() != 0 {
		t.Errorf("expected value=0, got %f", s.Value())
	}
	if s.Step() != 1 {
		t.Errorf("expected step=1, got %f", s.Step())
	}
}

func TestNewState_WithOptions(t *testing.T) {
	s := NewState(
		WithMin(10),
		WithMax(50),
		WithValue(30),
		WithStep(5),
	)

	if s.Min() != 10 {
		t.Errorf("expected min=10, got %f", s.Min())
	}
	if s.Max() != 50 {
		t.Errorf("expected max=50, got %f", s.Max())
	}
	if s.Value() != 30 {
		t.Errorf("expected value=30, got %f", s.Value())
	}
	if s.Step() != 5 {
		t.Errorf("expected step=5, got %f", s.Step())
	}
}

func TestNewState_ClampOnInit(t *testing.T) {
	// Value exceeds max
	s := NewState(WithMin(0), WithMax(100), WithValue(150))
	if s.Value() != 100 {
		t.Errorf("expected value to be clamped to 100, got %f", s.Value())
	}

	// Value below min
	s = NewState(WithMin(10), WithMax(100), WithValue(5))
	if s.Value() != 10 {
		t.Errorf("expected value to be clamped to 10, got %f", s.Value())
	}
}

func TestWithStep_InvalidValues(t *testing.T) {
	// Zero step should default to 1
	s := NewState(WithStep(0))
	if s.Step() != 1 {
		t.Errorf("expected step=1 for zero input, got %f", s.Step())
	}

	// Negative step should default to 1
	s = NewState(WithStep(-5))
	if s.Step() != 1 {
		t.Errorf("expected step=1 for negative input, got %f", s.Step())
	}
}

func TestSetValue_Clamping(t *testing.T) {
	s := NewState(WithMin(0), WithMax(100))

	s.SetValue(50)
	if s.Value() != 50 {
		t.Errorf("expected value=50, got %f", s.Value())
	}

	s.SetValue(150)
	if s.Value() != 100 {
		t.Errorf("expected value clamped to 100, got %f", s.Value())
	}

	s.SetValue(-10)
	if s.Value() != 0 {
		t.Errorf("expected value clamped to 0, got %f", s.Value())
	}
}

func TestSetMin_ReclampValue(t *testing.T) {
	s := NewState(WithMin(0), WithMax(100), WithValue(25))

	s.SetMin(50)
	if s.Min() != 50 {
		t.Errorf("expected min=50, got %f", s.Min())
	}
	if s.Value() != 50 {
		t.Errorf("expected value to be reclamped to 50, got %f", s.Value())
	}
}

func TestSetMax_ReclampValue(t *testing.T) {
	s := NewState(WithMin(0), WithMax(100), WithValue(75))

	s.SetMax(50)
	if s.Max() != 50 {
		t.Errorf("expected max=50, got %f", s.Max())
	}
	if s.Value() != 50 {
		t.Errorf("expected value to be reclamped to 50, got %f", s.Value())
	}
}

func TestSetStep_InvalidValue(t *testing.T) {
	s := NewState(WithStep(5))

	s.SetStep(0)
	if s.Step() != 5 {
		t.Errorf("expected step to remain 5 for zero input, got %f", s.Step())
	}

	s.SetStep(-10)
	if s.Step() != 5 {
		t.Errorf("expected step to remain 5 for negative input, got %f", s.Step())
	}

	s.SetStep(10)
	if s.Step() != 10 {
		t.Errorf("expected step=10, got %f", s.Step())
	}
}

func TestIncrement(t *testing.T) {
	s := NewState(WithMin(0), WithMax(100), WithValue(50), WithStep(10))

	s.Increment()
	if s.Value() != 60 {
		t.Errorf("expected value=60, got %f", s.Value())
	}

	// Increment near max
	s.SetValue(95)
	s.Increment()
	if s.Value() != 100 {
		t.Errorf("expected value clamped to 100, got %f", s.Value())
	}
}

func TestDecrement(t *testing.T) {
	s := NewState(WithMin(0), WithMax(100), WithValue(50), WithStep(10))

	s.Decrement()
	if s.Value() != 40 {
		t.Errorf("expected value=40, got %f", s.Value())
	}

	// Decrement near min
	s.SetValue(5)
	s.Decrement()
	if s.Value() != 0 {
		t.Errorf("expected value clamped to 0, got %f", s.Value())
	}
}

func TestPercentage(t *testing.T) {
	tests := []struct {
		name     string
		min      float64
		max      float64
		value    float64
		expected float64
	}{
		{"zero", 0, 100, 0, 0},
		{"half", 0, 100, 50, 0.5},
		{"full", 0, 100, 100, 1.0},
		{"quarter", 0, 200, 50, 0.25},
		{"offset range", 50, 150, 100, 0.5},
		{"negative range", -100, 100, 0, 0.5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewState(WithMin(tc.min), WithMax(tc.max), WithValue(tc.value))
			pct := s.Percentage()
			if math.Abs(pct-tc.expected) > 0.0001 {
				t.Errorf("expected percentage=%f, got %f", tc.expected, pct)
			}
		})
	}
}

func TestPercentage_ZeroRange(t *testing.T) {
	s := NewState(WithMin(50), WithMax(50), WithValue(50))
	if s.Percentage() != 0 {
		t.Errorf("expected 0 for zero range, got %f", s.Percentage())
	}
}

func TestSetFromPercentage(t *testing.T) {
	s := NewState(WithMin(0), WithMax(100))

	s.SetFromPercentage(0.5)
	if s.Value() != 50 {
		t.Errorf("expected value=50, got %f", s.Value())
	}

	s.SetFromPercentage(0.25)
	if s.Value() != 25 {
		t.Errorf("expected value=25, got %f", s.Value())
	}

	// Offset range
	s = NewState(WithMin(100), WithMax(200))
	s.SetFromPercentage(0.5)
	if s.Value() != 150 {
		t.Errorf("expected value=150, got %f", s.Value())
	}
}

func TestSetFromPercentage_Clamping(t *testing.T) {
	s := NewState(WithMin(0), WithMax(100))

	s.SetFromPercentage(1.5)
	if s.Value() != 100 {
		t.Errorf("expected value clamped to 100, got %f", s.Value())
	}

	s.SetFromPercentage(-0.5)
	if s.Value() != 0 {
		t.Errorf("expected value clamped to 0, got %f", s.Value())
	}
}

func TestRange(t *testing.T) {
	s := NewState(WithMin(10), WithMax(50))
	if s.Range() != 40 {
		t.Errorf("expected range=40, got %f", s.Range())
	}

	s = NewState(WithMin(-50), WithMax(50))
	if s.Range() != 100 {
		t.Errorf("expected range=100, got %f", s.Range())
	}
}

func TestFloatPrecision(t *testing.T) {
	s := NewState(WithMin(0), WithMax(1), WithValue(0.5), WithStep(0.1))

	s.Increment()
	if math.Abs(s.Value()-0.6) > 0.0001 {
		t.Errorf("expected value≈0.6, got %f", s.Value())
	}

	s.Decrement()
	if math.Abs(s.Value()-0.5) > 0.0001 {
		t.Errorf("expected value≈0.5, got %f", s.Value())
	}
}
