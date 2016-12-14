package main

import (
	"bytes"
	"fmt"
)

//go:generate stringer -type=Element

// Floors represents the number of floors
const Floors = 4

var numComponents int

// Element is the element of the generator or chip
type Element uint8

// Element types
const (
	Promethium Element = iota
	Cobalt
	Curium
	Ruthenium
	Plutonium
	Hydrogen
	Lithium
)

// Component represents a given component (generator or chip)
type Component struct {
	elem      Element
	generator bool
}

func (c Component) String() string {
	if c.generator {
		return fmt.Sprintf("%v [G]", c.elem)
	}
	return fmt.Sprintf("%v [C]", c.elem)
}

// Safe implements a test for each object, given the other objects on it's floor
// Generators are always Safe
// Chips must have a matching generator with them
func (c Component) Safe(floor []Component) bool {
	if c.generator {
		return true
	}
	for _, comp := range floor {
		if comp.generator && comp.elem == c.elem {
			return true
		}
	}
	return false
}

// GameState represents a snapshot of game state
type GameState struct {
	elevator int
	floors   [][]Component
}

func (gs GameState) String() string {
	var buffer bytes.Buffer
	for i := Floors - 1; i >= 0; i-- {
		buffer.WriteString(fmt.Sprintf("%d ", i))
		if gs.elevator == i {
			buffer.WriteString("X | ")
		} else {
			buffer.WriteString("  | ")
		}
		for _, comp := range gs.floors[i] {
			buffer.WriteString(comp.String())
			buffer.WriteString(" ")
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

// fingerprint returns a unique int that represents the game state
func (gs *GameState) fingerprint() int {
	return 0
}

// NewGameState Builds a new gamestate object
func NewGameState() *GameState {
	gs := &GameState{}
	gs.elevator = 0
	gs.floors = make([][]Component, 4)
	return gs
}

// InitialGame creates a Part 1 Game object
func InitialGame() *GameState {
	gs := NewGameState()
	gs.floors[0] = []Component{Component{Hydrogen, false}, Component{Lithium, false}}
	gs.floors[1] = []Component{Component{Hydrogen, true}}
	gs.floors[2] = []Component{Component{Lithium, true}}
	gs.floors[3] = []Component{}
	numComponents = 4
	return gs
}

func main() {
	gameState := InitialGame()
	fmt.Printf("%v\n", gameState)
}
