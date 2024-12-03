package main

import (
	"adventofcode/u"
	"fmt"
)

func main() {
	str := u.ReadFileStr("input.txt")
	res := 0
	do := true
	for _, groups := range u.RegexpGroups(str, `(do\(\)|don't\(\)|mul\((\d+),(\d+)\))`) {
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
		res += u.Num(groups[2]) * u.Num(groups[3])
	}
	fmt.Println(res)
}
