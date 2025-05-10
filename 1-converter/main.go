package main

import "fmt"

func main() {
	var from string
	var to string

	fmt.Printf("__ Конвертер валют __\n")
	fmt.Println("Введите начальную валюту USD (u) / EUR (e) / RUB (r): ")
	fmt.Scan(&from)

	for {

		if from != "e" && from != "r" && from != "u" {
			fmt.Println("Валюта введена неправильно!")
		} else {
			result := inputValue()
			if from == "u" {
				fmt.Println("Введите итоговую валюту EUR (e) или RUB (r): ")
			}
			if from == "r" {
				fmt.Println("Введите итоговую валюту EUR (e) или USD (u): ")
			}
			if from == "e" {
				fmt.Println("Введите итоговую валюту USD (u) или RUB (r): ")
			}

			fmt.Scan(&to)
			res := converter(from, to, result)
			fmt.Printf("Результат конвертации: %.2f\n", res)
			break
		}

	}
}

func inputValue() float64 {

	var amount float64

	for {
		fmt.Print("Введите сумму: ")
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("Ошибка: введите число!")

			var discard string
			fmt.Scanln(&discard)
			continue
		}
		return amount
	}

}

func converter(from, to string, amount float64) float64 {
	const UandErate = 1.13
	const RandUrate = 80.86

	var result float64

	switch {
	case from == "u" && to == "e":
		result = amount / UandErate
	case from == "u" && to == "r":
		result = amount * RandUrate
	case from == "e" && to == "u":
		result = amount * UandErate
	case from == "e" && to == "r":
		result = amount * UandErate * RandUrate
	case from == "r" && to == "e":
		result = amount * RandUrate * UandErate
	case from == "r" && to == "u":
		result = amount * RandUrate
	default:
		fmt.Println("Ошибка: неверный ввод валют")
		return 0
	}

	return result
}
