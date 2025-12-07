package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type diagram struct {
	dg      [][]byte
	visited [][]bool
	size    int
	beam    int
}

func (d *diagram) fireBeam(row, col int) {
	if row >= len(d.dg) || col >= len(d.dg[0]) || row < 0 || col < 0 {
		return
	}

	if d.visited[row][col] {
		return
	}

	d.visited[row][col] = true

	if d.dg[row][col] != '^' && d.dg[row][col] != 'S' {
		d.dg[row][col] = '|'
	}

	if d.dg[row][col] == '^' {
		d.beam++
		if col+1 < len(d.dg[row]) {
			d.dg[row][col+1] = '|'
		}
		if col-1 >= 0 {
			d.dg[row][col-1] = '|'
		}
		if row+1 < len(d.dg) && col+1 < len(d.dg[0]) {
			d.fireBeam(row+1, col+1)
		}
		if row+1 < len(d.dg) && col-1 >= 0 {
			d.fireBeam(row+1, col-1)
		}
	} else {
		if row+1 < len(d.dg) {
			d.fireBeam(row+1, col)
		}
	}
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

	diagram := diagram{}
	s := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		for i := range len(line) {
			if line[i] == 'S' {
				s = i
			}
		}
		diagram.dg = append(diagram.dg, []byte(line))
	}

	diagram.visited = make([][]bool, len(diagram.dg))
	for i := range diagram.visited {
		diagram.visited[i] = make([]bool, len(diagram.dg[0]))
	}

	diagram.fireBeam(0, s)

	for i := range len(diagram.dg) {
		fmt.Println(string(diagram.dg[i]))
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", diagram.beam)
}
