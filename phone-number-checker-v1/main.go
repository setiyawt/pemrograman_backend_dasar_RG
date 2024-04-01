package main

import (
	"fmt"
)

func PhoneNumberChecker(number string, result *string) {

	telkNum := []string{"0811", "0812", "0813", "0814", "0815", "62811", "62812", "62813", "62814", "62815"}
	for _, num := range telkNum {
		if len(num) <= len(number) && number[:len(num)] == num {
			*result = "Telkomsel"
			return
		}

	}

	indosatNum := []string{"0816", "0817", "0818", "0819", "62816", "62817", "62817", "62818", "62819"}
	for _, num := range indosatNum {
		if len(num) <= len(number) && number[:len(num)] == num {
			*result = "Indosat"
			return
		}
	}

	xlNum := []string{"0821", "0822", "0823", "62821", "62822", "62823"}
	for _, num := range xlNum {
		if len(num) <= len(number) && number[:len(num)] == num {
			*result = "XL"
			return
		}
	}

	triNum := []string{"0827", "0828", "0829", "62827", "62828", "62829"}
	for _, num := range triNum {
		if len(num) <= len(number) && number[:len(num)] == num {
			*result = "Tri"
			return
		}
	}

	asNum := []string{"0852", "0853", "62852", "62853"}
	for _, num := range asNum {
		if len(num) <= len(number) && number[:len(num)] == num {
			*result = "AS"
			return
		}
	}

	smartfrenNum := []string{"0881", "0882", "0883", "0884", "0885", "0886", "0887", "0888", "62881", "62882", "62883", "62884", "62885", "62886", "62887", "62888"}
	for _, num := range smartfrenNum {
		if len(num) <= len(number) && number[:len(num)] == num {
			*result = "Smartfren"
			return
		}
	}

	*result = "invalid"

}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "62816111187790"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
