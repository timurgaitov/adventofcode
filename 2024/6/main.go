package main

import (
	"adventofcode/utils"
	"fmt"
	"sync"
	"sync/atomic"
)

type dir struct {
	i, j int
}

type turn struct {
	i, j int
	dir  dir
}

func main() {
	M := utils.ReadLines("input.txt")
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
			i, j = i-dirs[dir1].i, j-dirs[dir1].j
			dir1 = (dir1 + 1) % 4
			continue
		}
		visited[dir{i, j}] = struct{}{}
		i, j = i+dirs[dir1].i, j+dirs[dir1].j
	}
	fmt.Println(len(visited))

	// part 2
	conur := 12
	obsCh := make(chan dir)
	count := atomic.Int64{}
	wg := &sync.WaitGroup{}
	wg.Add(conur)

	go func() {
		for obsI := 0; obsI < size; obsI++ {
			for obsJ := 0; obsJ < size; obsJ++ {
				obsCh <- dir{i: obsI, j: obsJ}
			}
		}
		close(obsCh)
	}()

	for range conur {
		go func() {
			defer wg.Done()
			for obs := range obsCh {
				visitedTurns := map[turn]struct{}{}
				dir2 := 0
				prevDir2 := 0
				for i, j := guardI, guardJ; !outOfMap(i, j, size); {
					if M[i][j] == '#' || i == obs.i && j == obs.j {
						i, j = i-dirs[dir2].i, j-dirs[dir2].j
						dir2 = (dir2 + 1) % 4
						continue
					}
					if dir2 != prevDir2 {
						prevDir2 = dir2
						tur := turn{i: i, j: j, dir: dirs[dir2]}
						if _, ok := visitedTurns[tur]; ok {
							count.Add(1)
							break
						}
						visitedTurns[tur] = struct{}{}
					}
					i, j = i+dirs[dir2].i, j+dirs[dir2].j
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println(count.Load())
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
