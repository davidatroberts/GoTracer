package property

import (
	"GoTracer/gtmath"
	"GoTracer/hitable"
)

// Lambertian struct for storing lambertian properties
type Lambertian struct {
	Albedo gtmath.Vector
}

// Scatter compute lambertian lighting
func (l *Lambertian) Scatter(rayIn gtmath.Ray, rec hitable.HitRecord) (success bool, attenuation gtmath.Vector, scattered gtmath.Ray) {
	target := gtmath.AddVec(gtmath.AddVec(rec.P, rec.Normal), gtmath.RandomVecInUnitSphere())

	scattered = gtmath.Ray{Origin: rec.P, Direction: gtmath.SubVec(target, rec.P)}
	attenuation = l.Albedo
	success = true
	return
}
