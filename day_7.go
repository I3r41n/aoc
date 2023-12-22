package main

import (
	"sort"
	"strings"

	"github.com/thoas/go-funk"
)

var cardsValues = map[string]int{
	"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10, "9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2,
}

var handType = map[string]int{
	"Five of a kind": 7, "Four of a kind": 6, "Full house": 5, "Three of a kind": 4, "Two pair": 3, "One pair": 2, "High card": 1,
}

type Hand struct {
	cards []int
	bet   int
	tipe  string
}

type Game []Hand

func (g Game) Len() int      { return len(g) }
func (g Game) Swap(i, j int) { g[i], g[j] = g[j], g[i] }
func (g Game) Less(i, j int) bool {
	t1 := handType[g[i].tipe]
	t2 := handType[g[j].tipe]
	if t1 != t2 {
		return t1 < t2
	}
	for idx, v := range g[i].cards {
		if v != g[j].cards[idx] {
			return v < g[j].cards[idx]
		}
	}
	return false
}

func solve_7_part1(url string) int {
	puzzle := readLinesFromFile(url)
	game := getGame(puzzle, false)

	winnings := getWinnings(game)
	return winnings
}

func solve_7_part2(url string) int {
	puzzle := readLinesFromFile(url)
	game := getGame(puzzle, true)

	winnings := getWinnings(game)
	return winnings
}

func getWinnings(game Game) (w int) {
	sort.Sort(game)

	for idx, v := range game {
		w = w + v.bet*(idx+1)
	}

	return w
}

func getGame(lines []string, useJokers bool) (game Game) {
	game = make(Game, len(lines))
	for idx, l := range lines {
		h := strings.Split(l, " ")
		cards := convertCardsValues(strings.Split(h[0], ""), useJokers)
		bet := convertStringToInt(h[1])
		tipe := getCardType(h[0], useJokers)

		game[idx] = Hand{cards, bet, tipe}
	}
	return game
}

func convertCardsValues(s []string, useJokers bool) (ints []int) {
	funk.ForEach(s, func(n string) {
		var v = cardsValues[n]
		if useJokers && n == "J" {
			v = 0
		}
		ints = append(ints, v)
	})

	return ints
}

func getCardType(c string, useJokers bool) string {
	jokers := strings.Count(c, "J")
	frequency := make(map[rune]int)

	for _, char := range c {
		frequency[char] = frequency[char] + 1
	}

	entries := (funk.Values(frequency)).([]int)
	sort.Ints(entries)
	lenght := len(entries)
	switch {
	case lenght == 1 || (useJokers && lenght == 2 && jokers > 0):
		return "Five of a kind"
	case (lenght == 2 && entries[1] == 4) || (useJokers && lenght == 3 && jokers+entries[2] >= 4):
		return "Four of a kind"
	case (lenght == 2 && entries[1] == 3) || (useJokers && lenght == 3 && jokers == 1):
		return "Full house"
	case (lenght == 3 && entries[2] == 3) || (useJokers && lenght == 4 && jokers > 0):
		return "Three of a kind"
	case (lenght == 3 && entries[2] == 2):
		return "Two pair"
	case (lenght == 4) || (useJokers && lenght == 5 && jokers == 1):
		return "One pair"
	default:
		return "High card"
	}
}
