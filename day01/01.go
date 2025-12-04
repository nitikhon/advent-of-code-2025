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
	input_file := flag.String("input", "example.txt", "Name of the input file")
	flag.Parse()

	fmt.Printf("Reading input from: %s\n", *input_file)

	input, err := os.Open(*input_file)
	if err != nil {
		log.Printf("Error opening input file: %v\n", err)
		return
	}
	defer input.Close()

	zero_counter := 0
	current_positon := 50 // starts at 50

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
			if current_positon + distance > 99 {
				current_positon = (current_positon + distance) % 100
			} else {
				current_positon += distance
			}
		} else {
			if current_positon - distance < 0 {
				current_positon = 100 - (distance - current_positon)
			} else {
				current_positon -= distance
			}
		}

		if current_positon == 0 {
			zero_counter++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", zero_counter)
}
