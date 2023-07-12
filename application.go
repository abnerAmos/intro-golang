package main

import (
	"fmt"
	"os"
)

func main() {
	showMenu()

	switch selectOption() {
	case 1:
		fmt.Println("Monitorando...")
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

// função: void
func showMenu() {
	name := "Abner Amos"
	version := 1.10
	fmt.Println("Hello,", name)
	fmt.Println("Version: ", version)

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
