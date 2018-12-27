package tuple

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	assert.Equal(t, Point(4, -4, 3), Tuple{4, -4, 3, 1}, "Wrong tuple was created by rt.Point")
}

func TestVector(t *testing.T) {
	assert.Equal(t, Vector(4, -4, 3), Tuple{4, -4, 3, 0}, "wrong tuple was created by rt.Vector")
}

func TestTuple_IsPoint(t *testing.T) {
	it := Tuple{4.3, -4.2, 3.1, 1}

	assert.True(t, it.IsPoint(), "this tuple (%+v) should be point", it)
	assert.False(t, it.IsVector(), "this tuple (%+v) shouldn't be vector", it)
}

func TestTuple_IsVector(t *testing.T) {
	it := Tuple{4.3, -4.2, 3.1, 0}

	assert.True(t, it.IsVector(), "this tuple (%+v) should be vector", it)
	assert.False(t, it.IsPoint(), "this tuple (%+v) shouldn't be point", it)
}

func TestTuple_Add(t *testing.T) {
	assert.Equal(t, Tuple{1, 1, 6, 1}, Tuple{3, -2, 5, 1}.Add(Tuple{-2, 3, 1, 0}))
	assert.PanicsWithValue(t, "cannot add 2 point tuples", func() {
		Point(1, 1, 1).Add(Point(1, 1, 1))
	}, "adding 2 point vectors should panic")
}

func TestTuple_Sub(t *testing.T) {
	assert.Equal(t,
		Vector(-2, -4, -6),
		Point(3, 2, 1).Sub(Point(5, 6, 7)),
		"wrong subtraction of 2 points",
	)
	assert.Equal(t,
		Point(-2, -4, -6),
		Point(3, 2, 1).Sub(Vector(5, 6, 7)),
		"wrong subtraction of point and vector",
	)
	assert.Equal(t,
		Vector(-2, -4, -6),
		Vector(3, 2, 1).Sub(Vector(5, 6, 7)),
		"wrong subtraction of 2 vectors",
	)
	assert.PanicsWithValue(t, "cannot subtract point from vector", func() {
		Vector(1, 2, 3).Sub(Point(4, 5, 6))
	}, "subtracting point from vector should panic")
}

func TestTuple_Neg(t *testing.T) {
	assert.Equal(t, Tuple{-1, 2, -3, 4}, Tuple{1, -2, 3, -4}.Neg())
}

func TestTuple_Mul(t *testing.T) {
	assert.Equal(t, Tuple{3.5, -7, 10.5, -14}, Tuple{1, -2, 3, -4}.Mul(3.5))
	assert.Equal(t, Tuple{0.5, -1, 1.5, -2}, Tuple{1, -2, 3, -4}.Mul(0.5))
}

func TestTuple_Div(t *testing.T) {
	assert.Equal(t, Tuple{0.5, -1, 1.5, -2}, Tuple{1, -2, 3, -4}.Div(2))
	assert.PanicsWithValue(t, "cannot divide tuple by 0", func() {
		Tuple{1, -2, 3, -4}.Div(0)
	}, "dividing tuple by 0 should panic")
}

func TestTuple_Magnitude(t *testing.T) {
	for _, test := range []struct {
		vector    Tuple
		magnitude float64
	}{
		{Vector(1, 0, 0), 1},
		{Vector(0, 1, 0), 1},
		{Vector(0, 0, 1), 1},
		{Vector(1, 2, 3), math.Sqrt(14)},
		{Vector(-1, -2, -3), math.Sqrt(14)},
	} {
		assert.Equal(t, test.magnitude, test.vector.Magnitude())
	}
}

func TestTuple_Normalize(t *testing.T) {
	assert.Equal(t, Vector(1, 0, 0), Vector(4, 0, 0).Normalize())
	assert.Equal(t, Vector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14)), Vector(1, 2, 3).Normalize())
	assert.Equal(t, 1.0, Vector(1, 2, 3).Normalize().Magnitude())
	assert.PanicsWithValue(t, "cannot normalize vector with 0 magnitude", func() {
		Vector(0, 0, 0).Normalize()
	}, "tuple with 0 magnitude should panic")
}

func TestTuple_Dot(t *testing.T) {
	assert.Equal(t, 20.0, Vector(1, 2, 3).Dot(Vector(2, 3, 4)))
}

func TestTuple_Cross(t *testing.T) {
	assert.Equal(t, Vector(-1, 2, -1), Vector(1, 2, 3).Cross(Vector(2, 3, 4)))
	assert.Equal(t, Vector(1, -2, 1), Vector(2, 3, 4).Cross(Vector(1, 2, 3)))
}
