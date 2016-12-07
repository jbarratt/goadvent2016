package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// supportsTLS returns true if it passes the constraints for supporting it
func supportsTLS(ip string) bool {
	runes := []rune(ip)
	max := len(runes) - 1
	brackets := 0
	foundAbba := false
	for idx, letter := range runes {
		switch letter {
		case '[':
			brackets++
		case ']':
			if brackets > 0 {
				brackets--
			}
		default:
			if (idx + 3) > max {
				break
			}
			if runes[idx] == runes[idx+3] && runes[idx+1] == runes[idx+2] && runes[idx] != runes[idx+1] {
				// OMG WE FOUND ABBA
				if brackets > 0 {
					return false
				}
				foundAbba = true
			}
		}
	}
	if brackets == 0 && foundAbba {
		return true
	}
	return false
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
	supported := 0

	for scanner.Scan() {
		row := scanner.Text()
		supports := supportsTLS(row)
		if supports {
			supported++
		}
		fmt.Printf("%s: %v\n", row, supports)
	}
	fmt.Printf("Total: %d\n", supported)
}
