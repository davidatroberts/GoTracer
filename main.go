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
		vec3 := gtmath.Vector{rec.Normal.X + 1, rec.Normal.Y + 1, rec.Normal.Z + 1}
		v := vec3.Mult(0.5)
		return v
	}

	unitDir := ray.Direction.UnitDirection()
	t := 0.5 * (unitDir.Y + 1.0)

	a := gtmath.Vector{1.0, 1.0, 1.0}
	b := gtmath.Vector{0.5, 0.7, 1.0}
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
	lowerLeftCorner := gtmath.Vector{-2.0, -1.0, -1.0}
	horizontal := gtmath.Vector{4.0, 0.0, 0.0}
	vertical := gtmath.Vector{0.0, 2.0, 0.0}
	origin := gtmath.Vector{0.0, 0.0, 0.0}

	s1 := hitable.Sphere{
		gtmath.Vector{0.0, 0.0, -1.0},
		0.5,
	}
	s2 := hitable.Sphere{
		gtmath.Vector{0.0, -100.5, -1.0},
		100,
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

			r := gtmath.Ray{origin, newLLCorner}
			col := colour(&r, &hitList)

			img.Put(uint(i), uint(j), col)
		}
	}

	img.Output()
}
