package main

import (
	"adventofcode/u"
	"fmt"
)

func main() {
	nums := u.Nums(u.ReadFileStr("input.txt"), " ")

	for blink := 0; blink < 25; blink++ {
		nextNums := make([]int, 0)
		for _, n := range nums {
			if n == 0 {
				nextNums = append(nextNums, 1)
				continue
			}

			str := fmt.Sprintf("%d", n)
			if len(str)%2 == 0 {
				nextNums = append(nextNums, u.Num(str[:len(str)/2]), u.Num(str[len(str)/2:]))
				continue
			}

			nextNums = append(nextNums, n*2024)
		}
		nums = nextNums
	}
	fmt.Println(len(nums))
}
