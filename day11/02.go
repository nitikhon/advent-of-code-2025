package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var nodes map[string][]string = make(map[string][]string)
var ans int = 0

func dfs(node, target string, memo map[string]int) int {
	if node == target {
		return 1
	}

	sum := 0
	for _, neighbor := range nodes[node] {
		sum += dfs(neighbor, target, memo)
	}

	memo[node] = sum
	return sum
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

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), ":")

		node := line[0]
		neighbor := strings.Fields(line[1])

		nodes[node] = append(nodes[node], neighbor...)
	}

	var wg sync.WaitGroup
	svr2fft, fft2dac, dac2out := 0, 0, 0
	wg.Add(3)
	go func() {
		defer wg.Done()
		memo := make(map[string]int)
		svr2fft = dfs("svr", "fft", memo)
	}()
	go func() {
		defer wg.Done()
		memo := make(map[string]int)
		fft2dac = dfs("fft", "dac", memo)
	}()
	go func() {
		defer wg.Done()
		memo := make(map[string]int)
		dac2out = dfs("dac", "out", memo)
	}()
	wg.Wait()

	fmt.Printf("The Password is %d\n", svr2fft*fft2dac*dac2out)
}
