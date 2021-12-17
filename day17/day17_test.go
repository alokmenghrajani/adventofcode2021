package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 45, Part1(`target area: x=20..30, y=-10..-5`))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 112, Part2(`target area: x=20..30, y=-10..-5`))
}
