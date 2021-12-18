package year2019day24

import (
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
	"github.com/alokmenghrajani/adventofcode2021/utils/inputs"
)

func Part1(input string) int {
	g := inputs.ToGrid(input, ".")

	seen := map[int]bool{}
	for {
		v := value(g)
		if seen[v] {
			return v
		}
		seen[v] = true
		next(g)
	}
}

func value(g *grids.Grid) int {
	r := 0
	t := 1
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if g.Get(i, j) == '#' {
				r += t
			}
			t = t << 1
		}
	}
	return r
}

func next(g *grids.Grid) {
	ng := grids.NewGrid('.')
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			t := count(g, i, j)
			if g.Get(i, j) == '#' {
				if t == 1 {
					ng.Set(i, j, '#')
				} else {
					ng.Set(i, j, '.')
				}
			} else {
				if t == 1 || t == 2 {
					ng.Set(i, j, '#')
				} else {
					ng.Set(i, j, '.')
				}
			}
		}
	}

	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			t := ng.Get(i, j)
			g.Set(i, j, t)
		}
	}
}

func count(g *grids.Grid, x, y int) int {
	r := 0
	if g.Get(x-1, y) == '#' {
		r++
	}
	if g.Get(x+1, y) == '#' {
		r++
	}
	if g.Get(x, y-1) == '#' {
		r++
	}
	if g.Get(x, y+1) == '#' {
		r++
	}
	return r
}
