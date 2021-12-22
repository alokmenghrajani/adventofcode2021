package year2019day20

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

type point struct {
	x, y int
}

type cell struct {
	typ    byte // one of [#.*A-Z]
	portal point
}

var portals = map[string]*[]point{}
var dist = grids.NewGrid(-1)
var g = grids.NewGrid(cell{typ: '#'})

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	for j, line := range lines {
		for i, r := range line {
			g.Set(i, j, cell{typ: byte(r)})
		}
	}

	xMin, xMax := g.SizeX()
	yMin, yMax := g.SizeY()
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			c := g.Get(x, y).(cell)
			if isLetter(c.typ) {
				c2 := g.Get(x, y-1).(cell)
				if isLetter(c2.typ) {
					// other letter is above
					if g.Get(x, y-2).(cell).typ == '.' {
						// open space is above
						s := string(c2.typ) + string(c.typ)
						setPortal(s, point{x, y - 2})
					} else if g.Get(x, y+1).(cell).typ == '.' {
						// open space is below
						s := string(c2.typ) + string(c.typ)
						setPortal(s, point{x, y + 1})
					} else {
						panic("meh")
					}
					continue
				}

				c2 = g.Get(x, y+1).(cell)
				if isLetter(c2.typ) {
					// other letter is below
					if g.Get(x, y-1).(cell).typ == '.' {
						// open space is above
						s := string(c.typ) + string(c2.typ)
						setPortal(s, point{x, y - 1})
					} else if g.Get(x, y+2).(cell).typ == '.' {
						// open space is below
						s := string(c.typ) + string(c2.typ)
						setPortal(s, point{x, y + 2})
					} else {
						panic("meh")
					}
					continue
				}

				c2 = g.Get(x-1, y).(cell)
				if isLetter(c2.typ) {
					// other letter is to the left
					if g.Get(x-2, y).(cell).typ == '.' {
						// open space is left
						s := string(c2.typ) + string(c.typ)
						setPortal(s, point{x - 2, y})
					} else if g.Get(x+1, y).(cell).typ == '.' {
						// open space is right
						s := string(c2.typ) + string(c.typ)
						setPortal(s, point{x + 1, y})
					} else {
						panic("meh")
					}
					continue
				}

				c2 = g.Get(x+1, y).(cell)
				if isLetter(c2.typ) {
					// other letter is to the right
					if g.Get(x-1, y).(cell).typ == '.' {
						// open space is left
						s := string(c.typ) + string(c2.typ)
						setPortal(s, point{x - 1, y})
					} else if g.Get(x+2, y).(cell).typ == '.' {
						// open space is right
						s := string(c.typ) + string(c2.typ)
						setPortal(s, point{x + 2, y})
					} else {
						panic("meh")
					}
					continue
				}
				panic("meh")
			}
		}
	}

	for k, v := range portals {
		if k == "AA" || k == "ZZ" {
			continue
		}
		if len(*v) != 2 {
			panic("meh")
		}
		p1 := (*v)[0]
		p2 := (*v)[1]

		c := g.Get(p1.x, p1.y).(cell)
		c.typ = '*'
		c.portal = p2
		g.Set(p1.x, p1.y, c)

		c = g.Get(p2.x, p2.y).(cell)
		c.typ = '*'
		c.portal = p1
		g.Set(p2.x, p2.y, c)
	}

	start := (*portals["AA"])[0]
	dist.Set(start.x, start.y, 0)

	done := false
	for !done {
		done = true

		// classic flood fill
		xMin, xMax := dist.SizeX()
		yMin, yMax := dist.SizeY()
		for x := xMin; x <= xMax; x++ {
			for y := yMin; y <= yMax; y++ {
				c := g.Get(x, y).(cell)
				if c.typ != '.' && c.typ != '*' {
					continue
				}
				done = done && check(x, y, x+1, y)
				done = done && check(x, y, x-1, y)
				done = done && check(x, y, x, y+1)
				done = done && check(x, y, x, y-1)
				if c.typ == '*' {
					done = done && check(x, y, c.portal.x, c.portal.y)
				}
			}
		}
	}

	end := (*portals["ZZ"])[0]
	return dist.Get(end.x, end.y).(int)
}

func check(fromX, fromY, toX, toY int) bool {
	c := g.Get(toX, toY).(cell)
	if c.typ != '.' && c.typ != '*' {
		return true
	}
	d1 := dist.Get(fromX, fromY).(int)
	if d1 == -1 {
		return true
	}
	d2 := dist.Get(toX, toY).(int)
	if d2 == -1 || d2 > d1+1 {
		dist.Set(toX, toY, d1+1)
		return false
	}
	return true
}

func isLetter(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func getPortal(l string) *[]point {
	v, ok := portals[l]
	if ok {
		return v
	}
	v = &[]point{}
	portals[l] = v
	return v
}

func setPortal(l string, p point) {
	v := getPortal(l)
	for _, p2 := range *v {
		if p2 == p {
			return
		}
	}
	*v = append(*v, p)
}
