package utils

type Pos struct {
	I, J int
}

var DirsSqClockwise = []Pos{{I: -1, J: 0}, {I: 0, J: 1}, {I: 1, J: 0}, {I: 0, J: -1}}
var DirsDiagClockwise = []Pos{{I: -1, J: 0}, {I: -1, J: 1}, {I: 0, J: 1}, {I: 1, J: 1}, {I: 1, J: 0}, {I: 1, J: -1}, {I: 0, J: -1}, {I: -1, J: -1}}
