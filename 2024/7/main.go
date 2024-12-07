package main

import (
	"adventofcode/u"
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"sync"
	"sync/atomic"
)

var availableOps = []op{
	add{},
	mul{},
	concat{},
}

type line struct {
	res  int
	nums []int
}

func main() {
	lines := u.ReadFileLines("input.txt")

	inp := make([]line, 0, len(lines))
	maxNums := 0
	for _, lin := range lines {
		split := u.Strs(lin, ": ")
		res, nums := u.Num(split[0]), u.Nums(split[1], " ")
		inp = append(inp, line{res: res, nums: nums})
		if len(nums) > maxNums {
			maxNums = len(nums)
		}
	}

	lens := make([]int, 0, maxNums)
	for range maxNums {
		lens = append(lens, 3)
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
				gen := combin.NewCartesianGenerator(lens[:len(ln.nums)])
				opsBuf := make([]int, len(ln.nums))
			GenLoop:
				for gen.Next() {
					ops := gen.Product(opsBuf)

					r := ln.nums[0]
					for i := 1; i < len(ln.nums); i++ {
						r = availableOps[ops[i-1]].Eval(r, ln.nums[i])
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
