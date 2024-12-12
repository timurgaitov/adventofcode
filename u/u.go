package u

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadFileStr(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func ReadFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	lines := make([]string, 0)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	return lines
}

func ReadCharMap(filename string) [][]rune {
	file, err0 := os.Open(filename)
	if err0 != nil {
		panic(err0)
	}
	defer func() {
		_ = file.Close()
	}()

	res := make([][]rune, 0)
	scan := bufio.NewScanner(file)
	expectedLen := -1
	for scan.Scan() {
		rd := bytes.NewReader(scan.Bytes())
		if expectedLen < 0 {
			expectedLen = rd.Len()
		}
		actualLen := 0
		rs := make([]rune, 0, expectedLen)
		for {
			r, _, err := rd.ReadRune()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			rs = append(rs, r)
			actualLen++
		}
		if actualLen != expectedLen {
			panic(fmt.Sprintf("wrong length, expected %d, actual %d", actualLen, expectedLen))
		}
		res = append(res, rs)
	}
	return res
}

func ReadIntMap(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	lines := make([][]int, 0)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		bytes := scan.Bytes()
		nums := make([]int, 0)
		for _, b := range bytes {
			nums = append(nums, int(b-'0'))
		}
		lines = append(lines, nums)
	}
	return lines
}

func Strs(str string, sepPattern string) []string {
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
