package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var arr = []int{11, 14, 3, 8, 18, 17, 43}
var expected = []int{3, 8, 11, 14, 17, 18, 43}

func TestQuicksort(t *testing.T) {
	assert.Equal(t, expected, Quicksort(arr))
}
