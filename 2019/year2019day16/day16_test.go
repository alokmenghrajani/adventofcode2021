package year2019day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, "48226158", Part1("12345678", 1))
	assert.Equal(t, "34040438", Part1("12345678", 2))
	assert.Equal(t, "03415518", Part1("12345678", 3))
	assert.Equal(t, "01029498", Part1("12345678", 4))

	assert.Equal(t, "24176176", Part1("80871224585914546619083218645595", 100))
	assert.Equal(t, "73745418", Part1("19617804207202209144916044189917", 100))
	assert.Equal(t, "52432133", Part1("69317163492948606335995924319873", 100))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "84462026", Part2("03036732577212944063491565474664", 100))
	assert.Equal(t, "78725270", Part2("02935109699940807407585447034323", 100))
	assert.Equal(t, "53553731", Part2("03081770884921959731165446850517", 100))
}
