package utils

import (
	"bufio"
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
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	lines := make([]string, 0)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	return lines
}

func ReadCharMap(filename string) [][]byte {
	cont, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var res [][]byte
	win := 0
	for l, r := 0, 0; r < len(cont); r++ {
		if cont[r] == '\r' {
			if win == 1 {
				panic("bad EOL format")
			}
			win = 1
			continue
		}
		if cont[r] == '\n' {
			res = append(res, cont[l:r-win])
			l = r + 1
			win = 0
		}
		if win == 1 {
			panic("bad EOL format")
		}
	}
	return res
}

func ReadNumMap(filename string) [][]byte {
	charMap := ReadCharMap(filename)
	for i := range charMap {
		for j := range charMap[i] {
			charMap[i][j] -= '0'
		}
	}
	return charMap
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
