package year2018day06

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

func Part1(input string) int {
	grid := grids.NewGrid(-1)
	nodesX := []int{}
	nodesY := []int{}

	lines := strings.Split(input, "\n")
	for n, line := range lines {
		pieces := strings.Split(line, ", ")
		x := utils.MustAtoi(pieces[0])
		nodesX = append(nodesX, x)

		y := utils.MustAtoi(pieces[1])
		nodesY = append(nodesY, y)

		grid.Set(x, y, n)
	}

	// fill the grid
	minX, maxX := grid.SizeX()
	minY, maxY := grid.SizeY()
	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			t := find(x, y, nodesX, nodesY)
			grid.Set(x, y, t)
		}
	}

	// exclude all the borders
	excluded := make([]bool, len(nodesX))
	for x := minX - 1; x <= maxX+1; x++ {
		t := grid.Get(x, minY-1).(int)
		if t != -1 {
			excluded[t] = true
		}
		t = grid.Get(x, maxY+1).(int)
		if t != -1 {
			excluded[t] = true
		}
	}
	for y := minY - 1; y <= maxY+1; y++ {
		t := grid.Get(minX-1, y).(int)
		if t != -1 {
			excluded[t] = true
		}
		t = grid.Get(maxX+1, y).(int)
		if t != -1 {
			excluded[t] = true
		}
	}

	// find the max
	max := 0
	for i := 0; i < len(excluded); i++ {
		if excluded[i] {
			continue
		}
		count := 0
		for x := minX - 1; x <= maxX+1; x++ {
			for y := minY - 1; y <= maxY+1; y++ {
				t := grid.Get(x, y).(int)
				if t == i {
					count++
				}
			}
		}
		if count > max {
			max = count
		}
	}

	return max
}

func find(x, y int, nodesX, nodesY []int) int {
	bestDistance := -1
	bestN := 0
	bestNode := -1
	for i := 0; i < len(nodesX); i++ {
		d := utils.Abs(x-nodesX[i]) + utils.Abs(y-nodesY[i])
		if d < bestDistance || bestDistance == -1 {
			bestDistance = d
			bestN = 1
			bestNode = i
		} else if d == bestDistance {
			bestN++
		}
	}
	if bestN > 1 {
		return -1
	}
	return bestNode
}

func Part2(input string) int {
	grid := grids.NewGrid(false)
	nodesX := []int{}
	nodesY := []int{}

	lines := strings.Split(input, "\n")
	for n, line := range lines {
		pieces := strings.Split(line, ", ")
		x := utils.MustAtoi(pieces[0])
		nodesX = append(nodesX, x)

		y := utils.MustAtoi(pieces[1])
		nodesY = append(nodesY, y)

		grid.Set(x, y, n)
	}

	// count
	r := 0
	minX, maxX := grid.SizeX()
	minY, maxY := grid.SizeY()
	for x := minX - 10000; x <= maxX+10000; x++ {
		for y := minY - 10000; y <= maxY+10000; y++ {
			t := sumDistances(x, y, nodesX, nodesY)
			if t < 10000 {
				r++
			}
		}
	}

	return r
}

func sumDistances(x, y int, nodesX, nodesY []int) int {
	r := 0
	for i := 0; i < len(nodesX); i++ {
		r += utils.Abs(x-nodesX[i]) + utils.Abs(y-nodesY[i])
	}
	return r
}
