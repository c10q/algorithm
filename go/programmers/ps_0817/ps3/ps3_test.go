package ps2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 0, solution([]string{"aabb", "baba"}))
	assert.Equal(t, 3, solution([]string{"cab", "adaaa", "e"}))
	assert.Equal(t, 3, solution([]string{"d", "a", "e", "d", "adbcc"}))
	assert.Equal(t, 1, solution([]string{"a"}))
}
