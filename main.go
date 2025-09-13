package main

import "fmt"

func main() {
	const (
		UsdToEur = 0.92
		UsdToRub = 95.50
		EurToRub = UsdToRub / UsdToEur
	)
}

func getUserInput() string {
	var currency string
	fmt.Print("Введите название валюты ")
	fmt.Scan(&currency)
	return currency
}

func calculateRate(amount float64, currency1 string, currency2 string) float64 {

}
