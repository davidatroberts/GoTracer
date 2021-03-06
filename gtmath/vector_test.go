package gtmath

import "testing"

func TestApprox(t *testing.T) {
	a := Vector{10, 10, 10}
	b := Vector{10.0000001, 10.0000001, 10.0000001}

	result := Approx(a, b)
	if !result {
		t.Errorf("Expected true, got false\n")
	}
}

func TestMagnitude(t *testing.T) {
	v := Vector{5, 10, 10}
	result := v.Magnitude()
	if result != 15.0 {
		t.Errorf("Expected 15.0, got: %f\n", result)
	}
}

func TestDot(t *testing.T) {
	a := Vector{3, 4, 5}
	b := Vector{6, 7, 8}

	result := Dot(a, b)
	if result != 86.0 {
		t.Errorf("Expected 86.0, got: %f\n", result)
	}
}

func TestSub(t *testing.T) {
	a := Vector{10, 40, 87}

	result := a.Sub(5.5)
	expected := Vector{4.5, 34.5, 81.5}
	if result != expected {
		t.Error("Error, unexpected answer", result)
	}

	b := Vector{30, 20, 50}
	result = a.Sub(b)
	expected = Vector{-20, 20, 37}
	if result != expected {
		t.Error("Error, unexpected answer: ", result)
	}
}

func TestAdd(t *testing.T) {
	a := Vector{10, 20, 30}

	result := a.Add(2.2)
	expected := Vector{12.2, 22.2, 32.2}
	if result != expected {
		t.Error("Error, unexpected answer: ", result)
	}

	b := Vector{30, 40, 50}
	result = a.Add(b)
	expected = Vector{40, 60, 80}
	if result != expected {
		t.Error("Error, unexpected answer: ", result)
	}
}

func TestMult(t *testing.T) {
	a := Vector{10, 20, 30}

	result := a.Mult(8)
	expected := Vector{80, 160, 240}
	if result != expected {
		t.Error("Error, unexpected answer: ", result)
	}

	result = a.Mult(0.5)
	expected = Vector{5, 10, 15}
	if result != expected {
		t.Error("Error, unexpected answer: ", result)
	}

	b := Vector{4.2, 5.5, 6.3}
	result = a.Mult(b)
	expected = Vector{42, 110, 189}
	if result != expected {
		t.Error("Error, unexpected answer: ", result)
	}
}

func TestDiv(t *testing.T) {
	a := Vector{10, 20, 30}

	result := a.Div(5)
	expected := Vector{2, 4, 6}
	if result != expected {
		t.Error("Error, unexpected answer: ", result)
	}

	b := Vector{2, 4, 8}
	result = a.Div(b)
	expected = Vector{5, 5, 3.75}
	if result != expected {
		t.Error("Error, unexpected answer: ", result)
	}
}

func TestNormalize(t *testing.T) {
	a := Vector{3, 1, 2}
	result := a.Normalize()

	expected := Vector{0.801783, 0.267261, 0.534522}
	if !Approx(result, expected) {
		t.Error("Error, not approximately answer: ", result)
	}
}
func TestCrossProduct(t *testing.T) {
	a := Vector{1, 3, 4}
	b := Vector{2, -5, 8}

	result := CrossProduct(&a, &b)
	expected := Vector{44, 0, -11}
	if result != expected {
		t.Error("Error, unexpected answer: ", result)
	}

}
