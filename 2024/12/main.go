package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	byteMap := utils.ReadByteMap("input.txt")

	q := utils.NewQueue()
	visited := make(map[utils.Pos]struct{})

	cost := 0
	cost2 := 0

	for i := range byteMap {
		for j := range byteMap[i] {
			start := utils.Pos{I: i, J: j}
			if _, ok := visited[start]; ok {
				continue
			}
			q.Add(start)

			area := 0
			perimeter := 0
			sides := 0

			for !q.Empty() {
				cur := q.Remove().(utils.Pos)
				if _, ok := visited[cur]; ok {
					continue
				}
				visited[cur] = struct{}{}
				area++

				for _, dir := range utils.DirsSqClockwise {
					pos := posDir(cur, dir)
					if otherArea(pos, cur, byteMap) {
						perimeter++
						continue
					}
					q.Add(pos)
				}

				sides += countConvexCorners(cur, byteMap)
				sides += countConcaveCorners(cur, byteMap)
			}

			//fmt.Printf("%s | area %4d | perimeter %4d | sides %4d\n", byteMap[start.I][start.J], area, perimeter, sides)
			cost += perimeter * area
			cost2 += sides * area
		}
	}

	fmt.Println(cost)
	fmt.Println(cost2)
}

func countConvexCorners(cur utils.Pos, byteMap [][]byte) (corners int) {
	dirs := utils.DirsSqClockwise
	for i := 0; i < len(dirs); i++ {
		d1 := otherArea(posDir(cur, roundDir(i, dirs)), cur, byteMap)
		d2 := otherArea(posDir(cur, roundDir(i+1, dirs)), cur, byteMap)
		if d1 && d2 {
			corners++
		}
	}
	return
}

func countConcaveCorners(cur utils.Pos, byteMap [][]byte) (corners int) {
	dirs := utils.DirsDiagClockwise
	for i := 0; i < len(dirs); i += 2 {
		d1 := otherArea(posDir(cur, roundDir(i, dirs)), cur, byteMap)
		d2 := otherArea(posDir(cur, roundDir(i+1, dirs)), cur, byteMap)
		d3 := otherArea(posDir(cur, roundDir(i+2, dirs)), cur, byteMap)
		if !d1 && d2 && !d3 {
			corners++
		}
	}
	return
}

func roundDir(d int, dirs []utils.Pos) utils.Pos {
	return dirs[d%len(dirs)]
}

func posDir(cur utils.Pos, dir utils.Pos) utils.Pos {
	return utils.Pos{
		I: cur.I + dir.I,
		J: cur.J + dir.J,
	}
}

func otherArea(p utils.Pos, cur utils.Pos, byteMap [][]byte) bool {
	return p.I < 0 || p.J < 0 || p.I >= len(byteMap) || p.J >= len(byteMap) ||
		byteMap[p.I][p.J] != byteMap[cur.I][cur.J]
}
