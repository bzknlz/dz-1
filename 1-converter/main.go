package main

import "fmt"

func main() {
	exchangeInfo := inputExchangeInfo()
	fmt.Print(exchangeInfo)
}

func inputExchangeInfo() string {
	fmt.Println("Введите информацию о валютном обмене:")
	var exchangeInfo string
	fmt.Scan(&exchangeInfo)
	return exchangeInfo
}

// func convertCurrency(amount float64, fromCurrency string, toCurrency string) float64 {
// }
