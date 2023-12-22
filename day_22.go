package main

import (
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strings"
)

type block struct {
	label byte
	b     coordZ
	e     coordZ
}

func (b block) inZ(z int) bool {
	zMin := min(b.b.z, b.e.z)
	zMax := max(b.b.z, b.e.z)

	return zMin <= z && z <= zMax
}

func (b block) overlaps(other block, excludes ...block) bool {
	xMin := min(b.b.x, b.e.x)
	xMax := max(b.b.x, b.e.x)
	yMin := min(b.b.y, b.e.y)
	yMax := max(b.b.y, b.e.y)
	zMin := min(b.b.z, b.e.z)
	zMax := max(b.b.z, b.e.z)

	otherxMin := min(other.b.x, other.e.x)
	otherxMax := max(other.b.x, other.e.x)
	otheryMin := min(other.b.y, other.e.y)
	otheryMax := max(other.b.y, other.e.y)
	otherzMin := min(other.b.z, other.e.z)
	otherzMax := max(other.b.z, other.e.z)

	return !slices.Contains(excludes, other) &&
		otherxMin <= xMax && xMin <= otherxMax &&
		otheryMin <= yMax && yMin <= otheryMax &&
		otherzMin <= zMax && zMin <= otherzMax
}

type Blocks []block

func (b Blocks) Len() int      { return len(b) }
func (b Blocks) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b Blocks) Less(i, j int) bool {
	z1, z2 := b[i].b.z, b[j].b.z
	x1, x2 := b[i].b.x, b[j].b.x
	y1, y2 := b[i].b.y, b[j].b.y

	if z1 != z2 {
		return z1 < z2
	}

	if x1 != x2 {
		return x1 < x2
	}

	return y1 < y2
}

func (B Blocks) atZ(z int, excludes ...block) Blocks {
	sameZ := make(Blocks, 0)

	for _, b := range B {
		if b.inZ(z) && !slices.Contains(excludes, b) {
			sameZ = append(sameZ, b)
		}
	}

	return sameZ
}

func (B Blocks) hasCollisionsAtZ(z int, b block) bool {
	for _, other := range B {
		if other.inZ(z) && b.overlaps(other) {
			return true
		}
	}

	return false
}

func (blocks Blocks) dropAll() int {
	var droppedCount int

	for idx, b := range blocks {
		res := blocks.drop(b)
		if res != b {
			blocks[idx] = res
			droppedCount++
		}
	}
	return droppedCount
}

func (B Blocks) drop(b block) block {
	z := min(b.b.z, b.e.z)

	for z -= 1; z > 0; z-- {
		b.b.z -= 1
		b.e.z -= 1

		if B.hasCollisionsAtZ(z, b) {
			b.b.z += 1
			b.e.z += 1
			break
		}

	}
	return b
}

func (B Blocks) isDesintegrable(b block) bool {
	zMax := max(b.b.z, b.e.z)
	onTop := make(Blocks, 0)

	moved := b
	moved.b.z, moved.e.z = moved.b.z+1, moved.e.z+1
	for _, other := range B {
		if moved.overlaps(other, b) {
			onTop = append(onTop, other)
		}
	}

	for _, n := range B.atZ(zMax, b) {
		n.b.z += 1
		n.e.z += 1

		onTop = slices.DeleteFunc(onTop, func(o block) bool {
			return o.overlaps(n)
		})
	}

	return len(onTop) == 0
}

func (B Blocks) AllDesintegrable() Blocks {
	all := make(Blocks, 0)

	for _, b := range B {
		if B.isDesintegrable(b) {
			all = append(all, b)
		}
	}
	// fmt.Printf("all: %v\n", all)
	return all
}

func (B Blocks) AllNonDesintegrable() Blocks {
	all := make(Blocks, 0)

	for _, b := range B {
		if !B.isDesintegrable(b) {
			all = append(all, b)
		}
	}
	fmt.Printf("all: %v\n", all)
	return all
}

var REGEX_Z_COORDS = regexp.MustCompile("(\\d+,\\d+,\\d+)~(\\d+,\\d+,\\d+)")

func solve_22_1(url string) int {
	lines := readLinesFromFile(url)

	blocks := getBlocks(lines)
	blocks.dropAll()

	return len(blocks.AllDesintegrable())
}

func solve_22_2(url string) int {
	lines := readLinesFromFile(url)

	blocks := getBlocks(lines)
	blocks.dropAll()

	dropped := 0

	for _, b := range blocks.AllNonDesintegrable() {
		idx := slices.Index(blocks, b)
		test := slices.Clone(blocks)
		test = slices.Delete(test, idx, idx+1)

		for {
			n := test.dropAll()
			if n == 0 {
				break
			}
			dropped += n
		}

	}
	return dropped
}

func getBlocks(lines []string) Blocks {
	blocks := make(Blocks, len(lines))

	for idx, l := range lines {
		s := REGEX_Z_COORDS.FindAllStringSubmatch(l, -1)
		c := convertStringArrayIntArray(strings.Split(s[0][1], ","))
		b := coordZ{c[0], c[1], c[2]}

		c = convertStringArrayIntArray(strings.Split(s[0][2], ","))
		e := coordZ{c[0], c[1], c[2]}

		blocks[idx] = block{byte('A' + idx), b, e}
	}

	sort.Sort(blocks)

	// blocks = *fall(&blocks)

	// fmt.Printf("blocks: %v\n", blocks)
	return blocks
}
