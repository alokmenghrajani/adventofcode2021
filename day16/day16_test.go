package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	convert("D2FE28")
	assert.Equal(t, []int{1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0}, buffer)

	convert("D2FE28")
	r := parse()
	assert.Equal(t, 6, r.version)
	assert.Equal(t, 4, r.typ)
	assert.Equal(t, 2021, r.value)

	convert("38006F45291200")
	r = parse()
	assert.Equal(t, 1, r.version)
	assert.Equal(t, 6, r.typ)
	assert.Equal(t, 2, len(r.sub))
	assert.Equal(t, 4, r.sub[0].typ)
	assert.Equal(t, 10, r.sub[0].value)
	assert.Equal(t, 4, r.sub[1].typ)
	assert.Equal(t, 20, r.sub[1].value)

	convert("EE00D40C823060")
	r = parse()
	assert.Equal(t, 7, r.version)
	assert.Equal(t, 3, r.typ)
	assert.Equal(t, 3, len(r.sub))
	assert.Equal(t, 4, r.sub[0].typ)
	assert.Equal(t, 1, r.sub[0].value)
	assert.Equal(t, 4, r.sub[1].typ)
	assert.Equal(t, 2, r.sub[1].value)
	assert.Equal(t, 4, r.sub[2].typ)
	assert.Equal(t, 3, r.sub[2].value)

	assert.Equal(t, 16, Part1("8A004A801A8002F478"))
	assert.Equal(t, 12, Part1("620080001611562C8802118E34"))
	assert.Equal(t, 23, Part1("C0015000016115A2E0802F182340"))
	assert.Equal(t, 31, Part1("A0016C880162017C3686B18A3D4780"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 3, Part2("C200B40A82"))
	assert.Equal(t, 54, Part2("04005AC33890"))
	assert.Equal(t, 7, Part2("880086C3E88112"))
	assert.Equal(t, 9, Part2("CE00C43D881120"))
	assert.Equal(t, 1, Part2("D8005AC2A8F0"))
	assert.Equal(t, 0, Part2("F600BC2D8F"))
	assert.Equal(t, 0, Part2("9C005AC2F8F0"))
	assert.Equal(t, 1, Part2("9C0141080250320F1802104A08"))
}
