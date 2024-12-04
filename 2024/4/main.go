package main

import (
	"adventofcode/u"
	"fmt"
)

func main() {
	M := u.ReadFileLines("input.txt")
	if len(M[0]) != len(M) {
		panic("not square")
	}
	size := len(M)

	res := 0
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			for _, d := range dirs {
				res += searchXMAS(M, x, y, size, 0, d)
			}
		}
	}
	fmt.Println(res)
}

type dir struct {
	x int
	y int
}

var dirs = []dir{
	{x: 0, y: 1},
	{x: 1, y: 1},
	{x: 1, y: 0},
	{x: 1, y: -1},
	{x: 0, y: -1},
	{x: -1, y: -1},
	{x: -1, y: 0},
	{x: -1, y: 1},
}

func searchXMAS(M []string, x, y, size int, cur int, d dir) int {
	if isOutOfBoundaries(x, y, size) {
		return 0
	}
	if M[x][y] != "XMAS"[cur] {
		return 0
	}
	if cur == len("XMAS")-1 {
		return 1
	}
	return searchXMAS(M, x+d.x, y+d.y, size, cur+1, d)
}

func isOutOfBoundaries(x, y int, size int) bool {
	return x < 0 || x >= size || y < 0 || y >= size
}
