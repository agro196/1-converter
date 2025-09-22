package main

import (
	"fmt"
	"strings"
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
	)

	rates := map[string]float64{
		"USD": 1.0,
		"EUR": UsdToEur,
		"RUB": UsdToRub,
	}

	supported := keys(rates)

	fmt.Println("__Конвертер валют__")
	for {
		src := strings.ToUpper(getCurrencyInput("Введите исходную валюту", supported))
		amt := getAmountInput("Введите сумму для конвертации")
		dst := strings.ToUpper(getCurrencyInput("Введите целевую валюту", supported))

		result, err := calculateRateMap(amt, src, dst, rates)
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
		fmt.Scan(&cur)
		cur = strings.TrimSpace(cur)
		if isSupported(strings.ToUpper(cur), options) {
			return cur
		}
		fmt.Println("Неизвестная валюта. Попробуйте снова (", strings.Join(options, ", "), ").")
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

func calculateRateMap(amount float64, from, to string, rates map[string]float64) (float64, error) {
	fromRate, okFrom := rates[from]
	toRate, okTo := rates[to]
	if !okFrom || !okTo {
		return 0, fmt.Errorf("валюта не поддерживается (from=%s, to=%s)", from, to)
	}
	if from == to {
		return amount, nil
	}
	return amount / fromRate * toRate, nil
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

func keys(m map[string]float64) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
