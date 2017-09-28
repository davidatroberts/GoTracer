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

func (c Camera) GetRay(u, v float64) gtmath.Ray {

}
