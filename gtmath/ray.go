package gtmath

// Ray with o,d
type Ray struct {
	Origin, Direction Vector
}

// PointAtP returns the point at t from origin to direction
func (r *Ray) PointAtP(t float64) Vector {
	return AddVec(r.Origin, (r.Direction.Mult(t)))
}
