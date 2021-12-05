package day05

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

func Part1(input string) int {
	g := grids.NewGrid(0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		pieces := strings.Split(line, " -> ")
		xy := strings.Split(pieces[0], ",")
		x1 := utils.MustAtoi(xy[0])
		y1 := utils.MustAtoi(xy[1])

		xy = strings.Split(pieces[1], ",")
		x2 := utils.MustAtoi(xy[0])
		y2 := utils.MustAtoi(xy[1])

		if x1 == x2 {
			if y1 < y2 {
				place(g, x1, y1, x2, y2, 0, 1)
			} else {
				place(g, x1, y1, x2, y2, 0, -1)
			}
		} else if y1 == y2 {
			if x1 < x2 {
				place(g, x1, y1, x2, y2, 1, 0)
			} else {
				place(g, x1, y1, x2, y2, -1, 0)
			}
		}
	}

	xMin, xMax := g.SizeX()
	yMin, yMax := g.SizeY()
	r := 0
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			if g.Get(i, j).(int) >= 2 {
				r++
			}
		}
	}

	return r
}

func Part2(input string) int {
	g := grids.NewGrid(0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		pieces := strings.Split(line, " -> ")
		xy := strings.Split(pieces[0], ",")
		x1 := utils.MustAtoi(xy[0])
		y1 := utils.MustAtoi(xy[1])

		xy = strings.Split(pieces[1], ",")
		x2 := utils.MustAtoi(xy[0])
		y2 := utils.MustAtoi(xy[1])

		dx := 0
		dy := 0
		if x1 < x2 {
			dx = 1
		} else if x1 > x2 {
			dx = -1
		}

		if y1 < y2 {
			dy = 1
		} else if y1 > y2 {
			dy = -1
		}

		place(g, x1, y1, x2, y2, dx, dy)
	}

	xMin, xMax := g.SizeX()
	yMin, yMax := g.SizeY()
	r := 0
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			if g.Get(i, j).(int) >= 2 {
				r++
			}
		}
	}

	return r
}

func place(g *grids.Grid, x1, y1, x2, y2, dx, dy int) {
	x := x1
	y := y1
	t := g.Get(x, y).(int)
	g.Set(x, y, t+1)
	for x != x2 || y != y2 {
		x += dx
		y += dy
		t := g.Get(x, y).(int)
		g.Set(x, y, t+1)
	}
}
