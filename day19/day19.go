package day19

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type point struct {
	x, y, z int
}

type best struct {
	overlaps  int
	transform Mat
}

type pair struct {
	id1  int
	id2  int
	best best
}

func (b *best) record(overlaps int, transform Mat) {
	if overlaps > b.overlaps {
		b.overlaps = overlaps
		b.transform = transform
	}
}

func (p point) add(other point) point {
	return point{
		x: p.x + other.x,
		y: p.y + other.y,
		z: p.z + other.z,
	}
}

func (p point) sub(other point) point {
	return point{
		x: p.x - other.x,
		y: p.y - other.y,
		z: p.z - other.z,
	}
}

var scanners = map[int]map[point]bool{}
var rotations []Mat

func Part1(input string) int {
	initRotations()

	for _, piece := range strings.Split(input, "\n\n") {
		lines := strings.Split(piece, "\n")
		l := lines[0]
		id := utils.MustAtoi(l[len("--- scanner ") : len(l)-4])
		points := map[point]bool{}
		for i := 1; i < len(lines); i++ {
			l := lines[i]
			pieces := strings.Split(l, ",")
			points[point{
				x: utils.MustAtoi(pieces[0]),
				y: utils.MustAtoi(pieces[1]),
				z: utils.MustAtoi(pieces[2]),
			}] = true
		}
		scanners[id] = points
	}

	// keep merging scanners into 0 until we have a single scanner left
	for len(scanners) > 1 {
		for id, scanner := range scanners {
			if id == 0 {
				continue
			}
			b := bestOverlap(0, id)
			if b.overlaps >= 12 {
				for p := range scanner {
					newP := b.transform.multiply(p)
					scanners[0][newP] = true
				}
				delete(scanners, id)
				break
			}
		}
	}

	return len(scanners[0])
}

func Part2(input string) int {
	initRotations()

	for _, piece := range strings.Split(input, "\n\n") {
		lines := strings.Split(piece, "\n")
		l := lines[0]
		id := utils.MustAtoi(l[len("--- scanner ") : len(l)-4])
		points := map[point]bool{}
		for i := 1; i < len(lines); i++ {
			l := lines[i]
			pieces := strings.Split(l, ",")
			points[point{
				x: utils.MustAtoi(pieces[0]),
				y: utils.MustAtoi(pieces[1]),
				z: utils.MustAtoi(pieces[2]),
			}] = true
		}
		scanners[id] = points
	}

	// keep merging scanners into 0 until we have a single scanner left
	positions := []point{}
	for len(scanners) > 1 {
		for id, scanner := range scanners {
			if id == 0 {
				continue
			}
			b := bestOverlap(0, id)
			if b.overlaps >= 12 {
				for p := range scanner {
					newP := b.transform.multiply(p)
					scanners[0][newP] = true
				}
				delete(scanners, id)
				positions = append(positions, point{b.transform.data[3], b.transform.data[7], b.transform.data[11]})
				break
			}
		}
	}

	max := 0
	for i := 0; i < len(positions); i++ {
		for j := 0; j < len(positions); j++ {
			max = utils.IntMax(max, dist(positions[i], positions[j]))
		}
	}

	return max
}

func dist(p1, p2 point) int {
	return utils.Abs(p1.x-p2.x) + utils.Abs(p1.y-p2.y) + utils.Abs(p1.z-p2.z)
}

func bestOverlap(scanner1, scanner2 int) best {
	b := best{}
	for orientation := 0; orientation < 24; orientation++ {
		mat := rotations[orientation]
		points1 := scanners[scanner1]
		points2 := scanners[scanner2]

		// compute all the possible deltas
		deltas := map[point]bool{}
		for p := range points1 {
			for p2 := range points2 {
				delta := p.sub(mat.multiply(p2))
				deltas[delta] = true
			}
		}

		// for all deltas, check how many points overlap
		for delta := range deltas {
			transform := NewMat([16]int{
				1, 0, 0, delta.x,
				0, 1, 0, delta.y,
				0, 0, 1, delta.z,
				0, 0, 0, 1,
			})
			transform = transform.dot(mat)

			overlaps := 0
			for p2 := range points2 {
				newP2 := transform.multiply(p2)
				if points1[newP2] {
					overlaps++
				}
			}
			if overlaps == 0 {
				panic("meh") // at least one overlap should happen because of the delta
			}
			b.record(overlaps, transform)
		}
	}
	return b
}

type Mat struct {
	data [16]int
}

func (m Mat) String() string {
	s := ""
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			s += fmt.Sprintf("%d ", m.data[i+j*3])
		}
		s += "\n"
	}
	s += "\n"

	return s
}

func NewMat(data [16]int) Mat {
	return Mat{data: data}
}

func (m Mat) multiply(p point) point {
	return point{
		x: m.data[0]*p.x + m.data[1]*p.y + m.data[2]*p.z + m.data[3]*1,
		y: m.data[4]*p.x + m.data[5]*p.y + m.data[6]*p.z + m.data[7]*1,
		z: m.data[8]*p.x + m.data[9]*p.y + m.data[10]*p.z + m.data[11]*1,
	}
}

func (m Mat) dot(other Mat) Mat {
	return Mat{data: [16]int{
		m.data[0]*other.data[0] + m.data[1]*other.data[4] + m.data[2]*other.data[8] + m.data[3]*other.data[12],
		m.data[0]*other.data[1] + m.data[1]*other.data[5] + m.data[2]*other.data[9] + m.data[3]*other.data[13],
		m.data[0]*other.data[2] + m.data[1]*other.data[6] + m.data[2]*other.data[10] + m.data[3]*other.data[14],
		m.data[0]*other.data[3] + m.data[1]*other.data[7] + m.data[2]*other.data[11] + m.data[3]*other.data[15],

		m.data[4]*other.data[0] + m.data[5]*other.data[4] + m.data[6]*other.data[8] + m.data[7]*other.data[12],
		m.data[4]*other.data[1] + m.data[5]*other.data[5] + m.data[6]*other.data[9] + m.data[7]*other.data[13],
		m.data[4]*other.data[2] + m.data[5]*other.data[6] + m.data[6]*other.data[10] + m.data[7]*other.data[14],
		m.data[4]*other.data[3] + m.data[5]*other.data[7] + m.data[6]*other.data[11] + m.data[7]*other.data[15],

		m.data[8]*other.data[0] + m.data[9]*other.data[4] + m.data[10]*other.data[8] + m.data[11]*other.data[12],
		m.data[8]*other.data[1] + m.data[9]*other.data[5] + m.data[10]*other.data[9] + m.data[11]*other.data[13],
		m.data[8]*other.data[2] + m.data[9]*other.data[6] + m.data[10]*other.data[10] + m.data[11]*other.data[14],
		m.data[8]*other.data[3] + m.data[9]*other.data[7] + m.data[10]*other.data[11] + m.data[11]*other.data[15],

		m.data[12]*other.data[0] + m.data[13]*other.data[4] + m.data[14]*other.data[8] + m.data[15]*other.data[12],
		m.data[12]*other.data[1] + m.data[13]*other.data[5] + m.data[14]*other.data[9] + m.data[15]*other.data[13],
		m.data[12]*other.data[2] + m.data[13]*other.data[6] + m.data[14]*other.data[10] + m.data[15]*other.data[14],
		m.data[12]*other.data[3] + m.data[13]*other.data[7] + m.data[14]*other.data[11] + m.data[15]*other.data[15],
	}}
}

func initRotations() {
	rot := NewMat([16]int{
		1, 0, 0, 0,
		0, 0, 1, 0,
		0, -1, 0, 0,
		0, 0, 0, 1,
	})
	flip := NewMat([16]int{
		-1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, -1, 0,
		0, 0, 0, 1,
	})
	bases := []Mat{
		// x = x
		NewMat([16]int{
			1, 0, 0, 0,
			0, 1, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		}),
		// x = y
		NewMat([16]int{
			0, 1, 0, 0,
			1, 0, 0, 0,
			0, 0, -1, 0,
			0, 0, 0, 1,
		}),
		// x = z
		NewMat([16]int{
			0, 0, -1, 0,
			0, 1, 0, 0,
			1, 0, 0, 0,
			0, 0, 0, 1,
		}),
	}

	rotations = []Mat{}
	for i := 0; i < 3; i++ {
		m := bases[i]
		for j := 0; j < 2; j++ {
			for k := 0; k < 4; k++ {
				rotations = append(rotations, m)
				m = m.dot(rot)
			}
			m = m.dot(flip)
		}
	}
}
