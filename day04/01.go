package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func isAccessible(rolls []string, x, y int) bool {
	cnt := 0
	pos := []int{-1, -1, 0, -1, 1, -1, -1, 0, 1, 0, -1, 1, 0, 1, 1, 1}
	for i := 0; i < len(pos); i += 2 {
		// fmt.Println(i, rolls[x][y])
		if 	x + pos[i] >= 0 && x + pos[i] < len(rolls) &&
			y + pos[i+1] >= 0 && y + pos[i+1] < len(rolls) && 
			rolls[x+pos[i]][y+pos[i+1]] == '@' {
			cnt++
		}
	}
	return cnt < 4
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

	cnt := 0
	rolls := []string{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		rolls = append(rolls, line)
	}

	for {
		for i := range len(rolls) {
			for j := range len(rolls[i]) {
				if rolls[i][j] == '@' && isAccessible(rolls, i, j) {
					cnt++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", cnt)
}
