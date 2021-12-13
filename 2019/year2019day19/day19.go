package year2019day19

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

func Part1(program string) int {
	g := grids.NewGrid('?')

	r := 0
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			in := make(chan int)
			out := make(chan int)
			newIntCode(program, in, out)

			in <- i
			in <- j
			output := <-out
			if output == 1 {
				g.Set(i, j, '#')
				r++
			} else {
				g.Set(i, j, '.')
			}
		}
	}
	g.Print()
	return r
}

func Part2(program string) int {
	g := grids.NewGrid(false)

	left := 0
	right := 0
	for j := 0; ; j++ {
		if check(program, left-1, j) {
			left--
		}
		if check(program, right+1, j) {
			right++
		}
		for i := left; i <= right; i++ {
			g.Set(i, j, true)
			if done(g, i, j) {
				return (i-99)*10000 + (j - 99)
			}
		}
		left += 2
		right += 2
	}
}

func done(g *grids.Grid, x, y int) bool {
	if !g.Get(x-99, y).(bool) || !g.Get(x, y-99).(bool) {
		// quick optimization
		return false
	}
	for i := x - 99; i <= x; i++ {
		for j := y - 99; j <= y; j++ {
			if !g.Get(i, j).(bool) {
				return false
			}
		}
	}
	return true
}

func check(program string, i, j int) bool {
	if i < 0 {
		return false
	}
	in := make(chan int)
	out := make(chan int)
	newIntCode(program, in, out)

	in <- i
	in <- j
	output := <-out
	return output == 1
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
