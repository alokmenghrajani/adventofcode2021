package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 26, Part("3,4,3,1,2", 18))
	assert.Equal(t, 5934, Part("3,4,3,1,2", 80))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 26984457539, Part("3,4,3,1,2", 256))
}
