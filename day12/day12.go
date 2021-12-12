package day12

import (
	"strings"
)

type cave struct {
	id          string
	isSmall     bool
	connectedTo map[string]*cave
	visited     bool
}

func Part1(input string) int {
	caves := map[string]*cave{}
	for _, line := range strings.Split(input, "\n") {
		pieces := strings.Split(line, "-")
		r1 := getOrCreate(caves, pieces[0])
		r2 := getOrCreate(caves, pieces[1])
		r1.connectedTo[r2.id] = r2
		r2.connectedTo[r1.id] = r1
	}

	start := caves["start"]
	start.visited = true
	return countPaths(start)
}

func Part2(input string) int {
	caves := map[string]*cave{}
	for _, line := range strings.Split(input, "\n") {
		pieces := strings.Split(line, "-")
		r1 := getOrCreate(caves, pieces[0])
		r2 := getOrCreate(caves, pieces[1])
		r1.connectedTo[r2.id] = r2
		r2.connectedTo[r1.id] = r1
	}

	start := caves["start"]
	start.visited = true
	return countPaths2(start, false)
}

func countPaths(current *cave) int {
	if current.id == "end" {
		return 1
	}

	r := 0
	for _, v := range current.connectedTo {
		if v.isSmall && v.visited {
			continue
		}
		v.visited = true
		r += countPaths(v)
		v.visited = false
	}
	return r
}

func countPaths2(current *cave, visitedTwice bool) int {
	if current.id == "end" {
		return 1
	}

	r := 0
	for _, v := range current.connectedTo {
		if v.isSmall && v.visited {
			if visitedTwice {
				continue
			}
			if v.id != "start" {
				r += countPaths2(v, true)
			}
		} else {
			v.visited = true
			r += countPaths2(v, visitedTwice)
			v.visited = false
		}
	}
	return r
}

func getOrCreate(caves map[string]*cave, id string) *cave {
	r, ok := caves[id]
	if ok {
		return r
	}
	r = &cave{id: id, connectedTo: map[string]*cave{}}
	if id[0] >= 'a' && id[0] <= 'z' {
		r.isSmall = true
	}
	caves[id] = r
	return r
}
