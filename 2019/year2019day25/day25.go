package year2019day25

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type intCode struct {
	mem          map[int]int
	ip           int
	input        chan int
	output       chan int
	relativeBase int
}

/*
      -. .-.
	   | |
       . .
       | |
       .-x-.-.-.
           |
		 .-.-.
		     |
		   .-.-.-.
			     |
				 .
*/

func Part1(program string) int {
	commands := []string{ // Hull Breach
		"north", // Arcade
		"take sand",
		"north", // Observatory
		"take space heater",
		"east", // Engineering
		"take semiconductor",
		"west",
		"south",
		"south",
		"east", // Corridor
		"take ornament",
		"south", // Science Lab
		"take festive hat",
		"east", // Storage
		"take asterisk",
		"south", // Navigation
		"east",  // Kitchen
		"take cake",
		"east",  // Sick Bay
		"south", // Gift Wrapping Center
		"north",
		"west",
		"west",
		"west", // Holodeck
		"take food ration",
		"east",
		"north",
		"west",
		"west", // Hot Chocolate Fountain
		"east",
		"north",
		"east", // Stables
		"east", // Crew Quarters
		"west",
		"west",
		"west",
		"west",  // Warp Drive Maintenance
		"north", // Hallway
		"north", // Security Checkpoint
		//"drop space heater",  // <
		//"drop semiconductor", // <
		"drop asterisk", // <
		//"drop ornament", // <
		"drop cake", // <
		// "drop festive hat", // < keep
		"drop food ration", // <
		"drop sand",        // >
		"inv",
		"west",
	}

	in := make(chan int, 1000)
	out := make(chan int)
	newIntCode(program, in, out)

	for _, command := range commands {
		for i := 0; i < len(command); i++ {
			in <- int(command[i])
		}
		in <- '\n'
	}
	last := 0
	for output := range out {
		fmt.Print(string(output))
		last = output
	}
	return last
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
