package year2019day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 8, Part1(`#########
#b.A.@.a#
#########`))

	assert.Equal(t, 86, Part1(`########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`))

	assert.Equal(t, 132, Part1(`########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`))

	// 	assert.Equal(t, 136, Part1(`#################
	// #i.G..c...e..H.p#
	// ########.########
	// #j.A..b...f..D.o#
	// ########@########
	// #k.E..a...g..B.n#
	// ########.########
	// #l.F..d...h..C.m#
	// #################`))

	assert.Equal(t, 81, Part1(`########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`))
}