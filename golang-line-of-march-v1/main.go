package main

import (
	"fmt"
	"sort"
)

func Sortheight(height []int) []int {
	fmt.Println("Before", height)
	sort.Ints(height)
	fmt.Println("After", height)
	return height // TODO: replace this
}

func main() {
	fmt.Println(Sortheight([]int{172, 170, 150, 165, 144, 155, 159}))
}
