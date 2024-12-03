package main

import (
	"adventofcode/u"
	"fmt"
	"slices"
)

func main() {
	lines := u.ReadFileLines("input.txt")
	count := 0
	for _, line := range lines {
		nums := u.Nums(line, " ")
		for without := 0; without < len(nums); without++ {
			nums2 := slices.Concat(nums[:without], nums[without+1:])
			if isStable(nums2) {
				count++
				break
			}
		}
	}
	fmt.Println(count)
}

func isStable(nums []int) bool {
	dir := 0
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < prev {
			if dir == 0 {
				dir = -1
			} else if dir == 1 {
				return false
			}
		} else {
			if dir == 0 {
				dir = 1
			} else if dir == -1 {
				return false
			}
		}
		diff := u.Abs(nums[i] - prev)
		if diff < 1 || diff > 3 {
			return false
		}
		prev = nums[i]
	}
	return true
}
