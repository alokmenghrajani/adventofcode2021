package day11

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

func Part1(input string, steps int) int {
	g := grids.NewGrid(0)
	for j, line := range strings.Split(input, "\n") {
		for i := 0; i < len(line); i++ {
			g.Set(i, j, utils.MustAtoi(line[i:i+1]))
		}
	}

	xMin, xMax := g.SizeX()
	yMin, yMax := g.SizeY()
	r := 0
	for step := 0; step < steps; step++ {
		// increase all the cells by one
		for x := xMin; x <= xMax; x++ {
			for y := yMin; y <= yMax; y++ {
				increase(g, x, y)
			}
		}

		// count flashes and reset
		for x := xMin; x <= xMax; x++ {
			for y := yMin; y <= yMax; y++ {
				if g.Get(x, y).(int) > 9 {
					r++
					g.Set(x, y, 0)
				}
			}
		}
	}
	return r
}

func Part2(input string) int {
	g := grids.NewGrid(0)
	for j, line := range strings.Split(input, "\n") {
		for i := 0; i < len(line); i++ {
			g.Set(i, j, utils.MustAtoi(line[i:i+1]))
		}
	}

	xMin, xMax := g.SizeX()
	yMin, yMax := g.SizeY()
	for step := 1; ; step++ {
		// increase all the cells by one
		for x := xMin; x <= xMax; x++ {
			for y := yMin; y <= yMax; y++ {
				increase(g, x, y)
			}
		}

		// count flashes and reset
		r := 0
		for x := xMin; x <= xMax; x++ {
			for y := yMin; y <= yMax; y++ {
				if g.Get(x, y).(int) > 9 {
					r++
					g.Set(x, y, 0)
				}
			}
		}
		if r == (xMax-xMin+1)*(yMax-yMin+1) {
			return step
		}
	}
}

func increase(g *grids.Grid, x, y int) {
	xMin, xMax := g.SizeX()
	if x < xMin || x > xMax {
		return
	}
	yMin, yMax := g.SizeY()
	if y < yMin || y > yMax {
		return
	}

	n := g.Get(x, y).(int)
	g.Set(x, y, n+1)
	if n+1 == 10 {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				increase(g, x+i, y+j)
			}
		}
	}
}
