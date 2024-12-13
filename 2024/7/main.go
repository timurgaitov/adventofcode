package main

import (
	"adventofcode/utils"
	"fmt"
	"iter"
	"slices"
	"sync"
	"sync/atomic"
)

type line struct {
	res  int
	nums []int
}

func main() {
	lines := utils.ReadLines("input.txt")

	inp := make([]line, 0, len(lines))
	for _, lin := range lines {
		split := utils.Strings(lin, ": ")
		res, nums := utils.Num(split[0]), utils.Nums(split[1], " ")
		inp = append(inp, line{res: res, nums: nums})
	}

	sum := 0
	for _, ln := range inp {
		if recur(ln.res, ln.nums, 0, 0) > 0 {
			sum += ln.res
		}
	}
	fmt.Println(sum)

	// part 2
	sum2 := atomic.Int64{}

	concur := 12
	lineCh := make(chan line)
	wg := sync.WaitGroup{}
	wg.Add(concur)

	go func() {
		for _, ln := range inp {
			lineCh <- ln
		}
		close(lineCh)
	}()

	for range concur {
		go func() {
			defer wg.Done()
			for ln := range lineCh {
			GenLoop:
				for ops := range gen(len(ln.nums)) {
					r := ln.nums[0]
					for i := 1; i < len(ln.nums); i++ {
						r = operators[ops[i-1]](r, ln.nums[i])
					}
					if r != ln.res {
						continue
					}
					sum2.Add(int64(r))
					break GenLoop
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println(sum2.Load())
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

func gen(n int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		cur := make([]int, n)
		for {
			if !yield(slices.Clone(cur)) {
				break
			}
			mem := 1
			for i := n - 1; i >= 0; i-- {
				cur[i] += mem
				mem--
				if cur[i] < len(operators) {
					break
				}
				cur[i] = 0
				mem++
			}
			if mem == 1 {
				break
			}
		}
	}
}

var operators = []func(x, y int) int{
	add,
	mul,
	concat,
}

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func concat(x, y int) int {
	return utils.Num(fmt.Sprintf("%d%d", x, y))
}
