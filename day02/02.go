package main

import (
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

	content, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	input := strings.Split(string(content), ",")

	var cnt int64 = 0

	for _, seq := range input {
		ids := strings.Split(seq, "-")

		start, err := strconv.ParseInt(ids[0], 10, 64)
		if err != nil {
			log.Fatalln("Error convert string to int:", err)
		}

		end, err := strconv.ParseInt(ids[1], 10, 64)
		if err != nil {
			log.Fatalln("Error convert string to int:", err)
		}

		for i := start; i <= end; i++ {
			num := strconv.FormatInt(i, 10)
			if len(num) > 1 {
				for j := 1; j < len(num); j++ {
					t := strings.Repeat(num[:j], len(num)/len(num[:j]))
					if t == num {
						cnt += i
						fmt.Println("invalid:", i)
						break
					}
				}
			}
		}
	}

	fmt.Printf("The Password is %d\n", cnt)
}
