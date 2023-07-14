package main

import (
	"fmt"

	"github.com/abneramos/intro-golang/bank-go/accounts"
)

func main() {
	// Terceira forma de atribuir valores a uma struct
	var account1 accounts.CurrentAccount
	account1.holderName = "Anna Beatriz"
	account1.agency = 1
	account1.account = 789
	account1.amount = 300.00

	var account2 accounts.CurrentAccount
	account2.holderName = "Juliana Menezes"
	account2.agency = 1
	account2.account = 789
	account2.amount = 100.00

	fmt.Println("Conta de:", account1.holderName)
	fmt.Println("Saldo Atual:", account1.amount)
	fmt.Println(account1.withdraw(301))
	fmt.Println("Saldo Atual:", account1.amount)
	fmt.Println()

	fmt.Println(account1.withdraw(100))
	fmt.Println("Saldo Atual:", account1.amount)
	fmt.Println()

	fmt.Println(account1.deposit(200))
	fmt.Println("Saldo Atual:", account1.amount)
	fmt.Println()

	fmt.Println(account1.transfer(50, &account2))
	fmt.Println("Saldo Atual:", account1.amount)
	fmt.Println("Saldo Atual:", account2.amount)
}
