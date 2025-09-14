package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	const (
		UsdToEur = 0.92
		UsdToRub = 95.50
		EurToRub = UsdToRub / UsdToEur
	)

	supported := []string{"USD", "EUR", "RUB"}

	fmt.Println("__Конвертер валют__")
	for {
		src := getCurrencyInput("Введите исходную валюту", supported)
		amt := getAmountInput("Введите сумму для конвертации")
		dst := getCurrencyInput("Введите целевую валюту", supported)

		result, err := calculateRate(amt, src, dst, UsdToEur, UsdToRub, EurToRub)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Printf("Результат: %.2f %s -> %.2f %s\n", amt, src, result, dst)
		}

		if !checkRepeatCalculation() {
			break
		}
	}
}

func getCurrencyInput(prompt string, options []string) string {
	for {
		fmt.Print(prompt, " (")
		printOptions(options)
		fmt.Print("): ")

		var cur string
		fmt.Scan(&cur) // читаем без TrimSpace — Scan сам уберет пробелы и \n

		if isSupported(cur, options) {
			return cur
		}
		fmt.Println("Неизвестная валюта. Попробуйте снова (USD, EUR, RUB).")
	}
}

func getAmountInput(prompt string) float64 {
	for {
		var amt float64
		fmt.Print(prompt, ": ")
		_, err := fmt.Scan(&amt)
		if err == nil && amt > 0 {
			return amt
		}
		fmt.Println("Нужно положительное число. Попробуйте снова.")
		clearStdin()
	}
}

func isSupported(cur string, options []string) bool {
	for _, v := range options {
		if cur == v {
			return true
		}
	}
	return false
}

func calculateRate(amount float64, from string, to string, UsdToEur, UsdToRub, EurToRub float64) (float64, error) {
	if from == to {
		return amount, nil
	}

	switch from {
	case "USD":
		switch to {
		case "EUR":
			return amount * UsdToEur, nil
		case "RUB":
			return amount * UsdToRub, nil
		}
	case "EUR":
		switch to {
		case "USD":
			return amount / UsdToEur, nil
		case "RUB":
			return amount * EurToRub, nil
		}
	case "RUB":
		switch to {
		case "USD":
			return amount / UsdToRub, nil
		case "EUR":
			usd := amount / UsdToRub
			return usd * UsdToEur, nil
		}
	}
	return 0, fmt.Errorf("пара %s -> %s не поддерживается", from, to)
}

func checkRepeatCalculation() bool {
	var choice string
	fmt.Print("Рассчитать ещё раз? y/n: ")
	fmt.Scan(&choice)
	return choice == "y" || choice == "Y"
}

func clearStdin() {
	var dump string
	fmt.Scanln(&dump)
}

func printOptions(opts []string) {
	for i, v := range opts {
		fmt.Print(v)
		if i != len(opts)-1 {
			fmt.Print(", ")
		}
	}
}
