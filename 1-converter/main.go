package main

import "fmt"

func main() {
	const USDToEUR float64 = 0.85
	const USDToRUB float64 = 80.0
	EURToRUB := USDToRUB / USDToEUR
	fmt.Println(EURToRUB)
}
