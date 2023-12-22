package main

import (
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

var fOp = map[string]func(int, int) bool{
	">": func(a, value int) bool { return a > value },
	"<": func(a, value int) bool { return a < value },
}

type xmas struct {
	x, m, a, s int
}

type Getter func(v xmas) int

var xmasAccess = map[string]Getter{
	"x": func(v xmas) int { return v.x },
	"m": func(v xmas) int { return v.m },
	"a": func(v xmas) int { return v.a },
	"s": func(v xmas) int { return v.s },
}

func getProperty(v xmas, property string) (int, bool) {
	f, ok := xmasAccess[property]
	if ok {
		return f(v), true
	}

	return 0, false
}

func toXMAS(str string) xmas {
	parts := strings.Split(str, ",")
	x, _ := strconv.Atoi(parts[0][strings.Index(parts[0], "=")+1:])
	m, _ := strconv.Atoi(parts[1][strings.Index(parts[1], "=")+1:])
	a, _ := strconv.Atoi(parts[2][strings.Index(parts[2], "=")+1:])
	s, _ := strconv.Atoi(parts[3][strings.Index(parts[3], "=")+1 : strings.Index(parts[3], "}")])

	return xmas{x, m, a, s}
}

type Rule struct {
	letter string
	res    string
	op     string
	value  int
	f      func(xmas) (string, bool)
}

type workflow map[string][]Rule

var rangeIndex = map[string]int{"x": 0, "m": 1, "a": 2, "s": 3}

func solve_1(url string) int {
	lines := readLinesFromFile(url)
	wf, idx := getWorkflows(&lines)

	accepted := make([]xmas, 0)
	rulesString := (lines)[idx:]
	getAcceptedXmas(&rulesString, &accepted, wf)

	sum := 0
	for _, a := range accepted {
		sum += a.x + a.m + a.a + a.s
	}
	return sum
}

func solve_2(url string) int {
	lines := readLinesFromFile(url)
	wf, _ := getWorkflows(&lines)

	return getCombinations(wf, "in", []coord{
		{1, 4000},
		{1, 4000},
		{1, 4000},
		{1, 4000},
	})
}

func getCombinations(wf workflow, id string, ranges []coord) int {
	if id == "A" {
		return funk.Reduce(ranges, func(acc int, cur coord) int {
			acc *= cur.y - cur.x + 1
			return acc
		}, 1).(int)
	}

	if id == "R" {
		return 0
	}

	res := 0
	current := wf[id]

	for _, r := range current {

		newRanges := make([]coord, len(ranges))
		copy(newRanges, ranges)

		switch {
		case r.letter == "":
			res += getCombinations(wf, r.res, ranges)
		case r.op == "<":
			newRanges[rangeIndex[r.letter]].y = r.value - 1
			ranges[rangeIndex[r.letter]].x = r.value
			res += getCombinations(wf, r.res, newRanges)

		case r.op == ">":
			newRanges[rangeIndex[r.letter]].x = r.value + 1
			ranges[rangeIndex[r.letter]].y = r.value
			res += getCombinations(wf, r.res, newRanges)
		}
	}

	return res
}

func getWorkflows(lines *[]string) (workflow, int) {
	wf := make(workflow, 0)

	for idx, l := range *lines {
		if l == "" {
			return wf, idx + 1
		}

		i := strings.Index(l, "{")
		name := l[0:i]
		rulesDescrition := strings.Split(l[strings.Index(l, "{")+1:strings.Index(l, "}")], ",")
		fRules := make([]Rule, 0)

		for _, rd := range rulesDescrition {
			dotsIdx := strings.Index(rd, ":")
			if dotsIdx == -1 {
				fRules = append(fRules, Rule{"", rd, "", 0, createFRule("", "", rd, 0)})
			} else {
				i, _ := strconv.Atoi(rd[2:dotsIdx])

				fRules = append(fRules, Rule{rd[0:1], rd[dotsIdx+1:], rd[1:2], i, createFRule(rd[0:1], rd[1:2], rd[dotsIdx+1:], i)})
			}
		}

		wf[name] = fRules
	}

	panic("no workflow?!")
}

func getAcceptedXmas(lines *[]string, xmas *[]xmas, wf workflow) {
	for _, l := range *lines {
		var res string = "in"
		piece := toXMAS(l)
		for !funk.Contains([]string{"A", "R"}, res) {
			for _, r := range wf[res] {
				if r, ok := r.f(piece); ok {
					res = r
					break
				}
			}
		}

		if res == "A" {
			(*xmas) = append((*xmas), piece)
		}
	}
}

func createFRule(varF, op, result string, value int) func(xmas) (string, bool) {
	if varF == "" {
		return func(p xmas) (string, bool) {
			return result, true
		}
	}

	return func(p xmas) (string, bool) {
		x, ok := getProperty(p, varF)
		if ok && fOp[op](x, value) {
			return result, true
		}
		return "", false
	}
}
