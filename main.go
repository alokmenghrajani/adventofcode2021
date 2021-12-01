package main

import (
	"fmt"
	"os"

	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day01"
	"github.com/alokmenghrajani/adventofcode2021/day01"
	"github.com/alokmenghrajani/adventofcode2021/utils"
)

// Usage: go run main.go <NN>
// assumes input is in day<NN>/input.txt
func main() {
	d := day()
	fmt.Printf("Running day %02d\n", d)

	switch d {
	case 1:
		fmt.Printf("part 1: %d\n", day01.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day01.Part2(utils.Readfile(2021, d)))

	// catching up on old events.
	// TODO: move this into its own repo.
	case 201901:
		fmt.Printf("part 1: %d\n", year2019day01.Part1(utils.Readfile(2019, 1)))
		fmt.Printf("part 2: %d\n", year2019day01.Part2(utils.Readfile(2019, 1)))
	default:
		panic(fmt.Errorf("no such day: %d", d))
	}
}

// Reads day from os.Args.
func day() int {
	if len(os.Args) == 1 {
		return 2
	}
	day := utils.Atoi(os.Args[1], -1)
	return day
}
