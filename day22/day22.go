package day22

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type point struct {
	x, y, z int
}

type cube struct {
	xMin, xMax, yMin, yMax, zMin, zMax int
}

func (c cube) volume() int {
	return utils.Abs(c.xMax-c.xMin+1) *
		utils.Abs(c.yMax-c.yMin+1) *
		utils.Abs(c.zMax-c.zMin+1)
}

func (c cube) intersects(other cube) bool {
	return !(other.xMin > c.xMax ||
		other.xMax < c.xMin ||
		other.yMin > c.yMax ||
		other.yMax < c.yMin ||
		other.zMin > c.zMax ||
		other.zMax < c.zMin)
}

func (c cube) intersection(other cube) *cube {
	if !c.intersects(other) {
		return nil
	}
	r := &cube{}
	r.xMin, r.xMax = intersectionSegment(c.xMin, c.xMax, other.xMin, other.xMax)
	r.yMin, r.yMax = intersectionSegment(c.yMin, c.yMax, other.yMin, other.yMax)
	r.zMin, r.zMax = intersectionSegment(c.zMin, c.zMax, other.zMin, other.zMax)

	if c.volume() < r.volume() {
		panic("meh")
	}
	if other.volume() < r.volume() {
		panic("meh")
	}

	return r
}

func intersectionSegment(min1, max1, min2, max2 int) (int, int) {
	// case 1: 111111
	//           22
	if max2 <= max1 && min2 >= min1 {
		return min2, max2
	}
	// case 2:   11
	//         22222
	if max2 >= max1 && min2 <= min1 {
		return min1, max1
	}
	// case 3: 1111
	//           2222
	if min2 <= max1 && min2 >= min1 {
		return min2, max1
	}
	// case 4:   1111
	//         2222
	if max2 >= min1 && max2 <= max1 {
		return min1, max2
	}

	panic("meh")
}

func Part1(input string) int {
	cubes := []*cube{}
	subtract := []*cube{}

	for _, line := range strings.Split(input, "\n") {
		var isOn bool
		if strings.HasPrefix(line, "on ") {
			isOn = true
			line = line[len("on "):]
		} else {
			isOn = false
			line = line[len("off "):]
		}
		pieces := strings.Split(line, ",")

		c := &cube{}
		var ok bool
		c.xMin, c.xMax, ok = parseAndClamp(pieces[0])
		if !ok {
			continue
		}
		c.yMin, c.yMax, ok = parseAndClamp(pieces[1])
		if !ok {
			continue
		}
		c.zMin, c.zMax, ok = parseAndClamp(pieces[2])
		if !ok {
			continue
		}

		addLater := []*cube{}
		for i := 0; i < len(subtract); i++ {
			t := c.intersection(*subtract[i])
			if t != nil {
				addLater = append(addLater, t)
			}
		}

		// check if this cube intersects with any existing cubes
		for i := 0; i < len(cubes); i++ {
			t := c.intersection(*cubes[i])
			if t != nil {
				subtract = append(subtract, t)
			}
		}
		cubes = append(cubes, addLater...)

		// record this cube
		if isOn {
			cubes = append(cubes, c)
		}
	}

	// calculate result
	r := 0
	for _, c := range cubes {
		r += c.volume()
	}
	for _, c := range subtract {
		r -= c.volume()
	}

	return r
}

func Part2(input string) int {
	cubes := []*cube{}
	subtract := []*cube{}

	for _, line := range strings.Split(input, "\n") {
		var isOn bool
		if strings.HasPrefix(line, "on ") {
			isOn = true
			line = line[len("on "):]
		} else {
			isOn = false
			line = line[len("off "):]
		}
		pieces := strings.Split(line, ",")

		c := &cube{}
		c.xMin, c.xMax = parse(pieces[0])
		c.yMin, c.yMax = parse(pieces[1])
		c.zMin, c.zMax = parse(pieces[2])

		addLater := []*cube{}
		for i := 0; i < len(subtract); i++ {
			t := c.intersection(*subtract[i])
			if t != nil {
				addLater = append(addLater, t)
			}
		}

		// check if this cube intersects with any existing cubes
		for i := 0; i < len(cubes); i++ {
			t := c.intersection(*cubes[i])
			if t != nil {
				subtract = append(subtract, t)
			}
		}
		cubes = append(cubes, addLater...)

		// record this cube
		if isOn {
			cubes = append(cubes, c)
		}
	}

	// calculate result
	r := 0
	for _, c := range cubes {
		r += c.volume()
	}
	for _, c := range subtract {
		r -= c.volume()
	}

	return r
}

func parse(s string) (int, int) {
	pieces := strings.Split(s[2:], "..")
	min := utils.MustAtoi(pieces[0])
	max := utils.MustAtoi(pieces[1])
	return min, max
}

// s is of the form .=N..N
func parseAndClamp(s string) (int, int, bool) {
	pieces := strings.Split(s[2:], "..")
	min := utils.MustAtoi(pieces[0])
	if min > 50 {
		return 0, 0, false
	}
	min = utils.IntMax(min, -50)

	max := utils.MustAtoi(pieces[1])
	if max < -50 {
		return 0, 0, false
	}
	max = utils.IntMin(max, 50)

	return min, max, true
}
