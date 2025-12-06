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

	cnt := 0
	ranges := []string{}
	ids := []int64{}
	isRange := true

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isRange = false
			continue
		}

		if isRange {
			ranges = append(ranges, line)
		} else {
			intID, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatalf("Error convert string id to int64 id:%v\n", err)
			}
			ids = append(ids, intID)
		}
	}

	for i := range ids {
		for j := range ranges {
			t := strings.Split(ranges[j], "-")

			start, err := strconv.ParseInt(t[0], 10, 64)
			if err != nil {
				log.Fatalf("Error convert string id to int64 id:%v\n", err)
			}

			end, err := strconv.ParseInt(t[1], 10, 64)
			if err != nil {
				log.Fatalf("Error convert string id to int64 id:%v\n", err)
			}

			if ids[i] >= start && ids[i] <= end {
				cnt++
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", cnt)
}
