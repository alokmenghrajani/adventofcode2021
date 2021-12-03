package year2019day02

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, "3500,9,10,70,2,3,11,0,99,30,40,50", testHelper("1,9,10,3,2,3,11,0,99,30,40,50"))
	assert.Equal(t, "2,0,0,0,99", testHelper("1,0,0,0,99"))
	assert.Equal(t, "2,3,0,6,99", testHelper("2,3,0,3,99"))
	assert.Equal(t, "2,4,4,5,99,9801", testHelper("2,4,4,5,99,0"))
	assert.Equal(t, "30,1,1,4,2,5,6,0,99", testHelper("1,1,1,4,99,5,6,0,99"))
}

func testHelper(s string) string {
	r := compute(s, -1, -1)

	t := []string{}
	for i := 0; i < len(r); i++ {
		t = append(t, strconv.Itoa(r[i]))
	}
	return strings.Join(t, ",")
}
