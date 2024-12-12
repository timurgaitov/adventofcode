package main

import (
	"adventofcode/u"
	"fmt"
)

func main() {
	chM := u.ReadCharMap("example.txt")

	q := u.NewQueue()
	visited := make(map[u.Pos]struct{})

	cost := 0

	for i := range chM {
		for j := range chM[i] {
			start := u.Pos{I: i, J: j}
			if _, ok := visited[start]; ok {
				continue
			}
			q.Add(start)

			area := 0
			perimeter := 0

			for !q.Empty() {
				cur := q.Remove().(u.Pos)
				if _, ok := visited[cur]; ok {
					continue
				}
				visited[cur] = struct{}{}
				area++

				for _, dir := range u.DirsSq {
					d := u.Pos{I: cur.I + dir.I, J: cur.J + dir.J}
					if d.I < 0 || d.J < 0 || d.I >= len(chM) || d.J >= len(chM) || chM[d.I][d.J] != chM[cur.I][cur.J] {
						perimeter++
						continue
					}
					q.Add(d)
				}
			}

			//fmt.Println(string(chM[start.I][start.J]), perimeter, area)
			cost += perimeter * area
		}
	}

	fmt.Println(cost)
	// 1342720 incorrect
}
