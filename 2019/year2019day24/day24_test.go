package year2019day24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 2129920, Part1(`....#
#..#.
#..##
..#..
#....`))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 99, Part2(`....#
#..#.
#..##
..#..
#....`, 10))
}
