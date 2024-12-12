package utils

import (
	"testing"
)

func TestStrs(t *testing.T) {
	str := "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"
	res1 := Strings(str, ": ")
	if res1[0] != "Game 2" || res1[1] != "1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue" {
		t.Fail()
	}
	res2 := Strings(res1[1], "; ")
	if len(res2) != 3 || res2[0] != "1 blue, 2 green" || res2[1] != "3 green, 4 blue, 1 red" || res2[2] != "1 green, 1 blue" {
		t.Fail()
	}
	res3 := Strings(res2[1], ", ")
	if len(res3) != 3 || res3[0] != "3 green" || res3[1] != "4 blue" || res3[2] != "1 red" {
		t.Fail()
	}
}
