package year2019day10

import (
	"math"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

type Asteroid struct {
	x, y   int
	hidden bool
}

func Part1(input string) (int, int, int) {
	g := grids.NewGrid(nil)
	asteroids := []*Asteroid{}

	for j, line := range strings.Split(input, "\n") {
		for i := 0; i < len(line); i++ {
			if line[i] == '#' {
				asteroid := &Asteroid{x: i, y: j}
				asteroids = append(asteroids, asteroid)
				g.Set(i, j, asteroid)
			}
		}
	}

	max := utils.MinInt
	var maxAsteroid *Asteroid
	for _, p := range asteroids {
		t := len(asteroids) - countHidden(g, asteroids, p) - 1
		if t > max {
			max = t
			maxAsteroid = p
		}
	}

	return maxAsteroid.x, maxAsteroid.y, max
}

func Part2(input string, baseX, baseY, shots int) (int, int, int) {
	asteroids := []*Asteroid{}

	for j, line := range strings.Split(input, "\n") {
		for i := 0; i < len(line); i++ {
			if line[i] == '#' {
				asteroid := &Asteroid{x: i, y: j, hidden: false}
				if i == baseX && j == baseY {
					asteroid.hidden = true
				}
				asteroids = append(asteroids, asteroid)
			}
		}
	}

	var bestAsteroid *Asteroid
	angle := math.Pi * 3 / 2
	for shot := 0; shot < shots; shot++ {
		// find the asteroid which is closest (but larger) to angle
		// ignore hidden asteroids
		// break ties by distance
		bestAngleDelta := float64(utils.MaxInt)
		bestDistance := utils.MaxInt
		for _, asteroid := range asteroids {
			if asteroid.hidden {
				continue
			}
			a, d := computeAngleAndDistance(baseX, baseY, asteroid.x, asteroid.y)
			for a < angle {
				a += 2 * math.Pi
			}
			aDelta := a - angle
			if (roughlyEqual(aDelta, bestAngleDelta) && d < bestDistance) ||
				aDelta < bestAngleDelta {
				bestAngleDelta = aDelta
				bestAsteroid = asteroid
				bestDistance = d
			}
		}
		bestAsteroid.hidden = true

		angle += bestAngleDelta + 0.00001
	}

	return bestAsteroid.x, bestAsteroid.y, bestAsteroid.x*100 + bestAsteroid.y
}

func roughlyEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.00001
}

func computeAngleAndDistance(x1, y1, x2, y2 int) (float64, int) {
	dx := x2 - x1
	dy := y2 - y1
	a := math.Atan2(float64(dy), float64(dx))
	d := dx*dx + dy*dy
	return a, d
}

func countHidden(g *grids.Grid, asteroids []*Asteroid, asteroid *Asteroid) int {
	for _, p := range asteroids {
		p.hidden = false
	}
	for _, p := range asteroids {
		if p == asteroid {
			continue
		}
		minX, maxX := g.SizeX()
		minY, maxY := g.SizeY()
		dx := p.x - asteroid.x
		dy := p.y - asteroid.y

		if dx == 0 {
			dy = dy / utils.Abs(dy)
		} else if dy == 0 {
			dx = dx / utils.Abs(dx)
		} else {
			gcd := utils.Gcd(utils.Abs(dx), utils.Abs(dy))
			dx = dx / gcd
			dy = dy / gcd
		}

		x := p.x + dx
		y := p.y + dy
		for x >= minX && x <= maxX && y >= minY && y <= maxY {
			a := g.Get(x, y)
			if a != nil {
				a.(*Asteroid).hidden = true
			}
			x += dx
			y += dy
		}
	}

	r := 0
	for _, p := range asteroids {
		if p.hidden {
			r++
		}
	}
	return r
}
