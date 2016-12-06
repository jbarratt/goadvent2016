package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func maxValue(freq map[rune]int) rune {
	maxVal := 0
	maxRune := '*'

	for key, value := range freq {
		if value > maxVal {
			maxVal = value
			maxRune = key
		}
	}
	return maxRune
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
		output[i] = maxValue(freqs[i])
	}
	fmt.Printf("Signal: %s\n", string(output))
}
