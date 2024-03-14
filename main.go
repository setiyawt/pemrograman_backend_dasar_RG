package main

import "fmt"

func main() {

	var words string
	fmt.Print("Input words: ")
	fmt.Scanf(words)

	modified := ""

	for i := 0; i < len(words); i++ {
		input := string(words[i])
		switch input {
		case "s", "r", "z":
			modified += "l"
		case "S", "R", "Z":
			modified += "L"
		default:
			modified += input
		}
	}

	fmt.Println("Output:", modified)
}
