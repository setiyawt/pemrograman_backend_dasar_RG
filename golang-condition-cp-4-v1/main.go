package main

import "fmt"

var VIP, regular, student, day int

func GetTicketPrice(VIP, regular, student, day int) float32 {
	ticket := (VIP + regular + student)
	VIP *= 30
	regular *= 20
	student *= 10
	price := (VIP + regular + student)
	if price > 100 {
		if day%2 != 0 {
			if ticket < 5 {
				discount := float32(price) - float32(price)*0.15
				return discount
			} else {
				discount := float32(price) - float32(price)*0.25
				return discount
			}

		} else {
			if ticket < 5 {
				discount := float32(price) - float32(price)*0.10
				return discount
			} else {
				discount := float32(price) - float32(price)*0.20
				return discount
			}
		}
	} else {
		return float32(price)
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Print("VIP: ")
	fmt.Scan(&VIP)
	fmt.Print("regular: ")
	fmt.Scan(&regular)
	fmt.Print("student: ")
	fmt.Scan(&student)
	fmt.Print("day: ")
	fmt.Scan(&day)
	result := GetTicketPrice(VIP, regular, student, day)
	fmt.Println(result)
}
