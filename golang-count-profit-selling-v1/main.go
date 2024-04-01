package main

import "fmt"

func CountProfit(data [][][2]int) []int {
	resultMap := make(map[int]int)
	for _, cabang := range data {
		fmt.Println("cabang", cabang)
		for bulanKe, bulan := range cabang {
			fmt.Println("bulan", bulan)

			income := bulan[0]
			expense := bulan[1]

			fmt.Println("income", income)
			fmt.Println("expense", expense)

			profit := income - expense
			resultMap[bulanKe+1] += profit
		}
		fmt.Println()
	}
	numOfMonth := 0
	for k := range resultMap {
		if k > numOfMonth {
			numOfMonth = k
		}
	}
	result := make([]int, numOfMonth)
	for k, v := range resultMap {
		result[k-1] = v
	}
	return result // TODO: replace this
}

func main() {
	fmt.Println(CountProfit([][][2]int{
		{{1000, 200}},
		{{500, 100}},
		{{450, 150}},
		{{100, 50}},
	}))
}
