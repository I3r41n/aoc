package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func solve_day3(url string) int {
	puzzle := readLinesFromFile(url)
	engine := getEngine(puzzle)
	parts := getParts(engine)

	filteredParts := funk.Filter(parts, func(p pos[string]) bool {
		return hasNeighbors(p, engine)
	})

	sum := funk.Reduce(filteredParts, func(acc int, cur pos[string]) int {
		return acc + convertStringToInt(cur.value)
	}, 0)

	return sum.(int)
}

func solve2_day3(url string) int {
	puzzle := readLinesFromFile(url)
	engine := getEngine(puzzle)
	parts := getParts(engine)

	filteredParts := funk.Filter(parts, func(p pos[string]) bool {
		return hasStarNeighbors(p, engine)
	})

	sum := funk.Reduce(getGears(engine, filteredParts.([]pos[string])), func(acc int, cur pos[string]) int {
		return acc + convertStringToInt(cur.value)
	}, 0)

	return sum.(int)
}

func getGears(engine [][]string, parts []pos[string]) (gears []pos[string]) {
	gears = []pos[string]{}

	for x, value := range engine {
		for y, c := range value {
			if c == "*" {
				gears = append(gears, pos[string]{getRatio(pos[string]{c, y, x}, parts), y, x})
			}
		}
	}

	return gears
}

func getRatio(gear pos[string], parts []pos[string]) string {
	x := gear.x
	y := gear.y

	var candidates = funk.Filter(parts, func(p pos[string]) bool {
		px := p.x
		py := p.y
		psize := len(p.value)

		return funk.Contains(Range(px-1, px+psize+1, 1), x) && funk.Contains([]int{y - 1, y, y + 1}, py)
	})
	c := candidates.([]pos[string])

	if len(c) == 2 {
		return strconv.Itoa(convertStringToInt(c[0].value) * convertStringToInt(c[1].value))
	}

	if len(c) > 2 {
		fmt.Println(c)
		panic("too many candidates")
	}

	return "0"
}

func getParts(engine [][]string) (parts []pos[string]) {
	parts = []pos[string]{}

	for x, value := range engine {
		var number = ""
		var initial = -1
		for y, c := range value {
			isNumber := funk.Contains(numbers, c)
			if isNumber {
				number = number + c
				if initial == -1 {
					initial = y
				}
			}
			if (!isNumber || len(engine[0]) == y+1) && initial != -1 {
				parts = append(parts, pos[string]{number, initial, x})
				number = ""
				initial = -1
			}
		}
	}

	return parts
}

func hasNeighbors(part pos[string], engine [][]string) bool {
	x := part.x
	y := part.y
	size := len(part.value)
	max := len(engine[0])

	for i := -1; i <= size; i++ {
		xi := x + i
		if y-1 >= 0 && xi >= 0 && xi < max && !funk.Contains(numbers, engine[y-1][xi]) && engine[y-1][xi] != "." {
			return true
		}
		if xi >= 0 && xi < max && !funk.Contains(numbers, engine[y][xi]) && engine[y][xi] != "." {
			return true
		}
		if y+1 < max && xi >= 0 && xi < max && !funk.Contains(numbers, engine[y+1][xi]) && engine[y+1][xi] != "." {
			return true
		}
	}

	return false
}

func hasStarNeighbors(part pos[string], engine [][]string) bool {
	x := part.x
	y := part.y
	size := len(part.value)
	max := len(engine[0])

	for i := -1; i <= size; i++ {
		xi := x + i
		if y-1 >= 0 && xi >= 0 && xi < max && engine[y-1][xi] == "*" {
			return true
		}
		if xi >= 0 && xi < max && engine[y][xi] == "*" {
			return true
		}
		if y+1 < max && xi >= 0 && xi < max && engine[y+1][xi] == "*" {
			return true
		}
	}

	return false
}

func getEngine(lines []string) (engine [][]string) {
	var x = 0

	engine = make([][]string, len(lines))

	funk.ForEach(lines, func(l string) {
		engine[x] = make([]string, 0, len(l))
		engine[x] = strings.Split(l, "")
		x += 1
	})

	return engine
}
