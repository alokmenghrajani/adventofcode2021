package year2019day06

import "strings"

type node struct {
	name     string
	children []*node
	value    int
	tagged   bool
	parent   *node
}

func Part1(input string) int {
	nodes := map[string]*node{}

	for _, line := range strings.Split(input, "\n") {
		pieces := strings.Split(line, ")")
		parent, exists := nodes[pieces[0]]
		if !exists {
			parent = &node{name: pieces[0], children: []*node{}, value: -1}
		}
		nodes[parent.name] = parent

		child, exists := nodes[pieces[1]]
		if !exists {
			child = &node{name: pieces[1], children: []*node{}, value: -1}
		}
		nodes[child.name] = child

		parent.children = append(parent.children, child)
		child.parent = parent
	}

	root := nodes["COM"]
	root.value = 0
	process(root)

	r := 0
	for _, n := range nodes {
		if n.value == -1 {
			panic("meh")
		}
		r += n.value
	}

	return r
}

func Part2(input string) int {
	nodes := map[string]*node{}

	for _, line := range strings.Split(input, "\n") {
		pieces := strings.Split(line, ")")
		parent, exists := nodes[pieces[0]]
		if !exists {
			parent = &node{name: pieces[0], children: []*node{}, value: -1}
		}
		nodes[parent.name] = parent

		child, exists := nodes[pieces[1]]
		if !exists {
			child = &node{name: pieces[1], children: []*node{}, value: -1}
		}
		nodes[child.name] = child

		parent.children = append(parent.children, child)
		child.parent = parent
	}

	root := nodes["COM"]
	root.value = 0
	process(root)

	n := nodes["YOU"]
	for {
		n.tagged = true
		if n.value == 0 {
			break
		}
		n = n.parent
	}
	n = nodes["SAN"]
	for !n.tagged {
		n = n.parent
	}
	// n is the least common ancestor
	r := (nodes["YOU"].value - n.value) + (nodes["SAN"].value - n.value) - 2

	return r
}

func process(n *node) {
	if n.value == -1 {
		panic("meh")
	}
	for _, c := range n.children {
		c.value = n.value + 1
		process(c)
	}
}
