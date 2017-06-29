package main

import (
	"GoTracer/gtmath"
	"GoTracer/output"
	"flag"
	"math"
)

func hitSphere(centre gtmath.Vector, radius float64, ray gtmath.Ray) float64 {
	oc := ray.Origin.Sub(centre)
	a := gtmath.Dot(ray.Direction, ray.Direction)
	b := 2.0 * gtmath.Dot(oc, ray.Direction)
	c := gtmath.Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	}

	return (-b - math.Sqrt(discriminant)) / (2.0 * a)
}

func colour(r gtmath.Ray) gtmath.Vector {
	t := hitSphere(gtmath.Vector{0.0, 0.0, -1.0}, 0.5, r)
	if t > 0.0 {
		rpo := r.PointAtOrigin(t)
		rpoV := rpo.Sub(gtmath.Vector{0.0, 0.0, -1.0})
		n := rpoV.UnitDirection()
		r := gtmath.Vector{n.X + 1, n.Y + 1, n.Z + 1}
		return r.UnitDirection()
	}

	unitDir := r.Direction.UnitDirection()
	t = 0.5 * (unitDir.Y + 1.0)

	a := gtmath.Vector{1.0, 1.0, 1.0}
	a = a.Mult(1.0 - t)
	b := gtmath.Vector{0.5, 0.7, 1.0}
	b = b.Mult(t)

	return a.Add(b)
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

	// render
	for y := 0; y < int(*hp); y++ {
		for x := 0; x < int(*wp); x++ {
			u := float64(x) / float64(*wp)
			v := float64(y) / float64(*hp)

			newLLCorner := lowerLeftCorner.Add(
				gtmath.AddVec(horizontal.Mult(u), vertical.Mult(v)))

			r := gtmath.Ray{origin, newLLCorner}
			col := colour(r)

			img.Put(uint(x), uint(y), col)
		}
	}

	img.Output()
}
