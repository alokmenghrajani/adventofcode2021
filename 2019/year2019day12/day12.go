package year2019day12

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type pos struct {
	X, Y, Z    int
	VX, VY, VZ int
}

func Part1(input string, steps int) int {
	moons := []pos{}
	for _, line := range strings.Split(input, "\n") {
		var p pos
		re := regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)
		utils.MustParseToStruct(re, line, &p)
		moons = append(moons, p)
	}

	for step := 0; step < steps; step++ {
		// apply gravity
		for i := 0; i < len(moons); i++ {
			for j := 0; j < len(moons); j++ {
				if j == i {
					continue
				}
				moons[i].VX += utils.Sign(moons[j].X - moons[i].X)
				moons[i].VY += utils.Sign(moons[j].Y - moons[i].Y)
				moons[i].VZ += utils.Sign(moons[j].Z - moons[i].Z)
			}
		}

		// apply velocity
		for i := 0; i < len(moons); i++ {
			moons[i].X += moons[i].VX
			moons[i].Y += moons[i].VY
			moons[i].Z += moons[i].VZ
		}
	}

	sum := 0
	for i := 0; i < len(moons); i++ {
		pot := utils.Abs(moons[i].X) + utils.Abs(moons[i].Y) + utils.Abs(moons[i].Z)
		kin := utils.Abs(moons[i].VX) + utils.Abs(moons[i].VY) + utils.Abs(moons[i].VZ)
		sum += pot * kin
	}

	return sum
}

func Part2(input string) int {
	moons := []pos{}
	for _, line := range strings.Split(input, "\n") {
		var p pos
		re := regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)
		utils.MustParseToStruct(re, line, &p)
		moons = append(moons, p)
	}

	seen := map[string]int{}
	stepX := 0
	for ; ; stepX++ {
		state := ""
		for i := 0; i < len(moons); i++ {
			state += fmt.Sprintf("%d-%d ", moons[i].X, moons[i].VX)
		}
		if seen[state] > 0 {
			stepX = stepX - (seen[state] - 1)
			break
		}
		seen[state] = stepX + 1

		// apply gravity
		for i := 0; i < len(moons); i++ {
			for j := 0; j < len(moons); j++ {
				if j == i {
					continue
				}
				moons[i].VX += utils.Sign(moons[j].X - moons[i].X)
			}
		}

		// apply velocity
		for i := 0; i < len(moons); i++ {
			moons[i].X += moons[i].VX
		}
	}

	seen = map[string]int{}
	stepY := 0
	for ; ; stepY++ {
		state := ""
		for i := 0; i < len(moons); i++ {
			state += fmt.Sprintf("%d-%d ", moons[i].Y, moons[i].VY)
		}
		if seen[state] > 0 {
			stepY = stepY - (seen[state] - 1)
			break
		}
		seen[state] = stepY + 1

		// apply gravity
		for i := 0; i < len(moons); i++ {
			for j := 0; j < len(moons); j++ {
				if j == i {
					continue
				}
				moons[i].VY += utils.Sign(moons[j].Y - moons[i].Y)
			}
		}

		// apply velocity
		for i := 0; i < len(moons); i++ {
			moons[i].Y += moons[i].VY
		}
	}

	seen = map[string]int{}
	stepZ := 0
	for ; ; stepZ++ {
		state := ""
		for i := 0; i < len(moons); i++ {
			state += fmt.Sprintf("%d-%d ", moons[i].Z, moons[i].VZ)
		}
		if seen[state] > 0 {
			stepZ = stepZ - (seen[state] - 1)
			break
		}
		seen[state] = stepZ + 1

		// apply gravity
		for i := 0; i < len(moons); i++ {
			for j := 0; j < len(moons); j++ {
				if j == i {
					continue
				}
				moons[i].VZ += utils.Sign(moons[j].Z - moons[i].Z)
			}
		}

		// apply velocity
		for i := 0; i < len(moons); i++ {
			moons[i].Z += moons[i].VZ
		}
	}

	t := utils.Gcd(stepX, stepY)
	r := (stepX * stepY) / t
	t = utils.Gcd(r, stepZ)
	r = (r * stepZ) / t

	return r
}
