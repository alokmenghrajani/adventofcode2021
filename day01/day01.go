package day01

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
	for i := 1; i < len(numbers); i++ {
		if numbers[i].N > numbers[i-1].N {
			r++
		}
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
	for i := 3; i < len(numbers); i++ {
		w1 := numbers[i-3].N + numbers[i-2].N + numbers[i-1].N
		w2 := numbers[i-2].N + numbers[i-1].N + numbers[i].N

		if w2 > w1 {
			r++
		}
	}

	return r
}
