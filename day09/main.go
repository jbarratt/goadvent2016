package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jbarratt/goadvent2016/advent"
)

// MarkerRegexp finds a decompression block with possible additional preamble
var MarkerRegexp = regexp.MustCompile(`(\w*)\((\d+)x(\d+)\)`)

// UncompressBunny implements the decompression algorithm found
// at Easter bunny HQ
func UncompressBunny(input string) (string, error) {
	// Implementation:
	// Scan the string looking for (IxJ) blocks
	// When one is found, write unprinted data before it to the buffer
	// Write the specified number of (len) copies after it to the buffer
	// resume scanning after the matched segment
	var buffer bytes.Buffer
	// buffer.WriteString("...")
	clean := advent.DropWhitespace(input)
	handled := 0
	for {
		matches := MarkerRegexp.FindStringSubmatch(clean[handled:])
		// No more multiplier blocks in the string
		if len(matches) == 0 {
			buffer.WriteString(clean[handled:])
			return buffer.String(), nil
		}

		// copy any remaning text before the match (may be "")
		buffer.WriteString(matches[1])

		segSize, _ := strconv.Atoi(matches[2])
		repeatCount, _ := strconv.Atoi(matches[3])
		handled += len(matches[0])
		for i := 0; i < repeatCount; i++ {
			buffer.WriteString(clean[handled:(handled + segSize)])
		}
		handled += segSize
	}
}

// UncompressBunnyTwo implements the decompression algorithm found
// at Easter bunny HQ. It returns the length of the decompressed text.
func UncompressBunnyTwo(input string) int {

	clean := advent.DropWhitespace(input)
	handled := 0
	totalLength := 0
	for {
		matches := MarkerRegexp.FindStringSubmatch(clean[handled:])

		// No more multiplier blocks in the string
		if len(matches) == 0 {
			totalLength += len(clean[handled:])
			return totalLength
		}

		// account for any remaning text before the match (may be "")
		totalLength += len(matches[1])

		segSize, _ := strconv.Atoi(matches[2])
		repeatCount, _ := strconv.Atoi(matches[3])
		handled += len(matches[0])
		length := UncompressBunnyTwo(clean[handled:(handled + segSize)])

		totalLength += length * repeatCount
		handled += segSize
	}
}

func main() {
	lines, err := advent.ReadFile(advent.FilenameArg())
	if err != nil {
		panic(err)
	}
	input := strings.Join(lines, "")
	uncompressed, err := UncompressBunny(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Uncompressed V1 bytes: %d\n", len(uncompressed))
	fmt.Printf("Uncompressed V2 bytes: %d\n", UncompressBunnyTwo(input))
}
