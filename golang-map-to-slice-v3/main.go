package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	var result [][]string
	for key, value := range mapData {
		pair := []string{key, value}
		result = append(result, pair)
	}
	return result // TODO: replace this
}

func main() {
	mapData :=
		map[string]string{
			"hello": "world",
			"John":  "Doe",
			"age":   "14",
		}
	sliceData := MapToSlice(mapData)
	fmt.Println(sliceData)
}
