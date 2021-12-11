package year2019day15

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

type intCode struct {
	mem          map[int]int
	ip           int
	input        chan int
	output       chan int
	relativeBase int
}

type cell struct {
	v        byte
	x        int
	y        int
	d        int
	toOrigin *cell
}

func Part1(program string) int {
	g := grids.NewGrid(&cell{v: '?', x: utils.MinInt, y: utils.MinInt, d: utils.MaxInt, toOrigin: nil})

	current := &cell{v: '.', x: 0, y: 0, d: 0, toOrigin: nil}
	g.Set(0, 0, current)

	in := make(chan int)
	out := make(chan int)
	newIntCode(program, in, out)

	for {
		// find the closest '?' from 0, 0
		target := find(g)
		if target == nil {
			panic("meh")
		}
		// find path to target
		next := path(g, current, target)
		if next == nil {
			panic(":(")
		}
		if next.x == current.x-1 {
			in <- 3
		} else if next.x == current.x+1 {
			in <- 4
		} else if next.y == current.y-1 {
			in <- 1
		} else if next.y == current.y+1 {
			in <- 2
		}

		r, ok := <-out
		if !ok {
			panic("meh")
		}

		switch r {
		case 0:
			next.v = '#'
		case 1:
			next.v = '.'
			current = next
		case 2:
			// success!
			return next.d
		}
	}
}

func Part2(program string) int {
	g := grids.NewGrid(&cell{v: '?', x: utils.MinInt, y: utils.MinInt, d: utils.MaxInt, toOrigin: nil})

	current := &cell{v: '.', x: 0, y: 0, d: 0, toOrigin: nil}
	g.Set(0, 0, current)

	in := make(chan int)
	out := make(chan int)
	newIntCode(program, in, out)

	var oxygen *cell
	for {
		// find the closest '?' from 0, 0
		target := find(g)
		if target == nil {
			// we are done mapping
			break
		}

		// find path to target
		next := path(g, current, target)
		if next == nil {
			panic(":(")
		}
		if next.x == current.x-1 {
			in <- 3
		} else if next.x == current.x+1 {
			in <- 4
		} else if next.y == current.y-1 {
			in <- 1
		} else if next.y == current.y+1 {
			in <- 2
		}

		r, ok := <-out
		if !ok {
			panic("meh")
		}

		switch r {
		case 0:
			next.v = '#'
		case 1:
			next.v = '.'
			current = next
		case 2:
			next.v = '.'
			oxygen = next
			current = next
		}
	}

	// compute distance from oxygen to all other cells
	xMin, xMax := g.SizeX()
	yMin, yMax := g.SizeY()
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			c := g.Get(i, j).(*cell)
			c.d = utils.MaxInt
		}
	}

	oxygen.d = 0
	queue := []*cell{oxygen}
	last := 0
	for len(queue) > 0 {
		head := queue[0]
		last = utils.IntMax(last, head.d)
		queue = queue[1:]
		t := g.Get(head.x-1, head.y).(*cell)
		if t.v == '.' && head.d+1 < t.d {
			t.d = head.d + 1
			queue = append(queue, t)
		}
		t = g.Get(head.x+1, head.y).(*cell)
		if t.v == '.' && head.d+1 < t.d {
			t.d = head.d + 1
			queue = append(queue, t)
		}
		t = g.Get(head.x, head.y-1).(*cell)
		if t.v == '.' && head.d+1 < t.d {
			t.d = head.d + 1
			queue = append(queue, t)
		}
		t = g.Get(head.x, head.y+1).(*cell)
		if t.v == '.' && head.d+1 < t.d {
			t.d = head.d + 1
			queue = append(queue, t)
		}
	}

	return last
}

func find(g *grids.Grid) *cell {
	// find a '?' neighboring a '.'
	bestDistance := utils.MaxInt
	bestX := utils.MinInt
	bestY := utils.MinInt
	var bestFrom *cell

	xMin, xMax := g.SizeX()
	yMin, yMax := g.SizeY()
	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			c := g.Get(i, j).(*cell)
			if c.v != '.' {
				continue
			}
			if c.d+1 >= bestDistance {
				continue
			}
			t := g.Get(i-1, j).(*cell)
			if t.v == '?' {
				bestDistance = c.d + 1
				bestX = i - 1
				bestY = j
				bestFrom = c
				continue
			}
			t = g.Get(i+1, j).(*cell)
			if t.v == '?' {
				bestDistance = c.d + 1
				bestX = i + 1
				bestY = j
				bestFrom = c

				continue
			}
			t = g.Get(i, j-1).(*cell)
			if t.v == '?' {
				bestDistance = c.d + 1
				bestX = i
				bestY = j - 1
				bestFrom = c

				continue
			}
			t = g.Get(i, j+1).(*cell)
			if t.v == '?' {
				bestDistance = c.d + 1
				bestX = i
				bestY = j + 1
				bestFrom = c
				continue
			}
		}
	}
	if bestDistance == utils.MaxInt {
		return nil
	}

	c := &cell{v: '?', x: bestX, y: bestY, d: bestDistance, toOrigin: bestFrom}
	g.Set(bestX, bestY, c)
	return c
}

func path(g *grids.Grid, from, to *cell) *cell {
	// check if from in the [origin-to] path
	t := to
	var previous *cell
	for t != nil {
		if t == from {
			return previous
		}
		previous = t
		t = t.toOrigin
	}

	// otherwise we head towards the origin
	return from.toOrigin
}

func newIntCode(program string, input, output chan int) *intCode {
	cpu := &intCode{
		mem:          map[int]int{},
		ip:           0,
		input:        input,
		output:       output,
		relativeBase: 0,
	}

	for i, opcode := range strings.Split(program, ",") {
		cpu.mem[i] = utils.MustAtoi(opcode)
	}

	go func() {
		cpu.run()
	}()

	return cpu
}

func (cpu *intCode) run() {
	for {
		opcode := cpu.mem[cpu.ip] % 100
		arg1Mode := (cpu.mem[cpu.ip] / 100) % 10
		arg2Mode := (cpu.mem[cpu.ip] / 1000) % 10
		arg3Mode := (cpu.mem[cpu.ip] / 10000) % 10

		switch opcode {
		case 1:
			arg1 := cpu.mem[cpu.ip+1]
			arg2 := cpu.mem[cpu.ip+2]
			arg3 := cpu.mem[cpu.ip+3]
			cpu.argModeWrite(arg3Mode, arg3, cpu.argMode(arg1Mode, arg1)+cpu.argMode(arg2Mode, arg2))
			cpu.ip += 4
		case 2:
			arg1 := cpu.mem[cpu.ip+1]
			arg2 := cpu.mem[cpu.ip+2]
			arg3 := cpu.mem[cpu.ip+3]
			cpu.argModeWrite(arg3Mode, arg3, cpu.argMode(arg1Mode, arg1)*cpu.argMode(arg2Mode, arg2))
			cpu.ip += 4
		case 3:
			i := <-cpu.input
			arg1 := cpu.mem[cpu.ip+1]
			cpu.argModeWrite(arg1Mode, arg1, i)
			cpu.ip += 2
		case 4:
			arg1 := cpu.mem[cpu.ip+1]
			i := cpu.argMode(arg1Mode, arg1)
			cpu.output <- i
			cpu.ip += 2
		case 5:
			arg1 := cpu.mem[cpu.ip+1]
			arg2 := cpu.mem[cpu.ip+2]
			v := cpu.argMode(arg1Mode, arg1)
			if v != 0 {
				cpu.ip = cpu.argMode(arg2Mode, arg2)
			} else {
				cpu.ip += 3
			}
		case 6:
			arg1 := cpu.mem[cpu.ip+1]
			arg2 := cpu.mem[cpu.ip+2]
			v := cpu.argMode(arg1Mode, arg1)
			if v == 0 {
				cpu.ip = cpu.argMode(arg2Mode, arg2)
			} else {
				cpu.ip += 3
			}
		case 7:
			arg1 := cpu.mem[cpu.ip+1]
			arg2 := cpu.mem[cpu.ip+2]
			arg3 := cpu.mem[cpu.ip+3]
			v1 := cpu.argMode(arg1Mode, arg1)
			v2 := cpu.argMode(arg2Mode, arg2)
			r := 0
			if v1 < v2 {
				r = 1
			}
			cpu.argModeWrite(arg3Mode, arg3, r)
			cpu.ip += 4
		case 8:
			arg1 := cpu.mem[cpu.ip+1]
			arg2 := cpu.mem[cpu.ip+2]
			arg3 := cpu.mem[cpu.ip+3]
			v1 := cpu.argMode(arg1Mode, arg1)
			v2 := cpu.argMode(arg2Mode, arg2)
			r := 0
			if v1 == v2 {
				r = 1
			}
			cpu.argModeWrite(arg3Mode, arg3, r)
			cpu.ip += 4
		case 9:
			arg1 := cpu.mem[cpu.ip+1]
			v := cpu.argMode(arg1Mode, arg1)
			cpu.relativeBase += v
			cpu.ip += 2
		case 99:
			close(cpu.output)
			return
		default:
			panic(fmt.Errorf("unknown op: %d", cpu.mem[cpu.ip]))
		}
	}

}

func (cpu intCode) argMode(mode, value int) int {
	switch mode {
	case 0:
		return cpu.mem[value]
	case 1:
		return value
	case 2:
		return cpu.mem[value+cpu.relativeBase]
	}
	panic(fmt.Errorf("unknown mode: %d", mode))
}

func (cpu intCode) argModeWrite(mode, addr, value int) {
	switch mode {
	case 0:
		cpu.mem[addr] = value
	case 1:
		panic(fmt.Errorf("invalid mode for writing: %d", mode))
	case 2:
		cpu.mem[addr+cpu.relativeBase] = value
	default:
		panic(fmt.Errorf("unknown mode: %d", mode))
	}
}
