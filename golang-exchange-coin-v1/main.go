package main

import "fmt"

func ExchangeCoin(amount int) []int {
	coins := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	change := make([]int, 0)

	for _, coin := range coins {
		for amount >= coin {
			change = append(change, coin)
			amount -= coin
		}
	}

	return change // TODO: replace this
}

func main() {
	fmt.Println(ExchangeCoin(0))
}
