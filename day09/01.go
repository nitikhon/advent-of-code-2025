package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Tile struct {
	col, row int
}

func main() {
	inputFile := flag.String("input", "example.txt", "Name of the input file")
	flag.Parse()

	fmt.Printf("Reading input from: %s\n", *inputFile)

	input, err := os.Open(*inputFile)
	if err != nil {
		log.Printf("Error opening input file: %v\n", err)
		return
	}
	defer input.Close()

	tiles := []Tile{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		row, err := strconv.Atoi(line[0])
		col, err := strconv.Atoi(line[1])
		if err != nil {
			log.Printf("Error convert row or col to number: %v\n", err)
			break
		}

		tiles = append(tiles, Tile{col: col, row: row})
	}

	// sorted by row to iterate from left part of the rectangle
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i].row < tiles[j].row
	})

	maxArea := 0
	for i := 0; i < len(tiles); i++ {
		for j := i; j < len(tiles); j++ {
			p1 := tiles[i]
			p2 := tiles[j]
			if p2.col > p1.col {
				area := (p2.row - p1.row + 1) * (p2.col - p1.col + 1)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", maxArea)
}
