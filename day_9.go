package main

import (
	"regexp"

	"github.com/thoas/go-funk"
)

func solve_9_part1(url string) int {
	return calculateDay9(url, false)
}

func solve9_part2(url string) int {
	return calculateDay9(url, true)
}

func calculateDay9(url string, second bool) int {
	puzzle := readLinesFromFile(url)
	sequences := getDay9(puzzle)
	nextElements := getSequencesElement(sequences, second)

	return funk.SumInt(*nextElements)
}
func getSequencesElement(sequences [][]int, first bool) *[]int {
	nexts := make([]int, 0)

	for idx := 0; idx < len(sequences); idx++ {
		nexts = append(nexts, getElement(&(sequences[idx]), first))
	}

	return &nexts
}

func getDiffs(sequence *[]int) *[]int {
	diffs := make([]int, len(*sequence)-1)
	for i := 0; i < len(*sequence)-1; i++ {
		diffs[i] = (*sequence)[i+1] - (*sequence)[i]
	}

	return &diffs
}

func getElement(sequence *[]int, first bool) int {
	diffs := getDiffs(sequence)
	idx := Ternary[int](0, len(*sequence)-1, first)
	m := Ternary[int](-1, 1, first)
	if All(*diffs, func(i int) bool { return i == 0 }) {
		return (*sequence)[len(*sequence)-1]
	}

	return (*sequence)[idx] + (m * getElement(diffs, first))
}

func getDay9(lines []string) (sequences [][]int) {
	for _, l := range lines {
		sequences = append(sequences, convertStringArrayIntArray(regexp.MustCompile("(-?\\d+)").FindAllString(l, -1)))
	}

	return sequences
}
