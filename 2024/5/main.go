package main

import (
	"adventofcode/u"
	"fmt"
	"sort"
	"strings"
)

func main() {
	lines := u.ReadFileLines("input.txt")

	rules := make(map[int]map[int]struct{})
	updates := make([][]int, 0)
	f := false
	for _, line := range lines {
		if line == "" {
			f = true
			continue
		}
		if f {
			updates = append(updates, u.Nums(line, ","))
		} else {
			split := strings.Split(line, "|")
			l, r := u.Num(split[0]), u.Num(split[1])
			if rules[l] == nil {
				rules[l] = make(map[int]struct{})
			}
			rules[l][r] = struct{}{}
		}
	}
	sum := 0
	incor := make([][]int, 0)
UpdatesLoop:
	for _, update := range updates {
		for i, p := range update {
			rul := rules[p]
			for j := 0; j < i; j++ {
				if _, ok := rul[update[j]]; ok {
					incor = append(incor, update)
					continue UpdatesLoop
				}
			}
		}
		sum += update[len(update)/2]
	}
	fmt.Println(sum)

	sum2 := 0
	for _, update := range incor {
		sort.SliceStable(update, func(i, j int) bool {
			if _, ok := rules[update[i]][update[j]]; ok {
				return true
			}
			if _, ok := rules[update[j]][update[i]]; ok {
				return false
			}
			return false
		})
		sum2 += update[len(update)/2]
	}
	fmt.Println(sum2)
}
