package main

import (
	"fmt"
	"os"

	"github.com/alokmenghrajani/adventofcode2021/2018/year2018day06"
	"github.com/alokmenghrajani/adventofcode2021/2018/year2018day13"
	"github.com/alokmenghrajani/adventofcode2021/2018/year2018day17"
	"github.com/alokmenghrajani/adventofcode2021/2018/year2018day24"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day01"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day02"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day03"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day04"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day05"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day06"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day08"
	"github.com/alokmenghrajani/adventofcode2021/day01"
	"github.com/alokmenghrajani/adventofcode2021/day02"
	"github.com/alokmenghrajani/adventofcode2021/day03"
	"github.com/alokmenghrajani/adventofcode2021/day04"
	"github.com/alokmenghrajani/adventofcode2021/day05"
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
	case 2:
		fmt.Printf("part 1: %d\n", day02.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day02.Part2(utils.Readfile(2021, d)))
	case 3:
		fmt.Printf("part 1: %d\n", day03.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day03.Part2(utils.Readfile(2021, d)))
	case 4:
		fmt.Printf("part 1: %d\n", day04.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day04.Part2(utils.Readfile(2021, d)))
	case 5:
		fmt.Printf("part 1: %d\n", day05.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day05.Part2(utils.Readfile(2021, d)))

	// catching up on old events.
	// TODO: move this into its own repo.
	case 201901:
		fmt.Printf("part 1: %d\n", year2019day01.Part1(utils.Readfile(2019, 1)))
		fmt.Printf("part 2: %d\n", year2019day01.Part2(utils.Readfile(2019, 1)))
	case 201902:
		fmt.Printf("part 1: %d\n", year2019day02.Part1(utils.Readfile(2019, 2)))
		fmt.Printf("part 2: %d\n", year2019day02.Part2(utils.Readfile(2019, 2)))
	case 201903:
		fmt.Printf("part 1: %d\n", year2019day03.Part1(utils.Readfile(2019, 3)))
		fmt.Printf("part 2: %d\n", year2019day03.Part2(utils.Readfile(2019, 3)))
	case 201904:
		fmt.Printf("part 1: %d\n", year2019day04.Part1(134564, 585159))
		fmt.Printf("part 2: %d\n", year2019day04.Part2(134564, 585159))
	case 201905:
		fmt.Printf("part 1: %s\n", year2019day05.Part(utils.Readfile(2019, 5), "1"))
		fmt.Printf("part 2: %s\n", year2019day05.Part(utils.Readfile(2019, 5), "5"))
	case 201906:
		fmt.Printf("part 1: %d\n", year2019day06.Part1(utils.Readfile(2019, 6)))
		fmt.Printf("part 2: %d\n", year2019day06.Part2(utils.Readfile(2019, 6)))
	case 201908:
		fmt.Printf("part 1: %d\n", year2019day08.Part1(utils.Readfile(2019, 8)))
		fmt.Printf("part 2: %s\n", year2019day08.Part2(25, 6, utils.Readfile(2019, 8)))
	case 201806:
		fmt.Printf("part 1: %d\n", year2018day06.Part1(utils.Readfile(2018, 6)))
		fmt.Printf("part 2: %d\n", year2018day06.Part2(utils.Readfile(2018, 6)))
	case 201813:
		fmt.Printf("part 1: %s\n", year2018day13.Part1(utils.Readfile(2018, 13)))
		fmt.Printf("part 2: %s\n", year2018day13.Part2(utils.Readfile(2018, 13)))
	case 201817:
		fmt.Printf("part 1: %d\n", year2018day17.Part1(utils.Readfile(2018, 17)))
		fmt.Printf("part 2: %d\n", year2018day17.Part2(utils.Readfile(2018, 17)))
	case 201824:
		fmt.Printf("part 1: %d\n", year2018day24.Part1(utils.Readfile(2018, 24)))
		fmt.Printf("part 2: %d\n", year2018day24.Part2(utils.Readfile(2018, 24)))
	default:
		panic(fmt.Errorf("no such day: %d", d))
	}
}

// Reads day from os.Args.
func day() int {
	if len(os.Args) == 1 {
		return 201908
	}
	day := utils.Atoi(os.Args[1], -1)
	return day
}
