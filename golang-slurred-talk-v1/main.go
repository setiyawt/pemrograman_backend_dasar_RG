package main

import "fmt"

func SlurredTalk(words *string) {
	var word = *words
	var newWords string
	for _, char := range word {
		if char == 'S' || char == 'R' || char == 'Z' {
			newWords += "L"
		} else if char == 's' || char == 'r' || char == 'z' {
			newWords += "l"
		} else {
			newWords += string(char)
		}
	}
	*words = newWords
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Semangat pagi semuanya, semoga sehat selalu. Sehingga selalu bisa senyum"
	SlurredTalk(&words)
	fmt.Println(words)
}
