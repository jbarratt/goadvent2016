package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Screen Type. Index the display x, y
type Screen struct {
	display [][]bool
	width   int
	height  int
}

func (s *Screen) rect(wide int, high int) {
	if wide > s.width || high > s.height {
		log.Fatalf("Can't draw bigger than the screen")
	}
	for x := 0; x < wide; x++ {
		for y := 0; y < high; y++ {
			s.display[x][y] = true
		}
	}
}

func (s *Screen) rotate(direction string, index int, value int) {
	if direction == "row" {
		s.rotateRow(index, value)
	} else {
		s.rotateColumn(index, value)
	}
}

func (s *Screen) rotateRow(index int, value int) {
	clone := make([]bool, s.width)

	for x := 0; x < s.width; x++ {
		clone[x] = s.display[x][index]
	}

	for x := 0; x < s.width; x++ {
		s.display[(x+value)%s.width][index] = clone[x]
	}

}

func (s *Screen) rotateColumn(index int, value int) {
	clone := make([]bool, s.height)

	for y := 0; y < s.height; y++ {
		clone[y] = s.display[index][y]
	}

	for y := 0; y < s.height; y++ {
		s.display[index][(y+value)%s.height] = clone[y]
	}
}

func (s *Screen) command(cmd string) {
	rect := regexp.MustCompile(`rect (\d+)x(\d+)`)
	rotate := regexp.MustCompile(`rotate (\w+) (\w)=(\d+) by (\d+)`)
	matches := rect.FindStringSubmatch(cmd)
	// Check to see if it's a rect command
	if len(matches) == 3 {
		wide, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatalf("'%s' is not a valid width", matches[1])
		}
		high, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalf("'%s' is not a valid height", matches[2])
		}
		s.rect(wide, high)
	}

	rotate = regexp.MustCompile(`rotate (\w+) \w=(\d+) by (\d+)`)
	matches = rotate.FindStringSubmatch(cmd)
	if len(matches) == 4 {
		index, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalf("'%s' is not a valid column", matches[2])
		}
		value, err := strconv.Atoi(matches[3])
		if err != nil {
			log.Fatalf("'%s' is not a valid value", matches[3])
		}
		s.rotate(matches[1], index, value)
	}
}

func (s *Screen) litPixels() int {
	lit := 0
	for x := 0; x < s.width; x++ {
		for y := 0; y < s.height; y++ {
			if s.display[x][y] {
				lit++
			}
		}
	}
	return lit
}

func (s *Screen) print() {
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			if s.display[x][y] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// NewScreen creates a new Screen object given a height and width
func NewScreen(width int, height int) *Screen {
	s := Screen{}
	s.display = make([][]bool, width)
	for i := range s.display {
		s.display[i] = make([]bool, height)
	}
	s.height = height
	s.width = width
	return &s
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

	scanner := bufio.NewScanner(file)
	screen := NewScreen(50, 6)

	for scanner.Scan() {
		row := scanner.Text()
		screen.command(row)
		fmt.Printf("%s\n", row)
		screen.print()
	}
	fmt.Printf("Pixels lit: %d\n", screen.litPixels())
}
