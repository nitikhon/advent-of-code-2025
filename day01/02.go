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
	current_posisiton := 50 // starts at 50

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		rotation := byte(line[0])
		distance, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Printf("Error convert distance to number: %v\n", err)
			break
		}

		if rotation == 'R' {
			if current_posisiton + distance > 99 {
				zero_counter += (distance + current_posisiton) / 100
				current_posisiton = ((current_posisiton + distance) % 100 + 100) % 100
			} else {
				current_posisiton += distance
			}
		} else {
			if current_posisiton - distance < 0 {
				if current_posisiton > 0 {
					zero_counter += 1 + (distance - current_posisiton) / 100
				} else {
					zero_counter += distance / 100
				}
				current_posisiton = ((current_posisiton - distance) % 100 + 100) % 100
			} else {
				if current_posisiton == distance {
					zero_counter++
				}
				current_posisiton -= distance
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", zero_counter)
}
