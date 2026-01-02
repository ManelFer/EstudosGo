// o desafio é criar uma calculadora simples que execute as quatro operações básicas: adição, subtração, multiplicação e divisão.
// depois criar um arquivo de teste para garantir que todas as operações estão funcionando corretamente.

package main

import (
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Addition: ", add(5, 3))
	fmt.Println("Subtraction: ", subtract(5, 3))
	fmt.Println("Multiplication: ", multiply(5, 3))
	result, err := divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division: ", result)
	}
}
