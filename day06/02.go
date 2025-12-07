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
	sheet := [][]byte{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		sheet = append(sheet, []byte(line))
	}

	operations := strings.Fields(string(sheet[len(sheet)-1]))
	k := 0

	var total int64
	if operations[k] == "*" {
		total = 1
	} else {
		total = 0
	}

	for i := range len(sheet[0]) {
		space := 0

		for j := range len(sheet) - 1 {
			if sheet[j][i] == ' ' {
				space++
			}
		}

		if space != len(sheet)-1 {
			var t strings.Builder
			for j := range len(sheet) - 1 {
				if sheet[j][i] != ' ' {
					t.WriteString(string(sheet[j][i]))
				}
			}
			num, err := strconv.ParseInt(t.String(), 10, 64)
			if err != nil {
				log.Fatalf("error converting string to int64: %v\n", err)
			}
			if operations[k] == "*" {
				total *= num
			} else {
				total += num
			}
		} else {
			cnt += total
			k++
			if operations[k] == "*" {
				total = 1
			} else {
				total = 0
			}
		}
	}
	
	cnt += total

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", cnt)
}
