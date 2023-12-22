package main

import (
	"regexp"
	"strings"

	"github.com/thoas/go-funk"
)

type race struct {
	time     int
	distance int
}

func solve_6_part1(url string) int {
	puzzle := readLinesFromFile(url)
	races := getRaces(puzzle)
	numberWaysToBeat := getNumberWaysToBeat(races)

	tmp := funk.Reduce(numberWaysToBeat, func(acc int, d int) int {
		return acc * d
	}, 1)

	return tmp.(int)
}

func solve_6_part2(url string) int {
	puzzle := readLinesFromFile(url)
	races := getRaceNoSpaces(puzzle)
	numberWaysToBeat := getWaysToBeat(races)

	return len(numberWaysToBeat)
}

func getRaces(lines []string) (races []race) {
	races = make([]race, 0)
	times := convertStringArrayIntArray(regexp.MustCompile("(\\d+)").FindAllString(lines[0], -1))
	distance := convertStringArrayIntArray(regexp.MustCompile("(\\d+)").FindAllString(lines[1], -1))

	for idx, t := range times {
		races = append(races, race{t, distance[idx]})
	}

	return races
}

func getRaceNoSpaces(lines []string) (r race) {

	times := convertStringToInt(strings.Join(regexp.MustCompile("(\\d+)").FindAllString(lines[0], -1), ""))
	distance := convertStringToInt(strings.Join(regexp.MustCompile("(\\d+)").FindAllString(lines[1], -1), ""))

	return race{times, distance}
}

func getNumberWaysToBeat(races []race) []int {
	waysToBeat := funk.Reduce(races, func(acc []int, cur race) []int {
		return append(acc, len(getWaysToBeat(cur)))
	}, make([]int, 0))

	return waysToBeat.([]int)
}

func getWaysToBeat(r race) []int {
	ways := make([]int, r.time+1)
	for h := 0; h < r.time+1; h++ {
		ways[h] = h * (r.time - h)
	}

	return funk.FilterInt(ways, func(s int) bool {
		return s > r.distance
	})
}
