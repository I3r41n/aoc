package main

import (
	"container/heap"
	"math"
	"strings"
)

type Coordinate coord

var (
	North Coordinate = Coordinate{x: 0, y: -1}
	South Coordinate = Coordinate{x: 0, y: 1}
	East  Coordinate = Coordinate{x: 1, y: 0}
	West  Coordinate = Coordinate{x: -1, y: 0}
)

var reverseDirection = map[Coordinate]Coordinate{
	North: South,
	South: North,
	East:  West,
	West:  East,
}

var direction = map[Coordinate]map[string]bool{
	North: {"|": true, "L": true, "J": true},
	South: {"|": true, "7": true, "F": true},
	East:  {"-": true, "L": true, "F": true},
	West:  {"-": true, "7": true, "J": true},
}

func solve10_1(url string) int {
	lines := readLinesFromFile(url)
	graph, _ := generatePipeGraph(&lines)

	count := 0
	for _, v := range *graph {
		if v > count {
			count = v
		}
	}
	return count
}

func solve10_2(url string) int {
	lines := readLinesFromFile(url)
	graph, matrix := generatePipeGraph(&lines)

	for y, l := range *matrix {
		for x := range l {
			_, ok := (*graph)[coord{x: x, y: y}]
			if ok {
				continue
			}
			(*matrix)[y][x] = "O"
		}
	}
	return getInsideNodesCount_1(matrix)
}

func getInsideNodesCount_1(matrix *[][]string) (inside int) {
	hWall := 0
	for y, l := range *matrix {
		var last = ""
		for x, c := range l {
			v := string(c)

			switch {
			case v == "O" && hWall%2 == 1:
				(*matrix)[y][x] = "I"
				inside += 1
			case v == "O":
				continue
			case v == "|":
				hWall += 1
				last = ""
			case (v == "7" && last == "L") || (v == "J" && last == "F"):
				hWall += 1
				last = ""
			case v == "L" || v == "F":
				last = v
			default:
				continue
			}
		}
	}
	// fmt.Printf("matrix: %v\n", matrix)

	return inside
}

func generatePipeGraph(lines *[]string) (*map[coord]int, *[][]string) {
	s, matrix := getMatrixAndS(lines)
	dist := make(map[coord]int)
	dist[*s] = 0
	visited := make([]*Node[string], 0)
	open := make(PriorityQueue[Node[string]], 1)
	open[0] = &Item[Node[string]]{node: &Node[string]{c: *s, s: (*matrix)[s.y][s.x]}, priority: 0, index: 0}

	heap.Init(&open)

	for len(open) > 0 {
		var current = heap.Pop(&open).(*Item[Node[string]])

		visited = append(visited, current.node)
		for _, v := range getConnectedPipes(current, matrix) {
			alt := dist[current.node.c] + 1
			d, ok := dist[v.node.c]
			if !ok {
				d = math.MaxInt
				dist[v.node.c] = math.MaxInt
				heap.Push(&open, &Item[Node[string]]{node: v.node, priority: math.MaxInt})
			}
			if d > alt {
				dist[v.node.c] = alt
				v.node.prev = current.node
				open.update(v, v.node, alt)
			}
		}
	}

	return &dist, matrix
}

func getPipeNeighbors(xc, yc int, matrix *[][]string) []*Node[string] {
	my, mx := len(*matrix), len((*matrix)[0])
	neighbors := []*Node[string]{}

	for _, j := range []int{-1, 0, 1} {
		for _, i := range []int{-1, 0, 1} {
			if i == j || (i != 0 && j != 0) {
				continue
			}

			xn, yn := xc+i, yc+j
			if xn >= mx || xn < 0 || yn >= my || yn < 0 {
				continue
			}

			s := (*matrix)[yn][xn]
			if s == "." {
				continue
			}
			neighbors = append(neighbors, &Node[string]{s: s, c: coord{xn, yn}})
		}
	}
	return neighbors
}

func getConnectedPipes(c *Item[Node[string]], matrix *[][]string) (connectedPipes []*Item[Node[string]]) {
	xc, yc := c.node.c.x, c.node.c.y

	neighbors := getPipeNeighbors(xc, yc, matrix)

	for _, n := range neighbors {
		dir := Coordinate{x: n.c.x - xc, y: n.c.y - yc}
		nConnects := direction[reverseDirection[dir]][n.s]
		cConnects := direction[dir][c.node.s]
		if nConnects && cConnects {
			connectedPipes = append(connectedPipes, &Item[Node[string]]{node: n})
		}
	}
	return connectedPipes
}

func getMatrixAndS(lines *[]string) (*coord, *[][]string) {
	var sCoord *coord
	matrix := make([][]string, len(*lines))

	for y, l := range *lines {
		matrix[y] = strings.Split(l, "")
		for x, s := range matrix[y] {
			if s == "S" {
				sCoord = &coord{y: y, x: x}
			}
		}
	}

	for _, p := range "|-LJ7F" {
		s := string(p)

		connectedPipes := getConnectedPipes(&Item[Node[string]]{node: &Node[string]{c: *sCoord, s: s}}, &matrix)

		if len(connectedPipes) == 2 {
			matrix[sCoord.y][sCoord.x] = s
			return sCoord, &matrix
		}
	}

	panic("couldnt find S")
}
