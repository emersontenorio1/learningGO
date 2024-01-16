package main

import (
	"fmt"
)

func main() {
	var nome string
	fmt.Println("Qual o seu nome?")
	fmt.Scan(&nome)
	var idade int
	fmt.Println("Digite sua idade?")
	fmt.Scan(&idade)
	var versao float32 = 1.1
	fmt.Println("Olá senhor, ", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")
	var comando int
	fmt.Scanf("%d", &comando)
}
