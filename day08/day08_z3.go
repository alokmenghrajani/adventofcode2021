package day08

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/day08/go-z3"
)

// Today's problem solved using a generic SMT solver (z3).
//
// Note: z3 installation requires checking out the go-z3 submodule and running make.
//
// z3 allows combining linear equations with boolean constraints. We represent the segments
// (a, b, c, ...) as powers of two (1, 2, 4, ...) and then express valid digits as sums of specific values.

var sevenSegBinary = []int{
	0b1110111, // 0
	0b0100100, // 1
	0b1011101, // 2
	0b1101101, // 3
	0b0101110, // 4
	0b1101011, // 5
	0b1111011, // 6
	0b0100101, // 7
	0b1111111, // 8
	0b1101111, // 9
}

func Part1WithZ3(input string) int {
	r := 0
	for _, line := range strings.Split(input, "\n") {
		pieces := strings.Split(line, " | ")
		inputs := strings.Split(pieces[0], " ")
		mapping := solveWithZ3(inputs)

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

func Part2WithZ3(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		pieces := strings.Split(line, " | ")
		inputs := strings.Split(pieces[0], " ")
		mapping := solveWithZ3(inputs)

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

// It should be possible to solve the problem with a SAT solver. For some reason, I was getting unsatisfiable
// results with github.com/crillab/gophersat. So I switched to z3, at least for now.
//
// Each segment (a, b, c, ...) is represented by a power of two (1, 2, 4, ...). We can then express the puzzle
// input as a combination of linear equations and boolean constraints; something z3 can solve.
func solveWithZ3(inputs []string) map[byte]byte {
	config := z3.NewConfig()
	ctx := z3.NewContext(config)
	config.Close()
	defer ctx.Close()
	s := ctx.NewSolver()
	defer s.Close()

	letters := "abcdefg"

	// define our variables
	vars := map[byte]*z3.AST{}
	for i := 0; i < 7; i++ {
		v := ctx.Const(ctx.Symbol(letters[i:i+1]), ctx.IntSort())
		vars[letters[i]] = v
	}

	// every letter has a value of 1 or 2 or 4 ...
	for i := 0; i < 7; i++ {
		m := ctx.False()
		for j := 0; j < 7; j++ {
			m = m.Or(vars[letters[i]].Eq(ctx.Int(1<<j, ctx.IntSort())))
		}
		s.Assert(m)
	}

	// every value has a letter assigned to it
	for j := 0; j < 7; j++ {
		m := ctx.True()
		for i := 0; i < 7; i++ {
			m = m.Or(vars[letters[i]].Eq(ctx.Int(1<<j, ctx.IntSort())))
		}
		s.Assert(m)
	}

	// map our inputs to a valid digit. Notice we don't need to special case 1,4,7,8.
	for _, input := range inputs {
		m := ctx.False()
		for i := 0; i <= 9; i++ {
			m2 := ctx.Int(0, ctx.IntSort())
			for j := 0; j < len(input); j++ {
				m2 = m2.Add(vars[input[j]])
			}
			m2 = m2.Eq(ctx.Int(sevenSegBinary[i], ctx.IntSort()))
			m = m.Or(m2)
		}
		s.Assert(m)
	}
	if v := s.Check(); v != z3.True {
		panic("unsolveable")
	}

	// Get the resulting model:
	m := s.Model()
	assignments := m.Assignments()
	m.Close()

	// Convert the model into something easier to work with
	r := map[byte]byte{}
	bitToN := map[int]int{1: 0, 2: 1, 4: 2, 8: 3, 16: 4, 32: 5, 64: 6}
	for i := 0; i < 7; i++ {
		v := assignments[letters[i:i+1]].Int()
		r[letters[i]] = letters[bitToN[v]]
	}
	return r
}
