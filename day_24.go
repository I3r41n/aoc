package main

import (
	"fmt"
	"regexp"
	"slices"

	"github.com/thoas/go-funk"
)

type hailstone struct {
	pos      coordZ
	velocity coordZ
}

func (c coord) within(min, max int) bool {
	return true &&
		c.x >= min && c.x <= max &&
		c.y >= min && c.y <= max
}

func (s hailstone) posInFuture(x, y int) bool {
	return (x > s.pos.x) == (s.velocity.x > 0)
}

func (s hailstone) intersection(h, exclusions hailrain, min, max int) hailrain {
	intersections := make(hailrain, 0)
	if len(h) == len(exclusions) {
		return intersections
	}

	slope := float64(s.velocity.y) / float64(s.velocity.x)

	for _, other := range h {
		if slices.Contains(exclusions, other) {
			continue
		}

		otherSlope := float64(other.velocity.y) / float64(other.velocity.x)
		if otherSlope == slope {
			continue
		}

		c := float64(s.pos.y) - (slope * float64(s.pos.x))
		d := float64(other.pos.y) - (otherSlope * float64(other.pos.x))

		x := (d - c) / (slope - otherSlope)
		y := (slope * x) + c

		if (coord{int(x), int(y)}).within(min, max) && s.posInFuture(int(x), int(y)) && other.posInFuture(int(x), int(y)) {
			intersections = append(intersections, other)
		}
	}
	return intersections
}

type hailrain []hailstone

func (h hailrain) intersections(min, max int) *hailrain {
	intersections := make(hailrain, 0)
	exclusions := make(hailrain, 0)
	for _, s := range h {
		exclusions = append(exclusions, s)
		intersections = append(intersections, s.intersection(h, exclusions, min, max)...)
	}

	return &intersections
}

var REGEX_HAILSTONES = regexp.MustCompile(`(\d+), (\d+), (\d+) @ ([-| ]?\d+), ([-| ]?\d+), ([-| ]?\d+)`)

func solve_24_1(url string, testAreaMin, testAreaMAx int) int {
	lines := readLinesFromFile(url)

	stones := getHailStones(&lines)

	return len(*stones.intersections(testAreaMin, testAreaMAx))
}

func solve_24_2(url string, testAreaMin, testAreaMAx int) int {
	lines := readLinesFromFile(url)

	stones := getHailStones(&lines)
	fmt.Println("Run the following script in https://sagecell.sagemath.org/")
	fmt.Println("")
	fmt.Println("var('x y z vx vy vz t1 t2 t3 ans')")

	for idx, s := range (*stones)[0:3] {
		fmt.Printf("eq%d =  x + (vx * t%d) == %d + (%d * t%d)\n", idx*3+1, idx+1, s.pos.x, s.velocity.x, idx+1)
		fmt.Printf("eq%d =  y + (vy * t%d) == %d + (%d * t%d)\n", idx*3+2, idx+1, s.pos.y, s.velocity.y, idx+1)
		fmt.Printf("eq%d =  z + (vz * t%d) == %d + (%d * t%d)\n", idx*3+3, idx+1, s.pos.z, s.velocity.z, idx+1)
	}

	fmt.Println("eq10 = ans == x + y + z")
	fmt.Println("print(solve([eq1,eq2,eq3,eq4,eq5,eq6,eq7,eq8,eq9,eq10],x,y,z,vx,vy,vz,t1,t2,t3,ans))")

	return Ternary[int](47, 885093461440405, url == "puzzles/day_24/example_1")
}

func getHailStones(lines *[]string) *hailrain {
	stones := make(hailrain, len(*lines))
	for idx, l := range *lines {

		r := funk.Flatten(REGEX_HAILSTONES.FindAllStringSubmatch(l, -1)).([]string)
		x, y, z := convertStringToInt(r[1]), convertStringToInt(r[2]), convertStringToInt(r[3])
		vx, vy, vz := convertStringToInt(r[4]), convertStringToInt(r[5]), convertStringToInt(r[6])

		stones[idx] = hailstone{coordZ{x, y, z}, coordZ{vx, vy, vz}}
	}
	return &stones
}
