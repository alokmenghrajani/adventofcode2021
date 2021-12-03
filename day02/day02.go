package day02

import (
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type inst struct {
	Dir string
	N   int
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	var insts []inst
	for _, line := range lines {
		var i inst
		re := regexp.MustCompile(`([a-z]+) (\d+)`)
		if utils.ParseToStruct(re, line, &i) {
			insts = append(insts, i)
		}
	}

	x := 0
	y := 0
	for i := 0; i < len(insts); i++ {
		switch insts[i].Dir {
		case "forward":
			x += insts[i].N
		case "down":
			y += insts[i].N
		case "up":
			y -= insts[i].N
		}
	}

	return x * y
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	var insts []inst
	for _, line := range lines {
		var i inst
		re := regexp.MustCompile(`([a-z]+) (\d+)`)
		if utils.ParseToStruct(re, line, &i) {
			insts = append(insts, i)
		}
	}

	x := 0
	y := 0
	aim := 0
	for i := 0; i < len(insts); i++ {
		switch insts[i].Dir {
		case "forward":
			x += insts[i].N
			y += aim * insts[i].N
		case "down":
			aim += insts[i].N
		case "up":
			aim -= insts[i].N
		}
	}

	return x * y
}
