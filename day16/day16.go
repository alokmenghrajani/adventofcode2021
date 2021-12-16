package day16

import (
	"strconv"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type packet struct {
	version  int
	typ      int
	value    int
	sub      []*packet
	consumed int
}

var buffer []int

func Part1(input string) int {
	convert(input)
	r := parse()

	return versionSum(r)
}

func Part2(input string) int {
	convert(input)
	r := parse()

	return compute(r)
}

func convert(input string) {
	buffer = []int{}
	for i := 0; i < len(input); i++ {
		v, err := strconv.ParseInt(input[i:i+1], 16, 8) // hmm...
		utils.PanicOnErr(err)
		for j := 3; j >= 0; j-- {
			buffer = append(buffer, (int(v)>>j)&1)
		}
	}
}

func parse() *packet {
	r := &packet{}

	r.version = consume(r, 3)
	r.typ = consume(r, 3)

	if r.typ == 4 {
		// literal
		for {
			t := consume(r, 5)

			r.value = (r.value << 4) | (t & 0xf)
			if (t & 0x10) == 0 {
				break
			}
		}
	} else {
		// operator
		lenType := consume(r, 1)
		if lenType == 0 {
			totalLength := consume(r, 15)

			subConsumed := 0
			for subConsumed != totalLength {
				if subConsumed > totalLength {
					panic("meh")
				}
				subPacket := parse()
				r.sub = append(r.sub, subPacket)
				subConsumed += subPacket.consumed
				r.consumed += subPacket.consumed // yuck!
			}
		} else if lenType == 1 {
			totalPackets := consume(r, 11)

			for i := 0; i < totalPackets; i++ {
				subPacket := parse()
				r.sub = append(r.sub, subPacket)
				r.consumed += subPacket.consumed // yuck!
			}
		} else {
			panic("meh")
		}
	}

	return r
}

func consume(p *packet, n int) int {
	r := 0
	for i := 0; i < n; i++ {
		r = (r << 1) | buffer[0]
		buffer = buffer[1:]
	}
	p.consumed += n
	return r
}

func versionSum(p *packet) int {
	s := p.version
	for i := 0; i < len(p.sub); i++ {
		s += versionSum(p.sub[i])
	}
	return s
}

func compute(p *packet) int {
	switch p.typ {
	case 4:
		return p.value
	case 0:
		sum := 0
		for i := 0; i < len(p.sub); i++ {
			sum += compute(p.sub[i])
		}
		return sum
	case 1:
		product := 1
		for i := 0; i < len(p.sub); i++ {
			product *= compute(p.sub[i])
		}
		return product
	case 2:
		min := utils.MaxInt
		for i := 0; i < len(p.sub); i++ {
			min = utils.IntMin(min, compute(p.sub[i]))
		}
		return min
	case 3:
		max := 0
		for i := 0; i < len(p.sub); i++ {
			max = utils.IntMax(max, compute(p.sub[i]))
		}
		return max
	case 5:
		v1 := compute(p.sub[0])
		v2 := compute(p.sub[1])
		if v1 > v2 {
			return 1
		}
		return 0
	case 6:
		v1 := compute(p.sub[0])
		v2 := compute(p.sub[1])
		if v1 < v2 {
			return 1
		}
		return 0
	case 7:
		v1 := compute(p.sub[0])
		v2 := compute(p.sub[1])
		if v1 == v2 {
			return 1
		}
		return 0
	}
	panic("meh")
}
