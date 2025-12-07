package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	numSheet := [][]int64{}
	opSheet := []string{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		tempStr := strings.Fields(line)

		if line[0] != '*' && line[0] != '+' {
			tempInt := []int64{}
			for i := range tempStr {
				num, err := strconv.ParseInt(tempStr[i], 10, 64)
				if err != nil {
					log.Fatalf("error convert string to int64: %v\n", err)
				}
				tempInt = append(tempInt, num)
			}
			numSheet = append(numSheet, tempInt)
		} else {
			opSheet = append(opSheet, tempStr...)
		}

	}

	for i := range opSheet {
		var total int64
		if opSheet[i] == "*" {
			total = 1
		} else {
			total = 0
		}

		for j := range numSheet {
			switch opSheet[i] {
			case "+":
				total += numSheet[j][i]
			case "*":
				total *= numSheet[j][i]
			}
		}

		cnt += total
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", cnt)
}
