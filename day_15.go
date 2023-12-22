package main

import (
	"strings"

	"github.com/thoas/go-funk"
)

func solve_day_15_1(url string) int {
	strs := readLineFromFile(url, ",")
	return funk.SumInt(getHashs(&strs))
}

func solve_day_15_2(url string) int {
	strs := readLineFromFile(url, ",")
	HASHMAP := getHASHMAP(&strs)

	focusingPower := 0
	for k, lenses := range *HASHMAP {
		for _, lense := range lenses {
			focusingPower += (1 + k) * lense[0] * lense[1]
		}
	}

	return focusingPower
}

func getHASHMAP(strs *[]string) *map[int]map[string][]int {
	HASHMAP := make(map[int]map[string][]int, 256)
	for i := 0; i <= 256; i++ {
		HASHMAP[i] = make(map[string][]int)
	}

	for _, v := range *strs {
		if strings.Index(v, "-") != -1 {
			label := v[0 : len(v)-1]
			hash := getHASH(label)
			lenses := HASHMAP[hash]

			if removed, ok := lenses[label]; ok {
				delete(lenses, label)

				for i, v := range lenses {
					if v[0] > removed[0] {
						lenses[i][0] = v[0] - 1
					}
				}
			}

			continue
		}
		s := strings.Split(v, "=")
		label := s[0]
		lenses := HASHMAP[getHASH(label)]

		idx := len(lenses) + 1
		if len(lenses[label]) > 0 {
			idx = lenses[label][0]
		}

		lenses[s[0]] = []int{idx, convertStringToInt(s[1])}
	}

	return &HASHMAP
}

func getHashs(strs *[]string) []int {
	hashes := make([]int, len(*strs))

	for idx, v := range *strs {
		hashes[idx] = getHASH(v)
	}

	return hashes
}

func getHASH(value string) int {
	return (funk.Reduce([]rune(value), func(acc int, cur rune) int {
		return ((acc + int(cur)) * 17) % 256
	}, 0)).(int)
}
