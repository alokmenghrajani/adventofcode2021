package year2019day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.True(t, valid(111111))
	assert.False(t, valid(223450))
	assert.False(t, valid(123789))
}

func TestPart2(t *testing.T) {
	assert.True(t, valid2(112233))
	assert.False(t, valid2(123444))
	assert.True(t, valid2(111122))
}
