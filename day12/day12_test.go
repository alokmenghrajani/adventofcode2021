package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 10, Part1(`start-A
start-b
A-c
A-b
b-d
A-end
b-end`))

	assert.Equal(t, 19, Part1(`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`))

	assert.Equal(t, 226, Part1(`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 36, Part2(`start-A
start-b
A-c
A-b
b-d
A-end
b-end`))

	assert.Equal(t, 103, Part2(`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`))

	assert.Equal(t, 3509, Part2(`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`))
}
