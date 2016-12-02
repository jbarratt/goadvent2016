package main

import (
	"fmt"
	"strings"
	"log"
	"strconv"
	"math"
)

type Heading uint8
const (
	NORTH Heading = iota
	EAST
	WEST
	SOUTH
)

type Rotation uint8
const (
	LEFT Rotation = iota
	RIGHT
)

type Position struct {
	x, y int
}

func parseInstruction(ins string) (Rotation, int) {
	var rot Rotation
	if(string(ins[0]) == "R") {
		rot = RIGHT
	} else if(string(ins[0]) == "L") {
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

func (pos *Position) update(head Heading, distance int, visited map[Position]bool) (bool) {
	for i:= 1; i <= distance; i++ {
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
		} else {
			visited[*pos] = true
		}
	}
	return false
}


func (pos Position) distance() int {
	return int(math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
}

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

func main() {
	paths := [...]string{"R8, R4, R4, R8", "R2, L3", "R2, R2, R2", "R5, L5, R5, R3", "L5, R1, R4, L5, L4, R3, R1, L1, R4, R5, L1, L3, R4, L2, L4, R2, L4, L1, R3, R1, R1, L1, R1, L5, R5, R2, L5, R2, R1, L2, L4, L4, R191, R2, R5, R1, L1, L2, R5, L2, L3, R4, L1, L1, R1, R50, L1, R1, R76, R5, R4, R2, L5, L3, L5, R2, R1, L1, R2, L3, R4, R2, L1, L1, R4, L1, L1, R185, R1, L5, L4, L5, L3, R2, R3, R1, L5, R1, L3, L2, L2, R5, L1, L1, L3, R1, R4, L2, L1, L1, L3, L4, R5, L2, R3, R5, R1, L4, R5, L3, R3, R3, R1, R1, R5, R2, L2, R5, L5, L4, R4, R3, R5, R1, L3, R1, L2, L2, R3, R4, L1, R4, L1, R4, R3, L1, L4, L1, L5, L2, R2, L1, R1, L5, L3, R4, L1, R5, L5, L5, L1, L3, R1, R5, L2, L4, L5, L1, L1, L2, R5, R5, L4, R3, L2, L1, L3, L4, L5, L5, L2, R4, R3, L5, R4, R2, R1, L5"}
	for i := range paths {
		fmt.Printf("Path: '%v'\nDistance: %v\n\n", paths[i], findDistance(paths[i]))
	}

}
