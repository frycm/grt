package color

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRGBf64_RGBA(t *testing.T) {
	type rgbaTuple struct{ r, g, b, a uint32 }
	newRGBATuple := func(r, g, b, a uint32) rgbaTuple {
		return rgbaTuple{r, g, b, a}
	}

	assert.Equal(t, rgbaTuple{a: math.MaxUint32}, newRGBATuple(RGBf64{0, 0, 0}.RGBA()))
	assert.Equal(t, rgbaTuple{r: math.MaxUint32, g: uint32(math.Round(0.5 * math.MaxUint32)), b: uint32(math.Round(0.25 * math.MaxUint32)), a: math.MaxUint32}, newRGBATuple(RGBf64{1, 0.5, 0.25}.RGBA()))
	assert.Equal(t, rgbaTuple{a: math.MaxUint32}, newRGBATuple(RGBf64{-2, 0, 0}.RGBA()))
	assert.Equal(t, rgbaTuple{r: math.MaxUint32, a: math.MaxUint32}, newRGBATuple(RGBf64{2, 0, 0}.RGBA()))
}

func TestRGBf64_Add(t *testing.T) {
	assert.Equal(t, RGBf64{1.6, 0.7, 1}, RGBf64{0.9, 0.6, 0.75}.Add(RGBf64{0.7, 0.1, 0.25}))
}

func TestRGBf64_Sub(t *testing.T) {
	assert.True(t, RGBf64{0.2, 0.5, 0.5}.EqualApprox(RGBf64{0.9, 0.6, 0.75}.Sub(RGBf64{0.7, 0.1, 0.25})))
}

func TestRGBf64_MulS(t *testing.T) {
	assert.Equal(t, RGBf64{0.4, 0.6, 0.8}, RGBf64{0.2, 0.3, 0.4}.MulS(2))
}

func TestRGBf64_Mul(t *testing.T) {
	assert.True(t, RGBf64{0.9, 0.2, 0.04}.EqualApprox(RGBf64{1, 0.2, 0.4}.Mul(RGBf64{0.9, 1, 0.1})))
}
