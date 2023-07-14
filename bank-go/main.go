package main

import (
	"fmt"

	"intro-golang/bank-go/accounts"
)

func main() {
	// Terceira forma de atribuir valores a uma struct
	var account1 accounts.CurrentAccount
	account1.HolderName = "Anna Beatriz"
	account1.Agency = 1
	account1.Account = 789
	account1.Amount = 300.00

	var account2 accounts.CurrentAccount
	account2.HolderName = "Juliana Menezes"
	account2.Agency = 1
	account2.Account = 789
	account2.Amount = 100.00

	fmt.Println("Conta de:", account1.HolderName)
	fmt.Println("Saldo Atual:", account1.Amount)
	fmt.Println(account1.Withdraw(301))
	fmt.Println("Saldo Atual:", account1.Amount)
	fmt.Println()

	fmt.Println(account1.Withdraw(100))
	fmt.Println("Saldo Atual:", account1.Amount)
	fmt.Println()

	fmt.Println(account1.Deposit(200))
	fmt.Println("Saldo Atual:", account1.Amount)
	fmt.Println()

	fmt.Println(account1.Transfer(50, &account2))
	fmt.Println("Saldo Atual:", account1.Amount)
	fmt.Println("Saldo Atual:", account2.Amount)
}
