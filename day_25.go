package main

import (
	"math/rand"
	"slices"
	"strings"

	"github.com/thoas/go-funk"
)

type v string
type e struct {
	v1 v
	v2 v
}

func solve_25_1(url string) int {
	lines := readLinesFromFile(url)

	vertices := []v{}
	edges := []e{}

	for _, l := range lines {
		s := strings.Split(l, ": ")
		c := strings.Split(s[1], " ")

		v1 := v(s[0])
		if !slices.Contains(vertices, v1) {
			vertices = append(vertices, v1)
		}

		for _, s := range c {
			v2 := v(s)
			if !slices.Contains(vertices, v2) {
				vertices = append(vertices, v2)
			}
			if !slices.Contains(edges, e{v1, v2}) && !slices.Contains(edges, e{v2, v1}) {
				edges = append(edges, e{v1, v2})
			}
		}
	}

	subsets := *Karger(&edges, &vertices)
	return len(subsets[0]) * len(subsets[1])
}

func Karger(edges *[]e, vertices *[]v) *[][]v {
	subsets := [][]v{}

	for countCuts(&subsets, edges) != 3 {
		subsets = [][]v{}

		for _, vert := range *vertices {
			subsets = append(subsets, []v{vert})
		}

		for len(subsets) > 2 {
			i := rand.Int() % len(*edges)
			idx1 := funk.IndexOf(subsets, func(vert []v) bool {
				return slices.Contains(vert, (*edges)[i].v1)
			})
			idx2 := funk.IndexOf(subsets, func(vert []v) bool {
				return slices.Contains(vert, (*edges)[i].v2)
			})

			if idx1 == idx2 {
				continue
			}

			subsets[idx1] = append(subsets[idx1], subsets[idx2]...)
			subsets = slices.Delete(subsets, idx2, idx2+1)
		}
	}

	return &subsets
}

func countCuts(subsets *[][]v, edges *[]e) (cuts int) {
	if len(*subsets) != 2 {
		return 0
	}

	for i := 0; i < len(*edges); i++ {
		if slices.Contains((*subsets)[0], (*edges)[i].v1) &&
			slices.Contains((*subsets)[0], (*edges)[i].v2) ||
			slices.Contains((*subsets)[1], (*edges)[i].v1) &&
				slices.Contains((*subsets)[1], (*edges)[i].v2) {
			continue
		}

		cuts += 1
	}

	return cuts
}
