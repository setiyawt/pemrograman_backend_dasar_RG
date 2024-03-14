package main

import "fmt"

var height, age int

func TicketPlayground(height, age int) int {
	if age > 12 {
		return 100000
	} else if age == 12 || height > 160 {
		return 60000
	} else if age >= 10 && age <= 11 || height > 150 {
		return 40000
	} else if age >= 8 && age <= 9 || height > 135 {
		return 25000
	} else if age >= 5 && age <= 7 || height > 120 {
		return 15000
	} else {
		return -1
	}

}

// gunakan untuk melakukan debug
func main() {
	fmt.Print("Height: ")
	fmt.Scan(&height)
	fmt.Print("Age: ")
	fmt.Scan(&age)
	result := TicketPlayground(height, age)
	fmt.Println(result)
}
