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

	var nome2 = "Robson2"
	var idade = 10
	var versao2 = 1.2
	fmt.Println("Olá, Sr.", nome2, " e sua idade é ", idade)
	fmt.Println("Está é a versão ", versao2)

	fmt.Println("A variavél nome é do tipo ", reflect.TypeOf(nome))
	fmt.Println("A variavél idade é do tipo ", reflect.TypeOf(idade))
	fmt.Println("A variavél versao2 é do tipo ", reflect.TypeOf(versao2))

	nome3 := "Robson2"
	idade3 := 10
	versao3 := 1.2
	fmt.Println("Olá, Sr.", nome3, " e sua idade é ", idade3)
	fmt.Println("Está é a versão ", versao3)

}
