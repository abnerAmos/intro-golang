package main

import (
	"fmt"
)

/* if else, não precisa de parênteses para declarar condição, e é necessário deixar explicito a comparação */
func ifElse() {

	if selectOpt() == 1 {
		fmt.Println("Monitorando...")
	} else if selectOpt() == 2 {
		fmt.Println("Exibindo Logs...")
	} else if selectOpt() == 0 {
		fmt.Println("Encerrando...")
	} else {
		fmt.Println("Comando não identificado.")
	}
}

func selectOpt() int {
	var command int
	fmt.Scan(&command) // Scan, identifica o valor da váriável sem necessidade de fazer replace.
	fmt.Println("Comando escolhido:", command)
	return command
}
