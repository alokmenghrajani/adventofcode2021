package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	r := Part1(`199
	200
	208
	210
	200
	207
	240
	269
	260
	263`)
	assert.Equal(t, 7, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`199
	200
	208
	210
	200
	207
	240
	269
	260
	263`)
	assert.Equal(t, 5, r)
}
