package day10

import (
	"sort"
	"strings"
)

func Part1(input string) int {
	r := 0

	for _, line := range strings.Split(input, "\n") {
		c, _ := process(line)
		switch c {
		case ')':
			r += 3
		case ']':
			r += 57
		case '}':
			r += 1197
		case '>':
			r += 25137
		}
	}

	return r
}

func Part2(input string) int {
	r := []int{}

	for _, line := range strings.Split(input, "\n") {
		c, stack := process(line)
		if c != 0 {
			continue
		}
		t := 0
		for i := len(stack) - 1; i >= 0; i-- {
			t = t * 5
			switch stack[i] {
			case '(':
				t += 1
			case '[':
				t += 2
			case '{':
				t += 3
			case '<':
				t += 4
			}
		}
		r = append(r, t)
	}
	sort.Ints(r)
	return r[len(r)/2]
}

func process(input string) (byte, []byte) {
	stack := []byte{}
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c == '(' || c == '[' || c == '{' || c == '<' {
			stack = append(stack, c)
		} else {
			head := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if c == ')' && head == '(' {
				continue
			}
			if c == ']' && head == '[' {
				continue
			}
			if c == '}' && head == '{' {
				continue
			}
			if c == '>' && head == '<' {
				continue
			}
			return c, stack
		}
	}
	return 0, stack
}
