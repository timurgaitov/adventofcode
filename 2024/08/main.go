package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	charMap := utils.ReadCharMap("input.txt")
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

	// part 2
	resMap2 := make([][]byte, len(charMap))
	for i := range resMap2 {
		resMap2[i] = make([]byte, len(charMap))
	}
	for i := 0; i < len(charMap); i++ {
		for j := 0; j < len(charMap); j++ {
			markAntinodes2(charMap, resMap2, i, j)
		}
	}
	count2 := 0
	for i := 0; i < len(resMap2); i++ {
		for j := 0; j < len(resMap2); j++ {
			if resMap2[i][j] == '#' {
				count2++
			}
		}
	}
	fmt.Println(count2)
}

func markAntinodes(charMap [][]byte, resMap [][]byte, curI, curJ int) {
	cur := charMap[curI][curJ]
	if cur == '.' {
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

func markAntinodes2(charMap [][]byte, resMap [][]byte, curI, curJ int) {
	cur := charMap[curI][curJ]
	if cur == '.' {
		return
	}
	resMap[curI][curJ] = '#'

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
			for harm := 1; ; harm++ {

				if !markAntinode(resMap, curI+diffI*harm, curJ+diffJ*harm) {
					break
				}
			}
			for harm := 1; ; harm++ {
				if !markAntinode(resMap, i-diffI*harm, j-diffJ*harm) {
					break
				}
			}
		}
	}
}

func markAntinode(resMap [][]byte, i, j int) bool {
	if i < 0 || j < 0 || i >= len(resMap) || j >= len(resMap) {
		return false
	}
	resMap[i][j] = '#'
	return true
}
