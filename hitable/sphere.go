package hitable

import (
	"GoTracer/gtmath"
	"math"
)

// Sphere it's a sphere
type Sphere struct {
	Centre gtmath.Vector
	Radius float64
}

// Hit returns true if ray hits the sphere
func (s *Sphere) Hit(ray gtmath.Ray, tMin, tMax float64, rec *HitRecord) bool {
	oc := ray.Origin.Sub(s.Centre)
	a := gtmath.Dot(ray.Direction, ray.Direction)
	b := gtmath.Dot(oc, ray.Direction)
	c := gtmath.Dot(oc, oc) - s.Radius*s.Radius
	disc := b*b - a*c
	if disc > 0 {
		tmp := (-b - math.Sqrt(b*b-a*c)) / a
		if tmp < tMax && tmp > tMin {
			rec.t = tmp
			rec.P = ray.PointAtP(rec.t)
			tmpNormal := rec.P.Sub(s.Centre)
			rec.Normal = tmpNormal.Div(s.Radius)
			return true
		}
		tmp = (-b + math.Sqrt(b*b-a*c)) / a
		if tmp < tMax && tmp > tMin {
			rec.t = tmp
			rec.P = ray.PointAtP(rec.t)
			tmpNormal := rec.P.Sub(s.Centre)
			rec.Normal = tmpNormal.Div(s.Radius)
			return true
		}
	}
	return false
}
