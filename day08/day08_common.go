package day08

import "sort"

var sevenSeg = []string{"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"}
var sevenSegReverse = map[string]int{}

func remap(input string, mapping map[byte]byte) int {
	if len(sevenSegReverse) == 0 {
		for i := 0; i < len(sevenSeg); i++ {
			sevenSegReverse[sevenSeg[i]] = i
		}
	}

	mappedInput := []byte{}
	for i := 0; i < len(input); i++ {
		mappedInput = append(mappedInput, mapping[input[i]])
	}
	sort.Slice(mappedInput, func(i, j int) bool {
		return mappedInput[i] < mappedInput[j]
	})
	r, ok := sevenSegReverse[string(mappedInput)]
	if !ok {
		panic("something is wrong!")
	}
	return r
}
