package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			fmt.Println(printLogs())
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

	sites := readFile()

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

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("Site esta com problemas; Status Code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func readFile() []string {

	var sites []string

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // faz a leitura do bite
		line = strings.TrimSpace(line)       // remo os espaço e quebras de linha da string

		sites = append(sites, line) // adiciona a linha de um site ao um index de um array

		if err == io.EOF { // erro que indica o fim dos dados de um arquivo "End Of File"
			break // interrompe o loop infinito
		}
	}

	file.Close()
	return sites
}

func arraySites() []string {
	sites := []string{
		"https://httpstat.us/Random/200,404",
		"https://httpstat.us/Random/200",
		"https://httpstat.us/Random/200,400"}

	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	logTime := time.Now().Format("02/01/2006 15:04:05")
	file.WriteString(logTime + " - " + site + " - Online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func printLogs() string {
	fmt.Println("Exibindo Logs...")

	file, err := ioutil.ReadFile("log.txt") // ReadFile le o arquivo e fecha em seguida automaticamente.
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	return string(file)
}
