package main

import "fmt"

var gender string
var height int

func BMICalculator(gender string, height int) float64 {
	if gender == "laki-laki" {
		return float64(height-100) - float64(height-100)*0.1
	} else if gender == "perempuan" {
		return float64(height-100) - float64(height-100)*0.15
	} else {
		return 0
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Print("gender: ")
	fmt.Scan(&gender)
	fmt.Print("height: ")
	fmt.Scan(&height)
	result := BMICalculator(gender, height)
	fmt.Println(result)
}
