package main

import (
	"fmt"
)

func main() {
	fmt.Println(globalVar)
}

// функция месозем берун аз мейн
// функцияи сохтагира дар мейн вызов мекнем
var globalVar int = 1200

func someFunction() {
	fmt.Println("message from some function")
	var number int
	number = 100
	fmt.Println(number)
}
