package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	file, _ := os.Open("input")

	scanner := bufio.NewScanner(file)

	tiles := []Tile{}
	for scanner.Scan() {
		tile := Tile{}
		noStr := strings.Split(strings.Split(scanner.Text(), " ")[1], ":")[0]
		fmt.Println(noStr)

		tile.Number, _ = strconv.Atoi(noStr)
		for scanner.Scan() {
			if len(scanner.Text()) == 0 {
				break
			}
			tile.Content = append(tile.Content, []byte(scanner.Text()))
		}
		tiles = append(tiles, tile)
	}

	fmt.Println("Tiles", len(tiles))

	answer := int64(1)
	neighbourMap := findStartTile(tiles)
	for tile, nbCount := range neighbourMap {
		if nbCount == 2 {
			fmt.Println("corner", tile, nbCount)
			answer *= int64(tile)
		}
	}
	fmt.Println("Answer", answer)
}

func findStartTile(tiles []Tile) map[int]int {
	neighbourCountMap := map[int]int{}

	for i, tile := range tiles {
		others := []Tile{}
		for j := range tiles {
			if i != j {
				others = append(others, tiles[j])
			}
		}
		neighbours := countNeighbours(tile, others)
		neighbourCountMap[tile.Number] = neighbours
	}
	return neighbourCountMap
}

func countNeighbours(tile Tile, others []Tile) int {
	count := 0
	for _, other := range others {
		if tile.IsNeighbour(other) {
			count++
		}
	}
	return count
}


