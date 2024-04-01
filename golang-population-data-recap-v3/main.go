package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]any {
	result := []map[string]any{}
	for _, v := range data {
		curResult := make(map[string]any)
		fmt.Println(v)
		tokens := strings.Split(v, ";")
		name := tokens[0]
		age, _ := strconv.Atoi(tokens[1])
		address := tokens[2]
		height, _ := strconv.ParseFloat(tokens[3], 64)
		isMarried, _ := strconv.ParseBool(tokens[4])
		fmt.Println("name", name)
		fmt.Println("age", age)
		fmt.Println("address", address)
		fmt.Println("height", height)
		fmt.Println("isMarried", isMarried)
		fmt.Println()

		curResult["name"] = name
		curResult["age"] = age
		curResult["address"] = address
		if tokens[3] != "" {
			curResult["height"] = height
		}
		if tokens[4] != "" {
			curResult["isMarried"] = isMarried
		}
		fmt.Println("CurResult", curResult)
		result = append(result, curResult)
	}
	return result
}

func main() {
	data := []string{"Budi;23;Jakarta;;",
		"Joko;38;Bandung;;true",
		"Susi;25;Bogor;;165.42;"}
	output := PopulationData(data)
	switch val := output[0]["age"].(type) {
	case int:
		fmt.Println(val + 1)
	}
}
