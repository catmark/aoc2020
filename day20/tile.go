package main

type Tile struct {
	Number int
	Flipped bool
	Rotation int
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

