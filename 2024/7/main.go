package main

import (
	"adventofcode/u"
	"fmt"
	"slices"
)

func main() {
	lines := u.ReadFileLines("input.txt")

	inp := make(map[int][]int)
	for _, lin := range lines {
		split := u.Strs(lin, ": ")
		result, nums := u.Num(split[0]), u.Nums(split[1], " ")
		inp[result] = nums
	}

	sum := 0
	for res, nums := range inp {
		if recur(res, nums, 0, 0) > 0 {
			sum += res
		}
	}
	fmt.Println(sum)

	// part 2
	sum2 := 0
	for res, nums := range inp {
		opN := len(nums) - 1
		buf := make([]op, opN)
		coll := collector{}
		genOps(opN, buf, 0, &coll)

	GenLoop:
		for _, ops := range coll.combs {
			r := nums[0]
			for i := 1; i < len(nums); i++ {
				r = ops[i-1].Eval(r, nums[i])
			}
			if r != res {
				continue
			}
			sum2 += r
			break GenLoop
		}
	}
	// wrong 61561126041523
	fmt.Println(sum2)
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

var opt = []op{add{}, mul{}, concat{}}

func genOps(n int, buf []op, depth int, coll *collector) {
	if depth == n {
		coll.combs = append(coll.combs, slices.Clone(buf))
		return
	}
	for i := 0; i < len(opt); i++ {
		buf[depth] = opt[i]
		genOps(n, buf, depth+1, coll)
	}
}

type collector struct {
	combs [][]op
}

type op interface {
	Eval(x, y int) int
}

type add struct{}

func (_ add) Eval(x, y int) int {
	return x + y
}

type mul struct{}

func (_ mul) Eval(x, y int) int {
	return x * y
}

type concat struct{}

func (_ concat) Eval(x, y int) int {
	return u.Num(fmt.Sprintf("%d%d", x, y))
}
