package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

const Dimension int = 12

func main() {
	file, _ := os.Open("input")

	scanner := bufio.NewScanner(file)

	tiles := []Tile{}
	for scanner.Scan() {
		tile := Tile{}
		noStr := strings.Split(strings.Split(scanner.Text(), " ")[1], ":")[0]

		tile.Number, _ = strconv.Atoi(noStr)
		for scanner.Scan() {
			if len(scanner.Text()) == 0 {
				break
			}
			tile.Content = append(tile.Content, []byte(scanner.Text()))
		}
		tiles = append(tiles, tile)
	}

	start := findStartTile(tiles)
	start.Print()

	next := start.RightNb(tiles)
	next.Print()

	grid := solvePuzzle(start, tiles)

	for _, line := range grid {
		fmt.Println(string(line))
	}

	snakes := countSnakes(grid)
	fmt.Println("FOund", snakes, "snakes")

	waves := countWaves(grid)

	fmt.Println("Rough", waves - 15 * snakes)

}

func findStartTile(tiles []Tile) Tile {
	for i, tile := range tiles {
		others := []Tile{}
		for j := range tiles {
			if i != j {
				others = append(others, tiles[j])
			}
		}
		neighbours := findNeighbours(tile, others)
		if len(neighbours) == 2 {
			for i := 0; i < 4; i++ {
				tile = tile.Rotate()
				r := tile.RightNb(others)
				b := tile.BottomNb(others)
				if r.Number > 0 && b.Number > 0 {
					return tile
				}
				tile = tile.Flip()
				r = tile.RightNb(others)
				b = tile.BottomNb(others)
				if r.Number > 0 && b.Number > 0 {
					return tile
				}
			}
		}
	}
	return Tile{}
}

func findNeighbours(tile Tile, others []Tile) []Tile {
	neighbours := []Tile{}
	for _, other := range others {
		if tile.IsNeighbour(other) {
			neighbours = append(neighbours, tile)
		}
	}
	return neighbours
}

func solvePuzzle(start Tile, tiles []Tile) [][]byte {
	grid := make([][]byte, 8*12)

	for i := 0; i < Dimension; i++ {
		tile := start
		for j := 0; j < Dimension; j++ {
			fmt.Println("Appending", tile.Number, start.Number)
			for y, line := range tile.Content {
				if y == 0 || y == 9 {
					continue
				}
				grid[8*i + y - 1] = append(grid[8*i + y - 1], line[1:9]...)
			}
			tile = tile.RightNb(tiles)
		}
		start = start.BottomNb(tiles)
	}
	return grid
}

func findSnake() [][]byte {
	file, _ := os.Open("snake")

	scanner := bufio.NewScanner(file)

	snake := [][]byte{}
	for scanner.Scan() {
		snake = append(snake, []byte(scanner.Text()))
	}
	return snake
}

func countSnakes(grid [][]byte) int {
	snake := findSnake()

	snakes := 0

	for y, line := range grid {
		for x, _ := range line {
			if y >= len(grid) - 2 || x >= len(grid) - 20 {
				continue
			}
			found := true
			for y1, snakeLine := range snake {
				for x1, snakeBit := range snakeLine {
					if snakeBit == '#' {
						if grid[y + y1][x + x1] != '#' {
							found = false
						}
					}
				}
			}
			if found {
				fmt.Println("Found snake at", x, y)
				snakes++
			}
		}
	}
	return snakes
}

func countWaves(grid [][]byte) int {
	count := 0
	for _, line := range grid {
		for _, bit := range line {
			if bit == '#' {
				count++
			}
		}
	}
	return count
}
