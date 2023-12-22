package main

import (
	"strings"

	"github.com/thoas/go-funk"
)

type stringNumber struct {
	str   string
	value string
}

var dict []stringNumber = []stringNumber{
	{"one", "1"},
	{"two", "2"},
	{"three", "3"},
	{"four", "4"},
	{"five", "5"},
	{"six", "6"},
	{"seven", "7"},
	{"eight", "8"},
	{"nine", "9"},
}

func getTotalCalibration(url string) (total int64) {
	lines := readLinesFromFile(url)

	funk.ForEach(lines, func(l string) {
		total += getCalibration(getNumbers(l))
	})

	return total
}

func getTotalCalibration_2(url string) (total int64) {
	lines := readLinesFromFile(url)

	funk.ForEach(lines, func(l string) {
		newL := replaceStringPerNumber(l, 0, "")
		total += getCalibration(getNumbers(newL))
	})

	return total
}

func getCalibration(ints []int64) int64 {
	first := ints[0]
	last := ints[len(ints)-1]

	return first*10 + last
}

func getNumbers(puzz string) (results []int64) {
	slip := strings.Split(puzz, "")

	funk.ForEach(slip, func(i string) {
		if v, b := IsDigit(i); b {
			results = append(results, v)
		}
	})
	return results
}

func replaceStringPerNumber(l string, pointer int, acc string) string {
	if len(l) <= pointer {
		return acc
	}

	found := funk.Find(dict, func(s stringNumber) bool {
		return strings.Index(l[pointer:], s.str) == 0
	})

	if found != nil {
		foundStringNumber := found.(stringNumber)
		return replaceStringPerNumber(l, pointer+1, acc+foundStringNumber.value)
	}

	return replaceStringPerNumber(l, pointer+1, acc+l[pointer:pointer+1])
}
