package year2019day02

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

func Part1(input string) int {
	r := compute(input, 12, 2)
	return r[0]
}

func Part2(input string) int {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			r := compute(input, noun, verb)
			if len(r) >= 1 {
				if r[0] == 19690720 {
					return 100*noun + verb
				}
			}
		}
	}
	panic("no solution found")
}

func compute(input string, noun, verb int) []int {
	memString := strings.Split(input, ",")
	mem := []int{}
	for _, s := range memString {
		mem = append(mem, utils.MustAtoi(s))
	}

	if noun != -1 {
		mem[1] = noun
	}
	if verb != -1 {
		mem[2] = verb
	}

	ip := 0
outer:
	for {
		switch mem[ip] {
		case 1:
			arg1 := mem[mem[ip+1]]
			arg2 := mem[mem[ip+2]]
			res := mem[ip+3]
			mem[res] = arg1 + arg2
		case 2:
			arg1 := mem[mem[ip+1]]
			arg2 := mem[mem[ip+2]]
			res := mem[ip+3]
			mem[res] = arg1 * arg2
		case 99:
			break outer
		default:
			return []int{}
		}
		ip += 4
	}
	return mem
}
