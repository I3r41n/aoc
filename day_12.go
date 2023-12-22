package main

import (
	"strings"

	"github.com/thoas/go-funk"
)

func solve_12_1(url string) int {
	lines := readLinesFromFile(url)

	return funk.SumInt(calculateArrangements(&lines, false))
}

func solve_12_2(url string) int {
	lines := readLinesFromFile(url)

	return funk.SumInt(calculateArrangements(&lines, true))
}

func calculateArrangements(lines *[]string, unfold bool) []int {
	arrangements := make([]int, len(*lines))
	for i, l := range *lines {

		l, g := prepareUnfoldLine(l, unfold)
		arrangements[i] = calculateLineArrangements(l, g)
	}

	return arrangements
}

func prepareUnfoldLine(l string, unfold bool) (string, []int) {
	s := strings.Split(l, " ")

	if !unfold {
		return s[0], convertStringArrayIntArray(strings.Split(s[1], ","))
	}

	unfolded := strings.Repeat(s[0]+"?", 5)
	groups := strings.Repeat(s[1]+",", 5)

	return unfolded[0 : len(unfolded)-1], convertStringArrayIntArray(strings.Split(groups[0:len(groups)-1], ","))

}

func calculateLineArrangements(l string, groups []int) int {
	var cache [][]int
	for i := 0; i < len(l); i++ {
		cache = append(cache, make([]int, len(groups)+1))
		for j := 0; j <= len(groups); j++ {
			cache[i][j] = -1
		}
	}
	return aux(0, 0, l, groups, cache)
}

func aux(idx, g int, l string, groups []int, cache [][]int) int {
	if idx >= len(l) {
		return Ternary[int](1, 0, g >= len(groups))
	}

	if cache[idx][g] > -1 {
		return cache[idx][g]
	}

	res := 0
	c := l[idx]
	if c == '.' {
		res = aux(idx+1, g, l, groups, cache)
	} else {
		if c == '?' {
			res += aux(idx+1, g, l, groups, cache)
		}

		if g < len(groups) && idx+groups[g] <= len(l) {
			afterGroupIdx := idx + groups[g]
			possiblegroup := l[idx:afterGroupIdx]
			if !strings.Contains(possiblegroup, ".") && (afterGroupIdx == len(l) || l[afterGroupIdx] != '#') {
				res += aux(afterGroupIdx+1, g+1, l, groups, cache)
			}
		}
	}
	cache[idx][g] = res
	return res
}
