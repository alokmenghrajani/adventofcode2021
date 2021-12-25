package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	r := run(`inp x
mul x -1`)
	assert.Equal(t, -9, r[9][1])

	r = run(`inp z
inp x
mul z 3
eql z x`)
	assert.Equal(t, 1, r[39][3])

	r = run(`inp z
inp x
mul z 3
eql z x`)
	assert.Equal(t, 0, r[99][3])

	r = run(`inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2`)
	assert.Equal(t, 1, r[9][3])
	assert.Equal(t, 0, r[9][2])
	assert.Equal(t, 0, r[9][1])
	assert.Equal(t, 1, r[9][0])
}
