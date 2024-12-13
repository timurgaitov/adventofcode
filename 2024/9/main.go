package main

import (
	"adventofcode/utils"
	"fmt"
)

const empty = -1

func main() {
	str := utils.ReadStr("input.txt")

	unz := unzip(str)
	ins := 0
Outer:
	for i := len(unz) - 1; i >= 0; i-- {
		if unz[i] == empty {
			continue
		}
		for unz[ins] != empty {
			ins++
			if ins >= len(unz) {
				break Outer
			}
		}
		emp := true
		for j := ins; j < len(unz); j++ {
			if unz[j] != empty {
				emp = false
				break
			}
		}
		if emp {
			break
		}
		unz[ins] = unz[i]
		unz[i] = empty
	}
	res := unz[:ins]

	checkSum := 0
	for i, n := range res {
		checkSum += i * n
	}
	fmt.Println(checkSum)

	// part 2
	files := readFiles(str)

	var res2 []file
	for i := 0; i < len(files); i++ {
		if files[i].id != empty {
			res2 = append(res2, files[i])
			continue
		}

		for j := len(files) - 1; j > i; j-- {
			if files[j].id == empty {
				continue
			}
			if files[j].len <= files[i].len {
				res2 = append(res2, file{
					pos: files[i].pos,
					id:  files[j].id,
					len: files[j].len,
				})

				files[i].len -= files[j].len
				files[i].pos += files[j].len

				files[j] = file{
					pos: files[j].pos,
					id:  empty,
					len: files[j].len,
				}
			}
		}
	}

	//printRes2(res2)

	checkSum2 := 0
	for _, r := range res2 {
		for j := range r.len {
			checkSum2 += r.id * (r.pos + j)
		}
	}
	fmt.Println(checkSum2)
}

func printRes2(files []file) {
	prevNotEmpty := 0
	for _, f := range files {
		for range f.pos - prevNotEmpty {
			fmt.Print(".")
		}
		for range f.len {
			fmt.Print(f.id)
		}
		prevNotEmpty = f.pos + f.len
	}
	fmt.Println()
}

func readFiles(str string) []file {
	unz := unzip(str)
	files := make([]file, 0)
	cur2 := unz[0]
	pos := 0
	ln := 1

	for i := 1; i < len(unz); i++ {
		if unz[i] != cur2 {
			id := -1
			if cur2 != empty {
				id = cur2
			}
			files = append(files, file{pos: pos, id: id, len: ln})
			cur2 = unz[i]
			pos = i
			ln = 0
		}
		ln++
	}
	id := -1
	if cur2 != empty {
		id = cur2
	}
	files = append(files, file{pos: pos, id: id, len: ln})
	return files
}

func unzip(str string) []int {
	res := make([]int, 0)
	cur := 0
	for i, ch := range str {
		if i%2 == 0 {
			for range ch - '0' {
				res = append(res, cur)
			}
		} else {
			for range ch - '0' {
				res = append(res, empty)
			}
			cur++
		}
	}
	return res
}

type file struct {
	pos int
	id  int
	len int
}

func (f file) String() string {
	id := fmt.Sprintf("%d", f.id)
	if id == "-1" {
		id = "."
	}
	return fmt.Sprintf("(%s, [%d, %d])", id, f.pos, f.pos+f.len-1)
}
