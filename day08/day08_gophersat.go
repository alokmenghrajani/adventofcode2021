package day08

import (
	"fmt"
	"strings"

	"github.com/crillab/gophersat/bf"
)

// Today's problem solved using a generic SAT solver.
//
// Expressing today's puzzle as a SAT problem is a little tricky. We create 49 boolean variables to represent
// whether a-maps-to-a, a-maps-to-b, etc. We then create a ton of constrains (~80k) which result in ~12k more
// variables. Finally, gophersat spits out the mapping although it's yucky that the input permutations isn't
// itself expressed as a SAT constrain.

func Part1WithGophersat(input string) int {
	r := 0
	for _, line := range strings.Split(input, "\n") {
		pieces := strings.Split(line, " | ")
		inputs := strings.Split(pieces[0], " ")
		mapping := solveWithGophersat(inputs)

		outputs := strings.Split(pieces[1], " ")
		for _, output := range outputs {
			n := remap(output, mapping)
			if n == 1 || n == 4 || n == 7 || n == 8 {
				r++
			}
		}
	}
	return r
}

func Part2WithGophersat(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		pieces := strings.Split(line, " | ")
		inputs := strings.Split(pieces[0], " ")
		mapping := solveWithGophersat(inputs)

		r := 0
		outputs := strings.Split(pieces[1], " ")
		for _, output := range outputs {
			r = r * 10
			r += remap(output, mapping)
		}
		sum += r
	}
	return sum
}

func solveWithGophersat(inputs []string) map[byte]byte {
	model := bf.True
	letters := "abcdefg"

	// express that every letter must map to one other letter.
	for i := 0; i < 7; i++ {
		vars := []string{}
		for j := 0; j < 7; j++ {
			vars = append(vars, fmt.Sprintf("%c-%c", letters[i], letters[j]))
		}
		model = bf.And(model, bf.Unique(vars...))
	}

	// express that only one letter may map to a specific letter.
	for i := 0; i < 7; i++ {
		vars := []string{}
		for j := 0; j < 7; j++ {
			vars = append(vars, fmt.Sprintf("%c-%c", letters[j], letters[i]))
		}
		model = bf.And(model, bf.Unique(vars...))
	}

	// convert each input into constrains
	for _, input := range inputs {
		terms := []bf.Formula{}
		for _, p := range permutations(input) {
			terms2 := []bf.Formula{}
			for _, s := range sevenSeg {
				if len(s) != len(input) {
					continue
				}
				terms3 := []bf.Formula{}
				for i := 0; i < len(input); i++ {
					terms3 = append(terms3, bf.Var(fmt.Sprintf("%c-%c", p[i], s[i])))
				}
				terms2 = append(terms2, bf.And(terms3...))
			}
			terms = append(terms, bf.Or(terms2...))
		}
		model = bf.And(model, bf.Or(terms...))
	}

	solution := bf.Solve(model)
	if solution == nil {
		panic("unsat")
	}

	r := map[byte]byte{}
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			if solution[fmt.Sprintf("%c-%c", letters[i], letters[j])] {
				r[letters[i]] = letters[j]
			}
		}
	}

	return r
}

// copy-pasta from https://www.golangprograms.com/golang-program-to-print-all-permutations-of-a-given-string.html
func permutations(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}
