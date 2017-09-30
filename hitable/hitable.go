package hitable

import (
	"GoTracer/gtmath"
)

// Material interface type for materials
type Material interface {
	Scatter(rayIn gtmath.Ray, hitRecord HitRecord) (success bool, attenuation gtmath.Vector, scattered gtmath.Ray)
}

// HitRecord stores record of a hit
type HitRecord struct {
	t         float64
	P, Normal gtmath.Vector
	Material  Material
}

// Hitable interface for things that can be hit
type Hitable interface {
	Hit(ray gtmath.Ray, tMin, tMax float64, rec *HitRecord) bool
}

// List list of things that can be hit
type List struct {
	List []Hitable
}

// Hit iterates through list and passes through
func (l *List) Hit(ray gtmath.Ray, tMin, tMax float64, rec *HitRecord) bool {
	// var tmpRecord HitRecord
	hitAnything := false
	closestSoFar := tMax
	for _, h := range l.List {
		tmpRecord := &HitRecord{}
		if h.Hit(ray, tMin, closestSoFar, tmpRecord) {
			hitAnything = true
			closestSoFar = tmpRecord.t

			rec.Normal = tmpRecord.Normal
			rec.P = tmpRecord.P
			rec.t = tmpRecord.t
		}
	}

	return hitAnything
}
