package year2019day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	quine := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	output := Part("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99", "")
	assert.Equal(t, quine, output)

	largeN := "1102,34915192,34915192,7,4,7,99,0"
	output = Part(largeN, "")
	assert.Equal(t, 16, len(output), output)

	output = Part("104,1125899906842624,99", "")
	assert.Equal(t, "1125899906842624", output)

}
