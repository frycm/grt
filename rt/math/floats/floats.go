package floats

import "math"

// Epsilon constant for approximate allowed difference of 2 floats to be considered same.
const Epsilon float64 = 0.00001

// EqualApprox compare if 2 floats (a and b) are within epsilon (e), ie. approximately same.
func EqualApprox(a, b, e float64) bool {
	return math.Abs(a-b) < e
}
