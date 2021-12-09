package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 15, Part1(`2199943210
3987894921
9856789892
8767896789
9899965678`))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 1134, Part2(`2199943210
3987894921
9856789892
8767896789
9899965678`))
}
