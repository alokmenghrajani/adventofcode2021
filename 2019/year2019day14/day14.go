package year2019day14

import (
	"math"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type reaction struct {
	inputs map[string]int
	output string
	amount int
}

func Part1(input string) int {
	reactions := map[string]*reaction{}

	for _, line := range strings.Split(input, "\n") {
		r := reaction{inputs: map[string]int{}}
		pieces := strings.Split(line, " => ")
		inputs := strings.Split(pieces[0], ", ")
		for _, input := range inputs {
			pieces := strings.Split(input, " ")
			r.inputs[pieces[1]] = utils.MustAtoi(pieces[0])
		}
		pieces = strings.Split(pieces[1], " ")
		r.output = pieces[1]
		r.amount = utils.MustAtoi(pieces[0])
		reactions[r.output] = &r
	}

	extra := map[string]int{}
	return compute(reactions, "FUEL", 1, &extra)
}

func Part2(input string) int {
	reactions := map[string]*reaction{}

	for _, line := range strings.Split(input, "\n") {
		r := reaction{inputs: map[string]int{}}
		pieces := strings.Split(line, " => ")
		inputs := strings.Split(pieces[0], ", ")
		for _, input := range inputs {
			pieces := strings.Split(input, " ")
			r.inputs[pieces[1]] = utils.MustAtoi(pieces[0])
		}
		pieces = strings.Split(pieces[1], " ")
		r.output = pieces[1]
		r.amount = utils.MustAtoi(pieces[0])
		reactions[r.output] = &r
	}

	extra := map[string]int{}
	ore := compute(reactions, "FUEL", 1, &extra)

	lowerBound := 1000000000000 / ore
	upperBound := lowerBound * 10
	for lowerBound < upperBound-1 {
		mid := (lowerBound + upperBound) / 2
		extra = map[string]int{}
		ore := compute(reactions, "FUEL", mid, &extra)
		if ore <= 1000000000000 {
			lowerBound = mid
		} else {
			upperBound = mid
		}
	}
	return lowerBound
}

func compute(reactions map[string]*reaction, output string, amount int, extra *map[string]int) int {
	oreNeeded := 0
	for {
		t := (*extra)[output]
		if t >= amount {
			// we have enough extra chemicals
			(*extra)[output] -= amount
			return oreNeeded
		}
		(*extra)[output] = 0
		amount = amount - t

		if output == "ORE" {
			return amount
		}

		r, ok := reactions[output]
		if !ok {
			panic("meh")
		}

		// calculate how many times we need to run the reaction
		nReactions := int(math.Ceil(float64(amount) / float64(r.amount)))

		for input, value := range r.inputs {
			oreNeeded += compute(reactions, input, nReactions*value, extra)
		}
		(*extra)[output] += nReactions * r.amount
	}
}
