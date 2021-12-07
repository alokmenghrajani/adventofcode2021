package day07

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

func Part1(input string) (int, int) {
	min := utils.MaxInt
	max := utils.MinInt
	positions := []int{}
	for _, pos := range strings.Split(input, ",") {
		p := utils.MustAtoi(pos)
		min = utils.IntMin(min, p)
		max = utils.IntMax(max, p)
		positions = append(positions, p)
	}

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
	min := utils.MaxInt
	max := utils.MinInt
	positions := []int{}
	for _, pos := range strings.Split(input, ",") {
		p := utils.MustAtoi(pos)
		min = utils.IntMin(min, p)
		max = utils.IntMax(max, p)
		positions = append(positions, p)
	}

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
