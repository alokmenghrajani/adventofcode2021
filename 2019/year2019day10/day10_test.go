package year2019day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	x, y, v := Part1(`.#..#
.....
#####
....#
...##`)
	assert.Equal(t, 3, x)
	assert.Equal(t, 4, y)
	assert.Equal(t, 8, v)

	x, y, v = Part1(`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`)
	assert.Equal(t, 5, x)
	assert.Equal(t, 8, y)
	assert.Equal(t, 33, v)

	x, y, v = Part1(`#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`)
	assert.Equal(t, 1, x)
	assert.Equal(t, 2, y)
	assert.Equal(t, 35, v)

	x, y, v = Part1(`.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`)
	assert.Equal(t, 6, x)
	assert.Equal(t, 3, y)
	assert.Equal(t, 41, v)

	x, y, v = Part1(`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`)
	assert.Equal(t, 11, x)
	assert.Equal(t, 13, y)
	assert.Equal(t, 210, v)
}

func TestPart2(t *testing.T) {
	input := `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

	x, y, _ := Part2(input, 11, 13, 1)
	assert.Equal(t, 11, x)
	assert.Equal(t, 12, y)

	x, y, _ = Part2(input, 11, 13, 2)
	assert.Equal(t, 12, x)
	assert.Equal(t, 1, y)

	x, y, _ = Part2(input, 11, 13, 3)
	assert.Equal(t, 12, x)
	assert.Equal(t, 2, y)

	x, y, _ = Part2(input, 11, 13, 10)
	assert.Equal(t, 12, x)
	assert.Equal(t, 8, y)

	x, y, _ = Part2(input, 11, 13, 20)
	assert.Equal(t, 16, x)
	assert.Equal(t, 0, y)

	x, y, _ = Part2(input, 11, 13, 50)
	assert.Equal(t, 16, x)
	assert.Equal(t, 9, y)

	x, y, _ = Part2(input, 11, 13, 100)
	assert.Equal(t, 10, x)
	assert.Equal(t, 16, y)

	x, y, _ = Part2(input, 11, 13, 199)
	assert.Equal(t, 9, x)
	assert.Equal(t, 6, y)

	x, y, z := Part2(input, 11, 13, 200)
	assert.Equal(t, 8, x)
	assert.Equal(t, 2, y)
	assert.Equal(t, 802, z)

	x, y, _ = Part2(input, 11, 13, 201)
	assert.Equal(t, 10, x)
	assert.Equal(t, 9, y)

	x, y, _ = Part2(input, 11, 13, 299)
	assert.Equal(t, 11, x)
	assert.Equal(t, 1, y)
}
