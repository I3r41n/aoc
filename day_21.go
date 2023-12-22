package main

import (
	"container/heap"
	"fmt"
	"math"
)

func solve_21_1(url string, steps int) int {
	lines := readLinesFromFile(url)
	matrix := getContraption(&lines)

	return len((*stepsIn(matrix, steps))[steps])
}

func getS(matrix *[][]string) coord {
	for y, v := range *matrix {
		for x, c := range v {
			if c == "S" {
				return coord{x, y}
			}
		}
	}
	panic("No S!!")
}

func solve_21_2(url string, steps int) int {
	lines := readLinesFromFile(url)
	matrix := getContraption(&lines)

	gridSize := len(*matrix)
	period := steps % gridSize

	fmt.Printf("period: %d gridSize: %d\n", period, gridSize)

	// scoord = scoord.add(Coordinate{len((*matrix)[0]) * (4 / 2), len(*matrix) * (4 / 2)})
	gardenPerSteps := *stepsIn(matrix, period+gridSize*2)

	valuesForLagrange := []int{
		len(gardenPerSteps[period]),
		len(gardenPerSteps[period+gridSize]),
		len(gardenPerSteps[period+gridSize*2]),
	}

	a := (valuesForLagrange[0] + valuesForLagrange[2] - 2*valuesForLagrange[1]) / 2
	b := valuesForLagrange[1] - valuesForLagrange[0] - a
	c := valuesForLagrange[0]

	x := (steps - period) / gridSize

	fmt.Printf("valuesForLagrange: %v\n", valuesForLagrange)
	fmt.Printf("x: %v\n", x)

	return a*x*x + b*x + c
}

func mod(a, n int) int {
	return ((a % n) + n) % n
}

func stepsIn(matrix *[][]string, goalSteps int) *map[int]map[coord]bool {
	maxY := len(*matrix)
	maxX := len((*matrix)[0])

	factor := int(math.Ceil(float64(goalSteps / maxY)))

	Scoord := getS(matrix).add(Coordinate{maxX * factor / 2, maxY * factor / 2})

	pq := make(PriorityQueue[Node[int]], 0)
	heap.Init(&pq)

	heap.Push(&pq, Node[int]{
		c: Scoord,
		s: 0,
	}.toItem(0))

	seen := map[int]map[coord]bool{
		0: {Scoord: true},
	}

	for s := 0; s < goalSteps; s++ {
		for current := range seen[s] {
			for _, dir := range []Coordinate{North, South, East, West} {
				newCoord := current.add(dir)

				correctedCoord := coord{mod(newCoord.x, maxX), mod(newCoord.y, maxY)}

				if !correctedCoord.isValid(maxX, maxY) ||
					(*matrix)[correctedCoord.y][correctedCoord.x] == "#" {
					continue
				}

				if _, ok := seen[s+1]; !ok {
					seen[s+1] = make(map[coord]bool)
				}

				seen[s+1][newCoord] = true
			}
		}
	}

	return &seen
}

func expandMap(matrix *[][]string, factor int) *[][]string {
	yMax := len(*matrix)
	for y := 0; y < yMax; y++ {
		repeat := (*matrix)[y]
		for j := 0; j < factor-1; j++ {
			(*matrix)[y] = append((*matrix)[y], repeat...)
		}
	}
	for j := 0; j < factor-1; j++ {
		for y := 0; y < yMax; y++ {
			(*matrix) = append((*matrix), (*matrix)[y])
		}
	}
	return matrix
}
