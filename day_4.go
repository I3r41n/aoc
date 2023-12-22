package main

import (
	"regexp"
	"strings"

	"github.com/thoas/go-funk"
)

func solve_4_part1(url string) int {
	cards := readFile(url)
	scores := getScratchCardsPrized(cards, score)

	return funk.SumInt(scores)
}

func solve_4_part2(url string) int {
	cards := readFile(url)
	scores := getScratchCardsPrized(cards, prize)

	return funk.SumInt(getPileScoreCards(scores))
}

func getScratchCardsPrized(cards string, fn func([]int) int) (scores []int) {
	lines := strings.Split(cards, "\n")
	scores = make([]int, len(lines))
	for i, l := range lines {
		m := l[strings.Index(cards, ": ")+1:]
		s := strings.Split(m, " | ")

		winners := convertStringArrayIntArray(regexp.MustCompile("(\\d+)").FindAllString(s[0], -1))
		got := convertStringArrayIntArray(regexp.MustCompile("(\\d+)").FindAllString(s[1], -1))
		prized := funk.JoinInt(winners, got, funk.InnerJoinInt)

		scores[i] = fn(prized)
	}

	return scores
}

func getPileScoreCards(scores []int) (won []int) {
	won = make([]int, len(scores))

	for i := range won {
		won[i] = 1
	}

	for i, v := range scores {
		for j := i + 1; j <= i+v; j++ {
			won[j] = won[j] + won[i]
		}
	}
	return won
}

func prize(prized []int) int {
	return len(prized)
}

func score(prized []int) int {
	if 0 == len(prized) {
		return 0
	}
	var score = 1
	for j := 1; j < len(prized); j++ {
		score *= 2
	}
	return score
}
