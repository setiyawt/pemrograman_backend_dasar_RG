package main

import "fmt"

var math, science, english, indonesia int

func GetPredicate(math, science, english, indonesia int) string {
	rata2 := (math + science + english + indonesia) / 4
	if rata2 == 100 {
		return "Sempurna"
	} else if rata2 >= 90 && rata2 < 100 {
		return "Sangat Baik"
	} else if rata2 >= 80 && rata2 < 90 {
		return "Baik"
	} else if rata2 >= 70 && rata2 < 80 {
		return "Cukup"
	} else if rata2 >= 60 && rata2 < 70 {
		return "Kurang"
	} else if rata2 < 60 {
		return "Sangat kurang"
	} else {
		return ""
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Print("math: ")
	fmt.Scan(&math)
	fmt.Print("science: ")
	fmt.Scan(&science)
	fmt.Print("english: ")
	fmt.Scan(&english)
	fmt.Print("indonesia: ")
	fmt.Scan(&indonesia)
	result := GetPredicate(math, science, english, indonesia)
	fmt.Println(result)

}
