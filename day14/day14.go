package day14

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type cacheKey struct {
	polymer string
	depth   int
}

type count = map[byte]int

func Part(input string, steps int) int {
	pieces := strings.Split(input, "\n\n")

	rules := map[string]byte{}
	for _, line := range strings.Split(pieces[1], "\n") {
		pieces := strings.Split(line, " -> ")
		rules[pieces[0]] = pieces[1][0]
	}

	cache := map[cacheKey]count{}
	counts := count{}
	for i := 0; i < len(pieces[0])-1; i++ {
		r := solve(pieces[0][i:i+2], &rules, &cache, steps)
		for k, v := range r {
			counts[k] += v
		}
		if i >= 1 {
			counts[pieces[0][i]]--
		}
	}

	max := utils.MinInt
	min := utils.MaxInt
	for _, v := range counts {
		min = utils.IntMin(min, v)
		max = utils.IntMax(max, v)
	}

	return max - min
}

func solve(input string, rules *map[string]byte, cache *map[cacheKey]count, depth int) count {
	if depth == 0 {
		r := count{}
		r[input[0]]++
		r[input[1]]++
		return r
	}
	rule, ok := (*rules)[input]
	if !ok {
		r := count{}
		r[input[0]]++
		r[input[1]]++
		return r
	}
	key := cacheKey{polymer: input, depth: depth}
	r, ok := (*cache)[key]
	if ok {
		return r
	}

	r1 := solve(string(input[0])+string(rule), rules, cache, depth-1)
	r2 := solve(string(rule)+string(input[1]), rules, cache, depth-1)
	r = count{}
	for k, v := range r1 {
		r[k] += v
	}
	for k, v := range r2 {
		r[k] += v
	}
	r[rule]--
	(*cache)[key] = r
	return r
}
