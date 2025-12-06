package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func int64Pow10(m int64) int64 {
    if m == 0 {
        return 1
    }

    if m == 1 {
        return 10
    }

    result := int64(10)
    for i := int64(2); i <= m; i++ {
        result *= 10
    }
    return result
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

	cnt := int64(0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		m := int64(line[0] - '0')
		batteries := []int64{m}

		for i := 1; i < len(line); i++ {
			num := int64(line[i] - '0')
			m = batteries[len(batteries)-1]
			for len(batteries) > 0 && num > batteries[len(batteries)-1] && len(batteries)-1+len(line)-i >= 12 {
				batteries = batteries[:len(batteries)-1]
			}

			if (len(batteries) < 12){
				batteries = append(batteries, num)
			}
		}

		j := int64(12)
		for i := range 12 {
			cnt += batteries[i] * int64Pow10(j)
			j--
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error during scanning: %v\n", err)
	}

	fmt.Printf("The Password is %d\n", cnt)
}
