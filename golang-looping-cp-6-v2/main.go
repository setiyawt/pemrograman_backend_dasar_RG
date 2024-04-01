package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	numStr := strconv.Itoa(numbers)
	var maxPair = 0
	var pair = 0
	for i := 1; i < len(numStr); i++ {
		firstDigit, _ := strconv.Atoi(string(numStr[i]))
		secondDigit, _ := strconv.Atoi(string(numStr[i-1]))

		sum := firstDigit + secondDigit
		if sum > maxPair {
			maxPair = sum

			pairNumb, _ := strconv.Atoi(string(numStr[i-1]) + string(numStr[i]))
			pair = pairNumb
		}
	}
	return pair
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(89083278))
}
