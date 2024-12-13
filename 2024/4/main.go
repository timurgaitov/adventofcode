package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	M := utils.ReadLines("input.txt")
	if len(M[0]) != len(M) {
		panic("not square")
	}
	size := len(M)

	res1 := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for _, d := range dirs {
				res1 += searchXMAS(M, i, j, size, 0, d)
			}
		}
	}
	fmt.Println(res1)

	res2 := 0
	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			res2 += searchMASCROSS(M, i, j)
		}
	}
	fmt.Println(res2)
}

type dir struct {
	i int
	j int
}

var dirs = []dir{
	{i: 0, j: 1},
	{i: 1, j: 1},
	{i: 1, j: 0},
	{i: 1, j: -1},
	{i: 0, j: -1},
	{i: -1, j: -1},
	{i: -1, j: 0},
	{i: -1, j: 1},
}

func searchXMAS(M []string, i, j, size int, cur int, d dir) int {
	if isOutOfBoundaries(i, j, size) {
		return 0
	}
	if M[i][j] != "XMAS"[cur] {
		return 0
	}
	if cur == len("XMAS")-1 {
		return 1
	}
	return searchXMAS(M, i+d.i, j+d.j, size, cur+1, d)
}

func searchMASCROSS(M []string, i, j int) int {
	if M[i][j] != 'A' {
		return 0
	}
	diag1 := fmt.Sprintf("%cA%c", M[i-1][j-1], M[i+1][j+1])
	diag2 := fmt.Sprintf("%cA%c", M[i+1][j-1], M[i-1][j+1])
	if diag1 != "MAS" && diag1 != "SAM" || diag2 != "MAS" && diag2 != "SAM" {
		return 0
	}
	return 1
}

func isOutOfBoundaries(i, j int, size int) bool {
	return i < 0 || i >= size || j < 0 || j >= size
}
