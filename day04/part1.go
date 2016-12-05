package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func checksum(name string) string {
	runeFreq := map[rune]int{}
	for _, letter := range name {
		if letter != '-' {
			runeFreq[letter]++
		}
	}
	pl := make(RuneCountList, len(runeFreq))
	i := 0
	for k, v := range runeFreq {
		pl[i] = RuneCount{k, v}
		i++
	}
	sort.Sort(pl)
	checkRunes := []rune{}
	for i := 0; i < 5; i++ {
		checkRunes = append(checkRunes, pl[i].Key)
	}
	return string(checkRunes)
}

// RuneCount tracks how many times a given rune has been seen
type RuneCount struct {
	Key   rune
	Value int
}

// RuneCountList is a list of rune counts
type RuneCountList []RuneCount

// Len returns the legnth of the list
func (p RuneCountList) Len() int { return len(p) }

// Swap swaps items to sort them
func (p RuneCountList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Less returns the smaller between two items
// implements the custom comparison logic
func (p RuneCountList) Less(i, j int) bool {
	// if the scores are ties fall back to normal alpabetizing
	if p[i].Value == p[j].Value {
		return p[i].Key < p[j].Key
	}
	// order reverse by count
	return p[i].Value >= p[j].Value
}

// RoomCode has information about the room in Easter Bunny HQ
type RoomCode struct {
	name     string
	sectorID int
	checksum string
}

// NewRoomCode takes a string of the room code and returns
// a populated struct or error if it cannot be parsed
func NewRoomCode(code string) (*RoomCode, error) {
	re := regexp.MustCompile(`([a-z-]+)-([0-9]+)\[([a-z]+)\]`)
	matches := re.FindStringSubmatch(code)
	if len(matches) < 4 {
		return nil, fmt.Errorf("Invalid Room Identifier: %v", matches)
	}
	value, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, errors.New("Sector id should be numeric")
	}
	return &RoomCode{name: matches[1],
		sectorID: value,
		checksum: matches[3]}, nil
}

// valid tests if a roomcode checksum matches
func (rc RoomCode) valid() bool {
	if checksum(rc.name) == rc.checksum {
		return true
	}
	return false
}

// decrypt does the ceaser cypher of every letter in the name string
// rotating by the sectorID.
func (rc RoomCode) decrypt() string {
	base := rune('a')
	decrypted := make([]rune, len(rc.name))
	for i, value := range rc.name {
		if value == '-' {
			decrypted[i] = ' '
		} else {
			decrypted[i] = (((rune(value) - base) + rune(rc.sectorID)) % rune(26)) + base
		}
	}
	return string(decrypted)
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
	total := 0

	for scanner.Scan() {
		row := scanner.Text()
		room, err := NewRoomCode(row)
		if err == nil {
			fmt.Printf("%v\n", room)
			if room.valid() {
				total += room.sectorID
				fmt.Printf("decoded: %s\n\n", room.decrypt())
			}
		} else {
			fmt.Printf("%v\n", err)
		}
	}
	fmt.Printf("Total of valid sectors: %d\n", total)
}
