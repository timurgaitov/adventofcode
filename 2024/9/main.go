package main

import (
	"adventofcode/u"
	"fmt"
	"strings"
)

func main() {
	str := u.ReadFileStr("input.txt")
	unzip := make([]string, 0)
	cur := 0
	for i, ch := range str {
		if i%2 == 0 {
			for range ch - '0' {
				unzip = append(unzip, fmt.Sprintf("%d", cur))
			}
		} else {
			for range ch - '0' {
				unzip = append(unzip, ".")
			}
			cur++
		}
	}
	fmt.Println(strings.Join(unzip, ""))
	ins := 0
Outer:
	for i := len(unzip) - 1; i >= 0; i-- {
		if unzip[i] == "." {
			continue
		}
		for unzip[ins] != "." {
			ins++
			if ins >= len(unzip) {
				break Outer
			}
		}
		empty := true
		for j := ins; j < len(unzip); j++ {
			if unzip[j] != "." {
				empty = false
				break
			}
		}
		if empty {
			break
		}
		unzip[ins] = unzip[i]
		unzip[i] = "."
	}
	res := unzip[:ins]
	fmt.Println(strings.Join(res, ""))

	checkSum := 0
	for i, ch := range res {
		checkSum += i * u.Num(ch)
	}
	fmt.Println(checkSum)
}
