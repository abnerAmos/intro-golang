package main

import (
	"fmt"

	"intro-golang/bank-go/accounts"
	"intro-golang/bank-go/holders"
)

// Efetua um pagamento utilizando interface
func bankSlip(conta verifyAccount, value float64) {
	conta.Withdraw(value)
}

// Interface em GO, utiliza métodos em comum de pacotes diferentes
// Reaproveita métodos em comum de "classes" para efetuar uma função.
type verifyAccount interface {
	Withdraw(value float64) string
}

func main() {

	// Composição modelo 1
	account1 := accounts.CurrentAccount{
		Holder: holders.Holder{
			Name:       "Abner Amos",
			Document:   "123.456.789-12",
			Occupation: "Desenvolvedor"},
		Agency:  1,
		Account: 123}

	account1.Deposit(100)

	// Composição modelo 2
	client := holders.Holder{Name: "Juliana Menezes", Document: "789.456.123-21", Occupation: "Enfermeira"}
	account2 := accounts.CurrentAccount{Holder: client, Agency: 1, Account: 321}

	fmt.Println(account1)
	fmt.Println(account2)

	account2.Deposit(200)

	clientSavings := holders.Holder{Name: "Juliana Menezes"}
	accountSavings := accounts.CurrentAccount{Holder: clientSavings, Agency: 1, Account: 321}
	fmt.Println(accountSavings.Deposit(5000))
	fmt.Println(accountSavings.GetAmount())

	// Efetua um pagamento utilizando interface.
	bankSlip(&accountSavings, 1000)
	fmt.Println(accountSavings.GetAmount())

	fmt.Println("Conta de:", account1.Holder.Name)
	fmt.Println("Saldo Atual:", account1.GetAmount())
	fmt.Println(account1.Withdraw(301))
	fmt.Println("Saldo Atual:", account1.GetAmount())
	fmt.Println()

	fmt.Println(account1.Withdraw(100))
	fmt.Println("Saldo Atual:", account1.GetAmount())
	fmt.Println()

	fmt.Println(account1.Deposit(200))
	fmt.Println("Saldo Atual:", account1.GetAmount())
	fmt.Println()

	fmt.Println(account1.Transfer(50, &account2))
	fmt.Println("Saldo Atual:", account1.GetAmount())
	fmt.Println("Saldo Atual:", account2.GetAmount())
}
