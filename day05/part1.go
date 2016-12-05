package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

func main() {
	base := []byte("cxdnnyjw")

	passwordLength := 8
	found := 0
	password := make([]rune, passwordLength)
	for i := 0; i < passwordLength; i++ {
		password[i] = rune('_')
	}

	for index := 0; true; index++ {

		// Get the md5sum into a string
		key := append(base, []byte(strconv.Itoa(index))...)
		stringSum := fmt.Sprintf("%x", md5.Sum(key))

		// Check if that string meets the criteria
		if stringSum[:5] == "00000" {

			fmt.Printf("%s: %s\n", key, stringSum)
			slot, err := strconv.Atoi(string(stringSum[5]))
			if err != nil || slot > (passwordLength-1) || password[slot] != rune('_') {
				continue
			}
			password[slot] = rune(stringSum[6])
			fmt.Printf("%s\n", string(password))
			found++
			if found == passwordLength {
				fmt.Printf("Found password: %s\n", string(password))
				os.Exit(0)
			}
		}
	}
}
