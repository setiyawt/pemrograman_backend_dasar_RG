package main

import "fmt"

func ReverseData(arr [5]int) [5]int {
	reversed := [5]int{}

	for i := 0; i < len(arr); i++ {
		reversed[len(arr)-1-i] = reverseNumber(arr[i])
	}

	return reversed
}

func reverseNumber(num int) int {
	reversed := 0
	for num != 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	return reversed
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(ReverseData(arr))
}
