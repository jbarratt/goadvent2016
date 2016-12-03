package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Triangle stores a triangle
type Triangle struct {
	a, b, c int
}

// Triangle.valid() tests the validity of a triangle based on the inequality
// theorem, which states that the sum of any 2 sides is greater than the
// other side.
func (t Triangle) valid() bool {
	return (t.a+t.b) > t.c && (t.b+t.c) > t.a && (t.c+t.a) > t.b
}

// parseRow takes a whitespace separated list of integers as strings
// and returns a slice of integers. Parse errors are silently ignored.
func parseRow(row string) []int {
	values := []int{}
	for _, digit := range strings.Fields(row) {
		value, err := strconv.Atoi(digit)
		if err == nil {
			values = append(values, value)
		}
	}
	return values
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrix := [][]int{}
	var possible int

	for scanner.Scan() {
		row := scanner.Text()
		matrix = append(matrix, parseRow(row))
	}
	for i := 0; i < len(matrix); i += 3 {
		for j := 0; j <= 2; j++ {
			triangle := Triangle{matrix[i][j], matrix[i+1][j], matrix[i+2][j]}
			if triangle.valid() {
				possible++
			}
		}
	}
	fmt.Printf("Possible: %d\n", possible)
}
