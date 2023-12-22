package main

import (
	"container/heap"
	"strings"
)

func solve_day_17_1(url string) int {
	lines := readLinesFromFile(url)
	heatMap := getHeatMap(&lines)

	start := coord{0, 0}
	end := coord{y: len(*heatMap) - 1, x: len((*heatMap)[0]) - 1}
	return dijkstra(start, end, 0, 3, heatMap)
}

func solve_day_17_2(url string) int {
	lines := readLinesFromFile(url)
	heatMap := getHeatMap(&lines)

	start := coord{0, 0}
	end := coord{y: len(*heatMap) - 1, x: len((*heatMap)[0]) - 1}
	return dijkstra(start, end, 4, 10, heatMap)
}

var turn = map[Coordinate]map[string]Coordinate{
	North: {"L": East, "R": West},
	West:  {"L": North, "R": South},
	South: {"R": East, "L": West},
	East:  {"L": South, "R": North},
}

func getHeatMap(lines *[]string) *[][]int {
	matrix := make([][]int, len(*lines))

	for y, l := range *lines {
		matrix[y] = convertStringArrayIntArray(strings.Split(l, ""))
	}

	return &matrix
}

type HeatMove struct {
	c        coord
	dir      Coordinate
	straight int
	heatLoss int
}

type HeatMoveEntry struct {
	c        coord
	dir      Coordinate
	straight int
}

func (h HeatMove) toEntry() HeatMoveEntry {
	return HeatMoveEntry{h.c, h.dir, h.straight}
}

func (h HeatMove) toItem(endPos coord) *Item[HeatMove] {
	return &Item[HeatMove]{
		node:     &h,
		priority: h.heatLoss + h.c.manhattan(endPos),
	}
}

func dijkstra(startingCoord coord, goalCoord coord, minMov, MaxMov int, heatMap *[][]int) int {
	maxX, maxY := len((*heatMap)[0]), len(*heatMap)

	pq := make(PriorityQueue[HeatMove], 0)
	heap.Init(&pq)

	heap.Push(&pq, HeatMove{
		c:        startingCoord,
		dir:      East,
		straight: 0,
		heatLoss: 0,
	}.toItem(goalCoord))

	heap.Push(&pq, HeatMove{
		c:        startingCoord,
		dir:      South,
		straight: 0,
		heatLoss: 0,
	}.toItem(goalCoord))

	seen := make(map[HeatMoveEntry]int)
	previous := make(map[HeatMoveEntry]HeatMoveEntry)

	seen[HeatMoveEntry{startingCoord, East, 0}] = 0
	seen[HeatMoveEntry{startingCoord, South, 0}] = 0

	for !pq.isEmpty() {
		item := heap.Pop(&pq).(*Item[HeatMove])
		current := *item.node

		if current.c == goalCoord && current.straight >= minMov {
			return current.heatLoss
		}

		if v, ok := seen[current.toEntry()]; ok {
			if v < current.heatLoss {
				continue
			}
		}

		for _, dir := range []Coordinate{North, South, East, West} {

			newCoord := current.c.add(dir)

			if !newCoord.isValid(maxX, maxY) ||
				dir == reverseDirection[current.dir] ||
				dir == current.dir && current.straight >= MaxMov ||
				dir != current.dir && current.straight < minMov {
				continue
			}

			s := 1
			if dir == current.dir {
				s = current.straight + 1
			}

			hm := HeatMove{
				c:        newCoord,
				dir:      dir,
				straight: s,
				heatLoss: current.heatLoss + (*heatMap)[newCoord.y][newCoord.x],
			}.toItem(goalCoord)

			if v, ok := seen[hm.node.toEntry()]; !ok || v > hm.node.heatLoss {
				heap.Push(&pq, hm)

				seen[hm.node.toEntry()] = hm.node.heatLoss
				previous[hm.node.toEntry()] = current.toEntry()
			}
		}

	}

	panic("no result ")
}
