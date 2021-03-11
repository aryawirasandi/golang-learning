package main

import (
	"fmt"
	"strings"
)

// check vowel
func isVowel(c string) bool {
	if c == "a" || c == "i" || c == "u" || c == "e" || c == "o" || c == "A" || c == "U" || c == "E" || c == "O" {
		return true
	}
	return false
}

func Vowel(s string) string {
	k := 0

	str := strings.Split(s, "")
	vowel := ""

	// split vowel
	for i := 0; i < len(str); i++ {
		if isVowel(str[i]) {
			k++
			vowel += str[i]
		}
	}

	for i := 0; i < len(str); i++ {
		if isVowel(str[i]) {
			k--
			str[i] = string(vowel[k])
		}
	}

	return strings.Join(str, "")
}

func main() {
	fmt.Println(Vowel("hello"))
}
