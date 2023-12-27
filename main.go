package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

type pos[T any] struct {
	value T
	x, y  int
}

type coord struct {
	x, y int
}

type coordZ struct {
	x, y, z int
}

func (c coordZ) add(o coordZ) coordZ {
	return coordZ{c.x + o.x, c.y + o.y, c.z + o.z}
}
func (c coord) add(d Coordinate) coord {
	return coord{c.x + d.x, c.y + d.y}
}

func (c coord) move(d Coordinate, len int) coord {
	xl, yl := d.x*len, d.y*len
	return coord{c.x + xl, c.y + yl}
}

func (c coord) manhattan(d coord) int {
	return manhattan(c.x, d.x, c.y, d.y)
}

func (c coord) isValid(maxX int, maxY int) bool {
	return c.x >= 0 && c.x < maxX && c.y >= 0 && c.y < maxY
}

func (c coord) toItem(priority int) *Item[coord] {
	return &Item[coord]{
		node:     &c,
		priority: priority,
	}
}

type Node[T any] struct {
	c    coord
	s    T
	prev *Node[T]
}

type Item[T any] struct {
	node     *T
	priority int
	index    int
}

func (n Node[T]) toItem(priority int) *Item[Node[T]] {
	return &Item[Node[T]]{
		node:     &n,
		priority: priority,
	}
}

type PriorityQueue[T any] []*Item[T]

func (pq PriorityQueue[T]) isEmpty() bool { return pq.Len() == 0 }
func (pq PriorityQueue[T]) Len() int      { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) update(item *Item[T], node *T, priority int) {
	item.node = node
	item.priority = priority
	heap.Fix(pq, item.index)
}

func readFile(url string) string {
	b, err := os.ReadFile(url)
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func readLinesFromFile(url string) []string {
	return readLineFromFile(url, "\n")
}

func readLineFromFile(url string, sep string) []string {
	return strings.Split(readFile(url), sep)
}

func IsDigit(c string) (int64, bool) {
	res, err := strconv.ParseInt(c, 10, 64)
	return res, err == nil
}

func Range(start, max, step int) []int {
	count := (max - start) / step
	nums := make([]int, count)
	for i := range nums {
		nums[i] = start + i*step
	}
	return nums
}

func convertStringToInt(str string) int {
	id, _ := strconv.Atoi(strings.TrimSpace(str))

	return id
}

func convertStringArrayIntArray(ns []string) (ints []int) {
	funk.ForEach(ns, func(n string) {
		ints = append(ints, convertStringToInt(n))
	})

	return ints
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func All[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}

func Ternary[T any](truthy T, falsy T, result bool) T {
	return (map[bool]T{true: truthy, false: falsy})[result]

}
func manhattan(x0, x1, y0, y1 int) int {
	return int(math.Abs(float64(x0-x1)) + math.Abs(float64(y0-y1)))
}

func transpose[T any](matrix *[][]T) *[][]T {
	transpose := make([][]T, len((*matrix)[0]))

	for i := range (*matrix)[0] {
		transpose[i] = make([]T, len(*matrix))
	}

	for i, rows := range *matrix {
		for j := range rows {
			transpose[j][i] = (*matrix)[i][j]
		}
	}
	return &transpose
}

func printMatrix[T any](matrix *[][]T) {
	for i, rows := range *matrix {
		for j := range rows {
			fmt.Print((*matrix)[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}

func ArrayDiff(array1 []string, array2 []string) map[int]string {
	c := make(map[int]string)
	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			c[i] = array2[i]
		}
	}

	return c
}
