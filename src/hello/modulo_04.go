package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	for {
		exibeIntroducao()
		exibeMenu()

		switch leComando() {
		case 1:
			iniciarMonitoramneto()
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

func iniciarMonitoramneto() {

	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, " foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, " o site esta com problema", resp.StatusCode)
	}
}
