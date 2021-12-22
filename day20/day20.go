package day20

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

func Part(input string, steps int) int {
	pieces := strings.Split(input, "\n\n")
	iea := pieces[0]
	if len(iea) != 512 {
		panic("meh")
	}
	def := 0
	g := grids.NewGrid(def)
	for j, line := range strings.Split(pieces[1], "\n") {
		for i, r := range line {
			if r == '#' {
				g.Set(i, j, 1)
			}
		}
	}

	for step := 0; step < steps; step++ {
		enhance(&g, &def, iea)
	}

	r := 0
	xMin, xMax := g.SizeX()
	yMin, yMax := g.SizeY()
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			r += g.Get(x, y).(int)
		}
	}

	return r
}

func enhance(g **grids.Grid, def *int, iea string) {
	// figure out new value for default
	if iea[511*(*def)] == '#' {
		*def = 1
	} else {
		*def = 0
	}

	ng := grids.NewGrid(*def)

	xMin, xMax := (*g).SizeX()
	xMin--
	xMax++
	yMin, yMax := (*g).SizeY()
	yMin--
	yMax++
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			n := countNeighbors(*g, x, y)
			if iea[n] == '#' {
				ng.Set(x, y, 1)
			} else {
				ng.Set(x, y, 0)
			}
		}
	}
	*g = ng
}

func countNeighbors(g *grids.Grid, x, y int) int {
	r := 0
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			r = (r << 1) + g.Get(x+i, y+j).(int)
		}
	}
	return r
}
