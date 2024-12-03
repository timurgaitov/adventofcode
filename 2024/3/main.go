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

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	strs := re.FindAllStringSubmatch(str, -1)
	res := 0
	for _, s := range strs {
		n1, _ := strconv.Atoi(s[1])
		n2, _ := strconv.Atoi(s[2])
		res += n1 * n2
	}
	fmt.Println(res)
}
