package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/thoas/go-funk"
)

var categories = [...]string{
	"soil",
	"fertilizer",
	"water",
	"light",
	"temperature",
	"humidity",
	"location",
}

var categoriesReverse = [...]string{
	"location",
	"humidity",
	"temperature",
	"light",
	"water",
	"fertilizer",
	"soil",
}

type route struct {
	orig int
	dest int
	rang int
}

func solve_5_part1(url string) int {
	puzzle := readLinesFromFile(url)
	locations := getFinalLocations(puzzle, false)

	return funk.MinInt(locations)
}

func solve_5_part2(url string) int {
	puzzle := readLinesFromFile(url)
	locations := getFinalLocationsParallel(puzzle)

	return locations
}

func solve_5_part2_reverse(url string) int {
	puzzle := readLinesFromFile(url)
	locations := getFinalLocationsReverse(puzzle)

	return locations
}

func getFinalLocationsParallel(lines []string) int {
	seeds := getSeeds(lines[0])
	maps := getMaps(lines[2:], false)

	totalSeeds := len(seeds)

	numSeedRanges := totalSeeds / 2

	var results = make(chan int, numSeedRanges)
	for i := 0; i < numSeedRanges; i++ {
		go func(i int) {
			startSeed := seeds[i*2]
			endSeed := startSeed + seeds[i*2+1]

			fmt.Printf("worker %d: startSeed: %d, endSeed: %d total:%d\n", i, startSeed, endSeed, endSeed-startSeed)

			lowestLocation := math.MaxInt
			for seed := startSeed; seed < endSeed; seed++ {
				location := plantOneSeed(seed, maps, false)
				if location < lowestLocation {
					lowestLocation = location
				}
			}

			fmt.Printf("worker %d: lowestLocation: %d\n", i, lowestLocation)
			results <- lowestLocation
		}(i)
	}

	lowestLocation := math.MaxInt
	for i := 0; i < numSeedRanges; i++ {
		location := <-results
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	return lowestLocation
}

func getFinalLocationsReverse(lines []string) int {
	seeds := getSeeds(lines[0])
	maps := getMaps(lines[2:], true)

	for i := 0; i < math.MaxInt; i++ {
		seed := plantOneSeed(i, maps, true)
		if validSeed(seed, seeds) {
			return i
		}
	}

	panic("couldn't found")
}

func validSeed(seed int, seeds []int) bool {
	for i := 0; i < len(seeds); i += 2 {

		startSeed := seeds[i]
		endSeed := startSeed + seeds[i+1]
		if seed >= startSeed && seed < endSeed {
			return true
		}
	}
	return false
}

func getFinalLocations(lines []string, reverse bool) []int {
	seeds := getSeeds(lines[0])
	maps := getMaps(lines[2:], reverse)
	return plantSeeds(seeds, maps)
}

func plantSeeds(seeds []int, maps map[string][]route) []int {
	var tmp = funk.Reduce(seeds, func(acc []int, cur int) []int {
		return append(acc, plantOneSeed(cur, maps, false))
	}, []int{})

	return tmp.([]int)
}

func plantOneSeed(seed int, maps map[string][]route, reverseCat bool) int {
	almanac := Ternary[[7]string](categoriesReverse, categories, reverseCat)

	for _, cat := range almanac {
		seed = plantSeed(maps[cat], seed)
	}

	return seed
}

func plantSeed(category []route, seed int) int {
	for _, k := range category {
		if k.orig <= seed && k.orig+k.rang > seed {
			return k.dest + (seed - k.orig)
		}
	}
	return seed
}

func getMaps(rules []string, reverse bool) (maps map[string][]route) {
	maps = make(map[string][]route)
	var (
		idx = 0
		cat = categories[idx]
	)
	for _, l := range rules {
		switch {
		case "" == l:
			idx = idx + 1
			cat = categories[idx]
		case strings.Contains(l, cat):
			maps[cat] = []route{}
		default:
			orig := (map[bool]int{true: 0, false: 1})[reverse]
			dest := (map[bool]int{true: 1, false: 0})[reverse]

			routes := convertStringArrayIntArray(regexp.MustCompile("(\\d+)").FindAllString(l, -1))
			maps[cat] = append(maps[cat], route{routes[orig], routes[dest], routes[2]})
		}
	}

	return maps
}

func getSeeds(s string) []int {
	return convertStringArrayIntArray(regexp.MustCompile("(\\d+)").FindAllString(s, -1))
}
