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

// NewTriangle returns a triangle object from a string of integer side lengths
func NewTriangle(sides string) *Triangle {
	t := Triangle{}
	for i, digit := range strings.Fields(sides) {
		value, err := strconv.Atoi(digit)
		if err == nil {
			switch i {
			case 0:
				t.a = value
			case 1:
				t.b = value
			case 2:
				t.c = value
			}
		}
	}
	return &t
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var possible int

	for scanner.Scan() {
		row := scanner.Text()
		triangle := NewTriangle(row)
		if triangle.valid() {
			possible++
		}
	}
	fmt.Printf("Possible: %d\n", possible)
}
