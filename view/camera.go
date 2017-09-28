package view

import "GoTracer/gtmath"

// Camera basic camera struct
type Camera struct {
	Origin, LowerLeftCorner, Horizontal, Vertical gtmath.Vector
}

// NewCamera create a basic camera
func NewCamera(origin, lowerLeftCorner, horizontal,
	vertical gtmath.Vector) *Camera {
	return &Camera{
		Origin:          origin,
		LowerLeftCorner: lowerLeftCorner,
		Horizontal:      horizontal,
		Vertical:        vertical,
	}
}

// GetRay gets ray at u,v
func (c Camera) GetRay(u, v float64) gtmath.Ray {
	return gtmath.Ray{
		Origin: c.Origin,
		Direction: c.LowerLeftCorner.Add(
			gtmath.AddVec(c.Horizontal.Mult(u), c.Vertical.Mult(v))),
	}
}
