package year2019day22

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

// 9039 => too high
func Part1(input string, size, card int) int {
	deck := process(input, size)
	for i := 0; i < size; i++ {
		if deck[i] == card {
			return i
		}
	}
	panic("unreachable")
}

func process(input string, size int) []int {
	deck := make([]int, size)
	deck2 := make([]int, size)
	for i := 0; i < size; i++ {
		deck[i] = i
	}

	for _, line := range strings.Split(input, "\n") {
		if line == "deal into new stack" {
			reverseSlice(deck)
		}
		if strings.HasPrefix(line, "cut ") {
			n := utils.MustAtoi(line[len("cut "):])
			if n < 0 {
				n += size
			}
			offset := size - n
			copy(deck2[offset:size], deck[0:n])
			copy(deck2[0:offset], deck[n:size])
			deck, deck2 = deck2, deck
		}
		if strings.HasPrefix(line, "deal with increment ") {
			n := utils.MustAtoi(line[len("deal with increment "):])
			offset := 0
			for i := 0; i < size; i++ {
				deck2[offset] = deck[i]
				offset = (offset + n) % size
			}
			deck, deck2 = deck2, deck
		}
	}
	return deck
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
