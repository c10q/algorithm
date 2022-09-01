package ps1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 4, solution(156))
	assert.Equal(t, 10, solution(319))
	assert.Equal(t, 7, solution(551))
	assert.Equal(t, 22, solution(1234))
}
