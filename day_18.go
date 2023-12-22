package main

import (
	"math"
	"strconv"
	"strings"
)

var convertDir = map[string]Coordinate{
	"R": East,
	"D": South,
	"L": West,
	"U": North,
}

var convertNumberDir = map[int]Coordinate{
	0: East,
	1: South,
	2: West,
	3: North,
}

type LavaPit []coord

func (polygon LavaPit) shoelace() int {
	total := len(polygon) - 1
	res := 0
	j := total
	for i := 0; i < total; i++ {
		res += (polygon[j].x + polygon[j].x) * (polygon[j].y - polygon[i].y)
		j = i
	}

	res = int(math.Abs(float64(res)) / 2.0)

	return res
}
func solve_day_18_1(url string) int {
	lines := readLinesFromFile(url)

	return digLavaPit1(&lines)
}

func solve_day_18_2(url string) int {
	lines := readLinesFromFile(url)

	return digLavaPit2(&lines)
}

func digLavaPit2(lines *[]string) int {
	instructions := make([]digInstruction, len(*lines))

	for _, l := range *lines {
		instruction := strings.Split(l, " ")
		size, _ := strconv.ParseInt(instruction[2][2:7], 16, 64)
		coordinate := convertNumberDir[convertStringToInt(instruction[2][7:8])]

		instructions = append(instructions, digInstruction{dir: coordinate, move: int(size)})
	}

	return dig(instructions)
}

func digLavaPit1(lines *[]string) int {
	instructions := make([]digInstruction, len(*lines))

	for _, l := range *lines {

		instruction := strings.Split(l, " ")
		size := convertStringToInt(instruction[1])
		coordinate := convertDir[instruction[0]]

		instructions = append(instructions, digInstruction{dir: coordinate, move: size})
	}

	return dig(instructions)
}

type digInstruction struct {
	dir  Coordinate
	move int
}

func dig(instructions []digInstruction) int {
	lavaPit := make(LavaPit, 0)
	pos := coord{0, 0}
	perimeter := 0

	for _, i := range instructions {
		lavaPit = append(lavaPit, pos)
		perimeter += i.move
		pos = pos.move(i.dir, i.move)
	}

	area := lavaPit.shoelace()
	return perimeter/2 + area + 1
}
