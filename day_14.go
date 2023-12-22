package main

import "strings"

func solve_day_14_1(url string) int {
	lines := readLinesFromFile(url)
	matrix := rollRocksNorth(&lines)

	return getRocksLoad(matrix)
}

func solve_day_14_2(url string) int {
	lines := readLinesFromFile(url)
	matrix := rollRocksForever(&lines)

	return getRocksLoad(matrix)
}

func getRocksLoad(matrix *[][]string) (load int) {
	weight := len(*matrix)

	for _, row := range *matrix {
		load += weight * strings.Count(strings.Join(row, ""), "O")
		weight--
	}

	return load
}

func rollRocksNorth(lines *[]string) *[][]string {
	maxY := len(*lines)
	maxX := len((*lines)[0])

	matrix := make([][]string, maxY)

	for y, l := range *lines {
		matrix[y] = strings.Split(l, "")
	}

	for y := 1; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			c := matrix[y][x]
			if c != "O" {
				continue
			}

			moveRockNorth(x, y, &matrix)
		}
	}

	return &matrix
}

func moveRockNorth(x, y int, matrix *[][]string) {
	if y == 0 || (*matrix)[y-1][x] != "." {
		return
	}

	(*matrix)[y][x] = "."
	(*matrix)[y-1][x] = "O"

	moveRockNorth(x, y-1, matrix)
}

func moveRockWest(x, y int, matrix *[][]string) {
	if x == 0 || (*matrix)[y][x-1] != "." {
		return
	}

	(*matrix)[y][x] = "."
	(*matrix)[y][x-1] = "O"

	moveRockWest(x-1, y, matrix)
}

func moveRockSouth(x, y int, matrix *[][]string) {
	if y == len(*matrix)-1 || (*matrix)[y+1][x] != "." {
		return
	}

	(*matrix)[y][x] = "."
	(*matrix)[y+1][x] = "O"

	moveRockSouth(x, y+1, matrix)
}

func moveRockEast(x, y int, matrix *[][]string) {
	if x == len(*matrix)-1 || (*matrix)[y][x+1] != "." {
		return
	}

	(*matrix)[y][x] = "."
	(*matrix)[y][x+1] = "O"

	moveRockEast(x+1, y, matrix)
}

func rollRocksForever(lines *[]string) *[][]string {
	maxY := len(*lines)
	matrix := make([][]string, maxY)

	for y, l := range *lines {
		matrix[y] = strings.Split(l, "")
	}

	seen := map[string]int{}
	idx := 1000000000

	for i := 0; i < 1000000000; i++ {
		rockCycle(&matrix)

		if idx == 1000000000 {
			strs := []string{}

			for _, row := range matrix {
				s := strings.Join(row, "")
				strs = append(strs, s)
			}

			key := strings.Join(strs, "")

			if j, ok := seen[key]; ok {
				idx = i + (1000000000-j)%(i-j) - 1
			} else {
				seen[key] = i
			}
		}

		if idx == i {
			break
		}

	}

	return &matrix
}

func rockCycle(matrix *[][]string) {
	maxY := len(*matrix)
	maxX := len((*matrix)[0])

	for y := 1; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			c := (*matrix)[y][x]
			if c != "O" {
				continue
			}

			moveRockNorth(x, y, matrix)
		}
	}

	for x := 1; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			c := (*matrix)[y][x]
			if c != "O" {
				continue
			}

			moveRockWest(x, y, matrix)
		}
	}

	for y := maxY - 2; y >= 0; y-- {
		for x := 0; x < maxX; x++ {
			c := (*matrix)[y][x]
			if c != "O" {
				continue
			}

			moveRockSouth(x, y, matrix)
		}
	}

	for x := maxX - 2; x >= 0; x-- {
		for y := 0; y < maxY; y++ {
			c := (*matrix)[y][x]
			if c != "O" {
				continue
			}

			moveRockEast(x, y, matrix)
		}
	}
}
