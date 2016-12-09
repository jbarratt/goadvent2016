package advent

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

// FilenameArg Checks for a single filename as a CLI arg and returns it
// Existing and giving primitive `usage` to the user otherwise
func FilenameArg() string {
	if len(os.Args) != 2 {
		log.Fatal("Call with filename of input data")
	}
	return os.Args[1]
}

// DropWhitespace returns a version of the string which has all whitespace
// removed.
func DropWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// ReadFile returns a file as a list of strings, all whitespace trimmed
func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return lines, nil
}
