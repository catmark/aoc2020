package main

import (
	"fmt"
)


type Tile struct {
	Number int
	Content [][]byte
}

func (t Tile) IsNeighbour(other Tile) bool {
	for _, side := range t.getSides() {
		for _, otherSide := range other.getSides() {
			if string(side) == string(otherSide) {
				return true
			}
		}
	}
	return false
}

func (t Tile) RightNb(others []Tile) Tile {
	for _, option := range others {
		for i := 0; i < 4; i++ {
			option = option.Rotate()
			if isRightNb(t, option) {
				return option
			}
			option = option.Flip()
			if isRightNb(t, option) {
				return option
			}
		}
	}
	return Tile{}
}

func isRightNb(t, option Tile) bool {
	if t.Number == option.Number {
		return false
	}
	right := make([]byte, len(t.Content))
	for i, line := range t.Content {
		right[i] = line[len(line) - 1]
	}
	left := make([]byte, 10)
	for i, line := range option.Content {
		left[i] = line[0]
	}
	return string(right) == string(left)
}

func (t Tile) BottomNb(others []Tile) Tile {
	for _, option := range others {
		for i := 0; i < 4; i++ {
			option = option.Rotate()
			if isBottomNb(t, option) {
				return option
			}
			option = option.Flip()
			if isBottomNb(t, option) {
				return option
			}
		}
	}
	return Tile{}
}

func isBottomNb(t, option Tile) bool {
	if t.Number == option.Number {
		return false
	}
	bottom := t.Content[len(t.Content) - 1]
	top := option.Content[0]
	return string(bottom) == string(top)
}

func (t Tile) getSides() [][]byte {
	sides := [][]byte{}
	top := t.Content[0]
	sides = append(sides, top)
	sides = append(sides, revert(top))
	bottom := t.Content[len(t.Content) - 1]
	sides = append(sides, bottom)
	sides = append(sides, revert(bottom))

	left := make([]byte, len(t.Content))
	right := make([]byte, len(t.Content))
	for i, line := range t.Content {
		left[i] = line[0]
		right[i] = line[len(line) - 1]
	}

	sides = append(sides, left)
	sides = append(sides, revert(left))
	sides = append(sides, right)
	sides = append(sides, revert(right))

	return sides
}

func revert(s []byte) []byte {
	l := len(s)
	rev := make([]byte, l)
	for i, b := range s {
		rev[l - i - 1] = b
	}
	return rev
}

func (t Tile) Rotate() Tile {
	nt := Tile{}
	nt.Number = t.Number
	nt.Content = make([][]byte, 10)
	for i := range nt.Content {
		nt.Content[i] = make([]byte, 10)
	}
	for y, line := range t.Content {
		for x, bit := range line {
			nt.Content[9-x][9-y] = bit
		}
	}
	return nt
}

func (t Tile) Flip() Tile {
	nt := Tile{}
	nt.Number = t.Number
	nt.Content = make([][]byte, 10)
	for i, line := range t.Content {
		nt.Content[i] = revert(line)
	}
	return nt
}

func (t Tile) Print() {
	fmt.Println("TileNumber", t.Number)
	for _, line := range t.Content {
		fmt.Println(string(line))
	}
}
