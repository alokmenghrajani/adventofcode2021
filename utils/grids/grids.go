package grids

import (
	"fmt"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

// TODO:
// - make this an arbitrary dimension thing
type Grid struct {
	minX, maxX, minY, maxY int
	data                   map[string]interface{}
	empty                  interface{}
}

func NewGrid(empty interface{}) *Grid {
	return &Grid{
		minX:  utils.MaxInt,
		maxX:  utils.MinInt,
		minY:  utils.MaxInt,
		maxY:  utils.MinInt,
		data:  map[string]interface{}{},
		empty: empty,
	}
}

func (g *Grid) SizeX() (int, int) {
	return g.minX, g.maxX
}

func (g *Grid) SizeY() (int, int) {
	return g.minY, g.maxY
}

func (g *Grid) Visited() int {
	return len(g.data)
}

func (g *Grid) Get(x, y int) interface{} {
	k := key(x, y)
	v, ok := g.data[k]
	if !ok {
		return g.empty
	}
	return v
}

func (g *Grid) GetRune(x, y int) rune {
	r := g.Get(x, y)
	return r.(rune)
}

func (g *Grid) Set(x, y int, v interface{}) {
	k := key(x, y)
	current, ok := g.data[k]
	if ok && v == current {
		return
	}
	if !ok && v == g.empty {
		return
	}
	if v == g.empty {
		delete(g.data, k)
	} else {
		g.data[k] = v
		g.minX = utils.IntMin(g.minX, x)
		g.maxX = utils.IntMax(g.maxX, x)
		g.minY = utils.IntMin(g.minY, y)
		g.maxY = utils.IntMax(g.maxY, y)
	}
}

func (g *Grid) Print() {
	for j := g.minY; j <= g.maxY; j++ {
		for i := g.minX; i <= g.maxX; i++ {
			fmt.Printf("%c", g.GetRune(i, j))
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (g *Grid) PrintN() {
	for j := g.minY; j <= g.maxY; j++ {
		for i := g.minX; i <= g.maxX; i++ {
			fmt.Printf("%d ", g.Get(i, j))
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (g *Grid) PrintWithFormatter(formatter func(interface{}, int, int) rune) {
	for j := g.minY; j <= g.maxY; j++ {
		for i := g.minX; i <= g.maxX; i++ {
			fmt.Printf("%c ", formatter(g.Get(i, j), i, j))
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (g *Grid) PrintWithStringFormatter(formatter func(interface{}, int, int) string) {
	for j := g.minY; j <= g.maxY; j++ {
		for i := g.minX; i <= g.maxX; i++ {
			fmt.Printf("%s ", formatter(g.Get(i, j), i, j))
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func key(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}
