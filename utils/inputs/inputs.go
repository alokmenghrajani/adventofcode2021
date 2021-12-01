package inputs

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils/grids"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

func ToInts(input string, sep string) []int {
	var r []int
	for _, line := range strings.Split(input, sep) {
		if line != "" {
			r = append(r, utils.MustAtoi(line))
		}
	}
	return r
}

func ToGrid(input string, empty interface{}) *grids.Grid {
	grid := grids.NewGrid(empty)

	for y, line := range strings.Split(input, "\n") {
		for x, rune := range line {
			grid.Set(x, y, rune)
		}
	}

	return grid
}
