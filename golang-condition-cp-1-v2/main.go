package main

import "fmt"

var score, absent int

func GraduateStudent(score int, absent int) string {
	if score >= 70 && absent < 5 {
		return "lulus"
	} else {
		return "tidak lulus"
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Print("Score: ")
	fmt.Scan(&score)
	fmt.Print("Absent: ")
	fmt.Scan(&absent)
	result := GraduateStudent(score, absent)
	fmt.Println(result)
}
