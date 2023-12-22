package main

import (
	"strings"
)

type Beam struct {
	c   coord
	dir Coordinate
}

var moveDirection = map[Coordinate]map[string]Coordinate{
	East:  {"\\": South, "/": North},
	West:  {"\\": North, "/": South},
	South: {"\\": East, "/": West},
	North: {"\\": West, "/": East},
}

func solve_day_16_1(url string) int {
	lines := readLinesFromFile(url)
	matrix := getContraption(&lines)

	return getEnergized(matrix, Beam{coord{x: 0, y: 0}, East})
}

func solve_day_16_2(url string) int {
	lines := readLinesFromFile(url)
	matrix := getContraption(&lines)

	return getMaxEnergized(matrix)
}

func getMaxEnergized(matrix *[][]string) int {
	maxY := len(*matrix)
	maxX := len((*matrix)[0])

	max := 0

	beams := []Beam{}

	for i := 0; i < maxY; i++ {
		beams = append(beams, Beam{coord{maxX - 1, i}, East})
		beams = append(beams, Beam{coord{0, i}, West})
	}

	for i := 0; i < maxX; i++ {
		beams = append(beams, Beam{coord{i, 0}, South})
		beams = append(beams, Beam{coord{i, maxY}, North})
	}

	for _, b := range beams {
		if e := getEnergized(matrix, b); e > max {
			max = e
		}
	}

	return max
}

func getEnergized(matrix *[][]string, beam Beam) int {
	energized := make(map[coord]map[Coordinate]bool, 0)

	runBeam(beam, energized, matrix)

	return len(energized)
}

func runBeam(beam Beam, energized map[coord]map[Coordinate]bool, matrix *[][]string) {
	maxY := len(*matrix)
	maxX := len((*matrix)[0])

	if beam.c.x < 0 || beam.c.y < 0 || beam.c.y >= maxY || beam.c.x >= maxX {
		return
	}

	if val, ok := energized[beam.c]; ok {
		if val[beam.dir] {
			return
		}
	} else {
		energized[beam.c] = map[Coordinate]bool{
			East:  false,
			West:  false,
			South: false,
			North: false,
		}
	}

	energized[beam.c][beam.dir] = true
	c := (*matrix)[beam.c.y][beam.c.x]

	switch {
	case c == "/" || c == "\\":
		beam.dir = moveDirection[beam.dir][c]

	case c == "-" && (beam.dir == North || beam.dir == South):
		runBeam(Beam{beam.c, East}, energized, matrix)
		runBeam(Beam{beam.c, West}, energized, matrix)

		return
	case c == "|" && (beam.dir == East || beam.dir == West):
		runBeam(Beam{beam.c, South}, energized, matrix)
		runBeam(Beam{beam.c, North}, energized, matrix)

		return
	}
	beam.c = beam.c.add(beam.dir)
	runBeam(beam, energized, matrix)
}

func getContraption(lines *[]string) *[][]string {
	maxY := len(*lines)
	matrix := make([][]string, maxY)

	for y, l := range *lines {
		matrix[y] = strings.Split(l, "")
	}

	return &matrix
}
