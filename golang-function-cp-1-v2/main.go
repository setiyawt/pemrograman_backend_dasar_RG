package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	var monthName string
	switch month {
	case 1:
		monthName = "January"
	case 2:
		monthName = "February"
	case 3:
		monthName = "March"
	case 4:
		monthName = "April"
	case 5:
		monthName = "May"
	case 6:
		monthName = "June"
	case 7:
		monthName = "July"
	case 8:
		monthName = "August"
	case 9:
		monthName = "September"
	case 10:
		monthName = "October"
	case 11:
		monthName = "November"
	case 12:
		monthName = "December"
	}

	var days string
	if day < 10 {
		days = "0" + fmt.Sprint(day)
	} else {
		days = fmt.Sprint(day)
	}
	result := fmt.Sprintf("%s-%s-%d", days, monthName, year)
	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
