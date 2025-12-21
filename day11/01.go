package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var nodes map[string][]string = make(map[string][]string)
var ans int = 0

func dfs(node string) {
	if node == "out" {
		ans++
		return
	}

	for _, node := range nodes[node] {
		dfs(node)
	}
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

	dfs("you")

	fmt.Printf("The Password is %d\n", ans)
}
