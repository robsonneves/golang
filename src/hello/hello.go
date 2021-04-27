package main

import (
	"fmt"
	"reflect"
)

func main() {

	var nome string = "Robson"
	var versao float32 = 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Está é a versão ", versao)

	fmt.Println("--------------------------------------")

	var nome2 = "Robson2"
	var idade = 10
	var versao2 = 1.2
	fmt.Println("Olá, Sr.", nome2, " e sua idade é ", idade)
	fmt.Println("Está é a versão ", versao2)

	fmt.Println("A variavél nome é do tipo ", reflect.TypeOf(nome))
	fmt.Println("A variavél idade é do tipo ", reflect.TypeOf(idade))
	fmt.Println("A variavél versao2 é do tipo ", reflect.TypeOf(versao2))

	fmt.Println("--------------------------------------")

	nome3 := "Robson2"
	idade3 := 10
	versao3 := 1.2
	fmt.Println("Olá, Sr.", nome3, " e sua idade é ", idade3)
	fmt.Println("Está é a versão ", versao3)

	fmt.Println("--------------------------------------")

	nome4 := "Robson"
	versao4 := 1.1
	fmt.Println("Olá, Sr.", nome4)
	fmt.Println("Está é a versão ", versao4)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variavél comando é ", &comando)
	fmt.Println("O comando escolhido foi ", comando)

	fmt.Println("--------------------------------------")


	

}
