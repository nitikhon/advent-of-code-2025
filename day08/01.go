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

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
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

	boxes := []Box{}
	ids := []string{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), ",")

		x, err := strconv.ParseInt(line[0], 10, 64)
		if err != nil {
			log.Printf("Error convert string to int64: %v\n", err)
		}

		y, err := strconv.ParseInt(line[1], 10, 64)
		if err != nil {
			log.Printf("Error convert string to int64: %v\n", err)
		}

		z, err := strconv.ParseInt(line[2], 10, 64)
		if err != nil {
			log.Printf("Error convert string to int64: %v\n", err)
		}

		id := fmt.Sprintf("%d|%d|%d", x, y, z)
		ids = append(ids, id)

		boxes = append(boxes, Box{id: id, x: x, y: y, z: z})
	}

	pairs := [][]Box{}

	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			pairs = append(pairs, []Box{boxes[i], boxes[j]})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		dx1 := float64(pairs[i][0].x - pairs[i][1].x)
		dy1 := float64(pairs[i][0].y - pairs[i][1].y)
		dz1 := float64(pairs[i][0].z - pairs[i][1].z)
		d1 := dx1*dx1 + dy1*dy1 + dz1*dz1

		dx2 := float64(pairs[j][0].x - pairs[j][1].x)
		dy2 := float64(pairs[j][0].y - pairs[j][1].y)
		dz2 := float64(pairs[j][0].z - pairs[j][1].z)
		d2 := dx2*dx2 + dy2*dy2 + dz2*dz2

		if d1 != d2 {
			return d1 < d2
		}

		if pairs[i][0].id != pairs[j][0].id {
			return pairs[i][0].id < pairs[j][0].id
		}

		return pairs[i][1].id < pairs[j][1].id
	})

	dsu := NewDSU(ids)

	// 10 for example.txt
	// 1000 for input.txt
	for i := range 1000 {
		dsu.Union(pairs[i][0].id, pairs[i][1].id)
	}

	groups := map[string][]string{}

	for _, b := range boxes {
		root := dsu.Find(b.id)
		groups[root] = append(groups[root], b.id)
	}

	groupsSize := []int{}
	for _, group := range groups {
		groupsSize = append(groupsSize, len(group))
	}

	sort.Slice(groupsSize, func(i, j int) bool {
		return groupsSize[i] > groupsSize[j]
	})

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", groupsSize[0]*groupsSize[1]*groupsSize[2])
}
