package main

import (
	"adventofcode/u"
	"fmt"
)

func main() {
	nums := u.Nums(u.ReadFileStr("input.txt"), " ")

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
		return rec(u.Num(str[:len(str)/2]), blink-1, dp) + rec(u.Num(str[len(str)/2:]), blink-1, dp)
	}

	return rec(n*2024, blink-1, dp)
}
