package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numeroGerado := gerarValor()
	fmt.Println("Jogo de adivinhação \nchute um numero de 0 a 100")
	quantidadeTentativas := 0
	for i := 0; i <= 10; i++ {
		quantidadeTentativas++
		tentativa := 0
		fmt.Print("Digite o numero ")
		_, err := fmt.Scan(&tentativa)
		if err != nil {
			fmt.Println("Erro: use somente numeros de 0 a 100")
			continue
		}
		if tentativa < 0 || tentativa > 100 {
			fmt.Println("Erro: use somente numeros de 0 a 100")
			continue
		}
		if tentativa == numeroGerado {
			fmt.Println("Você acertou em", quantidadeTentativas, "tentativas")
			break
		}
		if tentativa > numeroGerado {
			fmt.Println("Muito alto, chute um numero mais baixo")
		} else {
			fmt.Println("Muito baixo, chute um numero maior")
		}
	}
}

func gerarValor() int {
	numInteiro := rand.Intn(100)
	return numInteiro
}
