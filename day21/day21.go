package day21

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type player struct {
	points int
	pos    int
}

var dice = 1
var diceRolls = 0

func Part1(input string) int {
	pieces := strings.Split(input, "\n")
	player1 := player{
		points: 0,
		pos:    utils.MustAtoi(strings.Split(pieces[0], ": ")[1]),
	}
	player2 := player{
		points: 0,
		pos:    utils.MustAtoi(strings.Split(pieces[1], ": ")[1]),
	}

	toPlay := 0
	players := []*player{&player1, &player2}
	for {
		move(players[toPlay])
		if players[toPlay].points >= 1000 {
			break
		}
		toPlay = 1 - toPlay
	}
	return players[1-toPlay].points * diceRolls
}

func Part2(input string) int {
	pieces := strings.Split(input, "\n")
	player1 := player{
		points: 0,
		pos:    utils.MustAtoi(strings.Split(pieces[0], ": ")[1]),
	}
	player2 := player{
		points: 0,
		pos:    utils.MustAtoi(strings.Split(pieces[1], ": ")[1]),
	}

	s := state{toPlay: 0, player1: player1, player2: player2}
	w1, w2 := calculateWins(s)
	return utils.IntMax(w1, w2)
}

type state struct {
	toPlay  int
	player1 player
	player2 player
}

var cache = map[state][2]int{}

func calculateWins(s state) (int, int) {
	v, ok := cache[s]
	if ok {
		return v[0], v[1]
	}

	w1 := 0
	w2 := 0
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				if s.toPlay == 0 {
					newPos := s.player1.pos + i + j + k
					if newPos > 10 {
						newPos -= 10
					}
					newPoints := s.player1.points + newPos
					if newPoints >= 21 {
						w1++
					} else {
						s2 := state{
							toPlay:  1,
							player1: player{pos: newPos, points: newPoints},
							player2: s.player2,
						}
						r1, r2 := calculateWins(s2)
						w1 += r1
						w2 += r2
					}
				} else {
					newPos := s.player2.pos + i + j + k
					if newPos > 10 {
						newPos -= 10
					}
					newPoints := s.player2.points + newPos
					if newPoints >= 21 {
						w2++
					} else {
						s2 := state{
							toPlay:  0,
							player1: s.player1,
							player2: player{pos: newPos, points: newPoints},
						}
						r1, r2 := calculateWins(s2)
						w1 += r1
						w2 += r2
					}
				}
			}
		}
	}

	cache[s] = [2]int{w1, w2}
	return w1, w2
}

func move(p *player) {
	n1 := rollDice()
	n2 := rollDice()
	n3 := rollDice()
	p.pos += n1 + n2 + n3
	for p.pos > 10 {
		p.pos -= 10
	}
	p.points += p.pos
}

func rollDice() int {
	r := dice
	dice++
	if dice == 101 {
		dice = 1
	}
	diceRolls++
	return r
}
