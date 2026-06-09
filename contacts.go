package main

import "fmt"

func main() {
	agenda := make(map[string]string)
	agenda["joao"] = "99999-1111"
	fmt.Println("Agenda de contatos")
	var opcao int
loopPrincipal:
	for {
		fmt.Println("Digite o numero da opção que vc deseja:")
		fmt.Println("1 - Adicionar")
		fmt.Println("2 - buscar")
		fmt.Println("3 - apagar")
		fmt.Println("4 - listar todos")
		fmt.Println("5 - sair")
		fmt.Print("Digite: ")
		fmt.Scan(&opcao)
		if opcao < 1 || opcao > 5 {
			fmt.Println("Opção inválida! Digite um número de 1 a 5")
			continue
		}
		switch opcao {
		case 1:
			adicionarContato(agenda)
		case 2:
			buscaContato(agenda)
		case 3:
			deletarContato(agenda)
		case 4:
			listarContatos(agenda)
		case 5:
			fmt.Println("Obrigado")
			break loopPrincipal
		}
	}
}

func buscaContato(agenda map[string]string) {
	var nome string
	fmt.Print("Digite o nome: ")
	fmt.Scan(&nome)
	telefoneEncontrado, existe := agenda[nome]
	if existe {
		fmt.Println("Telefone de", nome, ":", telefoneEncontrado)
	} else {
		fmt.Println("Esse nome não existe na lista de contatos.")
	}
}
func adicionarContato(agenda map[string]string) {
	var nome string
	fmt.Print("Digite o nome: ")
	fmt.Scan(&nome)
	var numero string
	fmt.Print("Digite o numero no formato (3198888-7777): ")
	fmt.Scan(&numero)
	if len(numero) != 12 {
		fmt.Println("Telefone deve ter exatamente 12 dígitos")
	} else {
		agenda[nome] = numero
		fmt.Println("o contato: ", nome, numero, "foi adicionado com sucesso")
	}
}

func deletarContato(agenda map[string]string) string {
	var nome string
	fmt.Print("Digite o nome para apagar: ")
	fmt.Scan(&nome)
	_, existe := agenda[nome]
	if existe {
		delete(agenda, nome)
		fmt.Println("Contato", nome, "deletado com sucesso!")
	} else {
		fmt.Println("Erro: Esse contato não existe.")
	}
	return ""
}

func listarContatos(agenda map[string]string) string {
	if len(agenda) == 0 {
		fmt.Println("A agenda está vazia.")
		return ""
	}
	fmt.Println("\n--- Lista de Contatos ---")
	for nome, telefone := range agenda {
		fmt.Println("Nome:", nome, "- Telefone:", telefone)
	}
	return ""
}
