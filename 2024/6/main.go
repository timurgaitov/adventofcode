package main

import (
	"adventofcode/u"
	"fmt"
)

type dir struct {
	i, j int
}

func main() {
	M := u.ReadFileLines("input.txt")
	if len(M) != len(M[0]) {
		panic("not square")
	}
	size := len(M)

	dirs := []dir{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	visited := make(map[dir]struct{})
	i, j := findGuard(M, size)
	curDir := 0
	for {
		if outOfMap(i, j, size) {
			break
		}
		visited[dir{i, j}] = struct{}{}
		for !outOfMap(i+dirs[curDir].i, j+dirs[curDir].j, size) && M[i+dirs[curDir].i][j+dirs[curDir].j] == '#' {
			curDir = (curDir + 1) % 4
		}
		i, j = i+dirs[curDir].i, j+dirs[curDir].j
	}
	fmt.Println(len(visited))
}

func findGuard(M []string, size int) (int, int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if M[i][j] == '^' {
				return i, j
			}
		}
	}
	panic("guard not found")
}

func outOfMap(i, j, size int) bool {
	return i < 0 || j < 0 || i >= size || j >= size
}
