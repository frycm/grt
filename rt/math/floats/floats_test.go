package floats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualApprox(t *testing.T) {
	assert.True(t, EqualApprox(0.4, 0.4, Epsilon))
	assert.True(t, EqualApprox(0.4, 0.4+Epsilon/2, Epsilon))
	assert.False(t, EqualApprox(0.4, 0.5, Epsilon))
}
