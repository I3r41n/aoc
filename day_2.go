package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

var REGEX_GROUP = regexp.MustCompile("Game (\\d+): ")
var REGEX_SET = regexp.MustCompile("(\\d+ \\w+)")

const (
	RED       = "red"
	BLUE      = "blue"
	GREEN     = "green"
	MAX_RED   = 12
	MAX_BLUE  = 14
	MAX_GREEN = 13
)

func solve(url string) int {
	return solve_day2(url, possibleGames)
}

func solve2(url string) int {
	return solve_day2(url, minimumValues)
}

func solve_day2(url string, f func(games map[int]map[string]int) (values []int)) int {
	lines := readLinesFromFile(url)
	values := f(getGames(lines))
	return funk.SumInt(values)
}

func minimumValues(games map[int]map[string]int) (values []int) {
	for _, value := range games {
		values = append(values, value[RED]*value[BLUE]*value[GREEN])
	}

	return values
}

func possibleGames(games map[int]map[string]int) (ids []int) {
	for key, value := range games {
		if value[RED] <= MAX_RED && value[BLUE] <= MAX_BLUE && value[GREEN] <= MAX_GREEN {
			ids = append(ids, key)
		}

	}
	return ids
}

func getGames(lines []string) (games map[int]map[string]int) {
	games = make(map[int]map[string]int)

	funk.ForEach(lines, func(l string) {

		id, _ := strconv.Atoi(REGEX_GROUP.FindStringSubmatch(l)[1]) // game id
		game := funk.Flatten(REGEX_SET.FindAllStringSubmatch(l, -1))

		games[id] = make(map[string]int)

		funk.ForEach(game, func(s string) {
			number_color := strings.Split(s, " ")
			number, _ := strconv.Atoi(number_color[0])
			i, _ := games[id][number_color[1]]

			games[id][number_color[1]] = number

			if number < i {
				games[id][number_color[1]] = i
			}
		})
	})

	return games
}
