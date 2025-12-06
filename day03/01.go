package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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

	cnt := int64(0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		m1, i1 := int64(line[0] - '0'), 0
		for i := 1; i < len(line); i++ {
			num := int64(line[i] - '0')
			if num > m1 && i + 1 < len(line) {
				m1 = num
				i1 = i
			}
		}

		m2 := int64(0)
		for i := i1 + 1; i < len(line); i++ {
			num := int64(line[i] - '0')
			if num > m2 {
				m2 = num
			}
		}

		cnt += 10 * m1 + m2
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", cnt)
}
