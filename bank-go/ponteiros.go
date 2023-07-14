package main

import (
	"fmt"
	"intro-golang/bank-go/accounts"
)

type Ponteiro struct {
	holderName string
	agency     int
	account    int
	amount     float64
}

func mainPonteiro() {

	// Primeira forma de atribuir valores a uma struct
	account1 := accounts.CurrentAccount{
		HolderName: "Abner Amos",
		Agency:     1,
		Account:    123,
		Amount:     100.00}

	fmt.Println(account1)

	// // Segunda forma de atribuir valores a uma struct
	// account2 := accounts.CurrentAccount{
	// 	"Juliana Meneze",
	// 	1,
	// 	456,
	// 	200.00}

	// fmt.Println(account2)

	// Terceira forma de atribuir valores a uma struct
	var account3 accounts.CurrentAccount
	account3.HolderName = "Anna Beatriz"
	account3.Agency = 1
	account3.Account = 789
	account3.Amount = 300.00

	fmt.Println(account3)
	// Quarta forma (com ponteiros) de atribuir valores a uma struct
	var account4 *accounts.CurrentAccount
	account4 = new(accounts.CurrentAccount)
	account4.HolderName = "Alicia Silva"
	account4.Agency = 1
	account4.Account = 789
	account4.Amount = 400.00

	var account5 *accounts.CurrentAccount
	account5 = new(accounts.CurrentAccount)
	account5.HolderName = "Alicia Silva"
	account5.Agency = 1
	account5.Account = 789
	account5.Amount = 400.00

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
