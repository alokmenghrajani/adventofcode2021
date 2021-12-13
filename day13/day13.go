package day13

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

func Part1(input string) int {
	g := grids.NewGrid(false)
	pieces := strings.Split(input, "\n\n")
	for _, line := range strings.Split(pieces[0], "\n") {
		pieces := strings.Split(line, ",")
		x := utils.MustAtoi(pieces[0])
		y := utils.MustAtoi(pieces[1])
		g.Set(x, y, true)
	}
	xMin, xMax = g.SizeX()
	yMin, yMax = g.SizeY()
	xMin = 0
	yMin = 0

	for _, line := range strings.Split(pieces[1], "\n") {
		if strings.HasPrefix(line, "fold along y=") {
			n := utils.MustAtoi(line[len("fold along y="):])
			foldHorz(g, n)
		} else {
			n := utils.MustAtoi(line[len("fold along x="):])
			foldVert(g, n)
		}
		break
	}

	r := 0
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			if g.Get(i, j).(bool) {
				r++
			}
		}
	}
	return r
}

var xMin, xMax, yMin, yMax int

func Part2(input string) {
	g := grids.NewGrid(false)
	pieces := strings.Split(input, "\n\n")
	for _, line := range strings.Split(pieces[0], "\n") {
		pieces := strings.Split(line, ",")
		x := utils.MustAtoi(pieces[0])
		y := utils.MustAtoi(pieces[1])
		g.Set(x, y, true)
	}
	xMin, xMax = g.SizeX()
	yMin, yMax = g.SizeY()
	xMin = 0
	yMin = 0

	for _, line := range strings.Split(pieces[1], "\n") {
		if strings.HasPrefix(line, "fold along y=") {
			n := utils.MustAtoi(line[len("fold along y="):])
			foldHorz(g, n)
		} else {
			n := utils.MustAtoi(line[len("fold along x="):])
			foldVert(g, n)
		}
	}

	for j := yMin; j <= yMax; j++ {
		for i := xMin; i <= xMax; i++ {
			if g.Get(i, j).(bool) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func foldHorz(g *grids.Grid, n int) {
	for i := xMin; i <= xMax; i++ {
		for j := n; j <= yMax; j++ {
			if g.Get(i, j).(bool) {
				g.Set(i, 2*n-j, true)
				g.Set(i, j, false)
			}
		}
	}
	yMax = n
}

func foldVert(g *grids.Grid, n int) {
	for i := n; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			if g.Get(i, j).(bool) {
				g.Set(2*n-i, j, true)
				g.Set(i, j, false)
			}
		}
	}
	xMax = n
}
