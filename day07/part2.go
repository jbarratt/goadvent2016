package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Supernet represents a supernet (net component outside square brackets)
type Supernet string

// Hypernet represents a hypernet (net component inside square brackets)
type Hypernet string

// parseNets extracts a list of networks (the supernet and hypernet portions)
// from an ip string.
// "abc[def]ghi" would return supernet: ["abc" "ghi"] hypernet: ["def"]
func parseNets(ip string) ([]Supernet, []Hypernet) {
	supernet := []Supernet{}
	hypernet := []Hypernet{}
	start := 0
	for idx, letter := range ip {
		switch letter {
		case '[':
			supernet = append(supernet, Supernet(ip[start:idx]))
			start = idx + 1

		case ']':
			hypernet = append(hypernet, Hypernet(ip[start:idx]))
			start = idx + 1
		}

	}
	if start < len(ip) {
		supernet = append(supernet, Supernet(ip[start:]))
	}
	return supernet, hypernet
}

// Contains checks if a given Hypernet contains a BAB
func (net Hypernet) Contains(bab string) bool {
	return strings.Contains(string(net), bab)
}

// findBab determines of a given BaB is contained in any of the
// list of Hypernets
func findBab(bab string, nets []Hypernet) bool {
	for _, net := range nets {
		if net.Contains(bab) {
			return true
		}
	}
	return false
}

// extractAbaAsBab Finds all the ABA entries in the list of supernets
// and returns a slice of BAB's for matching inside Hypernets
func extractAbaAsBab(nets []Supernet) []string {
	babs := []string{}
	for _, net := range nets {
		netMax := len(net) - 1
		for idx := range net {
			if idx+2 > netMax {
				continue
			}
			if net[idx] == net[idx+2] && net[idx] != net[idx+1] {
				babs = append(babs, string([]byte{net[idx+1], net[idx], net[idx+1]}))
			}
		}
	}
	return babs
}

// supportsTLS returns true if it passes the constraints for supporting it
func supportsSSL(ip string) bool {
	supernets, hypernets := parseNets(ip)
	for _, bab := range extractAbaAsBab(supernets) {
		if findBab(bab, hypernets) {
			return true
		}
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
		supports := supportsSSL(row)
		if supports {
			supported++
		}
		fmt.Printf("%s: %v\n", row, supports)
	}
	fmt.Printf("Total: %d\n", supported)
}
