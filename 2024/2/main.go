package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	defer func() { _ = file.Close() }()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)

	count := 0
	for line, _, err := reader.ReadLine(); err == nil; line, _, err = reader.ReadLine() {
		strs := strings.Split(string(line), " ")
		nums := make([]int, 0, len(strs))
		for _, s := range strs {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}
		dir := 0
		prev := nums[0]
		unstable := false
		for i := 1; i < len(nums); i++ {
			if nums[i] < prev {
				if dir == 0 {
					dir = -1
				} else if dir == 1 {
					unstable = true
					break
				}
			} else {
				if dir == 0 {
					dir = 1
				} else if dir == -1 {
					unstable = true
					break
				}
			}
			diff := nums[i] - prev
			if diff < 0 {
				diff *= -1
			}
			if diff < 1 || diff > 3 {
				unstable = true
				break
			}
			prev = nums[i]
		}
		if unstable {
			continue
		}
		count++
	}
	fmt.Println(count)
}
