package color

import (
	"image/color"
	"math"

	"github.com/frycm/grt/rt/math/floats"
)

// RGBf64Model color model for RGBf64.
var RGBf64Model = color.ModelFunc(rgbF64Model)

func rgbF64Model(c color.Color) color.Color {
	if _, ok := c.(RGBf64); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: uint8(a >> 8)}
}

// RGBf64 of trace.
type RGBf64 struct {
	R, G, B float64
}

// RGBA returns the alpha-premultiplied red, green, blue and alpha values
// for the color.
// See color.Color for more.
func (c RGBf64) RGBA() (r, g, b, a uint32) {
	toColor := func(v float64) uint32 {
		switch {
		case v < 0:
			return 0
		case v > 1:
			return math.MaxUint32
		default:
			return uint32(math.Round(v * math.MaxUint32))
		}
	}
	return toColor(c.R), toColor(c.G), toColor(c.B), math.MaxUint32
}

// EqualApprox check if both colors are approximately equal.
func (c RGBf64) EqualApprox(o RGBf64) bool {
	return floats.EqualApprox(c.R, o.R, floats.Epsilon) &&
		floats.EqualApprox(c.G, o.G, floats.Epsilon) &&
		floats.EqualApprox(c.B, o.B, floats.Epsilon)
}

// Add another color.
func (c RGBf64) Add(o RGBf64) RGBf64 {
	return RGBf64{c.R + o.R, c.G + o.G, c.B + o.B}
}

// Sub subtract another color.
func (c RGBf64) Sub(o RGBf64) RGBf64 {
	return RGBf64{c.R - o.R, c.G - o.G, c.B - o.B}
}

// Mul multiply with another color (Hadamard or Schur product).
func (c RGBf64) Mul(o RGBf64) RGBf64 {
	return RGBf64{c.R * o.R, c.G * o.G, c.B * o.B}
}

// MulS multiply with scalar (Hadamard or Schur product).
func (c RGBf64) MulS(scalar float64) RGBf64 {
	return RGBf64{c.R * scalar, c.G * scalar, c.B * scalar}
}
