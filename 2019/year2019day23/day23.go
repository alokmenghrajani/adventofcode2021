package year2019day23

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type intCode struct {
	mem          map[int]int
	ip           int
	input        chan int
	output       chan int
	relativeBase int
	idleCount    int64
	id           *int
}

type nat struct {
	enabled bool
	x       int
	y       int
	lastY   int
}

func Part1(program string) int {
	cpus := make([]*intCode, 50)
	for i := 0; i < 50; i++ {
		in := make(chan int, 2)
		out := make(chan int, 3)
		cpu := newIntCode(program, in, out)
		in <- i
		cpus[i] = cpu
	}

	// read from all out channels
	r := func() int {
		for {
			for i := 0; i < 50; i++ {
				if len(cpus[i].output) == 3 {
					v1 := <-cpus[i].output
					v2 := <-cpus[i].output
					v3 := <-cpus[i].output

					if v1 == 255 {
						return v3
					}

					cpus[v1].input <- v2
					cpus[v1].input <- v3
				}
			}
		}
	}()
	return r
}

func Part2(program string) int {
	cpus := make([]*intCode, 50)
	for i := 0; i < 50; i++ {
		in := make(chan int, 1000)
		out := make(chan int, 1000)
		cpu := newIntCode(program, in, out)
		in <- i
		cpus[i] = cpu
	}

	// read from all out channels
	n := nat{}
	r := func() int {
		for {
			for i := 0; i < 50; i++ {
				if len(cpus[i].output) >= 3 {
					destination := <-cpus[i].output
					x := <-cpus[i].output
					y := <-cpus[i].output

					if destination == 255 {
						fmt.Printf("here: %d %d\n", x, y)
						n.x = x
						n.y = y
						n.enabled = true
					} else {
						cpus[destination].input <- x
						cpus[destination].input <- y
					}
				}
			}

			countIdle := 0
			for i := 0; i < 50; i++ {
				if cpus[i].idleCount >= 2 && len(cpus[i].input) == 0 {
					countIdle++
				}
			}
			if countIdle == 50 && n.enabled {
				fmt.Printf("nat sending %d %d\n", n.x, n.y)
				if n.lastY == n.y {
					return n.y
				}
				n.lastY = n.y
				cpus[0].input <- n.x
				cpus[0].input <- n.y
			}
		}
	}()

	return r
}

func newIntCode(program string, input, output chan int) *intCode {
	cpu := &intCode{
		mem:          map[int]int{},
		ip:           0,
		input:        input,
		output:       output,
		relativeBase: 0,
		idleCount:    0,
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
			var i int
			if cpu.id == nil {
				i = <-cpu.input
				cpu.id = &i
			} else {
				time.Sleep(time.Millisecond * 10) // this is dumb. IMHO this puzzle is broken...
				select {
				case i = <-cpu.input:
					atomic.StoreInt64(&cpu.idleCount, 0)
				default:
					i = -1
					atomic.AddInt64(&cpu.idleCount, 1)
				}
			}

			arg1 := cpu.mem[cpu.ip+1]
			cpu.argModeWrite(arg1Mode, arg1, i)
			cpu.ip += 2
		case 4:
			atomic.StoreInt64(&cpu.idleCount, 0)
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
