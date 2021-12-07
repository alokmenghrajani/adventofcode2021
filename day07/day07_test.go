package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	x, fuel := Part1("16,1,2,0,4,2,7,1,2,14")
	assert.Equal(t, 2, x)
	assert.Equal(t, 37, fuel)
}

func TestPart2(t *testing.T) {
	x, fuel := Part2("16,1,2,0,4,2,7,1,2,14")
	assert.Equal(t, 5, x)
	assert.Equal(t, 168, fuel)
}
