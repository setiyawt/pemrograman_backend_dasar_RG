package main

import "fmt"

func SchedulableDays(villager [][]int) []int {
	available := make(map[int]int)
	for _, v := range villager {
		for _, tanggal := range v {
			fmt.Println(tanggal)
			available[tanggal]++
		}
	}
	fmt.Println(available)
	result := []int{}
	for k, v := range available {
		if v == len(villager) {
			result = append(result, k)
		}
	}
	return result // TODO: replace this
}

func main() {
	fmt.Println(SchedulableDays([][]int{
		{7, 8, 9, 10, 11},
		{12, 13, 14, 15},
		{16, 17, 18, 19},
		{12, 19},
	}))
}
