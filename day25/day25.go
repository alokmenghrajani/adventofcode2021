package day25

import (
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
	"github.com/alokmenghrajani/adventofcode2021/utils/inputs"
)

var xMin, xMax, yMin, yMax int

func Part1(input string) int {
	g := inputs.ToGrid(input, ' ')
	g.Print()

	xMin, xMax = g.SizeX()
	yMin, yMax = g.SizeY()

	done := false
	r := 0
	for !done {
		done = true

		g2 := grids.NewGrid(' ')
		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				g2.Set(i, j, g.GetRune(i, j))
			}
		}
		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				if g.GetRune(i, j) == '>' && g.GetRune(nextX(i), j) == '.' {
					done = false
					g2.Set(i, j, '.')
					g2.Set(nextX(i), j, '>')
				}
			}
		}

		g = grids.NewGrid(' ')
		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				g.Set(i, j, g2.GetRune(i, j))
			}
		}
		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				if g2.GetRune(i, j) == 'v' && g2.GetRune(i, nextY(j)) == '.' {
					done = false
					g.Set(i, j, '.')
					g.Set(i, nextY(j), 'v')
				}
			}
		}
		g.Print()

		r++
	}

	return r
}

func nextX(x int) int {
	if x == xMax {
		return xMin
	}
	return x + 1
}

func nextY(y int) int {
	if y == yMax {
		return yMin
	}
	return y + 1
}
