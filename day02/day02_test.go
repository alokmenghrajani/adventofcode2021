package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	r := Part1(`forward 5
	down 5
	forward 8
	up 3
	down 8
	forward 2`)
	assert.Equal(t, 150, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`forward 5
	down 5
	forward 8
	up 3
	down 8
	forward 2`)
	assert.Equal(t, 900, r)
}
