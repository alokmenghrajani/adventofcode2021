package year2019day01

import (
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type number struct {
	N int
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	var numbers []number
	for _, line := range lines {
		var n number
		re := regexp.MustCompile(`(\d+)`)
		if utils.ParseToStruct(re, line, &n) {
			numbers = append(numbers, n)
		}
	}

	r := 0
	for i := 0; i < len(numbers); i++ {
		r += (numbers[i].N / 3) - 2
	}

	return r
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	var numbers []number
	for _, line := range lines {
		var n number
		re := regexp.MustCompile(`(\d+)`)
		if utils.ParseToStruct(re, line, &n) {
			numbers = append(numbers, n)
		}
	}

	r := 0
	for i := 0; i < len(numbers); i++ {
		r += compute(numbers[i].N)
	}

	return r
}

func compute(n int) int {
	v := (n / 3) - 2
	if v <= 0 {
		return 0
	}
	return v + compute(v)
}
