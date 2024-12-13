package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	byteMap := utils.ReadByteMap("input.txt")
	if len(byteMap) != len(byteMap[0]) {
		panic("not square")
	}

	resMap := make([][]byte, len(byteMap))
	for i := range resMap {
		resMap[i] = make([]byte, len(byteMap))
	}

	for i := 0; i < len(byteMap); i++ {
		for j := 0; j < len(byteMap); j++ {
			markAntinodes(byteMap, resMap, i, j)
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
	resMap2 := make([][]byte, len(byteMap))
	for i := range resMap2 {
		resMap2[i] = make([]byte, len(byteMap))
	}
	for i := 0; i < len(byteMap); i++ {
		for j := 0; j < len(byteMap); j++ {
			markAntinodes2(byteMap, resMap2, i, j)
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

func markAntinodes(byteMap [][]byte, resMap [][]byte, curI, curJ int) {
	cur := byteMap[curI][curJ]
	if cur == '.' {
		return
	}

	for i := 0; i < len(byteMap); i++ {
		for j := 0; j < len(byteMap); j++ {
			if i == curI && j == curJ {
				continue
			}
			if byteMap[i][j] != cur {
				continue
			}
			diffI := curI - i
			diffJ := curJ - j
			markAntinode(resMap, curI+diffI, curJ+diffJ)
			markAntinode(resMap, i-diffI, j-diffJ)
		}
	}
}

func markAntinodes2(byteMap [][]byte, resMap [][]byte, curI, curJ int) {
	cur := byteMap[curI][curJ]
	if cur == '.' {
		return
	}
	resMap[curI][curJ] = '#'

	for i := 0; i < len(byteMap); i++ {
		for j := 0; j < len(byteMap); j++ {
			if i == curI && j == curJ {
				continue
			}
			if byteMap[i][j] != cur {
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
