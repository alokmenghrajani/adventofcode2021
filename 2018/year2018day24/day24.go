package year2018day24

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/alokmenghrajani/adventofcode2021/utils"
)

type group struct {
	Id             int
	Units          int
	OriginalUnits  int
	Hit            int
	ParseLater     string
	ImmuneTo       []string
	WeakTo         []string
	Attack         int
	OriginalAttack int
	AttackType     string
	Initiative     int
	Selected       bool
	Targetting     int
}

func Part1(input string) int {
	groups := strings.Split(input, "\n\n")
	immuneSystemInput := strings.Split(groups[0], "\n")[1:]
	infectionInput := strings.Split(groups[1], "\n")[1:]

	var immuneSystem []group
	id := 1
	for _, line := range immuneSystemInput {
		g := parse(line)
		g.Id = id
		id++
		immuneSystem = append(immuneSystem, g)
	}

	var infection []group
	id = 1
	for _, line := range infectionInput {
		g := parse(line)
		g.Id = id
		id++
		infection = append(infection, g)
	}

	for {
		// fmt.Println("Immune System:")
		// for i := 0; i < len(immuneSystem); i++ {
		// 	fmt.Printf("Group %d contains %d units\n", immuneSystem[i].Id, immuneSystem[i].Units)
		// }
		// fmt.Println("Infection:")
		// for i := 0; i < len(infection); i++ {
		// 	fmt.Printf("Group %d contains %d units\n", infection[i].Id, infection[i].Units)
		// }
		// fmt.Println("")

		// check if an army is dead
		i := count(immuneSystem)
		j := count(infection)
		if i == 0 || j == 0 {
			return i + j
		}

		// target selection
		resetTarget(immuneSystem)
		resetTarget(infection)

		targetSelection("Immune System", immuneSystem, infection)
		targetSelection("Infection", infection, immuneSystem)
		// fmt.Println("")

		// attack
		attack(immuneSystem, infection)
		// fmt.Println("")
	}
}

func Part2(input string) int {
	groups := strings.Split(input, "\n\n")
	immuneSystemInput := strings.Split(groups[0], "\n")[1:]
	infectionInput := strings.Split(groups[1], "\n")[1:]

	var immuneSystem []group
	id := 1
	for _, line := range immuneSystemInput {
		g := parse(line)
		g.Id = id
		id++
		immuneSystem = append(immuneSystem, g)
	}

	var infection []group
	id = 1
	for _, line := range infectionInput {
		g := parse(line)
		g.Id = id
		id++
		infection = append(infection, g)
	}

	for boost := 1; ; boost++ {
		for i := 0; i < len(immuneSystem); i++ {
			immuneSystem[i].Units = immuneSystem[i].OriginalUnits
			immuneSystem[i].Attack = immuneSystem[i].OriginalAttack + boost
		}
		for i := 0; i < len(infection); i++ {
			infection[i].Units = infection[i].OriginalUnits
		}

		t := part2(immuneSystem, infection)
		if t > 0 {
			return t
		}
	}
}

func part2(immuneSystem []group, infection []group) int {
	for {
		// fmt.Println("Immune System:")
		// for i := 0; i < len(immuneSystem); i++ {
		// 	fmt.Printf("Group %d contains %d units\n", immuneSystem[i].Id, immuneSystem[i].Units)
		// }
		// fmt.Println("Infection:")
		// for i := 0; i < len(infection); i++ {
		// 	fmt.Printf("Group %d contains %d units\n", infection[i].Id, infection[i].Units)
		// }
		// fmt.Println("")

		// check if immuneSystem won
		i := count(immuneSystem)
		j := count(infection)
		if j == 0 {
			return i
		}
		if i == 0 {
			return 0
		}

		// target selection
		resetTarget(immuneSystem)
		resetTarget(infection)

		targetSelection("Immune System", immuneSystem, infection)
		targetSelection("Infection", infection, immuneSystem)
		// fmt.Println("")

		// attack
		if attack(immuneSystem, infection) == 0 {
			// stalemate
			return -1
		}
		// fmt.Println("")
	}
}

func parse(line string) group {
	var g group
	re := regexp.MustCompile(`(?P<Units>\d+) units each with (?P<Hit>\d+) hit points (?P<ParseLater>\(.*?\) )?with an attack that does (?P<Attack>\d+) (?P<AttackType>[a-z]+) damage at initiative (?P<Initiative>\d+)`)
	utils.MustParseToStruct(re, line, &g)

	if g.ParseLater != "" {
		// drop '(' and ') '
		t := g.ParseLater[1 : len(g.ParseLater)-2]

		parseLater := strings.Split(t, "; ")
		for _, p := range parseLater {
			if strings.HasPrefix(p, "weak to") {
				g.WeakTo = strings.Split(p[len("weak to "):], ", ")
			} else if strings.HasPrefix(p, "immune to") {
				g.ImmuneTo = strings.Split(p[len("immune to "):], ", ")
			} else {
				panic(fmt.Errorf("failed to parse: %s", p))
			}
		}
	}

	g.Targetting = -1
	g.OriginalUnits = g.Units
	g.OriginalAttack = g.Attack
	return g
}

func count(army []group) int {
	r := 0
	for i := 0; i < len(army); i++ {
		r += army[i].Units
	}
	return r
}

func resetTarget(army []group) {
	for i := 0; i < len(army); i++ {
		army[i].Selected = false
		army[i].Targetting = -1
	}

	sort.Slice(army, func(i, j int) bool {
		e1 := effectivePower(army[i])
		e2 := effectivePower(army[j])
		if e1 == e2 {
			return army[i].Initiative > army[j].Initiative
		}
		return e1 > e2
	})
}

func targetSelection(armyName string, army1 []group, army2 []group) {
	for i := 0; i < len(army1); i++ {
		selected := -1
		damage := 0
		if army1[i].Units == 0 {
			continue
		}
		for j := 0; j < len(army2); j++ {
			if army2[j].Selected || army2[j].Units == 0 {
				continue
			}
			d := calcDamage(army1[i], army2[j])
			if d > damage {
				damage = d
				selected = j
			}
		}
		if selected != -1 {
			army1[i].Targetting = selected
			army2[selected].Selected = true
			// fmt.Printf("%s group %d picked %d and will deal %d damage\n", armyName, army1[i].Id, army2[selected].Id, damage)
		}
	}
}

func calcDamage(group1 group, group2 group) int {
	v := effectivePower(group1)
	if utils.Contains(group2.ImmuneTo, group1.AttackType) {
		return 0
	}
	if utils.Contains(group2.WeakTo, group1.AttackType) {
		return v + v
	}
	return v
}

func attack(immuneSystem []group, infection []group) int {
	totalKilled := 0
outer:
	for initiative := len(immuneSystem) + len(infection); initiative > 0; initiative-- {
		// find army which has the right initiative
		for i := 0; i < len(immuneSystem); i++ {
			a := immuneSystem[i]
			if a.Initiative == initiative {
				if a.Targetting != -1 {
					totalKilled += doAttack("Immune System", a, &infection[a.Targetting])
				}
				continue outer
			}
		}
		for i := 0; i < len(infection); i++ {
			a := infection[i]
			if a.Initiative == initiative {
				if a.Targetting != -1 {
					totalKilled += doAttack("Infection", a, &immuneSystem[a.Targetting])
				}
				continue outer
			}
		}
		panic("meh")
	}

	return totalKilled
}

func doAttack(armyName string, group1 group, group2 *group) int {
	d := calcDamage(group1, *group2)
	kills := 0
	if d > 0 {
		kills = d / group2.Hit
		if kills > group2.Units {
			kills = group2.Units
		}
	}
	// fmt.Printf("%s group %d attacks defending group %d, killing %d units\n", armyName, group1.Id, group2.Id, kills)
	group2.Units = group2.Units - kills
	return kills
}

func effectivePower(g group) int {
	return g.Units * g.Attack
}
