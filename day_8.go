package main

import (
	"fmt"
	"regexp"
	"strings"
)

type RLInstruction struct {
	l string
	r string
}

type RLMap map[string]RLInstruction

type Day8 struct {
	steps []string
	rlMap RLMap
}

func solve_8_part1(url string) int {
	puzzle := readLinesFromFile(url)
	day_8 := getDay8(puzzle)
	steps := countSteps(day_8)

	return steps
}

func solve_8_part2(url string) int {
	puzzle := readLinesFromFile(url)
	day_8 := getDay8(puzzle)
	steps := countStepsGhosts(day_8)

	return steps
}

func getDay8(lines []string) (d Day8) {
	d.steps = strings.Split(lines[0], "")
	d.rlMap = make(RLMap)

	for i := 2; i < len(lines); i++ {
		inst := regexp.MustCompile("([A-Z1-9]{3}) = \\(([A-Z1-9]{3}), ([A-Z1-9]{3})\\)").FindAllStringSubmatch(lines[i], -1)
		d.rlMap[inst[0][1]] = RLInstruction{inst[0][2], inst[0][3]}
	}

	return d
}

func countSteps(d Day8) int {
	visited := make([]string, 0)
	return findZGoal(d, "AAA", &visited)
}

func countStepsGhosts(d Day8) (steps int) {
	As := make([]string, 0)
	for k := range d.rlMap {
		if strings.Contains(k, "A") {
			As = append(As, k)
		}
	}

	var results = make(chan int, len(As))
	for i := 0; i < len(As); i++ {
		go func(i int) {
			fmt.Printf("worker %d: start: %s\n", i, As[i])

			visited := make([]string, 0)
			res := findZGoal(d, As[i], &visited)

			fmt.Printf("worker %d: res: %d\n", i, res)
			results <- res
		}(i)
	}

	allResults := make([]int, len(As))
	for i := 0; i < len(As); i++ {
		allResults[i] = <-results
	}

	if len(allResults) == 1 {
		return allResults[0]
	}

	return LCM(allResults[0], allResults[1], allResults[2:]...)
}

func findZGoal(d Day8, initial string, visited *[]string) int {
	steps := len(*visited)
	if strings.Contains(initial, "Z") {
		fmt.Printf("goal %s\n", initial)

		return steps
	}

	(*visited) = append((*visited), initial)
	idx := steps % len(d.steps)

	next := Ternary[string](d.rlMap[initial].l, d.rlMap[initial].r, d.steps[idx] == "L")

	if initial == next && d.rlMap[initial].l == d.rlMap[initial].r {
		return -1
	}
	return findZGoal(d, next, visited)
}
