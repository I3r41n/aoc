package main

import (
	"strings"

	"github.com/thoas/go-funk"
)

func solve_13_1(url string) int {
	lines := readLinesFromFile(url)

	return funk.SumInt(calculateMirrorPatterns(&lines, false))
}

func solve_13_2(url string) int {
	lines := readLinesFromFile(url)

	return funk.SumInt(calculateMirrorPatterns(&lines, true))
}

func calculateMirrorPatterns(lines *[]string, smudged bool) (res []int) {
	matrix := make([][]string, 0)

	for _, l := range *lines {
		if len(l) == 0 {
			if ok, found := calculateMirrorPattern(&matrix, -1); !ok {
				panic("Not found mirror pattern")
			} else if smudged {
				res = append(res, fixMirror(matrix, found))
			} else {
				res = append(res, found)
			}

			matrix = make([][]string, 0)
			continue
		}

		matrix = append(matrix, strings.Split(l, ""))
	}

	return res
}

func fixMirror(matrix [][]string, avoid int) int {
	if ok, found := calculateMirrorPattern(&matrix, avoid); ok {
		return found
	}

	invert := map[string]string{
		".": "#",
		"#": ".",
	}
	var smudged [][]string

	for y, l := range matrix {
		for x, c := range l {
			smudged = make([][]string, len(matrix))
			for i := range matrix {
				smudged[i] = make([]string, len(matrix[i]))
				copy(smudged[i], matrix[i])
			}

			smudged[y][x] = invert[c]

			if ok, found := calculateMirrorPattern(&smudged, avoid); ok {
				return found
			}
		}

	}
	panic("Not found smudge")
}

func calculateMirrorPattern(matrix *[][]string, avoid int) (bool, int) {

	if res, ok := calculateVerticalMirrorPattern(matrix, avoid); ok {
		return true, res
	}
	tmatrix := transpose[string](matrix)
	// printMatrix(tmatrix)

	if res, ok := calculateVerticalMirrorPattern(tmatrix, avoid/100); ok {
		return true, res * 100
	}

	return false, -1
}

func calculateVerticalMirrorPattern(matrix *[][]string, avoid int) (int, bool) {
	pattern := (*matrix)
	lenght := len(pattern[0])

	for i := 0; i < lenght; i++ {

		if ok, columnsToLeft := foundReflection(lenght-1, i, pattern); ok && avoid != columnsToLeft {
			return columnsToLeft, true
		}

		if ok, columnsToLeft := foundReflection(lenght-1-i, 0, pattern); ok && avoid != columnsToLeft {
			return columnsToLeft, true
		}
	}

	return 0, false
}

func foundReflection(max int, x int, pattern [][]string) (bool, int) {
	for j := max; j > x; j-- {
		allMatch := All(pattern, func(s []string) bool {
			return s[j] == s[x]
		})
		if !allMatch {
			break
		}
		if j == x+1 {
			return true, x + 1
		}
		x++
	}
	return false, -1
}
