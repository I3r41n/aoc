package main

import (
	"strings"

	"github.com/thoas/go-funk"
)

func solve11_1(url string) int {
	return getExpandingGalaxyManSum(url, 1)
}

func solve11_2(url string) int {
	return getExpandingGalaxyManSum(url, 1000000-1)
}

func getExpandingGalaxyManSum(url string, expansion int) int {
	lines := readLinesFromFile(url)
	galaxies, emptyY, emptyX := buildGalaxy(&lines, expansion)
	distances := calculateManGalaxies(galaxies, emptyY, emptyX)

	return funk.SumInt(distances)
}

func calculateManGalaxies(g *[]coord, y *map[int]int, x *map[int]int) []int {
	dist := make([]int, 0)

	for i, f := range *g {
		for _, d := range (*g)[i+1:] {
			dist = append(dist, manhattan(f.x+(*x)[f.x], d.x+(*x)[d.x], f.y+(*y)[f.y], d.y+(*y)[d.y]))
		}
	}

	return dist
}

func buildGalaxy(lines *[]string, m int) (*[]coord, *map[int]int, *map[int]int) {
	emptyY := make(map[int]int, len(*lines))
	emptyX := make(map[int]int, len((*lines)[0]))
	galaxies := make([]coord, 0)
	auxEmptyX := make([]bool, len((*lines)[0]))

	for i := range (*lines)[0] {
		auxEmptyX[i] = true
	}

	for i := range *lines {
		emptyY[i] = (i + 1) * m
	}

	for y, l := range *lines {
		doesntHasGalaxy := strings.Index(l, "#") == -1

		if doesntHasGalaxy {
			continue
		}

		for idx := y; idx < len(*lines); idx++ {
			emptyY[idx] = emptyY[idx] - m
		}

		for x, c := range l {
			if c == '#' {
				auxEmptyX[x] = false
				galaxies = append(galaxies, coord{x: x, y: y})
			}
		}
	}

	aux := 0
	for i, b := range auxEmptyX {
		if b {
			aux += m
		}
		emptyX[i] = aux
	}

	return &galaxies, &emptyY, &emptyX
}
