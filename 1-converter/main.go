package main

import "fmt"

func main() {
	var startVal string
	fmt.Printf("__ Конвертер валют __\n")

	for {
		fmt.Println("Введите начальную валюту USD (u) / EUR (e) / RUB (r): ")
		fmt.Scan(&startVal)

		if startVal != "e" && startVal != "r" && startVal != "u" {
			fmt.Println("Валюта введена неправильно!")
		} else {
			result := inputValue(startVal)
			fmt.Printf("Результат конвертации: %.2f\n", result)
		}

	}
}

func inputValue(startVal string) float64 {
	var endVal string
	var amount float64

	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	fmt.Println("Введите итоговую валюту EUR (e) или RUB (r): ")
	fmt.Scan(&endVal)

	switch {
	case startVal == "u" && endVal == "e":
		return USDtoEUR(amount)
	case startVal == "u" && endVal == "r":
		return USDtoRUB(amount)
	case startVal == "e" && endVal == "u":
		return EURtoUSD(amount)
	case startVal == "e" && endVal == "r":
		return EURtoUSD(USDtoRUB(amount))
	case startVal == "r" && endVal == "e":
		return USDtoRUB(USDtoEUR(amount))
	case startVal == "r" && endVal == "u":
		return RUBtoUSD(amount)

	default:
		fmt.Println("Ошибка: неверный ввод валют")
		return 0
	}
}

func USDtoEUR(amount float64) float64 {
	const rate = 1.13
	return amount / rate
}

func EURtoUSD(amount float64) float64 {
	const rate = 1.13
	return amount * rate
}

func USDtoRUB(amount float64) float64 {
	const rate = 80.86
	return amount * rate
}

func RUBtoUSD(amount float64) float64 {
	const rate = 80.86
	return amount / rate
}
