package utils

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadStr(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func ReadLines(filename string) []string {
	byteMap := ReadByteMap(filename)
	lines := make([]string, 0, len(byteMap))
	for _, line := range byteMap {
		lines = append(lines, string(line))
	}
	return lines
}

func ReadByteMap(filename string) [][]byte {
	cont, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var resTmp [][]byte
	win := 0

	for l, r := 0, 0; ; r++ {
		if r == len(cont) {
			resTmp = append(resTmp, cont[l:r-win])
			break
		}
		if cont[r] == '\r' {
			if win == 1 {
				panic("bad EOL format")
			}
			win = 1
			continue
		}
		if cont[r] == '\n' {
			resTmp = append(resTmp, cont[l:r-win])
			l = r + 1
			win = 0
		}
		if win == 1 {
			panic("bad EOL format")
		}
	}
	lastEmpty := 0
	if len(resTmp[len(resTmp)-1]) == 0 {
		lastEmpty = 1
	}
	res := make([][]byte, len(resTmp)-lastEmpty)
	copy(res, resTmp)
	return res
}

func ReadNumMap(filename string) [][]byte {
	byteMap := ReadByteMap(filename)
	for i := range byteMap {
		for j := range byteMap[i] {
			byteMap[i][j] -= '0'
		}
	}
	return byteMap
}

func Strings(str string, sepPattern string) []string {
	re := regexp.MustCompile(sepPattern)

	return re.Split(strings.Trim(str, " \r\n"), -1)
}

func Nums(str string, sepPattern string) []int {
	re := regexp.MustCompile(sepPattern)
	numsStrs := re.Split(strings.Trim(str, " \r\n"), -1)

	nums := make([]int, 0, len(numsStrs))
	for _, numStr := range numsStrs {
		nums = append(nums, Num(numStr))
	}
	return nums
}

func RegexpGroups(str string, pattern string) [][]string {
	re := regexp.MustCompile(pattern)
	all := re.FindAllStringSubmatch(str, -1)

	result := make([][]string, 0, len(all))
	for _, groups := range all {
		result = append(result, groups)
	}
	return result
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Num(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func OutOfBound[T any](i int, j int, valMap [][]T) bool {
	return i < 0 || j < 0 || i >= len(valMap) || j >= len(valMap[i])
}

func RoundIndex[T any](i int, overSlice []T) int {
	l := len(overSlice)
	if i < 0 {
		panic("i < 0")
	}
	if l == 0 {
		panic("l == 0")
	}
	return i % l
}
