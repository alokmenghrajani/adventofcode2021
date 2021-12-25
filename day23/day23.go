package day23

// 16489 => too high

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/alokmenghrajani/adventofcode2021/utils"
// )

// type state struct {
// 	hallway [11]byte
// 	rooms   [4][2]byte
// 	energy  int
// }

// func (s state) clone() *state {
// 	s2 := state{}
// 	for i := 0; i < 11; i++ {
// 		s2.hallway[i] = s.hallway[i]
// 	}
// 	for i := 0; i < 4; i++ {
// 		for j := 0; j < 2; j++ {
// 			s2.rooms[i][j] = s.rooms[i][j]
// 		}
// 	}
// 	s2.energy = s.energy
// 	return &s2
// }

// func (s state) print() {
// 	for i := 0; i < 11; i++ {
// 		fmt.Print(string(s.hallway[i]))
// 	}
// 	fmt.Println()
// 	fmt.Print("##")
// 	for i := 0; i < 4; i++ {
// 		fmt.Print(string(s.rooms[i][1]))
// 		fmt.Print("#")
// 	}
// 	fmt.Println("#")

// 	fmt.Print(" #")
// 	for i := 0; i < 4; i++ {
// 		fmt.Print(string(s.rooms[i][0]))
// 		fmt.Print("#")
// 	}
// 	fmt.Println("#")
// 	fmt.Println()
// }

// func Part1(input string) int {
// 	lines := strings.Split(input, "\n")
// 	s := state{}
// 	for i := 3; i <= 9; i += 2 {
// 		s.rooms[(i-3)/2][1] = lines[2][i]
// 		s.rooms[(i-3)/2][0] = lines[3][i]
// 	}
// 	for i := 0; i < 11; i++ {
// 		s.hallway[i] = lines[1][i+1]
// 	}

// 	return solve(s)
// }

// func solve(s state) int {
// 	//	s.print()

// 	if s.isDone() {
// 		return s.energy
// 	}
// 	newStates := findStates(s)
// 	min := utils.MaxInt
// 	for _, newState := range newStates {
// 		min = utils.IntMin(min, solve(newState))
// 	}

// 	return min
// }

// func (s state) isDone() bool {
// 	for room := 0; room < 4; room++ {
// 		for roomOffset := 0; roomOffset < 2; roomOffset++ {
// 			if int(s.rooms[room][roomOffset]-'A') != room {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func (s state) canMoveToHallway(fromRoom, fromRoomOffset, toHallway int) (int, bool) {
// 	energy := 1
// 	e := s.rooms[fromRoom][fromRoomOffset]
// 	if e == '.' {
// 		return -1, false
// 	}
// 	// don't move once we have reached our final resting spot
// 	if int(s.rooms[fromRoom][0]-'A') == fromRoom {
// 		if int(s.rooms[fromRoom][1]-'A') == fromRoom {
// 			return -1, false
// 		}
// 		if s.rooms[fromRoom][1] == '.' {
// 			return -1, false
// 		}
// 	}

// 	// check that the cells are free
// 	if fromRoomOffset == 0 {
// 		energy++
// 		e2 := s.rooms[fromRoom][1]
// 		if e2 != '.' {
// 			return -1, false
// 		}
// 	}
// 	if toHallway == 2 || toHallway == 4 || toHallway == 6 || toHallway == 8 {
// 		return -1, false
// 	}
// 	x := (fromRoom + 1) * 2
// 	for {
// 		if s.hallway[x] != '.' {
// 			return -1, false
// 		}
// 		if x == toHallway {
// 			break
// 		}
// 		x += utils.Sign(toHallway - x)
// 		energy++
// 	}
// 	var cost int
// 	switch e {
// 	case 'A':
// 		cost = 1
// 	case 'B':
// 		cost = 10
// 	case 'C':
// 		cost = 100
// 	case 'D':
// 		cost = 1000
// 	default:
// 		panic("meh")
// 	}
// 	energy = energy * cost
// 	return energy, true
// }

// func (s state) canMoveToRoom(fromHallway, toRoom, toRoomOffset int) (int, bool) {
// 	energy := 1
// 	e := s.hallway[fromHallway]
// 	if e == '.' {
// 		return -1, false
// 	}
// 	// check room is destination
// 	if toRoom != int(e-'A') {
// 		return -1, false
// 	}
// 	// check room is empty or contains a friend
// 	if s.rooms[toRoom][1] != '.' {
// 		return -1, false
// 	}
// 	if toRoomOffset == 0 {
// 		if s.rooms[toRoom][0] != '.' {
// 			return -1, false
// 		}
// 	} else {
// 		if s.rooms[toRoom][0] != e {
// 			return -1, false
// 		}
// 	}
// 	// check that the cells are free
// 	x := fromHallway
// 	dest := (toRoom + 1) * 2
// 	for x != dest {
// 		x += utils.Sign(dest - x)
// 		energy++
// 		if s.hallway[x] != '.' {
// 			return -1, false
// 		}
// 	}
// 	var cost int
// 	switch e {
// 	case 'A':
// 		cost = 1
// 	case 'B':
// 		cost = 10
// 	case 'C':
// 		cost = 100
// 	case 'D':
// 		cost = 1000
// 	default:
// 		panic("meh")
// 	}
// 	energy = energy * cost
// 	return energy, true
// }

// func findStates(s state) []state {
// 	r := []state{}

// 	for hallway := 0; hallway < 11; hallway++ {
// 		for room := 0; room < 4; room++ {
// 			for roomOffset := 0; roomOffset < 2; roomOffset++ {
// 				e, ok := s.canMoveToHallway(room, roomOffset, hallway)
// 				if ok {
// 					s2 := s.clone()
// 					s2.hallway[hallway] = s2.rooms[room][roomOffset]
// 					s2.rooms[room][roomOffset] = '.'
// 					s2.energy += e
// 					r = append(r, *s2)
// 				}
// 				e, ok = s.canMoveToRoom(hallway, room, roomOffset)
// 				if ok {
// 					s2 := s.clone()
// 					s2.rooms[room][roomOffset] = s2.hallway[hallway]
// 					s2.hallway[hallway] = '.'
// 					s2.energy += e
// 					r = append(r, *s2)
// 				}
// 			}
// 		}
// 	}

// 	return r
// }
