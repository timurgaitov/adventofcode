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

	// A*a.x + B*b.x = p.x
	// A*a.y + B*b.y = p.y
	// A? B?
	// A = (p.y - B*b.y)/a.y					     <- A
	// (p.y - B*b.y) * a.x/a.y + B*b.x = p.x
	// p.y*a.x/a.y - B*a.x*b.y/a.y + B*b.x = p.x
	// B*(b.x - a.x*b.y/a.y) = p.x - p.y*a.x/a.y     | *a.y
	// B*(a.y*b.x - a.x*b.y) = p.x*a.y - p.y*a.x
	// B = (p.x*a.y - p.y*a.x)/(a.y*b.x - a.x*b.y)   <- B

	tokens1 := 0
	//tokens2 := 0

	for _, itm := range items {
		p := itm.Prize

		// part 2
		p.X += 10000000000000
		p.Y += 10000000000000

		a := itm.A
		b := itm.B

		B := (p.X*a.Y - p.Y*a.X) / (a.Y*b.X - a.X*b.Y)
		A := (p.Y - B*b.Y) / a.Y

		modB := (p.X*a.Y - p.Y*a.X) % (a.Y*b.X - a.X*b.Y)
		modA := (p.Y - B*b.Y) % a.Y

		// part 2
		if A < 0 || B < 0 || modA != 0 || modB != 0 {
			continue
		}
		//if A < 0 || B < 0 || A > 100 || B > 100 || modA != 0 || modB != 0 {
		//	continue
		//}

		tokens1 += 3*A + B
		//tokens2 += rec(itm)
		//if tokens1 != tokens2 {
		//	fmt.Println("wrong")
		//}
	}

	fmt.Println(tokens1)
	//fmt.Println(tokens2)
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
		//println("found")
		return cur.Tokens
	}
	//fmt.Println(cur)
	a := rec(press(cur, cur.A, 1, 0, 3))
	b := rec(press(cur, cur.B, 0, 1, 1))
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

func press(itm item, button xy, pressA, pressB, price int) item {
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
		prize := parseXY(utils.Strings(linePrize, ": ")[1])

		items = append(items, item{
			A:      parseXY(utils.Strings(lineA, ": ")[1]),
			B:      parseXY(utils.Strings(lineB, ": ")[1]),
			Prize:  prize,
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
