package main

import "fmt"

var score, absent int

func main() {

	fmt.Print("Score: ")
	fmt.Scan(&score)
	fmt.Print("Absent: ")
	fmt.Scan(&absent)
	GraduateStudent(score, absent)
}

func GraduateStudent(score int, absent int) {
	if score >= 70 && absent <= 5 {
		fmt.Println("lulus")
	} else {
		fmt.Println("tidak lulus")
	}
}
