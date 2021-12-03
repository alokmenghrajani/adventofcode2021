package day03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	size := len(lines[0])
	ones := make([]int, size)

	for _, line := range lines {
		if len(line) != size {
			panic(fmt.Errorf("invalid input: %s", line))
		}
		for i := 0; i < len(line); i++ {
			if line[i] == '1' {
				ones[i]++
			}
		}
	}

	gamma := 0
	epsilon := 0
	for i := 0; i < len(ones); i++ {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if ones[i] > len(lines)/2 {
			gamma = gamma | 0x01
		} else {
			epsilon = epsilon | 0x01
		}
	}

	return gamma * epsilon
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	size := len(lines[0])

	lines2 := make([]string, len(lines))
	copy(lines2, lines)

	bit := 0
	for len(lines2) > 1 {
		if bit == size {
			panic("meh")
		}

		ones := 0

		for _, line := range lines2 {
			if line[bit] == '1' {
				ones++
			}
		}

		lines3 := []string{}
		for _, line := range lines2 {
			if ones+ones >= len(lines2) {
				if line[bit] == '1' {
					lines3 = append(lines3, line)
				}
			} else if line[bit] == '0' {
				lines3 = append(lines3, line)
			}
		}
		lines2 = lines3
		bit++
	}

	oxygen, err := strconv.ParseInt(lines2[0], 2, 64)
	utils.PanicOnErr(err)

	lines2 = make([]string, len(lines))
	copy(lines2, lines)

	bit = 0
	for len(lines2) > 1 {
		if bit == size {
			panic("meh")
		}

		ones := 0

		for _, line := range lines2 {
			if line[bit] == '1' {
				ones++
			}
		}

		lines3 := []string{}
		for _, line := range lines2 {
			if ones+ones >= len(lines2) {
				if line[bit] == '0' {
					lines3 = append(lines3, line)
				}
			} else if line[bit] == '1' {
				lines3 = append(lines3, line)
			}
		}
		lines2 = lines3
		bit++
	}

	co2, err := strconv.ParseInt(lines2[0], 2, 64)
	utils.PanicOnErr(err)

	return int(oxygen * co2)
}
