package year2019day24

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
	"github.com/alokmenghrajani/adventofcode2021/utils/inputs"
)

type infiniteGrid struct {
	depth int
	bugs1 [5][5]bool
	bugs2 [5][5]bool
}

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

var infiniteGrids = map[int]*infiniteGrid{}

func Part2(input string, steps int) int {
	g := getGrid(0)
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			if r == '#' {
				setCell(g, x, y, 0, true)
			}
		}
	}

	for step := 0; step < steps; step++ {
		for _, g := range infiniteGrids {
			for x := 0; x < 5; x++ {
				for y := 0; y < 5; y++ {
					if x == 2 && y == 2 {
						continue
					}
					t := count2(g, x, y, step)
					if getCell(g, x, y, step) {
						if t == 1 {
							setCell(g, x, y, step+1, true)
						} else {
							setCell(g, x, y, step+1, false)
						}
					} else {
						if t == 1 || t == 2 {
							setCell(g, x, y, step+1, true)
						} else {
							setCell(g, x, y, step+1, false)
						}
					}

				}
			}
		}
	}

	r := 0
	for _, g := range infiniteGrids {
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				if x == 2 && y == 2 {
					continue
				}
				if getCell(g, x, y, steps) {
					r++
				}
			}
		}
	}
	return r
}

func count2(g *infiniteGrid, x, y, step int) int {
	r := 0
	// left
	if x == 0 {
		g2 := getGrid(g.depth - 1)
		if getCell(g2, 1, 2, step) {
			r++
		}
	} else if x == 3 && y == 2 {
		g2 := getGrid(g.depth + 1)
		for i := 0; i < 5; i++ {
			if getCell(g2, 4, i, step) {
				r++
			}
		}
	} else {
		if getCell(g, x-1, y, step) {
			r++
		}
	}

	// right
	if x == 4 {
		g2 := getGrid(g.depth - 1)
		if getCell(g2, 3, 2, step) {
			r++
		}
	} else if x == 1 && y == 2 {
		g2 := getGrid(g.depth + 1)
		for i := 0; i < 5; i++ {
			if getCell(g2, 0, i, step) {
				r++
			}
		}
	} else {
		if getCell(g, x+1, y, step) {
			r++
		}
	}

	// top
	if y == 0 {
		g2 := getGrid(g.depth - 1)
		if getCell(g2, 2, 1, step) {
			r++
		}
	} else if x == 2 && y == 3 {
		g2 := getGrid(g.depth + 1)
		for i := 0; i < 5; i++ {
			if getCell(g2, i, 4, step) {
				r++
			}
		}
	} else {
		if getCell(g, x, y-1, step) {
			r++
		}
	}

	// bottom
	if y == 4 {
		g2 := getGrid(g.depth - 1)
		if getCell(g2, 2, 3, step) {
			r++
		}
	} else if x == 2 && y == 1 {
		g2 := getGrid(g.depth + 1)
		for i := 0; i < 5; i++ {
			if getCell(g2, i, 0, step) {
				r++
			}
		}
	} else {
		if getCell(g, x, y+1, step) {
			r++
		}
	}

	return r

}

func getGrid(depth int) *infiniteGrid {
	r, ok := infiniteGrids[depth]
	if ok {
		return r
	}

	r = &infiniteGrid{
		depth: depth,
	}
	infiniteGrids[depth] = r
	return r
}

func setCell(g *infiniteGrid, x, y, step int, v bool) {
	if x == 2 && y == 2 {
		panic("meh")
	}
	if step%2 == 0 {
		g.bugs1[x][y] = v
	} else {
		g.bugs2[x][y] = v
	}
}

func getCell(g *infiniteGrid, x, y, step int) bool {
	if x == 2 && y == 2 {
		panic("meh")
	}
	if step%2 == 0 {
		return g.bugs1[x][y]
	} else {
		return g.bugs2[x][y]
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
