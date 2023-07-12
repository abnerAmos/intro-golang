package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// retorna o ponteiro do file
func readFilesWithOS() {

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(file)
}

// retorna os bites de file
func readFilesWithIOUtil() {

	file, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(string(file))
}

func readFilesWithBufio() {

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n') // faz a leitura do bite
		line = strings.TrimSpace(line)       // remo os espaço e quebras de linha da string

		fmt.Println(line)

		if err != io.EOF {
			break
		}
	}

	file.Close() // fecha o arquivo, após utilização.
}
