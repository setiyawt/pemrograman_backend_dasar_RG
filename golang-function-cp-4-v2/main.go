package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	var result string
	var firstMatch bool

	for _, d := range data {
		if strings.Contains(d, input) {
			if firstMatch {
				result += ","
			} else {
				firstMatch = true
			}
			result += d
		}
	}

	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("motor", "mobil APV", "mobil Avanza", "motor matic", "motor gede"))
}
