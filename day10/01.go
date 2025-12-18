package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func isMatchIndicator(currentState, indicator string) bool {
	for i := range currentState {
		if currentState[i] != indicator[i] {
			return false
		}
	}
	return true
}

// generate all possible combinations of buttonSeqs and find the minimum operations
func solve(indicator string, buttonSeqs [][]int) int {
	minOps := math.MaxInt

	for i := 0; i < (1 << len(buttonSeqs)); i++ {
		combination := [][]int{}
		for j := range buttonSeqs {
			if (i & (1 << j)) != 0 {
				combination = append(combination, buttonSeqs[j])
			}
		}

		currentState := make([]byte, len(indicator))
		for x := range currentState {
			currentState[x] = '.'
		}

		for _, ops := range combination {
			for k := range ops {
				if currentState[ops[k]] == '.' {
					currentState[ops[k]] = '#'
				} else {
					currentState[ops[k]] = '.'
				}
			}
		}

		if isMatchIndicator(string(currentState), indicator) {
			minOps = min(minOps, len(combination))
		}
	}

	return minOps
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

	ans := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), " ")

		// trim '[' and ']'
		line[0] = line[0][1:]
		line[0] = line[0][:len(line[0])-1]

		indicator := line[0]
		buttonSeqs := [][]int{}

		// discard last element (joltages) for now
		for i := 1; i < len(line)-1; i++ {
			// trim '(' and ')'
			line[i] = line[i][1:]
			line[i] = line[i][:len(line[i])-1]
			temp := []int{}
			for v := range strings.SplitSeq(line[i], ",") {
				num, err := strconv.Atoi(string(v))
				if err != nil {
					log.Printf("Error converting string to int: %v\n", err)
					return
				}
				temp = append(temp, num)
			}
			buttonSeqs = append(buttonSeqs, temp)
		}

		// solve by bruteforce
		minOps := solve(indicator, buttonSeqs)
		ans += minOps
	}

	fmt.Printf("The Password is %d\n", ans)
}
