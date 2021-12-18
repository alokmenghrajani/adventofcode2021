package day18

import (
	"fmt"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type pair struct {
	isLiteral bool
	value     int
	left      *pair
	right     *pair
}

func (p pair) String() string {
	if p.isLiteral {
		return fmt.Sprint(p.value)
	} else {
		return fmt.Sprintf("[%s,%s]", p.left.String(), p.right.String())
	}
}

func Part1(input string) int {
	v := sumAll(input)
	return mag(v)
}

func Part2(input string) int {
	m := utils.MinInt
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if j == i {
				continue
			}
			m = utils.IntMax(m, mag(sum(mustParse(lines[i]), mustParse(lines[j]))))
		}
	}
	return m
}

func sumAll(input string) *pair {
	lines := strings.Split(input, "\n")
	p := mustParse(lines[0])
	for i := 1; i < len(lines); i++ {
		v := mustParse(lines[i])
		p = sum(p, v)
	}
	return p
}

func parse(input string) (*pair, string) {
	if input[0] == '[' {
		left, i := parse(input[1:])
		if i[0] != ',' {
			panic(fmt.Errorf("parse failure: %s %s", input, i))
		}
		right, i := parse(i[1:])
		if i[0] != ']' {
			panic(fmt.Errorf("parse failure: %s %s", input, i))
		}
		return &pair{isLiteral: false, left: left, right: right}, i[1:]
	}
	offset := 0
	for input[offset] != ',' && input[offset] != ']' && offset < len(input) {
		offset++
	}
	return &pair{isLiteral: true, value: utils.MustAtoi(input[0:offset]), left: nil, right: nil}, input[offset:]
}

func mustParse(input string) *pair {
	p, s := parse(input)
	if s != "" {
		panic(fmt.Errorf("failed to parse: %s %s", input, s))
	}
	return p
}

func sum(v1, v2 *pair) *pair {
	newPair := &pair{isLiteral: false, left: v1, right: v2}
	done := false
	for !done {
		done = true

		// check if any pair is nested inside 4 pairs
		n := findDeep(newPair, 0)
		if n != nil {
			explode(newPair, n)
			done = false
			continue
		}
		// check if any value is more than 10
		v := findTenOrMore(newPair)
		if v != nil {
			split(v)
			done = false
		}
	}
	return newPair
}

func findDeep(node *pair, depth int) *pair {
	if node.isLiteral {
		return nil
	}
	if depth == 4 {
		return node
	}
	r := findDeep(node.left, depth+1)
	if r != nil {
		return r
	}
	return findDeep(node.right, depth+1)
}

func explode(root, v *pair) {
	// make a list of all the literals
	if v.isLiteral || !v.left.isLiteral || !v.right.isLiteral {
		panic("meh")
	}

	literals := []*pair{}
	listLiterals(root, &literals)

	// find v's position in the list
	pos := -1
	for i := 0; i < len(literals); i++ {
		if literals[i] == v.left {
			pos = i
			break
		}
	}
	if pos == -1 {
		panic("meh")
	}

	if pos > 0 {
		literals[pos-1].value += v.left.value
	}
	if pos+2 < len(literals) {
		literals[pos+2].value += v.right.value
	}
	v.isLiteral = true
	v.left = nil
	v.right = nil
	v.value = 0
}

func listLiterals(node *pair, literals *[]*pair) {
	if node.isLiteral {
		*literals = append(*literals, node)
	} else {
		listLiterals(node.left, literals)
		listLiterals(node.right, literals)
	}
}

func findTenOrMore(node *pair) *pair {
	if node.isLiteral {
		if node.value >= 10 {
			return node
		} else {
			return nil
		}
	}
	r := findTenOrMore(node.left)
	if r != nil {
		return r
	}
	return findTenOrMore(node.right)
}

func split(node *pair) {
	if !node.isLiteral {
		panic("meh")
	}
	node.isLiteral = false
	node.left = &pair{isLiteral: true, value: node.value / 2}
	node.right = &pair{isLiteral: true, value: node.value/2 + node.value%2}
}

func mag(p *pair) int {
	if p.isLiteral {
		return p.value
	}
	return 3*mag(p.left) + 2*mag(p.right)
}
