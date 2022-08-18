package ps2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	// assert.Equal(t, true, solution(9, [][]int{{3, 2, 6}, {5, 1, 4}, {1, 7, 13}}))
	assert.Equal(t, false, solution(20, [][]int{{5, 1, 15}, {14, 3, 18}, {3, 15, 21}, {9, 14, 52}}))
}
