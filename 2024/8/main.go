package main

import (
	"adventofcode/u"
	"fmt"
)

func main() {
	charMap := u.ReadCharMap("input.txt")
	if len(charMap) != len(charMap[0]) {
		panic("not square")
	}

	resMap := make([][]byte, len(charMap))
	for i := range resMap {
		resMap[i] = make([]byte, len(charMap))
	}

	for i := 0; i < len(charMap); i++ {
		for j := 0; j < len(charMap); j++ {
			markAntinodes(charMap, resMap, i, j)
		}
	}

	count := 0
	for i := 0; i < len(resMap); i++ {
		for j := 0; j < len(resMap); j++ {
			if resMap[i][j] == '#' {
				count++
			}
		}
	}
	fmt.Println(count)
}

func markAntinodes(charMap [][]byte, resMap [][]byte, curI, curJ int) {
	cur := charMap[curI][curJ]
	if cur == '.' || cur == '#' {
		return
	}

	for i := 0; i < len(charMap); i++ {
		for j := 0; j < len(charMap); j++ {
			if i == curI && j == curJ {
				continue
			}
			if charMap[i][j] != cur {
				continue
			}
			diffI := curI - i
			diffJ := curJ - j
			markAntinode(resMap, curI+diffI, curJ+diffJ)
			markAntinode(resMap, i-diffI, j-diffJ)
		}
	}
}

func markAntinode(resMap [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(resMap) || j >= len(resMap) {
		return
	}
	resMap[i][j] = '#'
}
