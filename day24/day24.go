package day24

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type cpu [4]int

func regToPos(r byte) int {
	if r >= 'w' && r <= 'z' {
		return int(r - 'w')
	}
	panic("meh")
}

func Part1(code string) int {
	pieces := strings.Split(code, "inp w\n")[1:]
	pairs := []int{0, 13, 1, 12, 2, 3, 4, 7, 5, 6, 8, 9, 10, 11}

	solution := [14]int{}
	for i := 0; i < len(pairs); i += 2 {
		code := "inp w\n" + pieces[pairs[i]] + "inp w\n" + pieces[pairs[i+1]]
		cpus := run(code)
		max := 0
		for k, v := range cpus {
			if v[regToPos('z')] == 0 {
				if k > max {
					max = k
				}
			}
		}
		solution[pairs[i]] = max / 10
		solution[pairs[i+1]] = max % 10
	}
	fmt.Println(solution)

	r := 0
	for i := 0; i < len(solution); i++ {
		r = r*10 + solution[i]
	}
	return r
}

func Part2(code string) int {
	pieces := strings.Split(code, "inp w\n")[1:]
	pairs := []int{0, 13, 1, 12, 2, 3, 4, 7, 5, 6, 8, 9, 10, 11}

	solution := [14]int{}
	for i := 0; i < len(pairs); i += 2 {
		code := "inp w\n" + pieces[pairs[i]] + "inp w\n" + pieces[pairs[i+1]]
		cpus := run(code)
		min := utils.MaxInt
		for k, v := range cpus {
			if v[regToPos('z')] == 0 {
				if k < min {
					min = k
				}
			}
		}
		solution[pairs[i]] = min / 10
		solution[pairs[i+1]] = min % 10
	}

	r := 0
	for i := 0; i < len(solution); i++ {
		r = r*10 + solution[i]
	}
	return r
}

func run(code string) map[int]cpu {
	cpus := map[int]cpu{}
	cpus[0] = cpu{}

	for _, line := range strings.Split(code, "\n") {
		if line == "" {
			continue
		} else if ok, s := HasPrefix(line, "inp "); ok {
			reg := s[0]

			// increase number of possible states
			newCpus := make(map[int]cpu, len(cpus)*10)
			for i := 1; i <= 9; i++ {
				for k, v := range cpus {
					c := cpu{}

					// magic
					for j := 0; j < 4; j++ {
						c[j] = v[j]
					}
					c[regToPos(reg)] = i
					newCpus[k*10+i] = c
				}
			}

			cpus = newCpus
		} else if ok, s := HasPrefix(line, "add "); ok {
			for k, v := range cpus {
				reg1 := s[0]
				v2 := parseReg(v, s[2:])
				v[regToPos(reg1)] = v[regToPos(reg1)] + v2
				cpus[k] = v
			}
		} else if ok, s := HasPrefix(line, "mul "); ok {
			for k, v := range cpus {
				reg1 := s[0]
				v2 := parseReg(v, s[2:])
				v[regToPos(reg1)] = v[regToPos(reg1)] * v2
				cpus[k] = v
			}
		} else if ok, s := HasPrefix(line, "div "); ok {
			for k, v := range cpus {
				reg1 := s[0]
				v2 := parseReg(v, s[2:])
				v[regToPos(reg1)] = v[regToPos(reg1)] / v2
				cpus[k] = v
			}
		} else if ok, s := HasPrefix(line, "mod "); ok {
			for k, v := range cpus {
				reg1 := s[0]
				v2 := parseReg(v, s[2:])
				v[regToPos(reg1)] = v[regToPos(reg1)] % v2
				cpus[k] = v
			}
		} else if ok, s := HasPrefix(line, "eql "); ok {
			for k, v := range cpus {
				reg1 := s[0]
				v2 := parseReg(v, s[2:])
				if v[regToPos(reg1)] == v2 {
					v[regToPos(reg1)] = 1
				} else {
					v[regToPos(reg1)] = 0
				}
				cpus[k] = v
			}
		} else {
			panic("meh")
		}
	}

	return cpus
}

func parseReg(registers [4]int, s string) int {
	if s == "w" || s == "x" || s == "y" || s == "z" {
		return registers[regToPos(s[0])]
	}
	return utils.MustAtoi(s)
}

func HasPrefix(s, prefix string) (bool, string) {
	if !strings.HasPrefix(s, prefix) {
		return false, ""
	}
	s = s[len(prefix):]
	return true, s
}

func v2() {
	input := [14]int{}
	w := 0
	x := 0
	y := 0
	z := 0

	w = input[0]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	// div z 1
	x += 11 // add x 11
	if x == w {
		x = 0
	} else {
		x = 1
	}
	// eql x 0
	y = 0     // mul y 0
	y += 25   // add y 25
	y = y * x // mul y x
	y++       // add y 1
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	y += 3    // add y 3
	y = y * x // mul y x
	z += y    // add z y

	w = input[1]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	// div z 1
	x += 14 // add x 14
	if x == w {
		x = 0
	} else {
		x = 1
	}
	// eql x 0
	y = 0     // mul y 0
	y += 25   // add y 25
	y = y * x // mul y x
	y++       // add y 1
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	y += 7    // add y 7
	y = y * x // mul y x
	z += y    // add z y

	w = input[2]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	// div z 1
	x += 13 // add x 13
	if x == w {
		x = 0
	} else {
		x = 1
	}
	// eql x 0
	y = 0     // mul y 0
	y += 25   // add y 25
	y = y * x // mul y x
	y++       // add y 1
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	y++       // add y 1
	y = y * x // mul y x
	z += y    // add z y

	w = input[3]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	z = z / 26 // div z 26
	x += -4    // add x 13
	if x == w {
		x = 0
	} else {
		x = 1
	}
	// eql x 0
	y = 0     // mul y 0
	y += 25   // add y 25
	y = y * x // mul y x
	y++       // add y 1
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	y += 6    // add y 6
	y = y * x // mul y x
	z += y    // add z y

	w = input[4]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	// div z 1
	x += 11 // add x 11
	if x == w {
		x = 0
	} else {
		x = 1
	}
	// eql x 0
	y = 0     // mul y 0
	y += 25   // add y 25
	y = y * x // mul y x
	y++       // add y 1
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	y += 14   // add y 14
	y = y * x // mul y x
	z += y    // add z y

	w = input[5]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	// div z 1
	x += 10 // add x 10
	if x == w {
		x = 0
	} else {
		x = 1
	}
	// eql x 0
	y = 0     // mul y 0
	y += 25   // add y 25
	y = y * x // mul y x
	y++       // add y 1
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	y += 7    // add y 7
	y = y * x // mul y x
	z += y    // add z y

	w = input[6]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	z = z / 26 // div z 26
	x += -4    // add x -4
	if x == w {
		x = 0
	} else {
		x = 1
	}
	// eql x 0
	y = 0     // mul y 0
	y += 25   // add y 25
	y = y * x // mul y x
	y++       // add y 1
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	y += 9    // add y 9
	y = y * x // mul y x
	z += y    // add z y

	w = input[7]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	z = z / 26 // div z 26
	x += -12   // add x -12
	if x == w {
		x = 0
	} else {
		x = 1
	}
	// eql x 0
	y = 0     // mul y 0
	y += 25   // add y 25
	y = y * x // mul y x
	y++       // add y 1
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	y += 9    // add y 9
	y = y * x // mul y x
	z += y    // add z y

	w = input[8]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	// div z 1
	x += 10 // add x 10
	if x == w {
		x = 0
	} else {
		x = 1
	}
	// eql x 0
	y = 0     // mul y 0
	y += 25   // add y 25
	y = y * x // mul y x
	y++       // add y 1
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	y += 6    // add y 6
	y = y * x // mul y x
	z += y    // add z y

	w = input[9]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	z = z / 26 // div z 26
	x += -11   // add x -11
	if x == w {
		x = 0
	} else {
		x = 1
	}
	// eql x 0
	y = 0     // mul y 0
	y += 25   // add y 25
	y = y * x // mul y x
	y++       // add y 1
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	y += 4    // add y 4
	y = y * x // mul y x
	z += y    // add z y

	w = input[10]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	// div z 1
	x += 12 // add x 12
	if x == w {
		x = 0
	} else {
		x = 1
	}
	// eql x 0
	y = 0     // mul y 0
	y += 25   // add y 25
	y = y * x // mul y x
	y++       // add y 1
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	// add y 0
	y = y * x // mul y x
	z += y    // add z y

	w = input[11]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	z = z / 26 // div z 26
	x += -1    // add x -1 // add x 12
	if x == w {
		x = 0
	} else {
		x = 1
	}
	// eql x 0
	y = 0     // mul y 0
	y += 25   // add y 25
	y = y * x // mul y x
	y++       // add y 1
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	y += 7    // add y 7
	y = y * x // mul y x
	z += y    // add z y

	w = input[12]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	z = z / 26 // div z 26
	// add x 0
	if x == w {
		x = 0
	} else {
		x = 1
	}
	// eql x 0
	y = 0     // mul y 0
	y += 25   // add y 25
	y = y * x // mul y x
	y++       // add y 1
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	y += 12   // add y 12
	y = y * x // mul y x
	z += y    // add z y

	w = input[13]
	x = 0      // mul x 0
	x = z      // add x z
	x = x % 26 // mod x 26
	z = z / 26 // div z 26
	x += -11   // add x -11
	if x == w {
		x = 0
		y = 1
	} else {
		x = 1
		y = 26
	}
	z = z * y // mul z y
	y = 0     // mul y 0
	y += w    // add y w
	y++       // add y 1
	y = y * x // mul y x
	z += y    // add z y

}
