package main

import (
	"fmt"
	"sort"
)

func SchedulableDays(date1 []int, date2 []int) []int {
	sort.Ints(date1)
	sort.Ints(date2)

	common := []int{}

	i, j := 0, 0
	for i < len(date1) && j < len(date2) {
		if date1[i] < date2[j] {
			i++
		} else if date1[i] > date2[j] {
			j++
		} else {
			common = append(common, date1[i])
			i++
			j++
		}
	}

	return common // TODO: replace this
}

func main() {
	fmt.Println(SchedulableDays([]int{1, 2, 3, 4}, []int{3, 4, 5}))
}
