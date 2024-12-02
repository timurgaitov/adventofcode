package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Lines(filename string) []string {
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

func Numbers(str string, sep string) []int {
	parts := strings.Split(str, sep)
	nums := make([]int, 0)
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
