package year2019day08

import (
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

func Part1(input string) int {
	layers := []string{}
	for len(input) > 0 {
		layer := input[0:150]
		layers = append(layers, layer)
		input = input[150:]
	}

	minValue := utils.MaxInt
	minOffset := -1
	for offset, layer := range layers {
		c := strings.Count(layer, "0")
		if c < minValue {
			minValue = c
			minOffset = offset
		}
	}

	layer := layers[minOffset]
	return strings.Count(layer, "1") * strings.Count(layer, "2")
}

func Part2(width, height int, input string) string {
	layers := []string{}
	for len(input) > 0 {
		layer := input[0 : width*height]
		layers = append(layers, layer)
		input = input[width*height:]
	}

	r := ""
	for i := 0; i < width*height; i++ {
		if i%width == 0 {
			r += "\n"
		}
		r += find(layers, i)
	}

	return r
}

func find(layers []string, offset int) string {
	for _, layer := range layers {
		switch layer[offset] {
		case '0':
			return " "
		case '1':
			return "*"
		case '2':
			continue
		}
	}
	panic("meh")
}
