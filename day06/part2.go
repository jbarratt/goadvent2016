package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func minValue(freq map[rune]int) rune {
	minVal := 500
	minRune := '*'

	for key, value := range freq {
		if value < minVal {
			minVal = value
			minRune = key
		}
	}
	return minRune
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Call with filename of input data")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	freqs := make([]map[rune]int, 8)
	output := make([]rune, 8)
	for i := 0; i < 8; i++ {
		freqs[i] = map[rune]int{}
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		for i := 0; i < 8; i++ {
			freqs[i][rune(row[i])]++
		}
	}
	for i := 0; i < 8; i++ {
		output[i] = minValue(freqs[i])
	}
	fmt.Printf("Signal: %s\n", string(output))
}
