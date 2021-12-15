package year2019day22

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, "9 8 7 6 5 4 3 2 1 0", helper(`deal into new stack`))
	assert.Equal(t, "3 4 5 6 7 8 9 0 1 2", helper(`cut 3`))
	assert.Equal(t, "6 7 8 9 0 1 2 3 4 5", helper(`cut -4`))
	assert.Equal(t, "0 7 4 1 8 5 2 9 6 3", helper(`deal with increment 3`))

	assert.Equal(t, "0 3 6 9 2 5 8 1 4 7", helper(`deal with increment 7
deal into new stack
deal into new stack`))

	assert.Equal(t, "3 0 7 4 1 8 5 2 9 6", helper(`cut 6
deal with increment 7
deal into new stack`))

	assert.Equal(t, "6 3 0 7 4 1 8 5 2 9", helper(`deal with increment 7
deal with increment 9
cut -2`))

	assert.Equal(t, "9 2 5 8 1 4 7 0 3 6", helper(`deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1`))
}

func helper(input string) string {
	deck := process(input, 10)
	r := []string{}
	for i := 0; i < 10; i++ {
		r = append(r, fmt.Sprint(deck[i]))
	}
	return strings.Join(r, " ")
}
