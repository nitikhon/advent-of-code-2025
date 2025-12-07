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
	dp      [][]int64
	size    int
	beam    int
}

func (d *diagram) fireBeam(row, col int) {
	if row >= len(d.dg) || col >= len(d.dg[0]) || row < 0 || col < 0 {
		d.beam++
		return
	}

	// fmt.Println(row, col)

	if d.visited[row][col] {
		return
	}

	d.visited[row][col] = true

	if d.dg[row][col] != '^' && d.dg[row][col] != 'S' {
		d.dg[row][col] = '|'
	}

	if d.dg[row][col] == '^' {
		// if col+1 < len(d.dg[row]) {
		// 	d.dg[row][col+1] = '|'
		// }
		// if col-1 >= 0 {
		// 	d.dg[row][col-1] = '|'
		// }
		d.fireBeam(row+1, col+1)
		d.fireBeam(row+1, col-1)
	} else {
		d.fireBeam(row+1, col)
	}
}

func (d *diagram) pathCount() {
	for i := range len(d.dg) {
		for j := range len(d.dg[0]) {
			switch d.dg[i][j] {
			case 'S':
				d.dp[i][j] = 1
			case '^':
				d.dp[i][j] = d.dp[i-1][j]
			case '|':
				if d.dg[i-1][j] == '|' || d.dg[i-1][j] == 'S' {
					d.dp[i][j] += d.dp[i-1][j]
				}
				if i-1 >= 0 && j-1 >= 0 && d.dg[i-1][j-1] == '^' {
					d.dp[i][j] += d.dp[i-1][j-1]
				}
				if i-1 >= 0 && j+1 < len(d.dg[0]) && d.dg[i-1][j+1] == '^' {
					d.dp[i][j] += d.dp[i-1][j+1]
				}
			default:
				d.dp[i][j] = 0
			}
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
	cnt := int64(0)

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
	diagram.dp = make([][]int64, len(diagram.dg))
	for i := range diagram.dp {
		diagram.dp[i] = make([]int64, len(diagram.dg[0]))
	}

	diagram.fireBeam(0, s)

	diagram.pathCount()

	for i := range len(diagram.dg) {
		fmt.Println(string(diagram.dg[i]))
	}

	for i := range len(diagram.dp) {
		fmt.Println(diagram.dp[i])
	}

	for i := range len(diagram.dp[len(diagram.dp) - 1]) {
		cnt += diagram.dp[len(diagram.dp) - 1][i]
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", cnt)
}
