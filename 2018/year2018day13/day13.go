package year2018day13

import (
	"fmt"
	"sort"

	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
	"github.com/alokmenghrajani/adventofcode2021/utils/inputs"
)

type cart struct {
	x, y       int
	dirX, dirY int
	turns      int
	crashed    bool
}

func Part1(input string) string {
	grid := inputs.ToGrid(input, ' ')
	carts := []cart{}

	minX, maxX := grid.SizeX()
	minY, maxY := grid.SizeX()
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			switch grid.Get(i, j) {
			case '<':
				c := cart{x: i, y: j, dirX: -1, dirY: 0, turns: 0, crashed: false}
				carts = append(carts, c)
				grid.Set(i, j, '-')
			case '>':
				c := cart{x: i, y: j, dirX: 1, dirY: 0, turns: 0, crashed: false}
				carts = append(carts, c)
				grid.Set(i, j, '-')
			case 'v':
				c := cart{x: i, y: j, dirX: 0, dirY: 1, turns: 0, crashed: false}
				carts = append(carts, c)
				grid.Set(i, j, '|')
			case '^':
				c := cart{x: i, y: j, dirX: 0, dirY: -1, turns: 0, crashed: false}
				carts = append(carts, c)
				grid.Set(i, j, '|')
			}
		}
	}

	for {
		// sort carts
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			}
			return carts[i].y < carts[j].y
		})

		// move carts
		for i := 0; i < len(carts); i++ {
			move(&carts[i], grid)

			// check collision
			for j := 0; j < len(carts); j++ {
				if j == i {
					continue
				}
				if carts[i].x == carts[j].x && carts[i].y == carts[j].y {
					return fmt.Sprintf("%d,%d", carts[i].x, carts[i].y)
				}
			}
		}
	}
}

func Part2(input string) string {
	grid := inputs.ToGrid(input, ' ')
	carts := []cart{}

	minX, maxX := grid.SizeX()
	minY, maxY := grid.SizeX()
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			switch grid.Get(i, j) {
			case '<':
				c := cart{x: i, y: j, dirX: -1, dirY: 0, turns: 0}
				carts = append(carts, c)
				grid.Set(i, j, '-')
			case '>':
				c := cart{x: i, y: j, dirX: 1, dirY: 0, turns: 0}
				carts = append(carts, c)
				grid.Set(i, j, '-')
			case 'v':
				c := cart{x: i, y: j, dirX: 0, dirY: 1, turns: 0}
				carts = append(carts, c)
				grid.Set(i, j, '|')
			case '^':
				c := cart{x: i, y: j, dirX: 0, dirY: -1, turns: 0}
				carts = append(carts, c)
				grid.Set(i, j, '|')
			}
		}
	}

	nCrashed := 0
	for nCrashed < len(carts)-1 {
		// sort carts
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			}
			return carts[i].y < carts[j].y
		})

		// move carts
		for i := 0; i < len(carts); i++ {
			if carts[i].crashed {
				continue
			}
			move(&carts[i], grid)

			// check collision
			for j := 0; j < len(carts); j++ {
				if j == i {
					continue
				}
				if carts[j].crashed {
					continue
				}
				if carts[i].x == carts[j].x && carts[i].y == carts[j].y {
					nCrashed += 2
					carts[i].crashed = true
					carts[j].crashed = true
				}
			}
		}
	}

	for i := 0; i < len(carts); i++ {
		if carts[i].crashed {
			continue
		}
		return fmt.Sprintf("%d,%d", carts[i].x, carts[i].y)
	}
	panic("meh")
}

func move(c *cart, g *grids.Grid) {
	// check if we need to change directions
	switch g.Get(c.x, c.y) {
	case '\\':
		c.dirX, c.dirY = c.dirY, c.dirX
	case '/':
		c.dirX, c.dirY = -c.dirY, -c.dirX
	case '+':
		switch c.turns % 3 {
		case 0:
			// turn left
			if c.dirY == 0 {
				c.dirX, c.dirY = -c.dirY, -c.dirX
			} else {
				c.dirX, c.dirY = c.dirY, c.dirX
			}
		case 2:
			// turn right
			if c.dirX == 0 {
				c.dirX, c.dirY = -c.dirY, -c.dirX
			} else {
				c.dirX, c.dirY = c.dirY, c.dirX
			}
		}
		c.turns++
	}
	c.x += c.dirX
	c.y += c.dirY
}
