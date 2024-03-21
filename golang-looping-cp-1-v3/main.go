package main

import "fmt"

var n int

func CountingNumber(n int) float64 {
	var total float64 = 0
	for i := 1; i <= n; i++ {
		add := float64(i)
		total += add
		if i < n {
			result := add + 0.5
			total += result
		}
	}
	return total
}

// gunakan untuk melakukan debug
func main() {
	fmt.Print("n: ")
	fmt.Scan(&n)
	fmt.Println(CountingNumber(n))
}
