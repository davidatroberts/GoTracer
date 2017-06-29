package gtmath

// Ray with o,d
type Ray struct {
	Origin, Direction Vector
}

// PointAtOrigin returns the point at t from origin to direction
func (r *Ray) PointAtOrigin(t float64) Vector {
	return AddVec(r.Origin, (r.Direction.Mult(t)))
}
