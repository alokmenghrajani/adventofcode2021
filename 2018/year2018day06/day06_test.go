package year2018day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 17, Part1(`1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 16, Part2(`1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`))
}
