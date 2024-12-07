package main

import (
	"adventofcode/u"
	"fmt"
)

func main() {
	lines := u.ReadFileLines("input.txt")

	sum := 0
	for _, l := range lines {
		split := u.Strs(l, ": ")
		result, nums := u.Num(split[0]), u.Nums(split[1], " ")

		if recur(result, nums, 0, 0) > 0 {
			sum += result
		}
	}
	fmt.Println(sum)
}

func recur(result int, nums []int, depth int, cur int) int {
	if depth >= len(nums) {
		if cur != result {
			return 0
		}
		return 1
	}

	return recur(result, nums, depth+1, cur+nums[depth]) + recur(result, nums, depth+1, cur*nums[depth])
}
