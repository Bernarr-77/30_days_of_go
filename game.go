package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numeroGerado := gerarValor()
	fmt.Println("Jogo de adivinhação \nchute um numero de 0 a 100")
	quantidadeTentativas := 0
	for {
		quantidadeTentativas++
		tentativa := 0
		fmt.Print("Digite o numero ")
		fmt.Scan(&tentativa)
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
