package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Heading tracks the cardinal direction the avatar is heading.
type Heading uint8

// NORTH, etc, (Compass Directions)
const (
	NORTH Heading = iota
	EAST
	WEST
	SOUTH
)

// Rotation identifies the direction (left or right) of a navigational step
type Rotation uint8

// LEFT, etc (Which direction we turn)
const (
	LEFT Rotation = iota
	RIGHT
)

// Position is the x, y location on the street grid the avatar is located at.
type Position struct {
	x, y int
}

// parseInstruction takes a single token ('L20' or 'R2')
// and returns a Rotation type and a distance to travel after rotating that direction
func parseInstruction(ins string) (Rotation, int) {
	var rot Rotation
	if string(ins[0]) == "R" {
		rot = RIGHT
	} else if string(ins[0]) == "L" {
		rot = LEFT
	} else {
		log.Fatal("Unable to parse instruction")
	}
	length, err := strconv.Atoi(ins[1:])
	if err != nil {
		log.Fatal(err)
	}
	return rot, length
}

// Heading.update takes a heading struct and a rotation, and sets a new heading
// based on it.
func (head *Heading) update(rot Rotation) {
	switch rot {
	case LEFT:
		switch *head {
		case NORTH:
			*head = WEST
		case WEST:
			*head = SOUTH
		case SOUTH:
			*head = EAST
		case EAST:
			*head = NORTH
		}
	case RIGHT:
		switch *head {
		case NORTH:
			*head = EAST
		case WEST:
			*head = NORTH
		case SOUTH:
			*head = WEST
		case EAST:
			*head = SOUTH
		}
	}
}

// Position.update takes a heading, a distance to travel, and a map of previously visited points.
// It updates the position based on heading and distance, and adds breadcrumbs at every
// point of travel in the `visited` map.
// If our journey ever overlaps a spot we have previously visited, stop the presses!
// Return `true` to indicate that the destination has been reached.
func (pos *Position) update(head Heading, distance int, visited map[Position]bool) bool {
	for i := 1; i <= distance; i++ {
		switch head {
		case NORTH:
			pos.y++
		case SOUTH:
			pos.y--
		case EAST:
			pos.x++
		case WEST:
			pos.x--
		}
		if visited[*pos] {
			return true
		}
		visited[*pos] = true
	}
	return false
}

// Position.distance returns manhattan distance of a Position object from the origin.
func (pos Position) distance() int {
	return int(math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
}

// findDistance takes a path string and returns an integer value.
// It implements the core algorithm to solve the first day of the Advent Of Code 2016
func findDistance(path string) int {
	var position Position
	var heading Heading
	var visited = make(map[Position]bool)
	visited[position] = true

	result := strings.Split(path, ", ")

	for i := range result {
		rotation, distance := parseInstruction(result[i])
		heading.update(rotation)
		atdestination := position.update(heading, distance, visited)
		if atdestination {
			return position.distance()
		}
	}
	return position.distance()
}

// main() includes pre-baked versions of the strings that were given as examples, as well as the real string.
func main() {
	paths := [...]string{"R8, R4, R4, R8", "R2, L3", "R2, R2, R2", "R5, L5, R5, R3", "L5, R1, R4, L5, L4, R3, R1, L1, R4, R5, L1, L3, R4, L2, L4, R2, L4, L1, R3, R1, R1, L1, R1, L5, R5, R2, L5, R2, R1, L2, L4, L4, R191, R2, R5, R1, L1, L2, R5, L2, L3, R4, L1, L1, R1, R50, L1, R1, R76, R5, R4, R2, L5, L3, L5, R2, R1, L1, R2, L3, R4, R2, L1, L1, R4, L1, L1, R185, R1, L5, L4, L5, L3, R2, R3, R1, L5, R1, L3, L2, L2, R5, L1, L1, L3, R1, R4, L2, L1, L1, L3, L4, R5, L2, R3, R5, R1, L4, R5, L3, R3, R3, R1, R1, R5, R2, L2, R5, L5, L4, R4, R3, R5, R1, L3, R1, L2, L2, R3, R4, L1, R4, L1, R4, R3, L1, L4, L1, L5, L2, R2, L1, R1, L5, L3, R4, L1, R5, L5, L5, L1, L3, R1, R5, L2, L4, L5, L1, L1, L2, R5, R5, L4, R3, L2, L1, L3, L4, L5, L5, L2, R4, R3, L5, R4, R2, R1, L5"}
	for i := range paths {
		fmt.Printf("Path: '%v'\nDistance: %v\n\n", paths[i], findDistance(paths[i]))
	}

}
