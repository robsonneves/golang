package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()

	/*if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa")
	} else {
		fmt.Println("Não conheço esse comando.")
	}*/

	switch leComando() {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 3:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando.")
		os.Exit(-1)
	}
}

func exibeIntroducao() {

	nome := "João"
	versao := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Está é a versão ", versao)
}

func exibeMenu() {

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variavél comando é ", &comando)
	fmt.Println("O comando escolhido foi ", comando)
	return comando
}
