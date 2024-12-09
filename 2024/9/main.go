package main

import (
	"adventofcode/u"
	"fmt"
	"slices"
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
	unzip2 := slices.Clone(unzip)
	//fmt.Println(strings.Join(unzip, ""))
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
	//fmt.Println(strings.Join(res, ""))

	checkSum := 0
	for i, ch := range res {
		checkSum += i * u.Num(ch)
	}
	//fmt.Println(checkSum)

	// part 2
	files := make([]file, 0)
	cur2 := unzip2[0]
	pos := 0
	ln := 1
	//fmt.Println(unzip2)

	for i := 1; i < len(unzip2); i++ {
		if unzip2[i] != cur2 {
			id := -1
			if cur2 != "." {
				id = u.Num(cur2)
			}
			files = append(files, file{pos: pos, id: id, len: ln})
			cur2 = unzip2[i]
			pos = i
			ln = 0
		}
		ln++
	}
	id := -1
	if cur2 != "." {
		id = u.Num(cur2)
	}
	files = append(files, file{pos: pos, id: id, len: ln})

	var res2 []file

	for i := 0; i < len(files); i++ {
		if files[i].id != -1 {
			res2 = append(res2, files[i])
			continue
		}
		emptySpace := files[i]

		for j := len(files) - 1; j >= 0; j-- {
			if files[j].id == -1 {
				continue
			}
			if files[j].len <= emptySpace.len {
				files[j].pos = emptySpace.pos
				res2 = append(res2, files[j])
				emptySpace.len -= files[j].len
				emptySpace.pos += files[j].len
				files[j] = file{id: -1, len: -1, pos: -1}
			}
		}
	}

	//fmt.Println(res2)

	checkSum2 := 0
	for _, r := range res2 {
		for j := range r.len {
			checkSum2 += r.id * (r.pos + j)
		}
	}
	fmt.Println(checkSum2)
	//10733824823427 wrong
}

type file struct {
	pos int
	id  int
	len int
}
