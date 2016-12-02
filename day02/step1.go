package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

// Solution for Day 2 of the advent of code

var state = map[rune][]int {
	'U': []int{0, 1, 2, 3, 1, 2, 3, 4, 5, 9},
	'L': []int{0, 1, 1, 2, 4, 4, 5, 7, 7, 8},
	'D': []int{0, 4, 5, 6, 7, 8, 9, 7, 8, 9},
	'R': []int{0, 2, 3, 3, 5, 6, 6, 8, 9, 9},
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
		fmt.Printf("%d", digit)
	}
	fmt.Printf("\n")
}
