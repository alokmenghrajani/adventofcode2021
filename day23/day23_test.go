package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 12521, Run1(`#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`))
}
