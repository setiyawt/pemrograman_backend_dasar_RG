package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {
	var nameList []string
	var currentName string

	for _, char := range names {
		if char == ' ' || char == ',' || char == ';' {
			if currentName != "" {
				nameList = append(nameList, currentName)
				currentName = ""
			}
		} else {
			currentName += string(char)
		}
	}

	if currentName != "" {
		nameList = append(nameList, currentName)
	}

	shortestName := nameList[0]

	for _, name := range nameList {
		if len(name) < len(shortestName) || (len(name) == len(shortestName) && strings.Compare(name, shortestName) < 0) {
			shortestName = name
		}
	}
	return shortestName // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
