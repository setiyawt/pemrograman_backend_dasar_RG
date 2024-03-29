package main

import (
	"fmt"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vowelCount := 0
	consonantCount := 0

	for _, char := range str {
		switch char {
		case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
			vowelCount++
		case ' ', '\t', '\n', '\r', ',':

		default:
			consonantCount++
		}
	}

	return vowelCount, consonantCount, vowelCount == 0 || consonantCount == 0
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("SEMANGAT PAGI, itu kata orang yang baru datang ke rumahku"))
}
