package main

import (
	"GoTracer/gtmath"
	"GoTracer/hitable"
	"GoTracer/output"
	"flag"
	"math"
)

func colour(ray *gtmath.Ray, world *hitable.List) gtmath.Vector {
	var rec hitable.HitRecord
	if world.Hit(*ray, 0.0, math.MaxFloat64, &rec) {
		vec3 := gtmath.Vector{
			X: rec.Normal.X + 1.0,
			Y: rec.Normal.Y + 1.0,
			Z: rec.Normal.Z + 1.0,
		}
		v := vec3.Mult(0.5)
		return v
	}

	unitDir := ray.Direction.UnitDirection()
	t := 0.5 * (unitDir.Y + 1.0)

	a := gtmath.Vector{X: 1.0, Y: 1.0, Z: 1.0}
	b := gtmath.Vector{X: 0.5, Y: 0.7, Z: 1.0}
	aa := a.Mult(1.0 - t)
	bb := b.Mult(t)
	return aa.Add(bb)

}

func main() {
	// read in the cmd line args
	wp := flag.Uint("width", 400, "width of the image")
	hp := flag.Uint("height", 200, "height of the image")
	filePath := flag.String("file", "image.png", "output file path")
	flag.Parse()

	img := output.NewImageOutputter(*wp, *hp, *filePath)

	// generate position vector
	lowerLeftCorner := gtmath.Vector{X: -2.0, Y: -1.0, Z: -1.0}
	horizontal := gtmath.Vector{X: 4.0, Y: 0.0, Z: 0.0}
	vertical := gtmath.Vector{X: 0.0, Y: 2.0, Z: 0.0}
	origin := gtmath.Vector{X: 0.0, Y: 0.0, Z: 0.0}

	s1 := hitable.Sphere{
		Centre: gtmath.Vector{X: 0.0, Y: 0.0, Z: -1.0},
		Radius: 0.5,
	}
	s2 := hitable.Sphere{
		Centre: gtmath.Vector{X: 0.0, Y: -100.5, Z: -1.0},
		Radius: 100,
	}
	var hitList hitable.List
	hitList.List = append(hitList.List, &s1, &s2)

	// render
	for j := 0; j < int(*hp); j++ {
		for i := 0; i < int(*wp); i++ {
			u := float64(i) / float64(*wp)
			v := float64(j) / float64(*hp)

			newLLCorner := lowerLeftCorner.Add(
				gtmath.AddVec(horizontal.Mult(u), vertical.Mult(v)))

			r := gtmath.Ray{Origin: origin, Direction: newLLCorner}
			col := colour(&r, &hitList)

			img.Put(uint(i), uint(j), col)
		}
	}

	img.Output()
}
