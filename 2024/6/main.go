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
	dir1 := 0
	for i, j := guardI, guardJ; !outOfMap(i, j, size); {
		if M[i][j] == '#' {
			// obstacle, go back, change direction
			i, j = i-dirs[dir1].i, j-dirs[dir1].j
			dir1 = (dir1 + 1) % 4
			continue
		}
		visited[dir{i, j}] = struct{}{}
		i, j = i+dirs[dir1].i, j+dirs[dir1].j
	}
	fmt.Println(len(visited))

	// part 2
	count := 0
	for obsI := 0; obsI < size; obsI++ {
		for obsJ := 0; obsJ < size; obsJ++ {
			visitedTurns := map[turn]struct{}{}
			dir2 := 0
			prevDir2 := 0
			for i, j := guardI, guardJ; !outOfMap(i, j, size); {
				if M[i][j] == '#' || i == obsI && j == obsJ {
					// obstacle, go back, change direction
					i, j = i-dirs[dir2].i, j-dirs[dir2].j
					dir2 = (dir2 + 1) % 4
					continue
				}
				if dir2 != prevDir2 {
					// turned
					prevDir2 = dir2
					tur := turn{i: i, j: j, dir: dirs[dir2]}
					if _, ok := visitedTurns[tur]; ok {
						count++
						break
					}
					visitedTurns[tur] = struct{}{}
				}
				i, j = i+dirs[dir2].i, j+dirs[dir2].j
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
