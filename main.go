package main

import (
	"GoTracer/gtmath"
	"GoTracer/hitable"
	"GoTracer/output"
	"GoTracer/view"
	"flag"
	"math"
	"math/rand"
)

func colour(ray *gtmath.Ray, world *hitable.List) gtmath.Vector {
	rec := &hitable.HitRecord{}
	if world.Hit(*ray, 0.001, math.MaxFloat64, rec) {
		return rec.Normal.Add(1.0).Mult(0.5)
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
	ns := flag.Uint("samples", 100, "Number of subsamples to take")
	filePath := flag.String("file", "image.png", "output file path")
	flag.Parse()

	img := output.NewImageOutputter(*wp, *hp, *filePath)

	// create the world objects
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

	// create the camera
	camera := view.Camera{
		Origin:          gtmath.Vector{X: 0.0, Y: 0.0, Z: 0.0},
		LowerLeftCorner: gtmath.Vector{X: -2.0, Y: -1.0, Z: -1.0},
		Horizontal:      gtmath.Vector{X: 4.0, Y: 0.0, Z: 0.0},
		Vertical:        gtmath.Vector{X: 0.0, Y: 2.0, Z: 0.0},
	}

	// render
	for j := 0; j < int(*hp); j++ {
		for i := 0; i < int(*wp); i++ {
			col := gtmath.Vector{X: 0, Y: 0, Z: 0}
			for s := 0; s < int(*ns); s++ {
				u := (float64(i) + rand.Float64()) / float64(*wp)
				v := (float64(j) + rand.Float64()) / float64(*hp)

				r := camera.GetRay(u, v)
				col = gtmath.AddVec(col, colour(&r, &hitList))
			}

			col = col.Div(int(*ns))
			img.Put(uint(i), uint(j), col)
		}
	}

	img.Output()
}
