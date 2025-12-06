package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type interval struct {
	start, end int64
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

	ranges := []interval{}
	cnt := int64(0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		t := strings.Split(line, "-")

		start, err := strconv.ParseInt(t[0], 10, 64)
		if err != nil {
			log.Fatalf("Error convert string id to int64 id:%v\n", err)
		}

		end, err := strconv.ParseInt(t[1], 10, 64)
		if err != nil {
			log.Fatalf("Error convert string id to int64 id:%v\n", err)
		}

		ranges = append(ranges, interval{start: start, end: end})
	}

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].start < ranges[j].start {
			return ranges[i].start < ranges[j].start
		}
		return ranges[i].end < ranges[j].end
	})

	curr := ranges[0]

	for i := 1; i < len(ranges); i++ {
		if ranges[i].start <= curr.end {
			curr.end = max(curr.end, ranges[i].end)
		} else {
			cnt += curr.end - curr.start + 1
			curr = ranges[i]
		}
	}
	
	cnt += curr.end - curr.start + 1

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", cnt)
}
