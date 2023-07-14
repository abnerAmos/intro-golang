package accounts

import (
	"fmt"
	"intro-golang/bank-go/holders"
)

// Ao declarar uma variável com a primeira letra minuscula será "private"
// Ao declarar uma variável com a primeira letra maiuscula será "public"
type SavingsAccount struct {
	Holder  holders.Holder
	Agency  int
	Account int
	amount  float64
}

func (c *SavingsAccount) GetAmount() float64 {
	return c.amount
}

/* Podemos parametrizar na função a conta que será apontada para saque "(c *CurrentAccount)";
* Desta forma a struct que foi apontada no mótodo pode chamar o método diretamente "account1.withdraw()". */
func (c *SavingsAccount) Withdraw(value float64) string {
	containsAmount := value > 0 && value <= c.amount
	if containsAmount {
		c.amount -= value
		fmt.Println("Valor do saque:", value)
		return "Saque efetuado com sucesso!"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *SavingsAccount) Deposit(value float64) (string, float64) {
	isPositive := value > 0
	if isPositive {
		c.amount += value
		return "Deposito efetuado com sucesso. Saldo atual:", c.amount
	} else {
		return "Deposito não efetuado. Saldo atual:", c.amount
	}
}

// TODO: - Só deve transferir para a própria conta

// func (c *SavingsAccount) Transfer(value float64, receiver *CurrentAccount) string {
// 	containsAmount := value > 0 && value <= c.amount
// 	if containsAmount {
// 		c.amount -= value
// 		receiver.Deposit(value)
// 		return "Transferência efetuada com sucesso"
// 	} else {
// 		return "Transferência não efetuada"
// 	}
// }
