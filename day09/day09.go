package day09

import (
	"sort"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

func Part1(input string) int {
	g := grids.NewGrid(utils.MaxInt)
	for j, line := range strings.Split(input, "\n") {
		for i, rune := range line {
			g.Set(i, j, utils.MustAtoi(string(rune)))
		}
	}

	xMin, xMax := g.SizeX()
	yMin, yMax := g.SizeY()
	r := 0
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			if isMin(g, i, j) {
				r += g.Get(i, j).(int) + 1
			}
		}
	}
	return r
}

func Part2(input string) int {
	g := grids.NewGrid(9)
	for j, line := range strings.Split(input, "\n") {
		for i, rune := range line {
			g.Set(i, j, utils.MustAtoi(string(rune)))
		}
	}

	bassins := []int{}
	for {
		i, j, done := findNextCell(g)
		if done {
			break
		}
		bassins = append(bassins, flood(g, i, j))
	}

	sort.Ints(bassins)
	bassins = bassins[len(bassins)-3:]
	return bassins[0] * bassins[1] * bassins[2]
}

func findNextCell(g *grids.Grid) (int, int, bool) {
	xMin, xMax := g.SizeX()
	yMin, yMax := g.SizeY()
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			v := g.Get(i, j).(int)
			if v != 9 && v != -1 {
				return i, j, false
			}
		}
	}
	return -1, -1, true
}

func flood(g *grids.Grid, x, y int) int {
	xMin, xMax := g.SizeX()
	yMin, yMax := g.SizeY()

	g.Set(x, y, -1)
	filled := 1
	done := false
	for !done {
		done = true
		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				v := g.Get(i, j).(int)
				if v == -1 {
					t := g.Get(i-1, j).(int)
					if t != 9 && t != -1 {
						g.Set(i-1, j, -1)
						filled++
						done = false
					}
					t = g.Get(i+1, j).(int)
					if t != 9 && t != -1 {
						g.Set(i+1, j, -1)
						filled++
						done = false
					}
					t = g.Get(i, j-1).(int)
					if t != 9 && t != -1 {
						g.Set(i, j-1, -1)
						filled++
						done = false
					}
					t = g.Get(i, j+1).(int)
					if t != 9 && t != -1 {
						g.Set(i, j+1, -1)
						filled++
						done = false
					}
				}
			}
		}
	}
	return filled
}

func isMin(g *grids.Grid, x, y int) bool {
	m := utils.MaxInt

	t := g.Get(x-1, y).(int)
	m = utils.IntMin(m, t)

	t = g.Get(x+1, y).(int)
	m = utils.IntMin(m, t)

	t = g.Get(x, y-1).(int)
	m = utils.IntMin(m, t)

	t = g.Get(x, y+1).(int)
	m = utils.IntMin(m, t)

	return g.Get(x, y).(int) < m
}
