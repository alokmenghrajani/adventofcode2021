package year2018day17

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

type cell struct {
	l, r bool
	c    byte
}

func Part1(input string) int {
	g := grids.NewGrid(cell{l: false, r: false, c: '.'})
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
				g.Set(i, j, cell{l: true, r: true, c: '#'})
			}
		}
	}

	yMin, yMax := g.SizeY()

	g.Set(500, 0, cell{c: '|'})
	putWater(g)

	xMin, xMax := g.SizeX()
	r := 0
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			if g.Get(i, j).(cell).c == '~' || g.Get(i, j).(cell).c == '|' {
				r++
			}
		}
	}
	return r
}

func Part2(input string) int {
	g := grids.NewGrid(cell{l: false, r: false, c: '.'})
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
				g.Set(i, j, cell{l: true, r: true, c: '#'})
			}
		}
	}

	yMin, yMax := g.SizeY()

	g.Set(500, 0, cell{c: '|'})
	putWater(g)

	xMin, xMax := g.SizeX()
	r := 0
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			if g.Get(i, j).(cell).c == '~' {
				r++
			}
		}
	}
	return r
}

// I wanted to write a nifty recursive thing but I didn't find a clean way to handle some edge cases. I'll have
// to lookup how others have done this. This implementation is super slow! The grid is implemented as a map, which
// doesn't help...
func putWater(g *grids.Grid) {
	// grid to improve runtime
	w := grids.NewGrid(false)
	w.Set(500, 0, true)

	done := false
	for !done {
		done = true
		xMin, xMax := w.SizeX()
		yMin, yMax := w.SizeY()
		_, yMaxReal := g.SizeY()
		for i := xMin - 1; i <= xMax+1; i++ {
			for j := yMin; j <= yMax+1; j++ {
				c := g.Get(i, j).(cell)
				switch c.c {
				case '|':
					if j < yMaxReal && g.Get(i, j+1).(cell).c == '.' {
						// propagate water downwards
						g.Set(i, j+1, cell{c: '|'})
						w.Set(i, j+1, true)
						done = false
					}
					if !c.l && g.Get(i-1, j).(cell).l {
						// propagate that left side is bounded
						c.l = true
						g.Set(i, j, c)
						done = false
					}
					if !c.r && g.Get(i+1, j).(cell).r {
						// propagate that right side is bounded
						c.r = true
						g.Set(i, j, c)
						done = false
					}
					if c.l && c.r {
						// convert to "solid" water
						c.c = '~'
						g.Set(i, j, c)
						done = false
					}
				case '.':
					// check that left side is | with solid under it
					if g.Get(i-1, j).(cell).c == '|' && g.Get(i-1, j+1).(cell).l && g.Get(i-1, j+1).(cell).r {
						g.Set(i, j, cell{c: '|'})
						w.Set(i, j, true)
						done = false
					}

					// same check for right side
					if g.Get(i+1, j).(cell).c == '|' && g.Get(i+1, j+1).(cell).l && g.Get(i+1, j+1).(cell).r {
						g.Set(i, j, cell{c: '|'})
						w.Set(i, j, true)
						done = false
					}
				}
			}
		}

		for i := xMax + 1; i >= xMin-1; i-- {
			for j := yMin; j <= yMax+1; j++ {
				c := g.Get(i, j).(cell)
				switch c.c {
				case '|':
					if j < yMaxReal && g.Get(i, j+1).(cell).c == '.' {
						// propagate water downwards
						g.Set(i, j+1, cell{c: '|'})
						w.Set(i, j+1, true)
						done = false
					}
					if !c.l && g.Get(i-1, j).(cell).l {
						// propagate that left side is bounded
						c.l = true
						g.Set(i, j, c)
						done = false
					}
					if !c.r && g.Get(i+1, j).(cell).r {
						// propagate that right side is bounded
						c.r = true
						g.Set(i, j, c)
						done = false
					}
					if c.l && c.r {
						// convert to "solid" water
						c.c = '~'
						g.Set(i, j, c)
						done = false
					}
				case '.':
					// check that left side is | with solid under it
					if g.Get(i-1, j).(cell).c == '|' && g.Get(i-1, j+1).(cell).l && g.Get(i-1, j+1).(cell).r {
						g.Set(i, j, cell{c: '|'})
						w.Set(i, j, true)
						done = false
					}

					// same check for right side
					if g.Get(i+1, j).(cell).c == '|' && g.Get(i+1, j+1).(cell).l && g.Get(i+1, j+1).(cell).r {
						g.Set(i, j, cell{c: '|'})
						w.Set(i, j, true)
						done = false
					}
				}
			}
		}
	}
}
