package day15

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

func Part(input string, repeat int) int {
	g := grids.NewGrid(100000)

	for j, line := range strings.Split(input, "\n") {
		l := len(line)
		for i := 0; i < l; i++ {
			n := utils.MustAtoi(line[i : i+1])
			for x := 0; x < repeat; x++ {
				for y := 0; y < repeat; y++ {
					t := n + x + y
					if t > 9 {
						t -= 9
					}
					g.Set(i+l*x, j+l*y, t)
				}
			}
		}
	}

	xMin, xMax := g.SizeX()
	yMin, yMax := g.SizeY()

	risk := grids.NewGrid(100000)
	risk.Set(0, 0, 0)
	done := false
	for !done {
		done = true
		for x := xMin; x <= xMax; x++ {
			for y := yMin; y <= yMax; y++ {
				r1 := risk.Get(x, y).(int)
				t := g.Get(x, y).(int)
				min := risk.Get(x-1, y).(int) + t
				min = utils.IntMin(min, risk.Get(x+1, y).(int)+t)
				min = utils.IntMin(min, risk.Get(x, y-1).(int)+t)
				min = utils.IntMin(min, risk.Get(x, y+1).(int)+t)
				if min < r1 {
					risk.Set(x, y, min)
					done = false
				}
			}
		}
	}

	return risk.Get(xMax, yMax).(int)
}
