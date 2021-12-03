package year2019day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 6, Part1(`R8,U5,L5,D3
U7,R6,D4,L4`))
	assert.Equal(t, 159, Part1(`R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`))
	assert.Equal(t, 135, Part1(`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 30, Part2(`R8,U5,L5,D3
U7,R6,D4,L4`))
	assert.Equal(t, 610, Part2(`R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`))
	assert.Equal(t, 410, Part2(`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`))
}
