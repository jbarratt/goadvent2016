package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Solution for Day 2 of the advent of code

var state = map[rune][]int{
	'U': []int{0, 1, 2, 1, 4, 5, 2, 3, 4, 9, 6, 7, 8, 11},
	'L': []int{0, 1, 2, 2, 3, 5, 5, 6, 7, 8, 10, 10, 11, 13},
	'D': []int{0, 3, 6, 7, 8, 5, 10, 11, 12, 9, 10, 13, 12, 13},
	'R': []int{0, 1, 3, 4, 4, 6, 7, 8, 9, 9, 11, 12, 12, 13},
}

func main() {
	// codes := [...]string{"ULL", "RRDDD", "LURDL", "UUUUD"}
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var digit = 5
	for scanner.Scan() {
		row := scanner.Text()
		for direction := range row {
			digit = state[rune(row[direction])][digit]
		}
		fmt.Printf("%x", digit)
	}
	fmt.Printf("\n")
}
