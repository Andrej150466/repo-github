package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, res float32
	var op string

	fmt.Print("Введите первое число: ")
	fmt.ScanLn(&a)

	fmt.Print("Введите второе число: ")
	fmt.ScanLn(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /): ")
	fmt.ScanLn(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		fmt.PrintLn("Операция выбрана неверно")
		os.Exit(1)
	}
		fmt.Printf("Результат выполнения операции: %f\n", res)
}