package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	left, right := readNums()
	slices.Sort(left)
	slices.Sort(right)
	if len(left) != len(right) {
		panic("lengths do not match")
	}
	// part 1
	sum := 0
	for i := 0; i < len(left); i++ {
		sum += distance(left[i], right[i])
	}
	fmt.Println(sum)
	// part 2
	sim := 0
	for l, r := 0, 0; l < len(left) && r < len(right); {
		countL := 0
		countR := 0
		cur := left[l]
		// skip smaller numbers on the right
		for ; r < len(right) && right[r] < cur; r++ {
		}
		for ; l < len(left) && left[l] == cur; countL, l = countL+1, l+1 {
		}
		for ; r < len(right) && right[r] == cur; countR, r = countR+1, r+1 {
		}
		sim += cur * countL * countR
	}
	fmt.Println(sim)
}

func distance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func readNums() (left []int, right []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)
	r := false
	for scan.Scan() {
		num := bytesToNum(scan.Bytes())
		if r {
			right = append(right, num)
		} else {
			left = append(left, num)
		}
		r = !r
	}
	return
}

func bytesToNum(bytes []byte) int {
	var num int
	cur := 1
	for i := len(bytes) - 1; i >= 0; i-- {
		num += int(bytes[i]-'0') * cur
		cur *= 10
	}
	return num
}
