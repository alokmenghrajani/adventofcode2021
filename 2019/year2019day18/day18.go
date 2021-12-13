package year2019day18

import (
	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
	"github.com/alokmenghrajani/adventofcode2021/utils/inputs"
)

type point struct {
	x, y int
}

type pointDistance struct {
	x, y, d int
}

type state struct {
	g       *grids.Grid
	pos     point
	keys    map[rune]point
	grabbed map[rune]bool
}

func Part1(input string) int {
	s := &state{
		g:       inputs.ToGrid(input, '#'),
		pos:     point{},
		keys:    map[rune]point{},
		grabbed: map[rune]bool{},
	}

	xMin, xMax := s.g.SizeX()
	yMin, yMax := s.g.SizeY()

	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			c := s.g.Get(i, j).(rune)
			if c == '@' {
				s.pos.x = i
				s.pos.y = j
				s.g.Set(i, j, '.')
			}
			if c >= 'a' && c <= 'z' {
				s.keys[c] = point{x: i, y: j}
			}
		}
	}

	return solve(s)
}

func solve(s *state) int {
	if len(s.grabbed) == len(s.keys) {
		return 0
	}

	keys := []pointDistance{}

	// find all the keys reachable
	g := grids.NewGrid(false)
	queue := []pointDistance{}
	queue = append(queue, pointDistance{x: s.pos.x, y: s.pos.y, d: 0})
	g.Set(s.pos.x, s.pos.y, true)

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		c := s.g.Get(head.x, head.y).(rune)
		if c >= 'a' && c <= 'z' && !s.grabbed[c] {
			// record that a key is available here
			keys = append(keys, head)
		} else {
			if canMoveTo(s, head.x-1, head.y) && !g.Get(head.x-1, head.y).(bool) {
				queue = append(queue, pointDistance{x: head.x - 1, y: head.y, d: head.d + 1})
				g.Set(head.x-1, head.y, true)
			}
			if canMoveTo(s, head.x+1, head.y) && !g.Get(head.x+1, head.y).(bool) {
				queue = append(queue, pointDistance{x: head.x + 1, y: head.y, d: head.d + 1})
				g.Set(head.x+1, head.y, true)
			}
			if canMoveTo(s, head.x, head.y-1) && !g.Get(head.x, head.y-1).(bool) {
				queue = append(queue, pointDistance{x: head.x, y: head.y - 1, d: head.d + 1})
				g.Set(head.x, head.y-1, true)
			}
			if canMoveTo(s, head.x, head.y+1) && !g.Get(head.x, head.y+1).(bool) {
				queue = append(queue, pointDistance{x: head.x, y: head.y + 1, d: head.d + 1})
				g.Set(head.x, head.y+1, true)
			}
		}
	}

	// find min solution
	best := utils.MaxInt
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		c := s.g.Get(key.x, key.y).(rune)
		if s.grabbed[c] {
			panic("meh")
		}
		s.grabbed[c] = true
		s.pos.x = key.x
		s.pos.y = key.y
		l := solve(s) + key.d
		delete(s.grabbed, c)

		best = utils.IntMin(best, l)
	}
	return best
}

func canMoveTo(s *state, x, y int) bool {
	c := byte(s.g.Get(x, y).(rune))
	if c == '#' {
		return false
	}
	if c == '.' {
		return true
	}
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return s.grabbed[rune(c-'A'+'a')]
	}
	panic("unreachable")
}
