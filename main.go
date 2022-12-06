package main

import (
	"fmt"
	"os"

	"github.com/alokmenghrajani/adventofcode2021/2018/year2018day06"
	"github.com/alokmenghrajani/adventofcode2021/2018/year2018day13"
	"github.com/alokmenghrajani/adventofcode2021/2018/year2018day15"
	"github.com/alokmenghrajani/adventofcode2021/2018/year2018day17"
	"github.com/alokmenghrajani/adventofcode2021/2018/year2018day24"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day01"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day02"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day03"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day04"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day05"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day06"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day07"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day08"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day09"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day10"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day11"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day12"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day13"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day14"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day15"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day16"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day17"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day18"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day19"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day20"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day21"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day22"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day23"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day24"
	"github.com/alokmenghrajani/adventofcode2021/2019/year2019day25"
	"github.com/alokmenghrajani/adventofcode2021/day01"
	"github.com/alokmenghrajani/adventofcode2021/day02"
	"github.com/alokmenghrajani/adventofcode2021/day03"
	"github.com/alokmenghrajani/adventofcode2021/day04"
	"github.com/alokmenghrajani/adventofcode2021/day05"
	"github.com/alokmenghrajani/adventofcode2021/day06"
	"github.com/alokmenghrajani/adventofcode2021/day07"
	"github.com/alokmenghrajani/adventofcode2021/day08"
	"github.com/alokmenghrajani/adventofcode2021/day09"
	"github.com/alokmenghrajani/adventofcode2021/day10"
	"github.com/alokmenghrajani/adventofcode2021/day11"
	"github.com/alokmenghrajani/adventofcode2021/day12"
	"github.com/alokmenghrajani/adventofcode2021/day13"
	"github.com/alokmenghrajani/adventofcode2021/day14"
	"github.com/alokmenghrajani/adventofcode2021/day15"
	"github.com/alokmenghrajani/adventofcode2021/day16"
	"github.com/alokmenghrajani/adventofcode2021/day17"
	"github.com/alokmenghrajani/adventofcode2021/day18"
	"github.com/alokmenghrajani/adventofcode2021/day19"
	"github.com/alokmenghrajani/adventofcode2021/day20"
	"github.com/alokmenghrajani/adventofcode2021/day21"
	"github.com/alokmenghrajani/adventofcode2021/day22"
	"github.com/alokmenghrajani/adventofcode2021/day24"
	"github.com/alokmenghrajani/adventofcode2021/day25"
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
	case 6:
		fmt.Printf("part 1: %d\n", day06.Part(utils.Readfile(2021, d), 80))
		fmt.Printf("part 2: %d\n", day06.Part(utils.Readfile(2021, d), 256))
	case 7:
		_, fuel := day07.Part1(utils.Readfile(2021, d))
		fmt.Printf("part 1: %d\n", fuel)
		_, fuel = day07.Part2(utils.Readfile(2021, d))
		fmt.Printf("part 2: %d\n", fuel)
	case 8:
		fmt.Println("Using z3")
		fmt.Printf("part 1: %d\n", day08.Part1WithZ3(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day08.Part2WithZ3(utils.Readfile(2021, d)))
		fmt.Println("Using gophersat")
		fmt.Printf("part 1: %d\n", day08.Part1WithGophersat(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day08.Part2WithGophersat(utils.Readfile(2021, d)))
	case 9:
		fmt.Printf("part 1: %d\n", day09.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day09.Part2(utils.Readfile(2021, d)))
	case 10:
		fmt.Printf("part 1: %d\n", day10.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day10.Part2(utils.Readfile(2021, d)))
	case 11:
		fmt.Printf("part 1: %d\n", day11.Part1(utils.Readfile(2021, d), 100))
		fmt.Printf("part 2: %d\n", day11.Part2(utils.Readfile(2021, d)))
	case 12:
		fmt.Printf("part 1: %d\n", day12.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day12.Part2(utils.Readfile(2021, d)))
	case 13:
		fmt.Printf("part 1: %d\n", day13.Part1(utils.Readfile(2021, d)))
		fmt.Println("part 2")
		day13.Part2(utils.Readfile(2021, d))
	case 14:
		fmt.Printf("part 1: %d\n", day14.Part(utils.Readfile(2021, d), 10))
		fmt.Printf("part 2: %d\n", day14.Part(utils.Readfile(2021, d), 40))
	case 15:
		fmt.Printf("part 1: %d\n", day15.Part(utils.Readfile(2021, d), 1))
		fmt.Printf("part 2: %d\n", day15.Part(utils.Readfile(2021, d), 5))
	case 16:
		fmt.Printf("part 1: %d\n", day16.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day16.Part2(utils.Readfile(2021, d)))
	case 17:
		fmt.Printf("part 1: %d\n", day17.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day17.Part2(utils.Readfile(2021, d)))
	case 18:
		fmt.Printf("part 1: %d\n", day18.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day18.Part2(utils.Readfile(2021, d)))
	case 19:
		fmt.Printf("part 1: %d\n", day19.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day19.Part2(utils.Readfile(2021, d)))
	case 20:
		fmt.Printf("part 1: %d\n", day20.Part(utils.Readfile(2021, d), 2))
		fmt.Printf("part 1: %d\n", day20.Part(utils.Readfile(2021, d), 50))
	case 21:
		fmt.Printf("part 1: %d\n", day21.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day21.Part2(utils.Readfile(2021, d)))
	case 22:
		fmt.Printf("part 1: %d\n", day22.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day22.Part2(utils.Readfile(2021, d)))
	case 23:
		//		fmt.Printf("part 1: %d\n", day23.Run1(utils.Readfile(2021, d)))
	case 24:
		fmt.Printf("part 1: %d\n", day24.Part1(utils.Readfile(2021, d)))
		fmt.Printf("part 2: %d\n", day24.Part2(utils.Readfile(2021, d)))
	case 25:
		fmt.Printf("part 1: %d\n", day25.Part1(utils.Readfile(2021, d)))

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
	case 201907:
		_, v := year2019day07.Part1(utils.Readfile(2019, 7))
		fmt.Printf("part 1: %d\n", v)
		_, v = year2019day07.Part2(utils.Readfile(2019, 7))
		fmt.Printf("part 2: %d\n", v)
	case 201908:
		fmt.Printf("part 1: %d\n", year2019day08.Part1(utils.Readfile(2019, 8)))
		fmt.Printf("part 2: %s\n", year2019day08.Part2(25, 6, utils.Readfile(2019, 8)))
	case 201909:
		fmt.Printf("part 1: %s\n", year2019day09.Part(utils.Readfile(2019, 9), "1"))
		fmt.Printf("part 1: %s\n", year2019day09.Part(utils.Readfile(2019, 9), "2"))
	case 201910:
		x, y, v := year2019day10.Part1(utils.Readfile(2019, 10))
		fmt.Printf("part 1: %d\n", v)
		_, _, v = year2019day10.Part2(utils.Readfile(2019, 10), x, y, 200)
		fmt.Printf("part 2: %d\n", v)
	case 201911:
		fmt.Printf("part 1: %d\n", year2019day11.Part1(utils.Readfile(2019, 11)))
		fmt.Println("part 2:")
		year2019day11.Part2(utils.Readfile(2019, 11))
	case 201912:
		fmt.Printf("part 1: %d\n", year2019day12.Part1(utils.Readfile(2019, 12), 1000))
		fmt.Printf("part 2: %d\n", year2019day12.Part2(utils.Readfile(2019, 12)))
	case 201913:
		fmt.Printf("part 1: %d\n", year2019day13.Part1(utils.Readfile(2019, 13)))
		fmt.Println("part 2")
		year2019day13.Part2(utils.Readfile(2019, 13))
	case 201914:
		fmt.Printf("part 1: %d\n", year2019day14.Part1(utils.Readfile(2019, 14)))
		fmt.Printf("part 2: %d\n", year2019day14.Part2(utils.Readfile(2019, 14)))
	case 201915:
		fmt.Printf("part 1: %d\n", year2019day15.Part1(utils.Readfile(2019, 15)))
		fmt.Printf("part 2: %d\n", year2019day15.Part2(utils.Readfile(2019, 15)))
	case 201916:
		fmt.Printf("part 1: %s\n", year2019day16.Part1(utils.Readfile(2019, 16), 100))
		//		fmt.Printf("part 2: %d\n", year2019day16.Part2(utils.Readfile(2019, 16)))
	case 201917:
		fmt.Printf("part 1: %d\n", year2019day17.Part1(utils.Readfile(2019, 17)))
		fmt.Printf("part 2: %d\n", year2019day17.Part2(utils.Readfile(2019, 17)))
	case 201918:
		fmt.Printf("part 1: %d\n", year2019day18.Part1(utils.Readfile(2019, 18)))
	case 201919:
		fmt.Printf("part 1: %d\n", year2019day19.Part1(utils.Readfile(2019, 19)))
		fmt.Printf("part 2: %d\n", year2019day19.Part2(utils.Readfile(2019, 19)))
	case 201920:
		fmt.Printf("part 1: %d\n", year2019day20.Part1(utils.Readfile(2019, 20)))
	case 201921:
		fmt.Printf("part 1: %d\n", year2019day21.Part1(utils.Readfile(2019, 21)))
		fmt.Printf("part 2: %d\n", year2019day21.Part2(utils.Readfile(2019, 21)))
	case 201922:
		fmt.Printf("part 1: %d\n", year2019day22.Part1(utils.Readfile(2019, 22), 10007, 2019))
		//fmt.Printf("part 2: %d\n", year2019day22.Part2(utils.Readfile(2019, 22)))
	case 201923:
		fmt.Printf("part 1: %d\n", year2019day23.Part1(utils.Readfile(2019, 23)))
		fmt.Printf("part 2: %d\n", year2019day23.Part2(utils.Readfile(2019, 23)))
	case 201924:
		fmt.Printf("part 1: %d\n", year2019day24.Part1(utils.Readfile(2019, 24)))
		fmt.Printf("part 2: %d\n", year2019day24.Part2(utils.Readfile(2019, 24), 200))
	case 201925:
		fmt.Printf("part 1: %d\n", year2019day25.Part1(utils.Readfile(2019, 25)))

	case 201806:
		fmt.Printf("part 1: %d\n", year2018day06.Part1(utils.Readfile(2018, 6)))
		fmt.Printf("part 2: %d\n", year2018day06.Part2(utils.Readfile(2018, 6)))
	case 201813:
		fmt.Printf("part 1: %s\n", year2018day13.Part1(utils.Readfile(2018, 13)))
		fmt.Printf("part 2: %s\n", year2018day13.Part2(utils.Readfile(2018, 13)))
	case 201815:
		fmt.Printf("part 1: %d\n", year2018day15.Part1(utils.Readfile(2018, 15)))
		fmt.Printf("part 2: %d\n", year2018day15.Part2(utils.Readfile(2018, 15)))
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
		return 201815
	}
	day := utils.Atoi(os.Args[1], -1)
	return day
}
