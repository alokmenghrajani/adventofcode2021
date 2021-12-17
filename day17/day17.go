package day17

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

func Part1(input string) int {
	input = input[len("target area: x="):]
	pieces := strings.Split(input, ", y=")
	pieces2 := strings.Split(pieces[0], "..")
	xMin := utils.MustAtoi(pieces2[0])
	xMax := utils.MustAtoi(pieces2[1])

	pieces2 = strings.Split(pieces[1], "..")
	yMin := utils.MustAtoi(pieces2[0])
	yMax := utils.MustAtoi(pieces2[1])

	best := utils.MinInt
	for x := 0; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			t, ok := compute(x, y, xMin, xMax, yMin, yMax)
			if ok {
				best = utils.IntMax(best, t)
			}
		}
	}
	return best
}

func Part2(input string) int {
	input = input[len("target area: x="):]
	pieces := strings.Split(input, ", y=")
	pieces2 := strings.Split(pieces[0], "..")
	xMin := utils.MustAtoi(pieces2[0])
	xMax := utils.MustAtoi(pieces2[1])

	pieces2 = strings.Split(pieces[1], "..")
	yMin := utils.MustAtoi(pieces2[0])
	yMax := utils.MustAtoi(pieces2[1])

	count := 0
	for x := 0; x < 1000; x++ {
		for y := -2000; y < 2000; y++ {
			_, ok := compute(x, y, xMin, xMax, yMin, yMax)
			if ok {
				count++
			}
		}
	}
	return count
}

func compute(vx, vy, xMin, xMax, yMin, yMax int) (int, bool) {
	x := 0
	y := 0
	top := 0
	for y >= yMin && x <= xMax {
		x += vx
		y += vy
		top = utils.IntMax(top, y)
		vx -= utils.Sign(vx)
		vy -= 1
		if x >= xMin && x <= xMax && y >= yMin && y <= yMax {
			return top, true
		}
	}
	return 0, false
}
