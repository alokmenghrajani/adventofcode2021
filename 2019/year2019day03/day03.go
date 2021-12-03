package year2019day03

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

func Part1(input string) int {
	g := grids.NewGrid(false)
	lines := strings.Split(input, "\n")
	plot(g, lines[0], true)
	r := plot(g, lines[1], false)
	return r
}

func plot(g *grids.Grid, input string, mark bool) int {
	cmds := strings.Split(input, ",")
	min := utils.MaxInt
	x := 0
	y := 0
	for _, cmd := range cmds {
		v := utils.MustAtoi(cmd[1:])
		for i := 0; i < v; i++ {
			switch cmd[0] {
			case 'R':
				min = utils.IntMin(min, move(g, &x, &y, 1, 0, mark))
			case 'L':
				min = utils.IntMin(min, move(g, &x, &y, -1, 0, mark))
			case 'U':
				min = utils.IntMin(min, move(g, &x, &y, 0, -1, mark))
			case 'D':
				min = utils.IntMin(min, move(g, &x, &y, 0, 1, mark))
			}
		}
	}
	return min
}

func move(g *grids.Grid, x, y *int, dx, dy int, mark bool) int {
	*x = *x + dx
	*y = *y + dy
	if mark {
		g.Set(*x, *y, true)
	}
	if g.Get(*x, *y).(bool) {
		return utils.Abs(*x) + utils.Abs(*y)
	}
	return utils.MaxInt
}

func Part2(input string) int {
	g := grids.NewGrid(utils.MaxInt)
	lines := strings.Split(input, "\n")
	plot2(g, lines[0], true)
	r := plot2(g, lines[1], false)
	return r
}

func plot2(g *grids.Grid, input string, mark bool) int {
	cmds := strings.Split(input, ",")
	min := utils.MaxInt
	x := 0
	y := 0
	step := 0
	for _, cmd := range cmds {
		v := utils.MustAtoi(cmd[1:])
		for i := 0; i < v; i++ {
			step++
			switch cmd[0] {
			case 'R':
				min = utils.IntMin(min, move2(g, &x, &y, 1, 0, mark, step))
			case 'L':
				min = utils.IntMin(min, move2(g, &x, &y, -1, 0, mark, step))
			case 'U':
				min = utils.IntMin(min, move2(g, &x, &y, 0, -1, mark, step))
			case 'D':
				min = utils.IntMin(min, move2(g, &x, &y, 0, 1, mark, step))
			}
		}
	}
	return min
}

func move2(g *grids.Grid, x, y *int, dx, dy int, mark bool, step int) int {
	*x = *x + dx
	*y = *y + dy

	if mark {
		if g.Get(*x, *y) == utils.MaxInt {
			g.Set(*x, *y, step)
		}
	} else {
		v := g.Get(*x, *y).(int)
		if v != utils.MaxInt {
			return step + v
		}
	}
	return utils.MaxInt
}
