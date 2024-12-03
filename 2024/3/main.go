package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	str := utils.ReadFileStr("input.txt")
	res := 0
	do := true
	for _, groups := range utils.RegexpGroups(`(do\(\)|don't\(\)|mul\((\d+),(\d+)\))`, str) {
		if groups[0] == "do()" {
			do = true
			continue
		}
		if groups[0] == "don't()" {
			do = false
			continue
		}
		if !do {
			continue
		}
		res += utils.Num(groups[2]) * utils.Num(groups[3])
	}
	fmt.Println(res)
}
