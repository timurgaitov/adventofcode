package main

import (
	"adventofcode/u"
	"fmt"
)

type dir struct {
	i, j int
}

type turn struct {
	i, j int
	dir  dir
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

	guardI, guardJ := findGuard(M, size)

	visited := make(map[dir]struct{})
	curDir := 0
	for i, j := guardI, guardJ; !outOfMap(i, j, size); {
		for M[i][j] == '#' {
			i, j = i-dirs[curDir].i, j-dirs[curDir].j
			curDir = (curDir + 1) % 4
			continue
		}
		visited[dir{i, j}] = struct{}{}
		i, j = i+dirs[curDir].i, j+dirs[curDir].j
	}
	fmt.Println(len(visited))

	// part 2
	count := 0
	for obsI := 0; obsI < size; obsI++ {
		for obsJ := 0; obsJ < size; obsJ++ {
			visitedTurns := map[turn]struct{}{}
			curDir = 0
			for i, j := guardI, guardJ; !outOfMap(i, j, size); {
				var tur turn
				turned := false
				for M[i][j] == '#' || i == obsI && j == obsJ {
					turned = true
					i, j = i-dirs[curDir].i, j-dirs[curDir].j
					curDir = (curDir + 1) % 4
					tur = turn{i: i, j: j, dir: dirs[curDir]}
					i, j = i+dirs[curDir].i, j+dirs[curDir].j
				}
				if turned {
					if _, ok := visitedTurns[tur]; ok {
						count++
						break
					}
					visitedTurns[tur] = struct{}{}
				}
				i, j = i+dirs[curDir].i, j+dirs[curDir].j
			}
		}
	}
	fmt.Println(count)
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
