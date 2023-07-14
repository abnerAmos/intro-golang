package accounts

import "fmt"

// Ao declarar uma variável com a primeira letra minuscula será "private"
// Ao declarar uma variável com a primeira letra maiuscula será "public"
type CurrentAccount struct {
	HolderName string
	Agency     int
	Account    int
	Amount     float64
}

/* Podemos parametrizar na função a conta que será apontada para saque "(c *CurrentAccount)";
* Desta forma a struct que foi apontada no mótodo pode chamar o método diretamente "account1.withdraw()". */
func (c *CurrentAccount) Withdraw(value float64) string {
	containsAmount := value > 0 && value <= c.Amount
	if containsAmount {
		c.Amount -= value
		fmt.Println("Valor do saque:", value)
		return "Saque efetuado com sucesso!"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *CurrentAccount) Deposit(value float64) (string, float64) {
	isPositive := value > 0
	if isPositive {
		c.Amount += value
		return "Deposito efetuado com sucesso. Saldo atual:", c.Amount
	} else {
		return "Deposito não efetuado. Saldo atual:", c.Amount
	}
}

func (c *CurrentAccount) Transfer(value float64, receiver *CurrentAccount) string {
	containsAmount := value > 0 && value <= c.Amount
	if containsAmount {
		c.Amount -= value
		receiver.Deposit(value)
		return "Transferência efetuada com sucesso"
	} else {
		return "Transferência não efetuada"
	}
}
