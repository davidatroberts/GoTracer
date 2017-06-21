package gtmath

import "testing"

func TestMagnitude(t *testing.T) {
	v := Vector{5, 10, 10}
	result := v.Magnitude()
	if result != 15.0 {
		t.Errorf("Expected 15.0, got: %f\n", result)
	}
}
