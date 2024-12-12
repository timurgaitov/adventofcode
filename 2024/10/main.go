package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
)

func main() {
	m := utils.ReadNumMap("input.txt")
	sum := 0
	sum2 := 0
	for i, line := range m {
		for j, num := range line {
			if num != 0 {
				continue
			}
			score, rating := dfs(m, i, j)
			sum += score
			sum2 += rating
		}
	}
	fmt.Println(sum)
	// part 2
	fmt.Println(sum2)
}

var dirs = []pos{{i: -1, j: 0}, {i: 1, j: 0}, {i: 0, j: -1}, {i: 0, j: 1}}

func dfs(m [][]byte, i, j int) (int, int) {
	stack := make([]pos, 10000)
	stackCur := 0
	stack[stackCur] = pos{i: i, j: j, trace: []pos{{i: i, j: j}}}

	count := make(map[lwpos]struct{})
	count2 := 0

	for stackCur >= 0 {
		cur := stack[stackCur]
		stackCur--

		for _, d := range dirs {
			cand := pos{i: cur.i + d.i, j: cur.j + d.j, trace: slices.Concat(nil, cur.trace, []pos{{i: cur.i + d.i, j: cur.j + d.j}})}

			if cand.i < 0 || cand.j < 0 || cand.i >= len(m) || cand.j >= len(m) {
				continue
			}
			if m[cand.i][cand.j]-m[cur.i][cur.j] != 1 {
				continue
			}
			if m[cand.i][cand.j] == 9 {
				//printTrace(m, cand.trace)
				count[lwpos{cand.i, cand.j}] = struct{}{}
				count2++
				continue
			}

			stackCur++
			stack[stackCur] = cand
		}
	}
	return len(count), count2
}

func printTrace(m [][]int, trace []pos) {
	for i := range m {
	JLoop:
		for j, v := range m[i] {
			for _, t := range trace {
				if t.i == i && t.j == j {
					fmt.Print(v)
					continue JLoop
				}
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

type lwpos struct {
	i, j int
}

type pos struct {
	i, j  int
	trace []pos
}

func (p pos) String() string {
	return fmt.Sprintf("(%d,%d)", p.i, p.j)
}
