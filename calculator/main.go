package main

import "fmt"

func minus(a int, b int) int {
	return a - b
}

func plus(a int, b int) int {
	return a + b
}

func divide(a int, b int) int {
	return a / b
}

func multiply(a int, b int) int {
	return a * b
}

func calculator(a int, b int, operation string) {
	var value int
	switch operation {
	case "+":
		value = plus(a, b)
	case "-":
		value = minus(a, b)
	case "*":
		value = multiply(a, b)
	case "/":
		value = divide(a, b)
	}

	fmt.Println(value)
}

func main() {
	var a int
	var b int
	var op string
	fmt.Println("********************")
	fmt.Println("Calculator 1.0")
	fmt.Println("********************")
	fmt.Println("Type a number: ")
	fmt.Scan(&a) //& hace referencia a la direccion de memoria donde va a guarrdar el valor
	fmt.Println("Type a number: ")
	fmt.Scan(&b) //& hace referencia a la direccion de memoria donde va a guarrdar el valor
	fmt.Println("Type an operation: + - * / ")
	fmt.Scan(&op) //& hace referencia a la direccion de memoria donde va a guarrdar el valor

	calculator(a, b, op)

}
