package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"flag"
)

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
	currPos := 50 // starts at 50

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		rotation := byte(line[0])
		distance, err := strconv.Atoi(line[1:])
		distance %= 100
		if err != nil {
			log.Printf("Error convert distance to number: %v\n", err)
			break
		}
		
		if rotation == 'R' {
			if currPos + distance > 99 {
				currPos = (currPos + distance) % 100
			} else {
				currPos += distance
			}
		} else {
			if currPos - distance < 0 {
				currPos = 100 - (distance - currPos)
			} else {
				currPos -= distance
			}
		}

		if currPos == 0 {
			cnt++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", cnt)
}
