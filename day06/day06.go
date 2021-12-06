package day06

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

func Part(input string, days int) int {
	fishes := make([]int, 9)
	for _, line := range strings.Split(input, ",") {
		n := utils.MustAtoi(line)
		fishes[n]++
	}

	nextDay := make([]int, 9)
	for day := 0; day < days; day++ {
		for i := 1; i <= 8; i++ {
			nextDay[i-1] = fishes[i]
		}
		nextDay[8] = fishes[0]
		nextDay[6] += fishes[0]

		fishes, nextDay = nextDay, fishes
	}

	r := 0
	for i := 0; i <= 8; i++ {
		r += fishes[i]
	}
	return r
}
