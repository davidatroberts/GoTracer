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
func (l Lambertian) Scatter(rayIn gtmath.Ray, rec hitable.HitRecord) (success bool, attenuation gtmath.Vector, scattered gtmath.Ray) {
	target := gtmath.AddVec(gtmath.AddVec(rec.P, rec.Normal), gtmath.RandomVecInUnitSphere())

	scattered = gtmath.Ray{Origin: rec.P, Direction: gtmath.SubVec(target, rec.P)}
	attenuation = l.Albedo
	success = true
	return
}

// Metal struct for storing metal properties
type Metal struct {
	Albedo gtmath.Vector
}

// Scatter compute metal lighting
func (m Metal) Scatter(rayIn gtmath.Ray, rec hitable.HitRecord) (success bool, attenuation gtmath.Vector, scattered gtmath.Ray) {
	reflected := gtmath.Reflect(rayIn.Direction.UnitDirection(), rec.Normal)
	scattered = gtmath.Ray{Origin: rec.P, Direction: reflected}
	attenuation = m.Albedo
	success = gtmath.Dot(scattered.Direction, rec.Normal) > 0
	return
}
