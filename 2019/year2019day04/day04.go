package year2019day04

import (
	"fmt"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

func Part1(min, max int) int {
	n := 0
	for i := min; i <= max; i++ {
		if valid(i) {
			n++
		}
	}
	return n
}

func Part2(min, max int) int {
	n := 0
	for i := min; i <= max; i++ {
		if valid2(i) {
			n++
		}
	}
	return n
}

func valid(n int) bool {
	s := fmt.Sprintf("%d", n)
	if len(s) != 6 {
		return false
	}
	hasDouble := false
	for i := 0; i < 5; i++ {
		if s[i] == s[i+1] {
			hasDouble = true
			break
		}
	}
	if !hasDouble {
		return false
	}

	for i := 0; i < 5; i++ {
		n1 := utils.MustAtoi(s[i : i+1])
		n2 := utils.MustAtoi(s[i+1 : i+2])
		if n2 < n1 {
			return false
		}
	}
	return true
}

func valid2(n int) bool {
	s := fmt.Sprintf("%d", n)
	if len(s) != 6 {
		return false
	}
	s = s + "_"

	hasDouble := false
	currentRune := byte(' ')
	currentLen := 0
	for i := 0; i < 6; i++ {
		if s[i] == currentRune {
			currentLen++
		} else {
			if currentLen == 2 {
				break
			}
			currentLen = 1
			currentRune = s[i]
		}
	}
	if currentLen == 2 {
		hasDouble = true
	}
	if !hasDouble {
		return false
	}

	for i := 0; i < 5; i++ {
		n1 := utils.MustAtoi(s[i : i+1])
		n2 := utils.MustAtoi(s[i+1 : i+2])
		if n2 < n1 {
			return false
		}
	}
	return true
}
