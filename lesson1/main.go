package main

import (
	"fmt"
)

func main() {
	var sum float64
	fmt.Println("Введите сумму в сомони: ")
	fmt.Scan(&sum)
	const tjsToUsd = 0.091
	const tjsToRub = 8.6522
	usd := sum * tjsToUsd
	rub := sum * tjsToRub
	fmt.Printf("%.2f сомони = %.2f доллар\n", sum, usd)
	fmt.Printf("%.2f сомони = %.2f рубль", sum, rub)
}

func someFunction() {
	fmt.Println("message from some function")
	var number int
	number = 100
	fmt.Println(number)
}
