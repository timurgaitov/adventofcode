package u

import "slices"

type Pos struct {
	i, j int
}

var DirsSq = []Pos{{i: -1, j: 0}, {i: 1, j: 0}, {i: 0, j: -1}, {i: 0, j: 1}}
var DirsDiag = slices.Concat(nil, DirsSq,
	[]Pos{{i: -1, j: -1}, {i: -1, j: 1}, {i: 1, j: 1}, {i: 1, j: -1}})
