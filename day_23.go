package main

import (
	"strings"

	"github.com/thoas/go-funk"
)

func solve_23_1(url string) int {
	lines := readLinesFromFile(url)
	matrix := getMatrix(&lines)

	maxY := len(*matrix)

	start := coord{strings.Index(strings.Join((*matrix)[0], ""), "."), 0}
	goal := coord{strings.Index(strings.Join((*matrix)[maxY-1], ""), "."), maxY - 1}

	return dfs(paths(matrix), start, goal, 0, make(map[coord]bool, 0))
}

func solve_23_2(url string) int {
	lines := readLinesFromFile(url)
	matrix := getMatrix(&lines)

	maxY := len(*matrix)

	start := coord{strings.Index(strings.Join((*matrix)[0], ""), "."), 0}
	goal := coord{strings.Index(strings.Join((*matrix)[maxY-1], ""), "."), maxY - 1}

	return dfs(getEdges(matrix), start, goal, 0, make(map[coord]bool, 0))
}

type edge struct {
	s      coord
	e      coord
	weight int
}

func getEdges(matrix *[][]string) *map[coord][]edge {
	edges := make(map[coord][]edge, 0)
	intersections := make(map[coord]bool)

	maxY := len(*matrix)
	maxX := len((*matrix)[0])

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if (*matrix)[y][x] != "#" {
				c := coord{x, y}
				n := 0

				for _, dir := range []Coordinate{North, South, East, West} {
					newCoord := c.add(dir)

					if !newCoord.isValid(maxX, maxY) || (*matrix)[newCoord.y][newCoord.x] == "#" {
						continue
					}

					n += 1
				}

				if n == 2 {
					continue
				}
				intersections[c] = true
				edges[c] = make([]edge, 0)
			}
		}
	}

	for k := range edges {
		edges[k] = *getPathsFromEdge(k, intersections, matrix)
	}
	return &edges
}

func getPathsFromEdge(start coord, edges map[coord]bool, matrix *[][]string) *[]edge {
	maxY := len(*matrix)
	maxX := len((*matrix)[0])

	paths := make([]edge, 0)
	for _, dir := range []Coordinate{North, South, East, West} {
		newCoord := start.add(dir)

		if newCoord.isValid(maxX, maxY) && (*matrix)[newCoord.y][newCoord.x] != "#" {
			paths = append(paths, *getPathFromEdge(start, newCoord, dir, 1, edges, matrix))
		}
	}

	return &paths
}

func getPathFromEdge(s, current coord, lastDir Coordinate, weight int, edges map[coord]bool, matrix *[][]string) *edge {
	maxY := len(*matrix)
	maxX := len((*matrix)[0])

	for _, dir := range []Coordinate{lastDir, turn[lastDir]["R"], turn[lastDir]["L"]} {
		newCoord := current.add(dir)
		if newCoord.isValid(maxX, maxY) && (*matrix)[newCoord.y][newCoord.x] != "#" {
			if _, found := edges[newCoord]; found {
				return &edge{s, newCoord, weight + 1}
			} else {
				return getPathFromEdge(s, newCoord, dir, weight+1, edges, matrix)
			}
		}
	}

	panic("No path")
}

func paths(matrix *[][]string) *map[coord][]edge {
	maxY := len(*matrix)
	maxX := len((*matrix)[0])

	edges := make(map[coord][]edge, 0)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if (*matrix)[y][x] != "#" {
				edges[coord{x, y}] = *getHikePaths(matrix, coord{x, y})
			}
		}
	}

	return &edges
}

func getHikePaths(matrix *[][]string, c coord) *[]edge {
	maxY := len(*matrix)
	maxX := len((*matrix)[0])

	n := make([]edge, 0)
	dirs := funk.Filter([]Coordinate{North, South, East, West}, func(cd Coordinate) bool {
		v := (*matrix)[c.y][c.x]
		switch {
		case v == ">":
			return cd == East
		case v == "^":
			return cd == North
		case v == "v":
			return cd == South
		case v == "<":
			return cd == West
		}
		return true
	}).([]Coordinate)

	for _, dir := range dirs {
		newCoord := c.add(dir)

		if !newCoord.isValid(maxX, maxY) || (*matrix)[newCoord.y][newCoord.x] == "#" {
			continue
		}

		n = append(n, edge{c, newCoord, 1})
	}
	return &n
}

func dfs(edges *map[coord][]edge, src, goal coord, steps int, seen map[coord]bool) int {
	if src == goal {
		return steps
	}

	maxSteps := 0

	for _, edge := range (*edges)[src] {
		if _, ok := seen[edge.e]; !ok {
			seen[edge.e] = true
			maxSteps = max(maxSteps, dfs(edges, edge.e, goal, edge.weight+steps, seen))
			delete(seen, edge.e)
		}
	}

	return maxSteps
}
