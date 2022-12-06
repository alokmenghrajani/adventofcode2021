package year2018day15

import (
	"fmt"
	"sort"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grid2d"
)

type raceType string

const (
	goblin raceType = "goblin"
	elf    raceType = "elf"
)

type cellType byte

const (
	cellWall     cellType = '#'
	cellEmpty    cellType = '.'
	cellCreature cellType = '*'
)

type creatureType struct {
	race      raceType
	hitPoints int
	attack    int
	x, y      int
}

type cell struct {
	creature *creatureType
	c        cellType
}

type gameState struct {
	grid      *grid2d.Grid[cell]
	creatures []*creatureType
	goblins   int
	elves     int
}

type distance struct {
	distance uint
	deltaX   int
	deltaY   int
}

// I hate globals but this is convenient...
var moves [][]int = [][]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}

func Part1(input string) int {
	gs := parseInput(input)

	fullRounds := 0
outer:
	for {
		gs.sortCreatures()
		for _, creature := range gs.creatures {
			if creature.hitPoints <= 0 {
				continue
			}
			if gs.elves == 0 || gs.goblins == 0 {
				break outer
			}
			gs.moveAndAttack(creature)
		}
		fullRounds++
	}
	sum := 0
	for _, creature := range gs.creatures {
		if creature.hitPoints <= 0 {
			continue
		}
		sum += creature.hitPoints
	}
	return fullRounds * sum
}

func Part2(input string) int {
	attack := 3
	for {
		attack++
		done, score := didWin(input, attack)
		if done {
			return score
		}
	}
}

func didWin(input string, newAttack int) (bool, int) {
	gs := parseInput(input)
	for _, creature := range gs.creatures {
		if creature.race == elf {
			creature.attack = newAttack
		}
	}

	fullRounds := 0
outer:
	for {
		gs.sortCreatures()
		for _, creature := range gs.creatures {
			if creature.hitPoints <= 0 {
				continue
			}
			if gs.elves == 0 || gs.goblins == 0 {
				break outer
			}
			elfKilled := gs.moveAndAttack(creature)
			if elfKilled {
				return false, 0
			}
		}
		fullRounds++
	}
	sum := 0
	for _, creature := range gs.creatures {
		if creature.hitPoints <= 0 {
			continue
		}
		sum += creature.hitPoints
	}
	return true, fullRounds * sum
}

func parseInput(input string) *gameState {
	rows := strings.Split(input, "\n")
	gs := &gameState{
		grid:      grid2d.NewGrid(len(rows[0]), len(rows), cell{creature: nil, c: cellWall}),
		creatures: []*creatureType{},
	}

	for j, line := range rows {
		for i, c := range line {
			switch c {
			case '.':
				gs.grid.Set(i, j, cell{c: cellEmpty})
			case '#':
				// nothing to do
			case 'G':
				newGoblin := &creatureType{
					race:      goblin,
					hitPoints: 200,
					attack:    3,
					x:         i,
					y:         j,
				}
				gs.grid.Set(i, j, cell{creature: newGoblin, c: cellCreature})
				gs.creatures = append(gs.creatures, newGoblin)
				gs.goblins++
			case 'E':
				newElf := &creatureType{
					race:      elf,
					hitPoints: 200,
					attack:    3,
					x:         i,
					y:         j,
				}
				gs.grid.Set(i, j, cell{creature: newElf, c: cellCreature})
				gs.creatures = append(gs.creatures, newElf)
				gs.elves++
			default:
				panic("unimplemented")
			}
		}
	}
	return gs
}

func (gs *gameState) sortCreatures() {
	sort.Slice(gs.creatures, func(i, j int) bool {
		return isFirst(gs.creatures[i].x, gs.creatures[i].y, gs.creatures[j].x, gs.creatures[j].y)
	})
}

func isFirst(x1, y1, x2, y2 int) bool {
	if y1 < y2 {
		return true
	}
	if y1 > y2 {
		return false
	}
	return x1 < x2
}

func (gs *gameState) moveAndAttack(creature *creatureType) bool {
	// check if there are any ennemies already in range
	opponent := gs.findOpponentInRange(creature)
	if opponent == nil {
		// implement move + find opponent again
		gs.move(creature)

		opponent = gs.findOpponentInRange(creature)
	}

	if opponent == nil {
		return false
	}

	// perform attack
	opponent.hitPoints -= creature.attack
	if opponent.hitPoints <= 0 {
		gs.grid.Set(opponent.x, opponent.y, cell{c: cellEmpty})
		switch creature.race {
		case goblin:
			gs.elves--
			return true
		case elf:
			gs.goblins--
		default:
			panic("unreachable")
		}
	}
	return false
}

func (gs *gameState) findOpponentInRange(creature *creatureType) *creatureType {
	x := creature.x
	y := creature.y
	minHit := utils.MaxInt
	var best *creatureType
	for _, move := range moves {
		c := gs.grid.Get(x+move[0], y+move[1])
		if c.c == cellCreature && c.creature.race != creature.race {
			if c.creature.hitPoints < minHit {
				best = c.creature
				minHit = c.creature.hitPoints
			}
		}
	}
	return best
}

func (gs *gameState) move(creature *creatureType) {
	// compute the distance to all the cells
	distances := gs.computeDistances(creature)

	// find list of targets
	targets := gs.targets(creature)

	// find closest target which is reachable
	value, x, y := filterClosest(distances, targets)
	if value == utils.MaxUint {
		// nowhere to move to
		return
	}

	// rebuild the path
	for i := 1; i < int(value); i++ {
		d := distances.Get(x, y)
		x = x - d.deltaX
		y = y - d.deltaY
	}

	if x != creature.x && y != creature.y {
		panic("unrechable")
	}

	// Update grid
	gs.grid.Set(x, y, cell{creature: creature, c: cellCreature})
	gs.grid.Set(creature.x, creature.y, cell{c: cellEmpty})

	// Update creature
	creature.x = x
	creature.y = y
}

func (gs *gameState) computeDistances(creature *creatureType) *grid2d.Grid[distance] {
	max := utils.MaxUint
	distances := grid2d.NewGrid(gs.grid.SizeX(), gs.grid.SizeY(), distance{
		distance: max,
	})
	distances.Set(creature.x, creature.y, distance{distance: 0})
	done := false
	for !done {
		done = true
		for j := 0; j < distances.SizeY(); j++ {
			for i := 0; i < distances.SizeX(); i++ {
				t := distances.Get(i, j)
				if t.distance == max {
					// don't propagate from unreachable cells
					continue
				}
				// update cell by looking at neighbor's value
				for _, move := range moves {
					c2 := gs.grid.Get(i+move[0], j+move[1])
					if c2.c != cellEmpty {
						// we only care about empty cells
						continue
					}

					t2 := distances.Get(i+move[0], j+move[1])
					if t2.distance > t.distance+1 {
						// we found a shorter path to this cell
						distances.Set(i+move[0], j+move[1], distance{
							distance: t.distance + 1,
							deltaX:   move[0],
							deltaY:   move[1],
						})
						done = false
					} else if t2.distance == t.distance+1 {
						// keep the path which is first in reading order
						if isFirst(-move[0], -move[1], -t2.deltaX, -t2.deltaY) {
							distances.Set(i+move[0], j+move[1], distance{
								distance: t.distance + 1,
								deltaX:   move[0],
								deltaY:   move[1],
							})
							done = false
						}
					}
				}
			}
		}
	}
	return distances
}

func printDistances(distances *grid2d.Grid[distance]) {
	r := distances.StringWithFormatter(func(d distance, _, _ int) string {
		if d.distance == utils.MaxUint {
			return "(-,-,-) "
		}
		return fmt.Sprintf("(%d,%d,%d) ", d.distance, d.deltaX, d.deltaY)
	})
	fmt.Println(r)
}

func (gs *gameState) targets(creature *creatureType) [][]int {
	r := [][]int{}
	for _, c := range gs.creatures {
		if c.hitPoints <= 0 {
			// ignore dead creatures
			continue
		}
		if c.race == creature.race {
			// ignore creatures of the same race
			continue
		}
		for _, move := range moves {
			i := c.x + move[0]
			j := c.y + move[1]
			c := gs.grid.Get(i, j)
			if c.c != cellEmpty {
				continue
			}
			r = append(r, []int{i, j})
		}
	}
	return r
}

func filterClosest(distances *grid2d.Grid[distance], targets [][]int) (uint, int, int) {
	var bestX, bestY int
	bestValue := utils.MaxUint
	for _, target := range targets {
		d := distances.Get(target[0], target[1])
		if d.distance < bestValue {
			bestX = target[0]
			bestY = target[1]
			bestValue = d.distance
		} else if d.distance == bestValue {
			if isFirst(target[0], target[1], bestX, bestY) {
				bestX = target[0]
				bestY = target[1]
				bestValue = d.distance
			}
		}
	}
	return bestValue, bestX, bestY
}
