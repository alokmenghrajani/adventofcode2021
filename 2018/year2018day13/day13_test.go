package year2018day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, "7,3", Part1(`/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "6,4", Part2(`/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`))
}
