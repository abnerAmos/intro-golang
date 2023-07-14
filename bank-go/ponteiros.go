package main

import "fmt"

type Ponteiro struct {
	holderName string
	agency     int
	account    int
	amount     float64
}

func mainPonteiro() {

	// Primeira forma de atribuir valores a uma struct
	account1 := CurrentAccount{
		holderName: "Abner Amos",
		agency:     1,
		account:    123,
		amount:     100.00}

	fmt.Println(account1)

	// Segunda forma de atribuir valores a uma struct
	account2 := CurrentAccount{
		"Juliana Meneze",
		1,
		456,
		200.00}

	fmt.Println(account2)

	// Terceira forma de atribuir valores a uma struct
	var account3 CurrentAccount
	account3.holderName = "Anna Beatriz"
	account3.agency = 1
	account3.account = 789
	account3.amount = 300.00

	fmt.Println(account3)
	// Quarta forma (com ponteiros) de atribuir valores a uma struct
	var account4 *CurrentAccount
	account4 = new(CurrentAccount)
	account4.holderName = "Alicia Silva"
	account4.agency = 1
	account4.account = 789
	account4.amount = 400.00

	var account5 *CurrentAccount
	account5 = new(CurrentAccount)
	account5.holderName = "Alicia Silva"
	account5.agency = 1
	account5.account = 789
	account5.amount = 400.00

	// Verificando o valor do conteudo sem ponteiro
	fmt.Println(*account4)
	fmt.Println(*account5)
	fmt.Println()

	// Verificando o valor do conteúdo com ponteiro "&"
	fmt.Println(account4)
	fmt.Println(account5)
	fmt.Println()

	// Verificando o valor do ponteiro
	fmt.Println(&account4)
	fmt.Println(&account5)
	fmt.Println()

	// Comparando valores do conteúdo: true
	fmt.Println(*account4 == *account5)

	// Comparando ponteiros: false
	fmt.Println(account4 == account5)
}
