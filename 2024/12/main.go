package main

import (
	"adventofcode/u"
	"fmt"
)

func main() {
	chM := u.ReadCharMap2("input.txt")

	q := u.NewQueue()
	visited := make(map[u.Pos]struct{})

	cost := 0
	cost2 := 0

	for i := range chM {
		for j := range chM[i] {
			start := u.Pos{I: i, J: j}
			if _, ok := visited[start]; ok {
				continue
			}
			q.Add(start)

			area := 0
			perimeter := 0
			corners := 0

			for !q.Empty() {
				cur := q.Remove().(u.Pos)
				if _, ok := visited[cur]; ok {
					continue
				}
				visited[cur] = struct{}{}
				area++

				for _, dir := range u.DirsSqClockwise {
					d := u.Pos{I: cur.I + dir.I, J: cur.J + dir.J}
					if otherArea(d, cur, chM) {
						perimeter++
						continue
					}
					q.Add(d)
				}

				corn := 0

				for k := 0; k < len(u.DirsSqClockwise); k++ {
					o1 := otherArea(u.Pos{
						I: cur.I + u.DirsSqClockwise[k].I,
						J: cur.J + u.DirsSqClockwise[k].J,
					}, cur, chM)
					o2 := otherArea(u.Pos{
						I: cur.I + u.DirsSqClockwise[(k+1)%len(u.DirsSqClockwise)].I,
						J: cur.J + u.DirsSqClockwise[(k+1)%len(u.DirsSqClockwise)].J,
					}, cur, chM)
					if o1 && o2 {
						corn++
					}
				}

				for k := 0; k < len(u.DirsDiagClockwise); k += 2 {
					o1 := otherArea(u.Pos{
						I: cur.I + u.DirsDiagClockwise[k].I,
						J: cur.J + u.DirsDiagClockwise[k].J,
					}, cur, chM)
					o2 := otherArea(u.Pos{
						I: cur.I + u.DirsDiagClockwise[(k+1)%len(u.DirsDiagClockwise)].I,
						J: cur.J + u.DirsDiagClockwise[(k+1)%len(u.DirsDiagClockwise)].J,
					}, cur, chM)
					o3 := otherArea(u.Pos{
						I: cur.I + u.DirsDiagClockwise[(k+2)%len(u.DirsDiagClockwise)].I,
						J: cur.J + u.DirsDiagClockwise[(k+2)%len(u.DirsDiagClockwise)].J,
					}, cur, chM)
					if !o1 && o2 && !o3 {
						corn++
					}
				}

				if corn > 0 {
					z := corn
					y := fmt.Sprintf("%v", cur)
					x := chM[cur.I][cur.J]
					_ = x
					_ = y
					_ = z
					corners += corn
				}
			}

			fmt.Println(chM[start.I][start.J], area, corners)
			cost += perimeter * area
			cost2 += corners * area
		}
	}

	fmt.Println(cost)
	fmt.Println(cost2)
}

func otherArea(p u.Pos, cur u.Pos, chM [][]string) bool {
	return p.I < 0 || p.J < 0 || p.I >= len(chM) || p.J >= len(chM) ||
		chM[p.I][p.J] != chM[cur.I][cur.J]
}
