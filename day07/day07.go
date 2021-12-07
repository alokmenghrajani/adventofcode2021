package day07

import (
	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/inputs"
)

func Part1(input string) (int, int) {
	positions := inputs.ToInts(input, ",")
	min, max := utils.SliceMinMax(positions)

	bestPos := -1
	best := utils.MaxInt
	for i := min; i <= max; i++ {
		t := 0
		for _, pos := range positions {
			t += utils.Abs(pos - i)
		}
		if t < best {
			best = t
			bestPos = i
		}
	}

	return bestPos, best
}

func Part2(input string) (int, int) {
	positions := inputs.ToInts(input, ",")
	min, max := utils.SliceMinMax(positions)

	bestPos := -1
	best := utils.MaxInt
	for i := min; i <= max; i++ {
		t := 0
		for _, pos := range positions {
			d := utils.Abs(pos - i)
			t += d * (d + 1) / 2
		}
		if t < best {
			best = t
			bestPos = i
		}
	}

	return bestPos, best
}
