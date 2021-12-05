package year2019day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, "1", Part("3,0,4,0,99", "1"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "0", Part("3,9,8,9,10,9,4,9,99,-1,8", "1"))
	assert.Equal(t, "1", Part("3,9,8,9,10,9,4,9,99,-1,8", "8"))

	assert.Equal(t, "1", Part("3,9,7,9,10,9,4,9,99,-1,8", "2"))
	assert.Equal(t, "0", Part("3,9,7,9,10,9,4,9,99,-1,8", "8"))
	assert.Equal(t, "0", Part("3,9,7,9,10,9,4,9,99,-1,8", "10"))

	assert.Equal(t, "0", Part("3,3,1108,-1,8,3,4,3,99", "3"))
	assert.Equal(t, "1", Part("3,3,1108,-1,8,3,4,3,99", "8"))

	assert.Equal(t, "1", Part("3,3,1107,-1,8,3,4,3,99", "4"))
	assert.Equal(t, "0", Part("3,3,1107,-1,8,3,4,3,99", "8"))
	assert.Equal(t, "0", Part("3,3,1107,-1,8,3,4,3,99", "9"))

	assert.Equal(t, "0", Part("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", "0"))
	assert.Equal(t, "1", Part("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", "12"))

	assert.Equal(t, "0", Part("3,3,1105,-1,9,1101,0,0,12,4,12,99,1", "0"))
	assert.Equal(t, "1", Part("3,3,1105,-1,9,1101,0,0,12,4,12,99,1", "12"))

	assert.Equal(t, "999", Part("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", "7"))
	assert.Equal(t, "1000", Part("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", "8"))
	assert.Equal(t, "1001", Part("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", "15"))
}
