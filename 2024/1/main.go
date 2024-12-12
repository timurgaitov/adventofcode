package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
)

func main() {
	str := utils.ReadStr("input.txt")
	nums := utils.Nums(str, `\s+`)
	left := make([]int, 0, len(nums)/2)
	right := make([]int, 0, len(nums)/2)
	for i, num := range nums {
		if i%2 == 0 {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := 0; i < len(left); i++ {
		sum += utils.Abs(left[i] - right[i])
	}
	fmt.Println(sum)

	sim := 0
	for l, r := 0, 0; l < len(left) && r < len(right); {
		countL := 0
		countR := 0
		cur := left[l]
		// skip smaller numbers on the right
		for ; r < len(right) && right[r] < cur; r++ {
		}
		for ; l < len(left) && left[l] == cur; countL, l = countL+1, l+1 {
		}
		for ; r < len(right) && right[r] == cur; countR, r = countR+1, r+1 {
		}
		sim += cur * countL * countR
	}
	fmt.Println(sim)
}
