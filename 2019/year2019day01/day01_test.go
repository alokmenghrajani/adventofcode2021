package year2019day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 2, Part1("12"))
	assert.Equal(t, 2, Part1("14"))
	assert.Equal(t, 654, Part1("1969"))
	assert.Equal(t, 33583, Part1("100756"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 2, Part2("14"))
	assert.Equal(t, 966, Part2("1969"))
	assert.Equal(t, 50346, Part2("100756"))
}
