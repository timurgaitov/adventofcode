package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	str := string(bytes)

	re := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d+),(\d+)\))`)
	strs := re.FindAllStringSubmatch(str, -1)
	res := 0
	do := true
	for _, s := range strs {
		if s[0] == "do()" {
			do = true
		} else if s[0] == "don't()" {
			do = false
		}
		if !do {
			continue
		}
		n1, _ := strconv.Atoi(s[2])
		n2, _ := strconv.Atoi(s[3])
		res += n1 * n2
	}
	fmt.Println(res)
}
