package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const quantityMonitoring = 3
const delay = 5

func main() {

	name, _ := nomeIdade() // "Underline" ignora o retorno que não vai utilizar
	version := 1.10
	fmt.Println("Hello,", name)
	fmt.Println("Version: ", version)

	/* Na linguagem GO não existe "while";
	*  para gerar um loop infinito, basta declara um for sem condição alguma */
	for {
		showMenu()

		switch selectOption() {
		case 1:
			monitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Encerrando...")
			os.Exit(0)
		default:
			fmt.Println("Comando não identificado.")
			os.Exit(-1)
		}
	}
}

// função: void
func showMenu() {
	fmt.Println("\n1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Encerrar")
}

// função: Integer
func selectOption() int {
	var command int
	fmt.Scan(&command) // Scan, identifica o valor da váriável sem necessidade de fazer replace.
	fmt.Println("Comando escolhido:", command)
	return command
}

// função: retornando mais de 1 tipo de variável
func nomeIdade() (string, int) {
	nome := "Abner"
	idade := 29
	return nome, idade
}

func monitoring() {
	fmt.Println("Monitorando...")

	sites := []string{
		"https://httpstat.us/Random/200,404",
		"https://httpstat.us/Random/200",
		"https://httpstat.us/Random/200,400"}

	for i := 0; i < quantityMonitoring; i++ {
		for i, site := range sites {
			fmt.Println("Verificando site", i, ":", site)
			verifySite(site)
		}
		fmt.Println()
		time.Sleep(delay * time.Second)
	}
}

func verifySite(site string) {

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site foi carregado com sucesso!")
	} else {
		fmt.Println("Site esta com problemas; Status Code:", resp.StatusCode)
	}
}
