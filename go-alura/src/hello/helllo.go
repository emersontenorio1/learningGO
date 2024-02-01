package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

func main() {
	exibeIntroducao()
	for {
		exibeMenu()

		fmt.Print("Digite uma opção: ")
		comando := lerComando()
		fmt.Println("O comando escolhido foi: ", comando)

		//exibenomes()

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
	fmt.Scan(&comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {

	/*var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"*/
	sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br",
		"https://www.caelum.com.br"}

	fmt.Println("Iniciando Monitoramento .....")
	for i := 0; i < 5; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(5 * time.Second)
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
func exibenomes() {
	nomes := []string{"Emerson", "Bernado", "Matheus"}
	reflect.TypeOf(nomes)
	fmt.Print(cap(nomes))
}
