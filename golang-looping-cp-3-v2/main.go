package main

import "fmt"

func CountingLetter(text string) int {
	var count = 0

	for i := 0; i < len(text); i++ {
		if text[i] == 'R' || text[i] == 'S' || text[i] == 'T' || text[i] == 'Z' || text[i] == 'r' || text[i] == 's' || text[i] == 't' || text[i] == 'z' {
			count++
		}
	}
	return count // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
