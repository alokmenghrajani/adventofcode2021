package year2018day15

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	r := Part1(`#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`)
	assert.Equal(t, 27730, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`)
	assert.Equal(t, 4988, r)
}

func TestPart1MoreTests1(t *testing.T) {
	r := Part1(`#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`)
	assert.Equal(t, 36334, r)
}

func TestReorderCreatures(t *testing.T) {
	gs := parseInput(`#######
#.G.E.#
#E.G.E#
#.G.E.#
#######`)
	gs.sortCreatures()

	s := gs.grid.StringWithFormatter(func(v cell, _, _ int) string {
		switch v.c {
		case cellWall:
			return "#"
		case cellEmpty:
			return "."
		case cellCreature:
			for i, c := range gs.creatures {
				if c == v.creature {
					return fmt.Sprintf("%d", i+1)
				}
			}
			panic("unreachable")
		}
		panic("unreachable")
	})
	fmt.Println(s)
	assert.Equal(t, `#######
#.1.2.#
#3.4.5#
#.6.7.#
#######
`, s)
}

func TestTargets(t *testing.T) {
	gs := parseInput(`#######
#E..G.#
#...#.#
#.G.#G#
#######`)
	gs.sortCreatures()
	targets := gs.targets(gs.creatures[0])

	assert.Equal(t, 6, len(targets))
	assert.Contains(t, targets, []int{3, 1})
	assert.Contains(t, targets, []int{5, 1})
	assert.Contains(t, targets, []int{5, 2})
	assert.Contains(t, targets, []int{2, 2})
	assert.Contains(t, targets, []int{3, 1})
	assert.Contains(t, targets, []int{3, 3})
}

func TestFilterClosest(t *testing.T) {
	gs := parseInput(`#######
#E..G.#
#...#.#
#.G.#G#
#######`)
	gs.sortCreatures()
	distances := gs.computeDistances(gs.creatures[0])
	targets := gs.targets(gs.creatures[0])
	v, x, y := filterClosest(distances, targets)

	assert.Equal(t, uint(2), v)
	assert.Equal(t, 3, x)
	assert.Equal(t, 1, y)
}

func TestMove(t *testing.T) {
	gs := parseInput(`#######
#.E...#
#.....#
#...G.#
#######`)
	gs.sortCreatures()
	gs.move(gs.creatures[0])
	s := gs.grid.StringWithFormatter(func(v cell, _, _ int) string {
		switch v.c {
		case cellWall:
			return "#"
		case cellEmpty:
			return "."
		case cellCreature:
			switch v.creature.race {
			case elf:
				return "E"
			case goblin:
				return "G"
			default:
				panic("unreachable")
			}
		}
		panic("unreachable")
	})
	fmt.Println(s)
	assert.Equal(t, `#######
#..E..#
#.....#
#...G.#
#######
`, s)
}
