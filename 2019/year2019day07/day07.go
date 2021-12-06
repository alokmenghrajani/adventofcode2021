package year2019day07

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type intCode struct {
	mem    []int
	ip     int
	input  chan int
	output chan int
}

func Part1(program string) ([]int, int) {
	bestValue := utils.MinInt
	var bestPhase []int

	for phaseA := 0; phaseA <= 4; phaseA++ {
		for phaseB := 0; phaseB <= 4; phaseB++ {
			if phaseB == phaseA {
				continue
			}
			for phaseC := 0; phaseC <= 4; phaseC++ {
				if phaseC == phaseA {
					continue
				}
				if phaseC == phaseB {
					continue
				}
				for phaseD := 0; phaseD <= 4; phaseD++ {
					if phaseD == phaseA {
						continue
					}
					if phaseD == phaseB {
						continue
					}
					if phaseD == phaseC {
						continue
					}
					for phaseE := 0; phaseE <= 4; phaseE++ {
						if phaseE == phaseA {
							continue
						}
						if phaseE == phaseB {
							continue
						}
						if phaseE == phaseC {
							continue
						}
						if phaseE == phaseD {
							continue
						}

						in := make(chan int, 2)
						in <- phaseA

						outAinB := make(chan int, 2)
						outAinB <- phaseB

						outBinC := make(chan int, 2)
						outBinC <- phaseC

						outCinD := make(chan int, 2)
						outCinD <- phaseD

						outDinE := make(chan int, 2)
						outDinE <- phaseE

						out := make(chan int, 2)

						newIntCode(program, in, outAinB)
						newIntCode(program, outAinB, outBinC)
						newIntCode(program, outBinC, outCinD)
						newIntCode(program, outCinD, outDinE)
						newIntCode(program, outDinE, out)

						in <- 0

						result := <-out

						if result > bestValue {
							bestValue = result
							bestPhase = []int{phaseA, phaseB, phaseC, phaseD, phaseE}
						}
					}
				}
			}
		}
	}

	return bestPhase, bestValue
}

func Part2(program string) ([]int, int) {
	bestValue := utils.MinInt
	var bestPhase []int

	for phaseA := 5; phaseA <= 9; phaseA++ {
		for phaseB := 5; phaseB <= 9; phaseB++ {
			if phaseB == phaseA {
				continue
			}
			for phaseC := 5; phaseC <= 9; phaseC++ {
				if phaseC == phaseA {
					continue
				}
				if phaseC == phaseB {
					continue
				}
				for phaseD := 5; phaseD <= 9; phaseD++ {
					if phaseD == phaseA {
						continue
					}
					if phaseD == phaseB {
						continue
					}
					if phaseD == phaseC {
						continue
					}
					for phaseE := 5; phaseE <= 9; phaseE++ {
						if phaseE == phaseA {
							continue
						}
						if phaseE == phaseB {
							continue
						}
						if phaseE == phaseC {
							continue
						}
						if phaseE == phaseD {
							continue
						}

						in := make(chan int, 2)
						in <- phaseA

						outAinB := make(chan int, 2)
						outAinB <- phaseB

						outBinC := make(chan int, 2)
						outBinC <- phaseC

						outCinD := make(chan int, 2)
						outCinD <- phaseD

						outDinE := make(chan int, 2)
						outDinE <- phaseE

						out := make(chan int, 2)

						newIntCode(program, in, outAinB)
						newIntCode(program, outAinB, outBinC)
						newIntCode(program, outBinC, outCinD)
						newIntCode(program, outCinD, outDinE)
						newIntCode(program, outDinE, out)
						in <- 0

						var lastResult int
						for n := range out {
							in <- n
							lastResult = n
						}

						if lastResult > bestValue {
							bestValue = lastResult
							bestPhase = []int{phaseA, phaseB, phaseC, phaseD, phaseE}
						}
					}
				}
			}
		}
	}

	return bestPhase, bestValue
}

func newIntCode(program string, input, output chan int) *intCode {
	cpu := &intCode{
		mem:    []int{},
		ip:     0,
		input:  input,
		output: output,
	}

	for _, opcode := range strings.Split(program, ",") {
		cpu.mem = append(cpu.mem, utils.MustAtoi(opcode))
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
			res := cpu.mem[cpu.ip+3]
			if arg3Mode != 0 {
				panic("meh")
			}
			cpu.mem[res] = argMode(cpu.mem, arg1Mode, arg1) + argMode(cpu.mem, arg2Mode, arg2)
			cpu.ip += 4
		case 2:
			arg1 := cpu.mem[cpu.ip+1]
			arg2 := cpu.mem[cpu.ip+2]
			res := cpu.mem[cpu.ip+3]
			if arg3Mode != 0 {
				panic("meh")
			}
			cpu.mem[res] = argMode(cpu.mem, arg1Mode, arg1) * argMode(cpu.mem, arg2Mode, arg2)
			cpu.ip += 4
		case 3:
			i := <-cpu.input
			arg1 := cpu.mem[cpu.ip+1]
			if arg1Mode != 0 {
				panic("meh")
			}
			cpu.mem[arg1] = i
			cpu.ip += 2
		case 4:
			arg1 := cpu.mem[cpu.ip+1]
			i := argMode(cpu.mem, arg1Mode, arg1)
			cpu.output <- i
			cpu.ip += 2
		case 5:
			arg1 := cpu.mem[cpu.ip+1]
			arg2 := cpu.mem[cpu.ip+2]
			v := argMode(cpu.mem, arg1Mode, arg1)
			if v != 0 {
				cpu.ip = argMode(cpu.mem, arg2Mode, arg2)
			} else {
				cpu.ip += 3
			}
		case 6:
			arg1 := cpu.mem[cpu.ip+1]
			arg2 := cpu.mem[cpu.ip+2]
			v := argMode(cpu.mem, arg1Mode, arg1)
			if v == 0 {
				cpu.ip = argMode(cpu.mem, arg2Mode, arg2)
			} else {
				cpu.ip += 3
			}
		case 7:
			arg1 := cpu.mem[cpu.ip+1]
			arg2 := cpu.mem[cpu.ip+2]
			arg3 := cpu.mem[cpu.ip+3]
			v1 := argMode(cpu.mem, arg1Mode, arg1)
			v2 := argMode(cpu.mem, arg2Mode, arg2)
			r := 0
			if v1 < v2 {
				r = 1
			}
			if arg3Mode != 0 {
				panic("meh")
			}
			cpu.mem[arg3] = r
			cpu.ip += 4
		case 8:
			arg1 := cpu.mem[cpu.ip+1]
			arg2 := cpu.mem[cpu.ip+2]
			arg3 := cpu.mem[cpu.ip+3]
			v1 := argMode(cpu.mem, arg1Mode, arg1)
			v2 := argMode(cpu.mem, arg2Mode, arg2)
			r := 0
			if v1 == v2 {
				r = 1
			}
			if arg3Mode != 0 {
				panic("meh")
			}
			cpu.mem[arg3] = r
			cpu.ip += 4
		case 99:
			close(cpu.output)
			return
		default:
			panic(fmt.Errorf("unknown op: %d", cpu.mem[cpu.ip]))
		}
	}

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
