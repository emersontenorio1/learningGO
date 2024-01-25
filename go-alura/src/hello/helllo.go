package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	for {
		exibeMenu()

		fmt.Print("Digite uma opção: ")
		comando := lerComando()
		fmt.Println("O comando escolhido foi: ", comando)

		/*if comando == 1{

		} else if comando == 2{

		}else if comando == 0{

		}else{
			fmt.Println("Não conheço este comando")
		}*/

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			break
		case 0:
			fmt.Println("Saindo do Programa......")
			os.Exit(0)
			break
		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	var nome string
	fmt.Println("Qual o seu nome?")
	fmt.Scan(&nome)
	//var idade int
	//fmt.Println("Digite sua idade?")
	//fmt.Scan(&idade)
	var versao float32 = 1.1
	fmt.Println("Olá senhor, ", nome /*"sua idade é", idade*/)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do Programa")
}
func lerComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Print("Qual site deseja verificar: ")
	var site string
	fmt.Scan(&site)
	fmt.Println("Iniciando Monitoramento .....")
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

}
