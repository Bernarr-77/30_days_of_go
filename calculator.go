package main

import (
	"fmt"
)

func main() {
	var num1 float64
	var num2 float64
	var operador string
	fmt.Println("Calculadora")
	fmt.Print("Digite o primeiro numero: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o operador (+,-,*,/): ")
	fmt.Scan(&operador)
	fmt.Print("Digite o segundo numero: ")
	fmt.Scan(&num2)

	var resultado float64

	switch operador {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Erro, Divisão por 0 nao é valida")
			return
		}
		resultado = num1 / num2
	default:
		fmt.Println("Erro: Operador inválido!")
		return
	}
	fmt.Println("A soma de ", num1, operador, num2, "=", resultado)
}
