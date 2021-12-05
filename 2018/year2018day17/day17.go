package year2018day17

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

func Part1(input string) int {
	g := grids.NewGrid('.')
	for _, line := range strings.Split(input, "\n") {
		pieces := strings.Split(line, ", ")
		var xMin, xMax, yMin, yMax int
		if pieces[0][0] == 'x' {
			t := utils.MustAtoi(pieces[0][2:])
			xMin = t
			xMax = t
			morePieces := strings.Split(pieces[1][2:], "..")
			yMin = utils.MustAtoi(morePieces[0])
			yMax = utils.MustAtoi(morePieces[1])
		} else {
			t := utils.MustAtoi(pieces[0][2:])
			yMin = t
			yMax = t
			morePieces := strings.Split(pieces[1][2:], "..")
			xMin = utils.MustAtoi(morePieces[0])
			xMax = utils.MustAtoi(morePieces[1])
		}
		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				g.Set(i, j, '#')
			}
		}
	}

	putWater(g, 500, 0)
	g.Print()

	r := 0
	xMin, xMax := g.SizeX()
	yMin, yMax := g.SizeY()
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			if g.Get(i, j) == '~' || g.Get(i, j) == '|' {
				r++
			}
		}
	}

	return r - 1
}

func putWater(g *grids.Grid, x, y int) {
	_, maxY := g.SizeY()
	if y > maxY {
		// we are done
		return
	}

	if g.Get(x, y+1) == '.' {
		// we have a space under us
		g.Set(x, y, '|')
		putWater(g, x, y+1)
	}
	if g.Get(x, y+1) == '~' || g.Get(x, y+1) == '#' {
		if isSolid(g, x, y, -1) && isSolid(g, x, y, 1) {
			// no overflow
			fill(g, x, y, -1)
			fill(g, x, y, 1)
		} else {
			// overflow
			overflow(g, x, y, -1)
			overflow(g, x, y, 1)
		}
	}
}

func isSolid(g *grids.Grid, x, y, dir int) bool {
	for {
		x += dir
		if g.Get(x, y) == '#' {
			return true
		}
		if g.Get(x, y+1) == '.' {
			return false
		}
	}
}

func fill(g *grids.Grid, x, y, dir int) {
	for {
		g.Set(x, y, '~')
		x += dir
		if g.Get(x, y) == '#' {
			return
		}
	}
}

func overflow(g *grids.Grid, x, y, dir int) {
	g.Set(x, y, '|')
	for {
		x += dir
		if g.Get(x, y) == '#' {
			return
		}
		g.Set(x, y, '|')
		if g.Get(x, y+1) == '.' {
			putWater(g, x, y+1)
			return
		}
	}
}
