package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	r := Part1(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`)
	assert.Equal(t, 198, r)
}

func TestPart2(t *testing.T) {
	r := Part2(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`)
	assert.Equal(t, 230, r)
}
