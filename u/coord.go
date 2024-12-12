package u

import "slices"

type Pos struct {
	I, J int
}

var DirsSq = []Pos{{I: -1, J: 0}, {I: 1, J: 0}, {I: 0, J: -1}, {I: 0, J: 1}}
var DirsDiag = slices.Concat(nil, DirsSq,
	[]Pos{{I: -1, J: -1}, {I: -1, J: 1}, {I: 1, J: 1}, {I: 1, J: -1}})
