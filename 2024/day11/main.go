package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	nums := utils.Nums(utils.ReadStr("input.txt"), " ")

	res := 0
	blink := 75
	dp := make(map[nblink]int)
	for _, num := range nums {
		res += rec(num, blink, dp)
	}
	fmt.Println(res)
}

type nblink struct {
	n, blink int
}

func rec(n int, blink int, dp map[nblink]int) (res int) {
	if blink == 0 {
		return 1
	}
	if r, ok := dp[nblink{n, blink}]; ok {
		return r
	}
	defer func() { dp[nblink{n, blink}] = res }()

	if n == 0 {
		return rec(1, blink-1, dp)
	}

	str := fmt.Sprintf("%d", n)
	if len(str)%2 == 0 {
		return rec(utils.Num(str[:len(str)/2]), blink-1, dp) + rec(utils.Num(str[len(str)/2:]), blink-1, dp)
	}

	return rec(n*2024, blink-1, dp)
}
