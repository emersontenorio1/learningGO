package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
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
			lerLogs()
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
	/*sites := []string{
	"https://random-status-code.herokuapp.com/",
	"https://www.alura.com.br",
	"https://www.caelum.com.br"}*/

	sites := leSiteDoArquivo()

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
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("ocorreu um erro na requisição", err)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSiteDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites

}

func exibenomes() {
	nomes := []string{"Emerson", "Bernado", "Matheus"}
	reflect.TypeOf(nomes)
	fmt.Print(cap(nomes))
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + " - Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func lerLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
