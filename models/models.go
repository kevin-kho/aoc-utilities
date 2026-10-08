package models

import "github.com/kevin-kho/aoc-utilities/common"

type Pos struct {
	X int
	Y int
}

// Calculates the Greatest Common Divisor between X and Y
// Useful for finding Primative Integer Vectors
func (p Pos) GetGcd() int {
	a, b := common.IntAbs(p.X), common.IntAbs(p.Y)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Returns the four deltas of a cartesian grid.
// Useful for DFS
func GetDeltas() []Pos {
	var res []Pos

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if common.IntAbs(x) == common.IntAbs(y) {
				continue
			}
			res = append(res, Pos{
				X: x,
				Y: y,
			})
		}
	}

	return res

}
