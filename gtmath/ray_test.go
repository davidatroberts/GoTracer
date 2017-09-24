package gtmath

import "testing"

func TestPointAtP(t *testing.T) {
	origin := Vector{0, 0, 0}
	direction := Vector{0, 1, 2}
	r := Ray{Origin: origin, Direction: direction}

	tvalue := 1.0
	result := r.PointAtP(tvalue)
	expected := Vector{0, 1, 2}
	if result != expected {
		t.Error("Error, unexpected answer: ", result)
	}

	tvalue = 0.5
	result = r.PointAtP(tvalue)
	expected = Vector{0, 0.5, 1}
	if result != expected {
		t.Error("Error, unexpected answer: ", result)
	}
}
