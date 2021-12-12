package year2019day16

import (
	"fmt"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

func Part1(input string, steps int) string {
	arr := []int{}
	for i := 0; i < len(input); i++ {
		arr = append(arr, utils.MustAtoi(input[i:i+1]))
	}

	pattern := []int{0, 1, 0, -1}
	for step := 0; step < steps; step++ {
		nextArr := []int{}
		for i := 0; i < len(arr); i++ {
			sum := 0
			for j := 0; j < len(arr); j++ {
				offset := ((j + 1) / (i + 1)) % 4
				t := arr[j] * pattern[offset]
				sum += t
			}
			nextArr = append(nextArr, utils.Abs(sum)%10)
		}
		arr = nextArr
	}

	s := ""
	for i := 0; i < 8; i++ {
		s += fmt.Sprint(arr[i])
	}
	return s
}

func Part2(input string, steps int) string {
	arr := make([]int, 0, len(input)*10000)
	for i := 0; i < len(input); i++ {
		arr = append(arr, utils.MustAtoi(input[i:i+1]))
	}
	for i := 0; i < 9999; i++ {
		arr = append(arr, arr[0:1000]...)
	}

	pattern := []int{0, 1, 0, -1}
	nextArr := make([]int, len(arr))
	for step := 0; step < steps; step++ {
		for i := 0; i < len(arr); i++ {
			sum := 0
			for j := 0; j < len(arr); j++ {
				offset := ((j + 1) / (i + 1)) % 4
				t := arr[j] * pattern[offset]
				sum += t
			}
			nextArr[i] = utils.Abs(sum) % 10
		}
		copy(arr, nextArr)
	}

	s := ""
	for i := 0; i < 8; i++ {
		s += fmt.Sprint(arr[i])
	}
	return s
}
