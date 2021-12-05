package year2019day05

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

func Part(program string, inputRaw string) string {
	input := []int{}
	for _, s := range strings.Split(inputRaw, "\n") {
		input = append(input, utils.MustAtoi(s))
	}

	output := compute(program, input)
	o := []string{}
	for _, n := range output {
		o = append(o, fmt.Sprint(n))
	}
	return strings.Join(o, ",")
}

func compute(program string, input []int) []int {
	memString := strings.Split(program, ",")
	mem := []int{}
	for _, s := range memString {
		mem = append(mem, utils.MustAtoi(s))
	}

	output := []int{}

	ip := 0
outer:
	for {
		opcode := mem[ip] % 100
		arg1Mode := (mem[ip] / 100) % 10
		arg2Mode := (mem[ip] / 1000) % 10
		arg3Mode := (mem[ip] / 10000) % 10

		switch opcode {
		case 1:
			arg1 := mem[ip+1]
			arg2 := mem[ip+2]
			res := mem[ip+3]
			if arg3Mode != 0 {
				panic("meh")
			}
			mem[res] = argMode(mem, arg1Mode, arg1) + argMode(mem, arg2Mode, arg2)
			ip += 4
		case 2:
			arg1 := mem[ip+1]
			arg2 := mem[ip+2]
			res := mem[ip+3]
			if arg3Mode != 0 {
				panic("meh")
			}
			mem[res] = argMode(mem, arg1Mode, arg1) * argMode(mem, arg2Mode, arg2)
			ip += 4
		case 3:
			i := input[0]
			input = input[1:]
			arg1 := mem[ip+1]
			if arg1Mode != 0 {
				panic("meh")
			}
			mem[arg1] = i
			ip += 2
		case 4:
			arg1 := mem[ip+1]
			i := argMode(mem, arg1Mode, arg1)
			output = append(output, i)
			ip += 2
		case 5:
			arg1 := mem[ip+1]
			arg2 := mem[ip+2]
			v := argMode(mem, arg1Mode, arg1)
			if v != 0 {
				ip = argMode(mem, arg2Mode, arg2)
			} else {
				ip += 3
			}
		case 6:
			arg1 := mem[ip+1]
			arg2 := mem[ip+2]
			v := argMode(mem, arg1Mode, arg1)
			if v == 0 {
				ip = argMode(mem, arg2Mode, arg2)
			} else {
				ip += 3
			}
		case 7:
			arg1 := mem[ip+1]
			arg2 := mem[ip+2]
			arg3 := mem[ip+3]
			v1 := argMode(mem, arg1Mode, arg1)
			v2 := argMode(mem, arg2Mode, arg2)
			r := 0
			if v1 < v2 {
				r = 1
			}
			if arg3Mode != 0 {
				panic("meh")
			}
			mem[arg3] = r
			ip += 4
		case 8:
			arg1 := mem[ip+1]
			arg2 := mem[ip+2]
			arg3 := mem[ip+3]
			v1 := argMode(mem, arg1Mode, arg1)
			v2 := argMode(mem, arg2Mode, arg2)
			r := 0
			if v1 == v2 {
				r = 1
			}
			if arg3Mode != 0 {
				panic("meh")
			}
			mem[arg3] = r
			ip += 4
		case 99:
			break outer
		default:
			panic(fmt.Errorf("unknown op: %d", mem[ip]))
		}
	}
	return output
}

func argMode(mem []int, mode, value int) int {
	switch mode {
	case 0:
		return mem[value]
	case 1:
		return value
	}
	panic(fmt.Errorf("unknown mode: %d", mode))
}
