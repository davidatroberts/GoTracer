package output

import (
	"GoTracer/gtmath"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

// Outputter stores image format
type Outputter interface {
	Put(x, y uint, colour gtmath.Vector)
	Output()
}

// ImageOutputter outputs to an image
type ImageOutputter struct {
	width, height uint
	path          string
	img           draw.Image
}

func vectorToRGBA(v gtmath.Vector) color.RGBA {
	r := uint8(math.Min(255.0, (v.X * 256.0)))
	g := uint8(math.Min(255.0, (v.Y * 256.0)))
	b := uint8(math.Min(255.0, (v.Z * 256.0)))
	return color.RGBA{r, g, b, 255}
}

// NewImageOutputter constructor for ImageOutPutter
func NewImageOutputter(width, height uint, path string) *ImageOutputter {
	return &ImageOutputter{
		width,
		height,
		path,
		image.NewRGBA(image.Rect(0, 0, int(width), int(height)))}
}

// Put puts colour at pos x,y into image
func (outp *ImageOutputter) Put(x, y uint, colour gtmath.Vector) {
	if x < outp.width && y < outp.height {
		col := vectorToRGBA(colour)
		outp.img.Set(int(x), int(y), col)
	}
}

// Output saves the image to path
func (outp *ImageOutputter) Output() error {
	f, err := os.Create(outp.path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = png.Encode(f, outp.img)
	if err != nil {
		return err
	}

	return nil
}
