package main

import (
	"fmt"
	"net/http"
	"os"
)

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
			verifySite()
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

func verifySite() {
	fmt.Println("Monitorando...")
	site := "https://httpstat.us/Random/200,404"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site foi carregado com sucesso!")
	} else {
		fmt.Println("Site esta com problemas; Status Code:", resp.StatusCode)
	}
}
