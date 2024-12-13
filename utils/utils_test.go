package utils

import (
	"slices"
	"testing"
)

func TestRoundIndex(t *testing.T) {
	sl := []int{1, 2, 3}
	sl2 := []int{1, 2, 3, 4}
	var res []int
	var res2 []int
	for i := range 6 {
		res = append(res, RoundIndex(i, sl))
		res2 = append(res2, RoundIndex(i, sl2))
	}
	if !slices.Equal(res, []int{0, 1, 2, 0, 1, 2}) {
		t.Fail()
	}
	if !slices.Equal(res2, []int{0, 1, 2, 3, 0, 1}) {
		t.Fail()
	}
}

func TestOutOfBound(t *testing.T) {
	valMap := [][]int{
		{1, 2},
		{1, 2, 3},
		{},
	}

	testInBound := [][]int{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
		{1, 2},
	}
	for _, test := range testInBound {
		if OutOfBound(test[0], test[1], valMap) {
			t.Fail()
		}
	}

	testOutOfBound := [][]int{
		{0, -1},
		{0, 2},
		{1, -1},
		{1, 3},
		{2, 0},
		{-1, 0},
		{3, 0},
	}
	for _, test := range testOutOfBound {
		if !OutOfBound(test[0], test[1], valMap) {
			t.Fail()
		}
	}
}

func TestStrings(t *testing.T) {
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
