package image

import (
	"image"
	"testing"

	"github.com/frycm/grt/rt/image/color"
	"github.com/stretchr/testify/assert"
)

func TestNewRGBf64(t *testing.T) {
	canvas := NewRGBf64(image.Rect(0, 0, 10, 20))

	assert.Equal(t, 10, canvas.Rect.Dx(), "wrong width")
	assert.Equal(t, 20, canvas.Rect.Dy(), "wrong depth")

	colorBlack := color.RGBf64{}
	for x := 0; x < 10; x++ {
		for y := 0; y < 20; y++ {
			assert.Equal(t, colorBlack, canvas.At(x, y), "wrong color at x: %d, y: %d", x, y)
		}
	}
}

func TestRGBf64_Set_At(t *testing.T) {
	canvas := NewRGBf64(image.Rect(0, 0, 10, 20))
	colorBlack := color.RGBf64{}
	colorRed := color.RGBf64{R: 1}

	canvas.Set(2, 3, colorRed)

	assert.Equal(t, colorRed, canvas.At(2, 3))
	assert.Equal(t, colorBlack, canvas.At(0, 0))
}
