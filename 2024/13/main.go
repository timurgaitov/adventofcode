package main

import (
	"adventofcode/utils"
	"fmt"
)

type xy struct {
	X, Y int
}

type item struct {
	A      xy
	B      xy
	Prize  xy
	Tokens int
	PressA int
	PressB int
	Mem    map[xy]int
}

func main() {
	lines := utils.ReadLines("input.txt")
	items := parseItems(lines)

	tokens := 0

	for _, itm := range items {
		tokens += rec(itm)
	}

	fmt.Println(tokens)
}

func rec(cur item) (ret int) {
	if mem, ok := cur.Mem[cur.Prize]; ok {
		return mem
	}
	defer func() {
		cur.Mem[cur.Prize] = ret
	}()
	if cur.PressA > 100 || cur.PressB > 100 {
		return 0
	}
	if cur.Prize.X < 0 || cur.Prize.Y < 0 {
		return 0
	}
	if cur.Prize.X == 0 && cur.Prize.Y == 0 {
		println("found")
		return cur.Tokens
	}
	//fmt.Println(cur)
	a := rec(pressButton(cur, cur.A, 1, 0, 3))
	b := rec(pressButton(cur, cur.B, 0, 1, 1))
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a > b {
		return b
	}
	return a
}

func pressButton(itm item, button xy, pressA, pressB, price int) item {
	return item{
		A: itm.A,
		B: itm.B,
		Prize: xy{
			X: itm.Prize.X - button.X,
			Y: itm.Prize.Y - button.Y,
		},
		Tokens: itm.Tokens + price,
		PressA: itm.PressA + pressA,
		PressB: itm.PressB + pressB,
		Mem:    itm.Mem,
	}
}

func parseItems(lines []string) []item {
	items := make([]item, 0)
	for i := 0; i < len(lines); i += 4 {
		lineA := lines[i]
		lineB := lines[i+1]
		linePrize := lines[i+2]
		items = append(items, item{
			A:      parseXY(utils.Strings(lineA, ": ")[1]),
			B:      parseXY(utils.Strings(lineB, ": ")[1]),
			Prize:  parseXY(utils.Strings(linePrize, ": ")[1]),
			Tokens: 0,
			Mem:    make(map[xy]int),
		})
	}
	return items
}

func parseXY(str string) xy {
	gr := utils.RegexpGroups(str, `X[\+=](\d+), Y[\+=](\d+)`)
	return xy{
		X: utils.Num(gr[0][1]),
		Y: utils.Num(gr[0][2]),
	}
}
