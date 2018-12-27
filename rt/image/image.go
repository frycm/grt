package image

import (
	"image"
	"image/color"

	rtColor "github.com/frycm/grt/rt/image/color"
)

// RGBf64 is an in-memory image whose At method returns rt.RGB values.
// Components are stored float64.
type RGBf64 struct {
	// Pix holds the image's pixels, in R, G, B order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*3].
	Pix []float64
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

// ColorModel see image.Image.ColorModel.
func (*RGBf64) ColorModel() color.Model {
	return rtColor.RGBf64Model
}

// Bounds see image.Image.Bounds.
func (p *RGBf64) Bounds() image.Rectangle {
	return p.Rect
}

// At see image.Image.At.
func (p *RGBf64) At(x, y int) color.Color {
	return p.RGBf64At(x, y)
}

// RGBf64At RGBf64 At implementation.
func (p *RGBf64) RGBf64At(x, y int) rtColor.RGBf64 {
	if !(image.Point{X: x, Y: y}.In(p.Rect)) {
		return rtColor.RGBf64{}
	}
	i := p.PixOffset(x, y)

	return rtColor.RGBf64{R: p.Pix[i+0], G: p.Pix[i+1], B: p.Pix[i+2]}
}

// NewRGBf64 returns a new RGBf64 image with the given bounds.
func NewRGBf64(r image.Rectangle) *RGBf64 {
	w, h := r.Dx(), r.Dy()
	buf := make([]float64, 3*w*h)
	return &RGBf64{buf, 3 * w, r}
}

// PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).
func (p *RGBf64) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*3
}

// Set see draw.Image.Set.
func (p *RGBf64) Set(x, y int, c color.Color) {
	if !(image.Point{X: x, Y: y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := rtColor.RGBf64Model.Convert(c).(rtColor.RGBf64)
	p.Pix[i+0] = c1.R
	p.Pix[i+1] = c1.G
	p.Pix[i+2] = c1.B
}
