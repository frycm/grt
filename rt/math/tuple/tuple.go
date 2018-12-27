package tuple

import (
	"math"

	"github.com/frycm/grt/rt/math/floats"
)

// Tuple is coordinate with type.
type Tuple struct {
	X, Y, Z, W float64
}

// Point creates new point tuple.
func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

// Vector creates new vector tuple.
func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

// IsPoint determine if this tuple is point.
func (t *Tuple) IsPoint() bool {
	return t.W == 1
}

// IsVector determine if this tuple is vector.
func (t *Tuple) IsVector() bool {
	return t.W == 0
}

// EqualApprox check if both tuples are approximately equal.
func (t Tuple) EqualApprox(o Tuple) bool {
	return floats.EqualApprox(t.X, o.X, floats.Epsilon) &&
		floats.EqualApprox(t.Y, o.Y, floats.Epsilon) &&
		floats.EqualApprox(t.Z, o.Z, floats.Epsilon) &&
		floats.EqualApprox(t.W, o.W, floats.Epsilon)
}

// Add another tuples.
// It will panic in case both tuples are point.
func (t Tuple) Add(o Tuple) Tuple {
	if t.IsPoint() && o.IsPoint() {
		panic("cannot add 2 point tuples")
	}
	return Tuple{t.X + o.X, t.Y + o.Y, t.Z + o.Z, t.W + o.W}
}

// Sub subtract another tuples.
// It will panic in case of subtracting point from vector.
func (t Tuple) Sub(o Tuple) Tuple {
	if t.IsVector() && o.IsPoint() {
		panic("cannot subtract point from vector")
	}
	return Tuple{t.X - o.X, t.Y - o.Y, t.Z - o.Z, t.W - o.W}
}

// Neg negates tuple.
func (t Tuple) Neg() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

// Mul multiply tuple by scalar.
func (t Tuple) Mul(scalar float64) Tuple {
	return Tuple{t.X * scalar, t.Y * scalar, t.Z * scalar, t.W * scalar}
}

// Div divide tuple by scalar.
func (t Tuple) Div(scalar float64) Tuple {
	if scalar == 0 {
		panic("cannot divide tuple by 0")
	}
	return Tuple{t.X / scalar, t.Y / scalar, t.Z / scalar, t.W / scalar}
}

// Magnitude of tuple.
func (t Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2) + math.Pow(t.W, 2))
}

// Normalize tuple.
// It will panic in case of tuple with 0 magnitude.
func (t Tuple) Normalize() Tuple {
	v := t.Magnitude()
	if v == 0 {
		panic("cannot normalize vector with 0 magnitude")
	}
	return Tuple{t.X / v, t.Y / v, t.Z / v, t.W / v}
}

// Dot product of tuples.
func (t Tuple) Dot(o Tuple) float64 {
	return t.X*o.X + t.Y*o.Y + t.Z*o.Z + t.W*o.W
}

// Cross product of tuples.
func (t Tuple) Cross(o Tuple) Tuple {
	return Vector(
		t.Y*o.Z-t.Z*o.Y,
		t.Z*o.X-t.X*o.Z,
		t.X*o.Y-t.Y*o.X,
	)
}
