package day04

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
	"github.com/alokmenghrajani/adventofcode2021/utils/grids"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n\n")
	rng := strings.Split(lines[0], ",")
	boards := []*grids.Grid{}
	for _, boardInput := range lines[1:] {
		b := parseBoard(boardInput)
		boards = append(boards, b)
	}

	for i := 0; i < len(rng); i++ {
		r := utils.MustAtoi(rng[i])
		for _, board := range boards {
			if place(board, r) {
				if didWin(board) {
					return value(board) * r
				}
			}
		}
	}
	panic("meh")
}

func Part2(input string) int {
	lines := strings.Split(input, "\n\n")
	rng := strings.Split(lines[0], ",")
	boards := []*grids.Grid{}
	inPlay := []bool{}
	for _, boardInput := range lines[1:] {
		b := parseBoard(boardInput)
		boards = append(boards, b)
		inPlay = append(inPlay, true)
	}

	win := 0
	lastWin := -1
	for i := 0; i < len(rng); i++ {
		if win == len(boards) {
			return lastWin
		}

		r := utils.MustAtoi(rng[i])
		for j, board := range boards {
			if !inPlay[j] {
				continue
			}
			if place(board, r) {
				if didWin(board) {
					win++
					lastWin = value(board) * r
					inPlay[j] = false
				}
			}
		}
	}
	panic("meh")
}

func place(board *grids.Grid, rng int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board.Get(i, j).(int) == rng {
				board.Set(i, j, -1)
				return true
			}
		}
	}
	return false
}

func value(board *grids.Grid) int {
	v := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			t := board.Get(i, j).(int)
			if t >= 0 {
				v += t
			}
		}
	}
	return v
}

func didWin(board *grids.Grid) bool {
outer1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board.Get(i, j).(int) >= 0 {
				continue outer1
			}
		}
		return true
	}

outer2:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board.Get(j, i).(int) >= 0 {
				continue outer2
			}
		}
		return true
	}

	return false
}

func parseBoard(input string) *grids.Grid {
	r := grids.NewGrid(-1)
	rows := strings.Split(input, "\n")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			s := rows[i][j*3 : j*3+2]
			if s[0] == ' ' {
				s = s[1:]
			}
			n := utils.MustAtoi(s)
			r.Set(j, i, n)
		}
	}
	return r
}
