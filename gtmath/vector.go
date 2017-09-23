package gtmath

import (
	"fmt"
	"math"
)

const epsilon = 0.000001

// Vector a 3D vector
type Vector struct {
	X, Y, Z float64
}

func (v *Vector) pNorm(p float64) float64 {
	x := math.Pow(v.X, p)
	y := math.Pow(v.Y, p)
	z := math.Pow(v.Z, p)

	return math.Pow(x+y+z, 1.0/p)
}

func (v *Vector) String() string {
	return fmt.Sprintf("%f, %f, %f", v.X, v.Y, v.Z)
}

// Magnitude returns the magnitude of the vector
func (v *Vector) Magnitude() float64 {
	return v.pNorm(2)
}

// SquaredLength returns squared length
func (v *Vector) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Mult returns v*s
func (v *Vector) Mult(i interface{}) Vector {
	switch value := i.(type) {
	case Vector:
		return Vector{
			X: v.X * value.X,
			Y: v.Y * value.Y,
			Z: v.Z * value.Z,
		}
	case float64:
		return Vector{
			X: v.X * value,
			Y: v.Y * value,
			Z: v.Z * value,
		}
	case int:
		return Vector{
			X: v.X * float64(value),
			Y: v.Y * float64(value),
			Z: v.Z * float64(value),
		}
	}

	return *v
}

// Div returns v/s
func (v *Vector) Div(i interface{}) Vector {
	switch value := i.(type) {
	case Vector:
		return Vector{
			X: v.X / value.X,
			Y: v.Y / value.Y,
			Z: v.Z / value.Z,
		}
	case float64:
		return Vector{
			X: v.X / value,
			Y: v.Y / value,
			Z: v.Z / value,
		}
	case int:
		return Vector{
			X: v.X / float64(value),
			Y: v.Y / float64(value),
			Z: v.Z / float64(value),
		}
	}

	return *v
}

// Add returns v+s
func (v *Vector) Add(i interface{}) Vector {
	switch value := i.(type) {
	case Vector:
		return AddVec(*v, value)
	case float64:
		return Vector{
			X: v.X + value,
			Y: v.Y + value,
			Z: v.Z + value,
		}
	}

	return *v
}

// Sub returns v-s
func (v *Vector) Sub(i interface{}) Vector {
	switch value := i.(type) {
	case Vector:
		return SubVec(*v, value)
	case float64:
		return Vector{
			X: v.X - value,
			Y: v.Y - value,
			Z: v.Z - value,
		}
	}

	return *v
}

// Normalize returns normalized vector
func (v *Vector) Normalize() Vector {
	mag := v.Magnitude()
	return v.Mult(1 / mag)
}

// UnitDirection returns unit direction of vector
func (v *Vector) UnitDirection() Vector {
	mag := v.Magnitude()
	return v.Mult(1.0 / mag)
}

// AddVec returns a+b
func AddVec(a, b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

// SubVec returns a-b
func SubVec(a, b Vector) Vector {
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

// Dot returns dot product
func Dot(a, b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// CrossProduct returns cross product
func CrossProduct(a, b *Vector) Vector {
	xx := a.Y*b.Z - a.Z*b.Y
	yy := a.Z*b.X - a.X*b.Z
	zz := a.X*b.Y - a.Y*b.X
	return Vector{
		X: xx,
		Y: yy,
		Z: zz,
	}
}

// Approx returns true if a-b < epsilon
func Approx(a, b Vector) bool {
	diff := a.Sub(b)
	diff.X = math.Abs(diff.X)
	diff.Y = math.Abs(diff.Y)
	diff.Z = math.Abs(diff.Z)

	if diff.X < epsilon && diff.Y < epsilon && diff.Z < epsilon {
		return true
	}

	return false
}
