package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 739785, Part1(`Player 1 starting position: 4
Player 2 starting position: 8`))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 444356092776315, Part2(`Player 1 starting position: 4
Player 2 starting position: 8`))
}
